// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md
package main

import (
	"harvestor/config"
	"log"
	"os"
)

func main() {
	filename := getFileName()
	conf, err := config.ReadFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Debug Log
	log.Println("INFO ||| conf.Database.DbName :", conf.Database.DbName)
	log.Println("INFO ||| conf.Walker.EntryPoint :", conf.Walker.EntryPoint)
	log.Println("INFO ||| conf.HttpClient.ApiUrl :", conf.HttpClient.ApiUrl)
	log.Println("INFO ||| conf.HttpClient.ObjectSource :", conf.HttpClient.ObjectSource)
	log.Println("INFO ||| conf.HttpClient.TimeOut :", conf.HttpClient.TimeOut)
}

// helper function to read args
func getFileName() string {
	args := os.Args
	if len(os.Args) < 1 {
		example := "(example : /app/harvestor_config.yml)"
		err := "Application requires an argument as a string to a config file, none has been provided ||| " + example
		log.Fatal(err)
	}
	log.Println("args :", args)
	filename := args[1]
	log.Println("filename :", filename)
	return filename
}
