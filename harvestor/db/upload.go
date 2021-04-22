package db

import (
	"gorm.io/datatypes"
	"time"
)

type Upload struct {
	ID                int            `gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	UploadID          int            `json:"id" gorm:"uniqueIndex"`
	FileIdentifier    string         `json:"fileIdentifier" gorm:"uniqueIndex"`
	Bucket            string         `json:"bucket" gorm:"type:varchar(64)"`
	Exif              datatypes.JSON `json:"exif"`
	DateTimeDigitized *string        `json:"dateTimeDigitized"`
	FileID            int            `json:"file_id" sql:"not null" gorm:"unique_index:idx_upload_file"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	File              File           `gorm:"foreignkey:FileID"`
}

// Define all interfaces for this struct
type IUpload interface {
	GetID() int
	GetFileIdentifier() string
	GetBucket() string
	GetDateTimeDigitized() *string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Create() error
}

// Implementation
func (u Upload) GetID() int {
	return u.ID
}
func (u Upload) GetUploadID() int {
	return u.UploadID
}
func (u Upload) GetFileID() int {
	return u.FileID
}
func (u Upload) GetFileIdentifier() string {
	return u.FileIdentifier
}
func (u Upload) GetBucket() string {
	return u.Bucket
}
func (u Upload) GetDateTimeDigitized() *string {
	s := "2019-11-06T10:21:31"
	return &s
	//return u.DateTimeDigitized
}
func (u Upload) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u Upload) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

// Create uplod record in DB
func CreateUpload(u *Upload) error {
	db := GetHarvesterDB()
	return db.Create(u).Error
}

// Get upload record from DB by file record from DB
func GetUploadByFile(file *File) (*Upload, error) {
	db := GetHarvesterDB()
	var upload Upload
	if err := db.Where("file_id = ?", file.GetID()).Joins("File").Find(&upload).Error; err != nil {
		return &upload, err
	}
	return &upload, nil
}
