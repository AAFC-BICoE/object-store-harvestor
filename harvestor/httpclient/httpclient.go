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
	httpClient.RetryMax = conf.HttpClient.GetRetryMax()
	httpClient.RetryWaitMin = time.Duration(conf.HttpClient.GetRetryWaitMin()) * time.Second
	httpClient.HTTPClient.Timeout = time.Duration(conf.HttpClient.GetTimeOut()) * time.Second
	logger := l.NewLogger()
	httpClient.Logger = logger
	logger.Info("Harvestor Http Client has been initialized !!!")
}

func Run() {
	processNewFiles()
	// covering rare case
	processStuckedFiles()
}

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
		processNewFile(&file)
	}
}

func processNewFile(file *db.File) {
	// try to upload
	upload, err := uplaodImage(file)
	// if all good set the status of the file as "uploaded"
	if err == nil {
		db.SetFileStatus(file, "uploaded")
		// try to post meta
		_, err := postMeta(&upload)
		// if all good set the status of the file as "completed"
		if err == nil {
			db.SetFileStatus(file, "completed")
		}
	}
}

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
	// if all good set the status of the file as "completed"
	if err == nil {
		db.SetFileStatus(file, "completed")
	}

}
