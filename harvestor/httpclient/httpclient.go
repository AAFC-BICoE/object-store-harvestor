package httpclient

import (
	c "github.com/hashicorp/go-retryablehttp"
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"time"
)

var httpClient *c.Client

func InitHttpClient() {
	httpClient = c.NewClient()
	conf := config.GetConf()
	// set max retry in case server fails to process a request
	httpClient.RetryMax = conf.HttpClient.GetRetryMax()
	// set wait in seconds before makeing the same request after server fails
	httpClient.RetryWaitMin = time.Duration(conf.HttpClient.GetRetryWaitMin()) * time.Second
	// set timeout in seconds after which http client should give up
	httpClient.HTTPClient.Timeout = time.Duration(conf.HttpClient.GetTimeOut()) * time.Second
	// Getting logger
	logger := l.NewLogger()
	// Assign our custom logger to http client logger
	httpClient.Logger = logger
	// we are good here
	logger.Info("Harvestor Http Client has been initialized !!!")
}

func Run() {
	// dealing with brand new files
	processNewFiles()
	// covering rare case when new files were not done E2E
	processStuckedFiles()
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

// Support for sidecar yml files to post managed metadata for derivatives
func processSideCarManagedMeta(file *db.File, meta *db.Meta) error {
	// init conf
	conf := config.GetConf()
	if !conf.SideCar.IsEnabled() {
		return nil
	}
	return postSideCarManagedMeta(file, meta)

}

// process singe new file
func processNewFile(file *db.File) {
	// try to upload
	upload, err := uplaodImage(file)
	// if all good set the status of the file as "uploaded"
	if err == nil {
		db.SetFileStatus(file, "uploaded")
		// try to post meta
		meta, err := postMeta(&upload)
		// if all good set the status of the file as "completed"
		if err == nil {
			db.SetFileStatus(file, "meta-good")
		}
		if processSideCarManagedMeta(file, &meta) == nil {
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
	meta, err := postMeta(upload)
	// if all good set the status of the file as "completed"
	if err == nil && processSideCarManagedMeta(file, &meta) == nil {
		db.SetFileStatus(file, "completed")
	}

}
