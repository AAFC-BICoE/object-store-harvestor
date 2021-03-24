package db

import (
	"os"
	"time"
)

type Sidecard struct {
	ID        int       `json:"id" gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	Path      string    `json:"path" gorm:"uniqueIndex"`
	FileID    int       `json:"file_id" sql:"not null" gorm:"unique_index:idx_sidecard_file"`
	CreatedAt time.Time `json:"created_at" gorm:"index:idx_sidecard_created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	File      File      `gorm:"foreignkey:FileID"`
}

// Define all interfaces for this struct
type ISidecard interface {
	GetID() int
	GetPath() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Create(path string, info os.FileInfo) error
}

// Implementation
func (s Sidecard) GetID() int {
	return s.ID
}

func (s Sidecard) GetPath() string {
	return s.Path
}

func (s Sidecard) GetFileID() int {
	return s.FileID
}

func (s Sidecard) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s Sidecard) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func CreateSidecard(s *Sidecard) error {
	db := GetHarvesterDB()
	err := db.Create(s).Error
	return err
}
