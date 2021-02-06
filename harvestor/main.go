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
	var logger = l.NewLogger()

	filename := getFileName()
	conf, err := config.ReadFromFile(filename)
	if err != nil {
		logger.Error(err)
	}
	// Debug Log
	logger.Debug("conf.Database.DbName :", conf.Database.DbName)
	logger.Debug("conf.Walker.EntryPoint :", conf.Walker.EntryPoint)
	logger.Debug("conf.HttpClient.ApiUrl :", conf.HttpClient.ApiUrl)
	logger.Debug("conf.HttpClient.ObjectSource :", conf.HttpClient.ObjectSource)
	logger.Debug("conf.HttpClient.TimeOut :", conf.HttpClient.TimeOut)
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
