// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md

// Package config provides a functionality to write and read from SQLite DB
// SQLite DB is used as a persistent storage to track image files and actions against object store api
package db

import (
	l "harvestor/logger"
	"time"
)

// File struct to store only image files
type File struct {
	ID         int       `json:"id" gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	Path       string    `json:"path" gorm:"uniqueIndex"`
	Name       string    `json:"name"`
	ModTime    time.Time `json:"mod_at"`
	Status     string    `json:"status" gorm:"type:varchar(64)"`    // status of the file against object store api
	UploadType string    `json:"file_type" gorm:"type:varchar(64)"` // original or derivative (bio cluster only for now)
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Define all interfaces for this struct
type IFile interface {
	GetID() int
	GetPath() string
	GetName() string
	GetModTime() time.Time
	GetStatus() string
	GetUploadType() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// Implementation
func (f File) GetID() int {
	return f.ID
}

func (f File) GetPath() string {
	return f.Path
}

func (f File) GetName() string {
	return f.Name
}

func (f File) GetModTime() time.Time {
	return f.ModTime
}

func (f File) GetStatus() string {
	return f.Status
}

func (f File) GetUploadType() string {
	return f.UploadType
}

func (f File) GetCreatedAt() time.Time {
	return f.CreatedAt
}

func (f File) GetUpdatedAt() time.Time {
	return f.UpdatedAt
}

// Create DB file record
func CreateFile(f *File) error {
	db := GetHarvesterDB()
	return db.Create(f).Error
}

// Set current status for the media file against object store api
func SetFileStatus(f *File, status string) error {
	// get logger
	var logger = l.NewLogger()
	db := GetHarvesterDB()
	f.Status = status
	err := db.Save(f).Error
	if err == nil {
		logger.Info("File record has been '"+status+"' for : ", f.GetPath())
	}
	return err
}

// Set current uplod type for the media file against object store api
func SetUploadType(f *File, uploadType string) error {
	// get logger
	var logger = l.NewLogger()
	db := GetHarvesterDB()
	f.UploadType = uploadType
	err := db.Save(f).Error
	if err == nil {
		logger.Info("File record with upload type : '"+uploadType+"' for : ", f.GetPath())
	}
	return err
}

// check by absolute path if the file exist in DB already
func DoesFileExist(absolutePath string) bool {
	var files []File
	db := GetHarvesterDB()
	db.Where("path = ?", absolutePath).Find(&files)
	return len(files) != 0
}

// get all files with status "new"
func GetNewFiles(files *[]File) {
	var logger = l.NewLogger()
	db := GetHarvesterDB()
	db.Where("status = ?", "new").Find(files)
	logger.Debug("Found for upload total files : ", len(*files))
}

// get all files with status "uploaded", but no metadata
func GetStuckedFiles(files *[]File) {
	var logger = l.NewLogger()
	db := GetHarvesterDB()
	db.Where("status = ?", "uploaded").Find(files)
	logger.Debug("Found total stucked files : ", len(*files))
}

// look up file record in DB by path
func GetFileByPath(path string) (*File, error) {
	var file File
	db := GetHarvesterDB()
	err := db.Where("path = ?", path).First(&file).Error
	return &file, err
}
