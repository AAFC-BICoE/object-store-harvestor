// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md

// TODO For now we do not modify config while the app is running
// viper has OnConfigChange event which we can trigger to reload our config during the app run
// Do not see yet this case. May be later on with UI for the user

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

type Configuration struct {
	Database   DatabaseConfiguration   // SQLite DB config
	Walker     FileWalkerConfiguration // Media File walker config
	HttpClient HttpClientConfiguration // Http Client config
	Logger     LoggerConfiguration     // Logger config
	App        AppConfiguration        // App config
}

// validation helpers
var ValidConfigFileExtension = "yml"

// define global empty configuration
var conf Configuration

// Getting Configuration struct
func GetConf() Configuration {
	return conf
}

// Loading from yml config file into our struct
func Load(filename string) {
	err := readFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
}

// Reading from file
func readFromFile(filename string) error {
	// get path and file
	path, file := filepath.Split(filename)

	// validation on supported extention
	if !isValidConfigFile(file) {
		err := fmt.Errorf("config file:(%q) is not supported with extention:(%q)", file, getFileExtension(file))
		return err
	}

	// define file name
	name := getFileName(file)
	// define file extension
	extension := getFileExtension(file)
	// Debug for now
	//log.Println("---------------------- C o n f i g ----------------------------")
	//log.Println("||| config file path :", path)
	//log.Println("||| config file name :", name)
	//log.Println("||| config file extension :", extension)
	//log.Println("---------------------------------------------------------------")

	// init new viper
	v := viper.New()
	// passing file to viper
	// config file name without extension
	v.SetConfigName(name)
	// config file extension
	v.SetConfigType(extension)
	// config file path
	v.AddConfigPath(path)
	// in case we have .env
	// read value ENV variable
	v.AutomaticEnv()

	// Reading from yml file
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	// Unmarshal to predefined struct
	err = v.Unmarshal(&conf)
	if err != nil {
		return err
	}

	// returning Configuration
	return nil
}

// check if the config type supported
func isValidConfigFile(file string) bool {
	return getFileExtension(file) == ValidConfigFileExtension
}

// get file extension
func getFileExtension(filename string) string {
	return strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
}

// get file name without extension
func getFileName(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}
