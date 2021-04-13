package db

import (
	"time"
)

type Meta struct {
	ID        int       `gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	MetaID    string    `json:"id" gorm:"uniqueIndex"`
	FileID    int       `json:"file_id" sql:"not null" gorm:"unique_index:idx_meta_file"`
	UploadID  int       `json:"upload_id" sql:"not null" gorm:"unique_index:idx_meta_upload"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	File      File      `gorm:"foreignkey:FileID"`
}

// Define all interfaces for this struct
type IMeta interface {
	GetID() int
	GetMetaID() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Create() error
}

// Implementation
func (m Meta) GetID() int {
	return m.ID
}

func (m Meta) GetMetaID() string {
	return m.MetaID
}

func (m Meta) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m Meta) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// Create meta record in DB
func CreateMeta(m *Meta) error {
	db := GetHarvesterDB()
	return db.Create(m).Error
}

// Get meta record from DB by file record from DB
func GetMetaByFile(file *File) (*Meta, error) {
	db := GetHarvesterDB()
	var meta Meta
	if err := db.Where("file_id = ?", file.GetID()).Joins("File").Find(&meta).Error; err != nil {
		return &meta, err
	}
	return &meta, nil
}
