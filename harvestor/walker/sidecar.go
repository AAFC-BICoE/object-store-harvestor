package walker

import (
	"fmt"
	"github.com/spf13/viper"
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"os"
	"path/filepath"
	"strings"
)

type SidecarFile struct {
	AcMetadataCreator string
	AcCreator         string
	AcDerivedFrom     string
	ManagedAttributes map[string]string
	Tags              string
}

func GetSidecarFile(filename string) (*SidecarFile, error) {
	// init logger
	var logger = l.NewLogger()
	var scf SidecarFile

	// get path and file
	path, file := filepath.Split(filename)

	// validation on supported extention
	if !isValidConfigFile(file) {
		err := fmt.Errorf("config file:(%q) is not supported with extention:(%q)", file, getFileExtension(file))
		logger.Error(err)
	}

	// define file name
	name := getFileName(file)
	// define file extension
	extension := getFileExtension(file)
	// init new viper
	v := viper.New()
	// passing file to viper
	// config file name without extension
	v.SetConfigName(name)
	// config file extension
	v.SetConfigType(extension)
	// config file path
	v.AddConfigPath(path)

	// Reading from yml file
	err := v.ReadInConfig()
	if err != nil {
		logger.Errorf("unable to read in config, %v", err)
		return &scf, err
	}

	// Unmarshal to predefined struct
	err = v.Unmarshal(&scf)
	if err != nil {
		logger.Errorf("unable to decode into struct, %v", err)
		return &scf, err
	}

	return &scf, err
}

// validation helpers
var ValidConfigFileExtension = "yml"

// get absolute path for sidecard
func GetSideCarPathByFilePath(path string) string {
	sideCarPath := strings.TrimSuffix(path, filepath.Ext(path))
	return sideCarPath + ".yml"
}

// check if we have a side card file
func HasSideCar(path string) bool {
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	if !conf.SideCar.IsEnabled() {
		return true
	}
	sideCarPath := GetSideCarPathByFilePath(path)
	_, err := os.Lstat(sideCarPath)
	if err != nil {
		return false
	}
	logger.Debug("side car absolute path : ", sideCarPath, " exist")
	return true
}

func CreateSideCarByFile(file *db.File) error {
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	if !conf.SideCar.IsEnabled() {
		return nil
	}
	logger.Debug("Walker:SideCar found :", file.GetPath())
	if HasSideCar(file.GetPath()) {
		var s = db.Sidecar{
			Path:   GetSideCarPathByFilePath(file.GetPath()),
			Status: "new",
			FileID: file.GetID(),
		}
		err := db.CreateSidecar(&s)
		if err != nil {
			errMsg := "SideCar record CAN NOT be stored in DB for :"
			logger.Error(errMsg, s.GetPath(), err)
			return err
		}
		logger.Info("SideCar record has been stored in DB for :", s.GetPath())
		//GetSidecarFile(s.GetPath())
		logger.Debug("DB side car record : ", logger.PrettyGoStruct(s))
		return err
	}
	return nil
}

// check if the config type supported
func isValidConfigFile(file string) bool {
	return getFileExtension(file) == ValidConfigFileExtension
}

// get file name without extension
func getFileName(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}
