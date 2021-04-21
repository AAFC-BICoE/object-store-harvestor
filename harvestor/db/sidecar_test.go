package db

import (
	"github.com/stretchr/testify/assert"
	"harvestor/config"
	"testing"
	"time"
)

func TestDbSideCarNotExist(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// testing file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.yml"
	// check if exist
	assert.Equal(t, true, doesSidecarNotExist(filePath))
}

func TestDbSideCarCreate(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// testing file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.yml"
	// now
	now := time.Now()
	// create db record
	sr := Sidecar{
		ID:               1,
		Path:             filePath,
		OriginalFileID:   1,
		DerivativeFileID: 2,
		SidecarStatus:    "new",
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	err := CreateSidecarRecord(&sr)
	// checking errors
	assert.Nil(t, err)
	assert.Equal(t, now, sr.GetUpdatedAt())
	assert.Equal(t, now, sr.GetCreatedAt())
	assert.Equal(t, 1, sr.GetID())
	assert.Equal(t, "new", sr.GetStatus())
}

func TestDbSideCarGet(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// testing file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.yml"
	// get db record
	sr, err := GetSidecarByPath(filePath)
	// checking errors
	assert.Nil(t, err)
	assert.Equal(t, 1, sr.GetID())
	assert.Equal(t, "new", sr.GetStatus())
}

func TestDbSideCarExist(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// testing file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.yml"
	// check if exist
	assert.Equal(t, false, doesSidecarNotExist(filePath))
}

func TestDbNewSideCars(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// init an empty slice of new sidecars
	var srs []Sidecar
	// getting all new files
	GetNewSidecars(&srs)
	// checking number of new sidecar
	assert.Equal(t, 1, len(srs))
	for _, sr := range srs {
		assert.Equal(t, 1, sr.GetOriginalFile().GetID())
		assert.Equal(t, 2, sr.GetDerivativeFile().GetID())
		// marking as new sidecar
		err := SetSidecarStatus(&sr, "new")
		// checking errors
		assert.Nil(t, err)
	}

}
