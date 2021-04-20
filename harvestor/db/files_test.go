package db

import (
	"github.com/stretchr/testify/assert"
	"harvestor/config"
	"testing"
	"time"
)

func TestDbFileNotExist(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// testing file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.jpg"
	// check if exist
	assert.Equal(t, false, DoesFileExist(filePath))
}

func TestDbFileCreateOriginal(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// testing file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.jpg"
	// now
	now := time.Now()
	// create db record
	f := File{
		1,
		filePath,
		"screen_meta.jpg",
		now,
		"new",
		"original",
		now,
		now,
	}
	err := CreateFile(&f)
	// checking errors
	assert.Nil(t, err)
	assert.Equal(t, now, f.GetUpdatedAt())
	assert.Equal(t, now, f.GetCreatedAt())
	assert.Equal(t, now, f.GetModTime())
}

func TestDbFileCreateDerivative(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// testing file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.png"
	// now
	now := time.Now()
	// create db record
	f := File{
		2,
		filePath,
		"screen_meta.png",
		now,
		"new",
		"derivative",
		now,
		now,
	}
	err := CreateFile(&f)
	// checking errors
	assert.Nil(t, err)
	assert.Equal(t, now, f.GetUpdatedAt())
	assert.Equal(t, now, f.GetCreatedAt())
	assert.Equal(t, now, f.GetModTime())
}

func TestDbFileExist(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// testing file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.jpg"
	// check if exist
	assert.Equal(t, true, DoesFileExist(filePath))
}

func TestDbNewFiles(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// init an empty slice of new files
	var frs []File
	// getting all new files
	GetNewFiles(&frs)
	// checking number of new files
	assert.Equal(t, 2, len(frs))
	for _, fr := range frs {
		// marking as Stucked file
		err := SetFileStatus(&fr, "uploaded")
		// checking errors
		assert.Nil(t, err)
	}

}

func TestDbStuckedFiles(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// checking db pointer
	assert.NotNil(t, db)
	// init an empty slice of stucked files
	var frs []File
	// get stucked files
	GetStuckedFiles(&frs)
	// checking number of upload files
	assert.Equal(t, 2, len(frs))
	// set complite
	for _, fr := range frs {
		err := SetFileStatus(&fr, "completed")
		// checking errors
		assert.Nil(t, err)
		// set upload type
		if fr.GetUploadType() == "original" {
			err = SetUploadType(&fr, "original")
			// checking errors
			assert.Nil(t, err)
		}
	}
}

func TestDbGetFileByPath(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// checking db pointer
	assert.NotNil(t, db)
	// file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.jpg"
	// init an empty slice of stucked files
	fr, err := GetFileByPath(filePath)
	// checking errors
	assert.Nil(t, err)
	// asserts
	assert.Equal(t, 1, fr.GetID())
	assert.Equal(t, "screen_meta.jpg", fr.GetName())
	assert.Equal(t, "completed", fr.GetStatus())
	assert.Equal(t, "original", fr.GetUploadType())
}
