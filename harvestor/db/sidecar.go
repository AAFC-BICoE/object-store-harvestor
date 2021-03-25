package db

import (
	"os"
	"time"
)

type Sidecar struct {
	ID     int    `json:"id" gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	Path   string `json:"path" gorm:"uniqueIndex"`
	FileID int    `json:"file_id" sql:"not null" gorm:"unique_index:idx_sidecard_file"`
	Status string `json:"status" gorm:"type:varchar(64)"`
	//CreatedAt time.Time `json:"created_at" gorm:"index:idx_sidecard_created_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	File      File      `gorm:"foreignkey:FileID"`
}

// Define all interfaces for this struct
type ISidecar interface {
	GetID() int
	GetPath() string
	GetStatus() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Create(path string, info os.FileInfo) error
}

// Implementation
func (s Sidecar) GetID() int {
	return s.ID
}

func (s Sidecar) GetPath() string {
	return s.Path
}

func (s Sidecar) GetFileID() int {
	return s.FileID
}

func (s Sidecar) GetStatus() string {
	return s.Status
}

func (s Sidecar) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s Sidecar) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func CreateSidecar(s *Sidecar) error {
	db := GetHarvesterDB()
	err := db.Create(s).Error
	return err
}

func GetSideCarByFile(file *File) (*Sidecar, error) {
	db := GetHarvesterDB()
	var sidecar Sidecar
	if err := db.Where("file_id = ?", file.GetID()).Find(&sidecar).Error; err != nil {
		return &sidecar, err
	}
	return &sidecar, nil
}

/*
type SidecarFile struct {
	AcMetadataCreator string
	DcCreator         string
	AcDerivedFrom     string
	ManagedAttributes []Attribute
	Tags              []Tag
}
*/
