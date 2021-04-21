package db

import (
	"github.com/stretchr/testify/assert"
	"harvestor/config"
	"testing"
	"time"
)

func TestDbEmptyMeta(t *testing.T) {
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
	m, err := GetMetaByFile(fr)
	// checking errors
	assert.Nil(t, err)
	// asserting
	assert.Equal(t, 0, m.GetID())
	assert.Equal(t, "", m.GetMetaID())
}

func TestDbCreateMeta(t *testing.T) {
	// now
	now := time.Now()
	m := Meta{
		ID:        1,
		MetaID:    "bd56bc1a-07d0-4e01-a95f-c96c94aa3659",
		UploadID:  1,
		FileID:    1,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := CreateMeta(&m)
	// checking errors
	assert.Nil(t, err)
	assert.Equal(t, now, m.GetUpdatedAt())
	assert.Equal(t, now, m.GetCreatedAt())
}
