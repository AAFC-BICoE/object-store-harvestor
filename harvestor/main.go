// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md
package main

import (
	"harvestor/config"
	l "harvestor/logger"
	"os"
)

func main() {

	// Getting our Configuration
	filename := getFileName()
	config.Load(filename)
	// Create new logger
	var logger = l.NewLogger()

	// Debug Log
	conf := config.GetConf()
	logger.Debug("conf.Database.DbName : ", conf.Database.DbName)
	logger.Debug("conf.Walker.EntryPoint : ", conf.Walker.EntryPoint)
	logger.Debug("conf.HttpClient.ApiUrl : ", conf.HttpClient.ApiUrl)
	logger.Debug("conf.HttpClient.ObjectSource : ", conf.HttpClient.ObjectSource)
	logger.Debug("conf.HttpClient.TimeOut : ", conf.HttpClient.TimeOut)
	logger.Debug("conf.App.Name : ", conf.App.Name)
	logger.Debug("conf.App.Release : ", conf.App.Release)
	logger.Debug("conf.App.Env : ", conf.App.Env)
	logger.Debug("conf.Loger.Level : ", conf.Logger.Level)
}

// helper function to read args
func getFileName() string {
	var logger = l.NewLogger()

	args := os.Args
	if len(os.Args) < 1 {
		example := "(example : /app/harvestor_config.yml)"
		err := "Application requires an argument as a string to a config file, none has been provided ||| " + example
		logger.Error(err)
	}
	logger.Debug("args :", args)
	filename := args[1]
	logger.Debug("filename :", filename)
	return filename
}
