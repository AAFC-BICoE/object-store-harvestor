package walker

import (
	_ "github.com/stretchr/testify/assert"
	_ "harvestor/config"
	_ "harvestor/db"
	"testing"
)

func TestOriginalFile(t *testing.T) {
	/*
		file := "../harvestor_config.yml"
		config.Load(file)
		db.Init()
		sqlitedb := db.GetHarvesterDB()
		assert.NotNil(t, sqlitedb)
		// testing file path
		filePath := "/tmp/data-test/images/2021/03/07/screen_meta.jpg"
		// check if exist
		assert.Equal(t, false, db.DoesFileExist(filePath))
		// create db record
		fr, err := createFileRecordOriginal(filePath)
		// checking errors
		assert.Nil(t, err)
		assert.Equal(t, filePath, fr.GetPath())
		assert.Equal(t, "screen_meta", getFileName(fr.GetName()))
		assert.Equal(t, "jpg", getFileExtension(fr.GetName()))
		// trying again db record
		_, err = createFileRecordOriginal(filePath)
		// should do nothing
		assert.Nil(t, err)
		// clean up
		sqlitedb.Exec("DROP TABLE files;")
	*/
}

func TestDerivativeFile(t *testing.T) {
	/*
		file := "../harvestor_config.yml"
		config.Load(file)
		db.Init()
		sqlitedb := db.GetHarvesterDB()
		assert.NotNil(t, sqlitedb)
		// testing file path
		filePath := "/tmp/data-test/images/2021/03/07/screen_meta.png"
		// check if exist
		assert.Equal(t, false, db.DoesFileExist(filePath))
		// create db record
		fr, err := createFileRecordDerivative(filePath)
		// checking errors
		assert.Nil(t, err)
		assert.Equal(t, filePath, fr.GetPath())
		assert.Equal(t, "screen_meta", getFileName(fr.GetName()))
		assert.Equal(t, "png", getFileExtension(fr.GetName()))
		// clean up
		sqlitedb.Exec("DROP TABLE files;")
	*/
}
