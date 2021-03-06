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
	var files []db.File
	db.GetNewFiles(&files)
	for _, file := range files {
		upload, err := uplaodImage(&file)
		// if all good set the status of the file as "uploaded"
		if err == nil {
			db.SetFileStatus(&file, "uploaded")
			_, err := postMeta(&upload)
			// if all good set the status of the file as "completed"
			if err == nil {
				db.SetFileStatus(&file, "completed")
			}
		}
	}

}
