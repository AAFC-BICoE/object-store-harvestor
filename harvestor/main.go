// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md
package main

import (
	"harvestor/config"
	"harvestor/db"
	"harvestor/orchestrator"
	"log"
	"os"
	_ "time"
)

func main() {
	// Getting our Configuration
	filename := getFileName()
	config.Load(filename)

	// DB Init
	db.Init()

	// Running orchestrator
	orchestrator.Run()

	// If you need to ssh to the container before exit
	// Uncomment the following line (it will be up for 3 min)
	// time.Sleep(300 * time.Second) // sleep for 5 min before exiting
}

// helper function to read args
func getFileName() string {
	args := os.Args
	if len(args) == 1 {
		example := "(example : /app/harvestor_config.yml)"
		err := "Application requires an argument as a string to a config file, none has been provided ||| " + example
		log.Fatal(err)
	}
	log.Println("args :", args)
	filename := args[1]
	log.Println("filename :", filename)
	return filename
}
