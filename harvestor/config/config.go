// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Media File walker + SQLite DB configs
type Configuration struct {
	Walker   FileWalkerConfiguration
	Database DatabaseConfiguration
}

// Validation helpers
var (
	ValidConfigFileExtensions                    = []string{"yml"}
	validConfigFileExtensionsMap map[string]bool = make(map[string]bool)
)

func ReadFromFile(filename string) (Configuration, error) {
	// init Configuration
	var conf Configuration
	// get path and file
	path, file := filepath.Split(filename)
	// validation on valid cpnfiguration
	if !isValidConfigFile(file) {
		err := fmt.Errorf("config file:(%q) is not supported with extention:(%q)", file, getFileExtension(file))
		return conf, err
	}

	confName := getFileName(file)
	confExt := getFileExtension(file)
	// Debug for now
	log.Println("----------------------------------------")
	log.Println("||| path :", path)
	log.Println("||| confName :", confName)
	log.Println("||| confExt :", confExt)
	log.Println("----------------------------------------")

	// passing file to viper
	v := viper.New()
	v.SetConfigName(confName) // config file name without extension
	v.SetConfigType(confExt)  // config file extension
	v.AddConfigPath(path)     // config file path
	// in case we have .env
	v.AutomaticEnv() // read value ENV variable

	// Reading from yml file
	err := v.ReadInConfig()
	if err != nil {
		return conf, err
	}

	// Unmarshal to predefined struct
	err = v.Unmarshal(&conf)
	if err != nil {
		return conf, err
	}

	// returning Configuration
	return conf, nil
}

// check if the config file exist
func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// check if the config type supported
func isValidConfigFile(file string) bool {
	for _, ext := range ValidConfigFileExtensions {
		validConfigFileExtensionsMap[ext] = true
	}

	ext := getFileExtension(file)
	return validConfigFileExtensionsMap[ext]
}

// get file extension
func getFileExtension(filename string) string {
	return strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
}

// get file name without extension
func getFileName(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}
