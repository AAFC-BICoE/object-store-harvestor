package db

import (
	"harvestor/config"
	l "harvestor/logger"
	"os"
	"time"
)

type File struct {
	ID        int       `json:"id" gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	Path      string    `json:"path" gorm:"uniqueIndex"`
	Name      string    `json:"name"`
	ModTime   time.Time `json:"mod_at" gorm:"index:idx_file_mod_time"`
	Status    string    `json:"status" gorm:"type:varchar(64)"`
	CreatedAt time.Time `json:"created_at" gorm:"index:idx_file_created_at"`
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
	Create(path string, info os.FileInfo) error
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

func (f File) Create(path string, info os.FileInfo) error {
	// get logger
	var logger = l.NewLogger()
	// get config
	conf := config.GetConf()
	// get DB instance
	db := GetHarvesterDB()
	// define absolute path
	absolutePath := conf.Walker.Path() + string(os.PathSeparator) + path
	// validation | check if the record already exist
	if doesNotExist(absolutePath) {
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
			return err
		}
		logger.Info("File record has been stored in DB for :", f.GetPath())
		logger.Debug("DB File record : ", logger.PrettyGoStruct(f))
		return err
	}
	return nil
}

// After upload change status from "new" to "complete"
func SetFileStatus(f *File, status string) error {
	// get logger
	var logger = l.NewLogger()
	db := GetHarvesterDB()
	f.Status = status
	err := db.Save(f).Error
	if err == nil {
		logger.Info("File record has been "+status+" for : ", f.GetPath())
	}
	return err
}

// check by absolute path if the file exist in DB already
func doesNotExist(absolutePath string) bool {
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
