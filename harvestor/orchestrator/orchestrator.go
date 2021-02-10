package orchestrator

import (
	"harvestor/config"
	l "harvestor/logger"
	"harvestor/walker"
)

func Run() {
	// Create new logger
	var logger = l.NewLogger()

	walker.Run()
	// Debug Log
	conf := config.GetConf()
	//logger.Debug("conf.Database.DBFile() : ", conf.Database.DBFile())
	//logger.Debug("conf.Walker.Path() : ", conf.Walker.Path())
	//logger.Debug("conf.HttpClient.GetBaseApiUrl() : ", conf.HttpClient.GetBaseApiUrl())
	//logger.Debug("conf.HttpClient.GetUri() : ", conf.HttpClient.GetUri())
	//logger.Debug("conf.HttpClient.GetTimeOut() : ", conf.HttpClient.GetTimeOut())
	logger.Debug("conf.App.GetName() : ", conf.App.GetName())
	logger.Debug("conf.App.GetRelease() : ", conf.App.GetRelease())
	logger.Debug("conf.App.GetEnvironment() : ", conf.App.GetEnvironment())
	logger.Debug("conf.Loger.GetLevel() : ", conf.Logger.GetLevel())
}
