package walker

import (
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"os"
	"path/filepath"
	"strings"
)

/*
type Sidecard struct {
	AcMetadataCreator string
	DcCreator         string
	AcDerivedFrom     string
	ManagedAttributes []Attribute
	Tags              []Tag
}
*/

// validation helpers
var ValidConfigFileExtension = "yml"

// get absolute path for sidecard
func GetSideCardPathByFilePath(path string) string {
	sideCardPath := strings.TrimSuffix(path, filepath.Ext(path))
	return sideCardPath + ".yml"
}

// check if we have a side card file
func HasSideCard(path string) bool {
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	if !conf.SideCard.IsEnabled() {
		return true
	}
	sideCardPath := GetSideCardPathByFilePath(
		conf.Walker.Path() + string(os.PathSeparator) + path)
	logger.Debug("sideCardAbsolutePath : ", sideCardPath)
	_, err := os.Lstat(sideCardPath)
	if err != nil {
		return false
	}
	return true
}

func CreateSideCardByFile(file *db.File) error {
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	if !conf.SideCard.IsEnabled() {
		return nil
	}
	if HasSideCard(file.GetPath()) {
		var s = db.Sidecard{
			Path:   GetSideCardPathByFilePath(file.GetPath()),
			FileID: file.GetID(),
		}
		logger.Debug("DB SideCard record : ", logger.PrettyGoStruct(s))
		err := db.CreateSidecard(&s)
		if err != nil {
			errMsg := "SideCard record CAN NOT be stored in DB for :"
			logger.Error(errMsg, s.GetPath(), err)
			return err
		}
		logger.Info("SideCard record has been stored in DB for :", s.GetPath())
		logger.Debug("DB SideCard record : ", logger.PrettyGoStruct(s))
		return err
	}
	return nil
}
