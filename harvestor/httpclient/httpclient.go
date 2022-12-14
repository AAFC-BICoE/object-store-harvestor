package httpclient

import (
	"crypto/tls"
	"crypto/x509"
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"io/ioutil"
	"net/http"
	"time"

	c "github.com/hashicorp/go-retryablehttp"
)

var httpClient *c.Client

func InitHttpClient() {

	// Getting logger
	logger := l.NewLogger()

	conf := config.GetConf()
	httpClient = c.NewClient()

	// depending on the config, allow to append custom certificate
	// inspired by https://forfuncsake.github.io/post/2017/08/trust-extra-ca-cert-in-go-app/
	if conf.HttpClient.GetLocalCertFile() != "" {
		tlsConfig, err := InstallLocalCA(conf.HttpClient.GetLocalCertFile(), logger)
		if err != nil {
			logger.Error("Can't setup LocalCA: ", err, " ... Continuing without it.")
		}

		defaultTransport := httpClient.HTTPClient.Transport
		defaultTransport.(*http.Transport).TLSClientConfig = tlsConfig
	}

	// set max retry in case server fails to process a request
	httpClient.RetryMax = conf.HttpClient.GetRetryMax()
	// set wait in seconds before makeing the same request after server fails
	httpClient.RetryWaitMin = time.Duration(conf.HttpClient.GetRetryWaitMin()) * time.Second
	// set timeout in seconds after which http client should give up
	httpClient.HTTPClient.Timeout = time.Duration(conf.HttpClient.GetTimeOut()) * time.Second

	// Assign our custom logger to http client logger
	httpClient.Logger = logger
	// we are good here
	logger.Info("Harvestor Http Client has been initialized !!!")
}

func InstallLocalCA(localCertFile string, logger *l.StandardLogger) (*tls.Config, error) {
	// Get the SystemCertPool if available otherwise create a new one
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Read in the cert file
	certs, err := ioutil.ReadFile(localCertFile)
	if err != nil {
		return nil, err
	}

	// Append the cert to the SystemCertPool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		logger.Println("No certs appended, using system certs only")
	}

	// Create a new tls.Config with our augmented rootCAs
	return &tls.Config{
		RootCAs: rootCAs,
	}, nil
}

// http client for Bio Cluster
func ClusterRun() {
	// Phase I
	// Upload files and Meta
	// dealing with brand new files
	processNewFiles()
	// covering rare case when new files were not done E2E
	processStuckedFiles()

	// Phase II
	// Building relations
	// dealing with brand new sidecars
	processNewRelations()
}

// http client for PC
func PcRun() {
	// Phase I
	// Upload files and Meta
	// dealing with brand new files
	processNewFiles()
	// covering rare case when new files were not done E2E
	processStuckedFiles()
}

func processNewRelations() {
	// Getting logger
	logger := l.NewLogger()
	// init an empty slice of new sidecars
	var sidecars []db.Sidecar
	// getting all new files
	db.GetNewSidecars(&sidecars)
	// checking if there are any new sidecars
	if len(sidecars) == 0 {
		logger.Info("No new sidecars. Harvester Http Client has nothing to upload for relations !!!")
	}
	// looping new sidecars
	for _, sidecar := range sidecars {
		logger.Info("= = = = = Building relationships E2E current sidecar : ", sidecar.GetPath(), " = = = = =")
		logger.Debug(" sidecar : ", logger.PrettyGoStruct(sidecar))
		processNewSidecar(&sidecar)
		logger.Info("= = = = = E2E is DONE for current sidecar : ", sidecar.GetPath(), " = = = = =")
	}
}

// process singe sidecar file
func processNewSidecar(sidecar *db.Sidecar) {
	// no longer we need managed meta here
	// just post derivatives
	_ = processSideCarDerivative(sidecar)
}

// A wrapper to loop all new files
func processNewFiles() {
	// Getting logger
	logger := l.NewLogger()
	// init an empty slice of new files
	var files []db.File
	// getting all new files
	db.GetNewFiles(&files)
	// checking if there are any new files
	if len(files) == 0 {
		logger.Info("No new files. Harvester Http Client has nothing to upload !!!")
	}
	// looping new files
	for _, file := range files {
		logger.Info("= = = = = Starting E2E current file : ", file.GetName(), " = = = = =")
		processNewFile(&file)
		logger.Info("= = = = = E2E is DONE for current file : ", file.GetName(), " = = = = =")
	}
}

// Support for sidecar yml files to post derivatives
func processSideCarDerivative(sidecar *db.Sidecar) error {
	err := postSideCarDerivative(sidecar)
	if err == nil {
		db.SetSidecarStatus(sidecar, "completed")
	}
	return err

}

// process singe new file
func processNewFile(file *db.File) {
	// try to upload
	upload, err := uplaodImage(file)
	// if all good set the status of the file as "uploaded"
	if err == nil {
		db.SetFileStatus(file, "uploaded")
		// try to post meta
		if file.GetUploadType() == "original" {
			_, err = postMeta(&upload)
		}
		//meta, err := postMeta(&upload)
		// if all good set the status of the file as "completed"
		if err == nil {
			db.SetFileStatus(file, "completed")
		}
	}
}

// A wrapper to loop all stucked files
func processStuckedFiles() {
	// Getting logger
	logger := l.NewLogger()
	// init an empty slice of Stucked files
	var files []db.File
	// getting all Stucked files
	db.GetStuckedFiles(&files)
	// checking if there are any new files
	if len(files) != 0 {
		logger.Warning("Found Stucked files. This is a rare case, please check the logs")
	}
	// looping Stucked files
	for _, file := range files {
		processStuckedFile(&file)
	}
}

// process singe stucked file
func processStuckedFile(file *db.File) {
	// Getting logger
	logger := l.NewLogger()
	// get upload by file
	upload, err := db.GetUploadByFile(file)
	if err != nil {
		logger.Error("file with id : ", file.GetID(), " has error on getting upload from DB :", err)
	}
	logger.Warning("Will try to post Meta for stucked file :", file.GetPath())
	// trying to post meta
	_, err = postMeta(upload)
	//meta, err := postMeta(upload)
	// if all good set the status of the file as "completed"
	if err == nil {
		db.SetFileStatus(file, "completed")
	}

}
