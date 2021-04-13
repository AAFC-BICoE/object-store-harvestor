package db

import (
	_ "github.com/stretchr/testify/assert"
	_ "harvestor/config"
	_ "os"
	"testing"
)

func TestDbFile(t *testing.T) {
	/*
		file := "../harvestor_config.yml"
		config.Load(file)
		Init()
		db := GetHarvesterDB()
		assert.NotNil(t, db)
		// testing file path
		filePath := "/tmp/data-test/images/2021/03/07/screen_files.png"
		// check if exist
		assert.Equal(t, false, DoesFileExist(filePath))
		// get stats
		fileStat, err := os.Stat(filePath)
		// checking errors
		assert.Nil(t, err)
		// create db record
		fr, err := CreateFile(filePath, fileStat)
		// checking errors
		assert.Nil(t, err)
		// empty db file record
		f := File{}
		assert.NotEqual(t, f.GetID(), fr.GetID())
		assert.NotEqual(t, f.GetPath(), fr.GetPath())
		assert.NotEqual(t, f.GetModTime(), fr.GetModTime())
		assert.NotEqual(t, f.GetStatus(), fr.GetStatus())
		assert.NotEqual(t, f.GetCreatedAt(), fr.GetCreatedAt())
		assert.NotEqual(t, f.GetUpdatedAt(), fr.GetUpdatedAt())
		assert.Equal(t, f.GetUploadType(), fr.GetUploadType())

		// init an empty slice of new files
		var frs []File
		// getting all new files
		GetNewFiles(&frs)
		// checking number of new files
		assert.Equal(t, 1, len(frs))

		// set status uploaded
		err = SetFileStatus(&fr, "uploaded")
		// checking errors
		assert.Nil(t, err)

		GetStuckedFiles(&frs)
		// checking number of upload files
		assert.Equal(t, 1, len(frs))
		// set upload type
		err = SetUploadType(&fr, "original")
		// checking errors
		assert.Nil(t, err)
		// set status completed
		err = SetFileStatus(&fr, "completed")
		// checking errors
		assert.Nil(t, err)

		frReload, err := GetFileByPath(fr.GetPath())
		// checking errors
		assert.Nil(t, err)
		// checking status
		assert.Equal(t, "completed", frReload.GetStatus())
		// checking status
		assert.Equal(t, "original", frReload.GetUploadType())
	*/
}

func TestDbUpload(t *testing.T) {
	// empty db file record
	/*
		u := Upload{
			1,
			111,
			"bd56bc1a-07d0-4e01-a95f-c96c94aa3659",
			"cnc",
			{
				"User Comment":      "Screenshot",
				"Exif Image Height": "1225 pixels",
				"Exif Image Width":  "1636 pixels",
			},
			"file_id": 1,
		}
		err := CreateUpload(u)
		// checking errors
		assert.Nil(t, err)
	*/
}
