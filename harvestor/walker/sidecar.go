package walker

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"os"
	"path/filepath"
	"strings"
)

// validation helpers
var ValidConfigFileExtension = "yml"

type SidecarFile struct {
	AcMetadataCreator  string
	DcCreator          string
	AcDigitizationDate string
	Original           string
	UploadWithFilename string
	Derivative         string
	ManagedAttributes  map[string]string
}

// Define all interfaces for this struct
type ISidecairFile interface {
	GetOriginal() string
	GetDerivative() string
	GetUploadWithFilename() string
}

// Implementation
func (sf SidecarFile) GetOriginal() string {
	return sf.Original
}
func (sf SidecarFile) GetDerivative() string {
	return sf.Derivative
}

func (sf SidecarFile) GetUploadWithFilename() string {
	return sf.UploadWithFilename
}

func processSidecarFromWalker(path string, info os.FileInfo) error {
	// init logger
	var logger = l.NewLogger()
	// get config
	conf := config.GetConf()
	// define absolute path
	absolutePath := conf.Walker.Path() + string(os.PathSeparator) + path
	// Read the content of yml sidecar file
	sidecarFile, err := GetSidecarFile(absolutePath)
	if err != nil {
		return err
	}
	logger.Debug(" sidecarFile : ", logger.PrettyGoStruct(sidecarFile))
	// about to create a record in DB for original file from sidecar
	logger.Debug(" getAbsoluteOriginalPath(sidecarFile, absolutePath) : ", getAbsoluteOriginalPath(sidecarFile, absolutePath))
	fileRecordOriginal, err := createFileRecordOriginal(getAbsoluteOriginalPath(sidecarFile, absolutePath))
	if err != nil {
		logger.Error(" createFileRecordOriginal : ", err)
		return err
	}
	logger.Debug(" sidecarFile : ", logger.PrettyGoStruct(sidecarFile))
	// about to create a record in DB for derivative file from sidecar
	fileRecordDerivative, err := createFileRecordDerivative(getAbsoluteDerivativePath(sidecarFile, absolutePath))
	if err != nil {
		logger.Error(" createFileRecordDerivative : ", err)
		return err
	}
	logger.Debug(" sidecarFile : ", logger.PrettyGoStruct(sidecarFile))
	// about to create a record in DB for sidecar
	s := &db.Sidecar{
		Path:             absolutePath,
		OriginalFileID:   fileRecordOriginal.GetID(),
		DerivativeFileID: fileRecordDerivative.GetID(),
		SidecarStatus:    "new",
	}
	// create sidecar record in the db
	err = db.CreateSidecarRecord(s)
	if err != nil {
		logger.Fatal("Can't create db record for Sidecar : ", absolutePath, " details : ", err)
		return err
	}

	logger.Debug(" fileRecordOriginal : ", logger.PrettyGoStruct(fileRecordOriginal))
	logger.Debug(" fileRecordDerivative : ", logger.PrettyGoStruct(fileRecordDerivative))
	logger.Debug(" sidecarFile : ", logger.PrettyGoStruct(sidecarFile))
	logger.Debug(" Sidecar : ", logger.PrettyGoStruct(s))
	return err
}

func getAbsoluteOriginalPath(sidecarFile *SidecarFile, path string) string {
	return filepath.Dir(path) +
		string(os.PathSeparator) +
		sidecarFile.Original
}

func getAbsoluteDerivativePath(sidecarFile *SidecarFile, path string) string {
	return filepath.Dir(path) +
		string(os.PathSeparator) +
		sidecarFile.Derivative
}

// load content of the sidecar into the struct
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

	// lets run simple validation
	if len(scf.GetOriginal()) == 0 {
		e := errors.New("Original is not set in : " + filename)
		return &scf, e
	}
	if len(scf.GetDerivative()) == 0 {
		e := errors.New("Derivative is not set in : " + filename)
		return &scf, e
	}

	return &scf, err
}

// get absolute path for sidecard
func GetSideCarPathByFilePath(path string) string {
	sideCarPath := strings.TrimSuffix(path, filepath.Ext(path))
	return sideCarPath + ".yml"
}

// get content sidecar yml file
func GetSidecarYmlFile(sidecar *db.Sidecar) (*SidecarFile, error) {
	return GetSidecarFile(sidecar.GetPath())
}

// get content sidecar yml file
func GetSidecarYmlFileByFile(file *db.File) (*SidecarFile, error) {
	path := GetSideCarPathByFilePath(file.GetPath())
	return GetSidecarFile(path)
}

func GetUploadMediaFileName(f *db.File) string {
	// get logger
	var logger = l.NewLogger()
	// check if it's original
	if f.GetUploadType() != "original" {
		// no need to replace
		return f.Name
	}
	// get sidecar file from DB
	s, err := db.GetSidecarByOriginalFile(f)
	// checking errors
	if err != nil {
		logger.Error("Error : GetSidecarByOriginalFile : details : ", err)
		return f.Name
	}
	// get content from sidecar file
	sf, err := GetSidecarFile(s.GetPath())
	// checking errors
	if err != nil {
		logger.Error("Error : GetSidecarFile : details : ", err)
		return f.Name
	}
	// UploadWithFilename is optional
	// checking if it's set
	if len(sf.GetUploadWithFilename()) == 0 {
		return f.Name
	}

	return sf.GetUploadWithFilename()
}

// check if we have a side card file
func HasSideCar(path string) bool {
	// init logger
	var logger = l.NewLogger()
	sideCarPath := GetSideCarPathByFilePath(path)
	_, err := os.Lstat(sideCarPath)
	if err != nil {
		return false
	}
	logger.Debug("side car absolute path : ", sideCarPath, " exist")
	return true
}

// check if the config type supported
func isValidConfigFile(file string) bool {
	return getFileExtension(file) == ValidConfigFileExtension
}
