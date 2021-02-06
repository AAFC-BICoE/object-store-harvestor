// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md

package config

import (
	"fmt"
	"github.com/spf13/viper"
	l "harvestor/logger"
	"path/filepath"
	"strings"
)

type Configuration struct {
	Database   DatabaseConfiguration   // SQLite DB config
	Walker     FileWalkerConfiguration // Media File walker config
	HttpClient HttpClientConfiguration // Http Client config
}

// Validation helpers
var (
	ValidConfigFileExtensions                    = []string{"yml"}
	validConfigFileExtensionsMap map[string]bool = make(map[string]bool)
)

func ReadFromFile(filename string) (Configuration, error) {
	var logger = l.NewLogger()

	// define empty configuration
	var conf Configuration

	// get path and file
	path, file := filepath.Split(filename)

	// validation on supported extention
	if !isValidConfigFile(file) {
		err := fmt.Errorf("config file:(%q) is not supported with extention:(%q)", file, getFileExtension(file))
		return conf, err
	}

	// define file name
	name := getFileName(file)
	// define file extension
	extension := getFileExtension(file)
	// Debug for now
	logger.Debug("----------------------------------------")
	logger.Debug("||| config file path :", path)
	logger.Debug("||| config file name :", name)
	logger.Debug("||| config file extension :", extension)
	logger.Debug("----------------------------------------")

	// init new viper
	v := viper.New()
	// passing file to viper
	v.SetConfigName(name)      // config file name without extension
	v.SetConfigType(extension) // config file extension
	v.AddConfigPath(path)      // config file path
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
