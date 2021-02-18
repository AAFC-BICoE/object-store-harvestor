package db

import (
	"gorm.io/gorm"
	"harvestor/config"
	l "harvestor/logger"
	"os"
	"time"
)

type File struct {
	ID        int       `json:"id" gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	Path      string    `json:"path" gorm:"uniqueIndex"`
	Name      string    `json:"name"`
	ModTime   time.Time `json:"mod_at" gorm:"index:idx_mod_time"`
	Status    string    `json:"status" gorm:"type:varchar(64)"`
	CreatedAt time.Time `json:"created_at" gorm:"index:idx_created_at"`
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
	var logger = l.NewLogger()
	conf := config.GetConf()
	db := GetHarvesterDB()
	// Create new or populate an existing one
	err := db.FirstOrCreate(&f,
		File{
			Path:    conf.Walker.Path() + string(os.PathSeparator) + path,
			Name:    info.Name(),
			ModTime: info.ModTime(),
			Status:  "new",
		}).Error
	if err != nil {
		logger.Error("File record CAN NOT BE created:", err)
		return err
	}
	logger.Info("File record has been processed for : ", f.GetPath())
	return err
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	f.CreatedAt = time.Now()
	return
}
