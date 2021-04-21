package db

import (
	"harvestor/config"
	l "harvestor/logger"
	"os"
	"time"
)

type Sidecar struct {
	ID               int       `json:"id" gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	Path             string    `json:"path" gorm:"uniqueIndex"`
	OriginalFileID   int       `json:"original_file_id" gorm:"unique_index:idx_sidecard_original_file"`
	DerivativeFileID int       `json:"derivative_file_id" gorm:"unique_index:idx_sidecard_derivative_file"`
	SidecarStatus    string    `json:"sidecar_status" gorm:"type:varchar(64)"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	OriginalFile     File      `gorm:"foreignkey:OriginalFileID"`
	DerivativeFile   File      `gorm:"foreignkey:DerivativeFileID"`
}

// Define all interfaces for this struct
type ISidecar interface {
	GetID() int
	GetPath() string
	GetStatus() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetOriginalFile() File
	GetDerivativeFile() File
	CreateSidecar(path string, info os.FileInfo) (*Sidecar, error)
	CreateSidecarRecord(s *Sidecar) error
}

// Implementation
func (s Sidecar) GetID() int {
	return s.ID
}

func (s Sidecar) GetPath() string {
	return s.Path
}

func (s Sidecar) GetOriginalFile() File {
	return s.OriginalFile
}

func (s Sidecar) GetDerivativeFile() File {
	return s.DerivativeFile
}

func (s Sidecar) GetStatus() string {
	return s.SidecarStatus
}

func (s Sidecar) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s Sidecar) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func CreateSidecar(path string, info os.FileInfo) (*Sidecar, error) {
	// get config
	conf := config.GetConf()
	// define absolute path
	absolutePath := conf.Walker.Path() + string(os.PathSeparator) + path
	// define basic sidecar
	s := &Sidecar{
		Path:          absolutePath,
		SidecarStatus: "new",
	}
	// create sidecar record in the db
	err := CreateSidecarRecord(s)

	return s, err
}

func CreateSidecarRecord(s *Sidecar) error {
	// double check if the record already there or not
	if len(s.GetPath()) != 0 && doesSidecarNotExist(s.GetPath()) {
		db := GetHarvesterDB()
		err := db.Create(s).Error
		return err
	}
	return nil
}

// get all sidecars with status "new"
func GetNewSidecars(sidecars *[]Sidecar) {
	var logger = l.NewLogger()
	db := GetHarvesterDB()
	db.Where("sidecar_status = ?", "new").Joins("OriginalFile").Joins("DerivativeFile").Find(sidecars)
	logger.Debug("Found for upload total sidecars : ", len(*sidecars))
}

// check by absolute path if the sidecar exist in DB already
func doesSidecarNotExist(absolutePath string) bool {
	var sidecars []Sidecar
	db := GetHarvesterDB()
	db.Where("path = ?", absolutePath).Find(&sidecars)
	return len(sidecars) == 0
}

// change status of sidecar
func SetSidecarStatus(s *Sidecar, status string) error {
	// get logger
	var logger = l.NewLogger()
	db := GetHarvesterDB()
	s.SidecarStatus = status
	err := db.Save(s).Error
	if err == nil {
		logger.Info("Sidecar record has been '"+status+"' for : ", s.GetPath())
	}
	return err
}

// look up Sidecar record in DB by path
func GetSidecarByPath(path string) (*Sidecar, error) {
	var s Sidecar
	db := GetHarvesterDB()
	err := db.Where("path = ?", path).First(&s).Error
	return &s, err
}
