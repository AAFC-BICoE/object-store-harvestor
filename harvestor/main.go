// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md
package main

import (
	"harvestor/config"
	"harvestor/db"
	"harvestor/httpclient"
	"harvestor/orchestrator"
	"log"
	"os"
)

// main returns nothing
// this is the entry point for the app
func main() {
	// Getting our Configuration
	filename := getFileName()
	// Load yml config file
	config.Load(filename)
	// DB Init
	db.Init()
	// httpclient Init
	httpclient.InitHttpClient()
	// Running orchestrator
	orchestrator.Run()
}

// very simple function to process arg
// could be doene with Package flag
// just to simple to use OS
// helper function to read args
func getFileName() string {
	// assign all args
	args := os.Args
	// check if any args are present
	// first arg is always an executable file
	if len(args) == 1 {
		example := "(example : /app/harvestor_config.yml)"
		err := "Application requires an argument as a string to a config file, none has been provided ||| " + example
		log.Fatal(err)
	}
	// assign the arg as a yml config file name
	filename := args[1]
	// return it
	return filename
}
