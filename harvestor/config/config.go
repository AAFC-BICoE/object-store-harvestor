// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md

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
}

// Validation helpers
var (
	ValidConfigFileExtensions                    = []string{"yml"}
	validConfigFileExtensionsMap map[string]bool = make(map[string]bool)
)

// define global empty configuration
var conf Configuration

func Load(filename string) {
	err := readFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
}

func Reload(filename string) {
	err := readFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConf() Configuration {
	return conf
}

func GetLoggerLevel() string {
	return strings.ToLower(conf.Logger.Level)
}

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
	log.Println("---------------------- C o n f i g ----------------------------")
	log.Println("||| config file path :", path)
	log.Println("||| config file name :", name)
	log.Println("||| config file extension :", extension)
	log.Println("---------------------------------------------------------------")

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
