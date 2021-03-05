package db

import (
	"database/sql"
	"gorm.io/datatypes"
	"os"
	"time"
)

type Upload struct {
	ID                int            `gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	UploadID          int            `json:"id" gorm:"uniqueIndex"`
	FileIdentifier    string         `json:"fileIdentifier" gorm:"uniqueIndex"`
	Bucket            string         `json:"bucket" gorm:"type:varchar(64)"`
	Status            string         `json:"status" gorm:"type:varchar(64)"`
	Exif              datatypes.JSON `json:"exif"`
	DateTimeDigitized time.Time      `json:"dateTimeDigitized" gorm:"default:null"`
	FileID            int            `json:"file_id" sql:"not null" gorm:"unique_index:idx_upload_file"`
	CreatedAt         time.Time      `json:"created_at" gorm:"index:idx_upload_created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	File              File           `gorm:"foreignkey:FileID"`
}

// Define all interfaces for this struct
type IUpload interface {
	GetID() int
	GetFileIdentifier() string
	GetBucket() string
	GetDateTimeDigitized() sql.NullTime
	GetStatus() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Create(path string, info os.FileInfo) error
}

// Implementation
func (u Upload) GetID() int {
	return u.ID
}
func (u Upload) GetFileIdentifier() string {
	return u.FileIdentifier
}
func (u Upload) GetBucket() string {
	return u.Bucket
}
func (u Upload) GetDateTimeDigitized() time.Time {
	return u.DateTimeDigitized
}
func (u Upload) GetStatus() string {
	return u.Status
}
func (u Upload) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u Upload) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
func (u Upload) Create() error {
	db := GetHarvesterDB()
	err := db.Create(&u).Error
	return err
}
