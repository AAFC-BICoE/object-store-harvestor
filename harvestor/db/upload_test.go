package db

import (
	"github.com/stretchr/testify/assert"
	"harvestor/config"
	"testing"
	"time"
)

func TestDbEmptyUpload(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	// init db
	Init()
	db := GetHarvesterDB()
	// cheking db pointer
	assert.NotNil(t, db)
	// file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.jpg"
	// init an empty slice of stucked files
	fr, err := GetFileByPath(filePath)
	// checking errors
	assert.Nil(t, err)
	// asserts
	u, err := GetUploadByFile(fr)
	// checking errors
	assert.Nil(t, err)
	// asserting
	assert.Equal(t, 0, u.GetID())
	assert.Equal(t, 0, u.GetUploadID())
	assert.Equal(t, 0, u.GetFileID())
	assert.Equal(t, "", u.GetFileIdentifier())
	assert.Equal(t, "", u.GetBucket())
}

func TestDbCreateUpload(t *testing.T) {
	// now
	now := time.Now()
	u := Upload{
		ID:                1,
		UploadID:          1,
		FileIdentifier:    "bd56bc1a-07d0-4e01-a95f-c96c94aa3659",
		Bucket:            "cnc",
		DateTimeDigitized: nil,
		FileID:            1,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
	err := CreateUpload(&u)
	// checking errors
	assert.Nil(t, err)
	assert.Equal(t, now, u.GetUpdatedAt())
	assert.Equal(t, now, u.GetCreatedAt())
}
