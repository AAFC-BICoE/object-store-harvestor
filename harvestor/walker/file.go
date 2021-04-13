package walker

import (
	"harvestor/db"
	l "harvestor/logger"
	"os"
	"path/filepath"
	"strings"
)

func createFileRecordOriginal(filePath string) (*db.File, error) {
	return createFileRecord(filePath, "original")
}

func createFileRecordDerivative(filePath string) (*db.File, error) {
	return createFileRecord(filePath, "derivative")
}

func createFileRecord(filePath string, uploadType string) (*db.File, error) {
	// init logger
	var logger = l.NewLogger()
	// init empty file db struct
	var file db.File
	// get stats
	fileStat, err := os.Stat(filePath)
	if err != nil {
		logger.Fatal("Can't get stats for : ", filePath, " details : ", err)
		return &file, err
	}
	// Create File struct
	file.Path = filePath
	file.Name = fileStat.Name()
	file.ModTime = fileStat.ModTime()
	file.Status = "new"
	file.UploadType = uploadType
	logger.Debug("file db record for file : ", logger.PrettyGoStruct(file))
	if db.DoesFileExist(file.GetPath()) {
		logger.Debug("file exist, do nothing for : ", file.GetPath())
		return &file, nil
	}
	return &file, db.CreateFile(&file)
}

// get file name without extension
func getFileName(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}

// get file extension
func getFileExtension(filename string) string {
	return strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
}
