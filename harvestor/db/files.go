package db

import (
	_ "harvestor/config"
	l "harvestor/logger"
	"os"
	"time"
)

type File struct {
	ID        int       `json:"id" gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	Path      string    `json:"path" gorm:"uniqueIndex"`
	Name      string    `json:"name"`
	ModTime   time.Time `json:"mod_at"`
	Status    string    `json:"status" gorm:"type:varchar(64)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Define all interfaces for this struct
type IFile interface {
	GetID() int
	GetPath() string
	GetName() string
	GetModTime() time.Time
	GetStatus() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	CreateFile(path string, info os.FileInfo) (*File, error)
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

func (f File) GetCreatedAt() time.Time {
	return f.CreatedAt
}

func (f File) GetUpdatedAt() time.Time {
	return f.UpdatedAt
}

func CreateFile(absolutePath string, info os.FileInfo) (*File, error) {
	var f File
	// get logger
	var logger = l.NewLogger()
	// get config
	//conf := config.GetConf()
	// get DB instance
	db := GetHarvesterDB()
	// double check if the record already there or not
	if doesFileNotExist(absolutePath) {
		// Create new one
		err := db.FirstOrCreate(&f,
			File{
				Path:    absolutePath,
				Name:    info.Name(),
				ModTime: info.ModTime(),
				Status:  "new",
			}).Error
		if err != nil {
			errMsg := "File record CAN NOT be stored in DB for :"
			logger.Error(errMsg, f.GetPath(), err)
			return &f, err
		}
		logger.Info("File record has been stored in DB for :", f.GetPath())
		logger.Debug("DB File record : ", logger.PrettyGoStruct(f))
		return &f, err
	}
	return &f, nil
}

// After upload change status from "new" to "complete"
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

// check by absolute path if the file exist in DB already
func doesFileNotExist(absolutePath string) bool {
	var files []File
	db := GetHarvesterDB()
	db.Where("path = ?", absolutePath).Find(&files)
	return len(files) == 0
}

// get all files with status "new"
func GetNewFiles(files *[]File) {
	var logger = l.NewLogger()
	db := GetHarvesterDB()
	db.Where("status = ?", "new").Find(files)
	logger.Debug("Found for upload total files : ", len(*files))
}

// get all files with status "uploaded"
func GetStuckedFiles(files *[]File) {
	var logger = l.NewLogger()
	db := GetHarvesterDB()
	db.Where("status = ?", "uploaded").Find(files)
	logger.Debug("Found total stucked files : ", len(*files))
}

func GetFileByPath(path string) (*File, error) {
	var file File
	db := GetHarvesterDB()
	err := db.Where("path = ?", path).First(&file).Error
	return &file, err
}

/*
func GetFileBySidecar(sidecar *Sidecar) (*File, error) {
	var file File
	db := GetHarvesterDB()
	err := db.Where("id = ?", sidecar.GetFileID()).First(&file).Error
	return &file, err
}
*/
