package db

import (
	"time"
)

type Meta struct {
	ID        int       `gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	MetaID    string    `json:"id" gorm:"uniqueIndex"`
	UploadID  int       `json:"file_id" sql:"not null" gorm:"unique_index:idx_meta_upload"`
	CreatedAt time.Time `json:"created_at" gorm:"index:idx_meta_created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Upload    Upload    `gorm:"foreignkey:UploadID"`
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

func CreateMeta(m *Meta) error {
	db := GetHarvesterDB()
	err := db.Create(m).Error
	return err
}
