package httpclient

import (
	"github.com/liamylian/jsontime"
	"github.com/stretchr/testify/assert"
	"harvestor/config"
	"harvestor/db"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUpload(t *testing.T) {
	// custom json for all time formats
	var json = jsontime.ConfigWithCustomTimeFormat
	// config file
	file := "../harvestor_config.yml"
	// load config
	config.Load(file)
	// get conf
	conf := config.GetConf()
	// helper var
	now := time.Now()
	// getting upload struct to return from mock server
	dtd := "2019-11-06T10:21:31"
	u := db.Upload{
		ID:                5,
		UploadID:          55,
		FileIdentifier:    "d64d9cdc-644b-4946-9b86-91ef820810a8",
		Bucket:            "cnc",
		FileID:            1,
		DateTimeDigitized: &dtd,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
	// struct to json
	jData, err := json.Marshal(u)
	// create mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	}))
	// defer closing the server until the end of the test
	defer ts.Close()
	// define mock serve base pai url
	conf.HttpClient.BaseApiUrl = ts.URL
	// disable keycloak for now
	conf.Keycloak.Enabled = false
	// init db
	db.Init()
	// file path
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.jpg"
	// init an empty slice of stucked files
	f, err := db.GetFileByPath(filePath)
	// checking errors
	assert.Nil(t, err)
	// init http client
	InitHttpClient()
	// testing upload against mock server
	_, err = uplaodImage(f)
	// asserting
	assert.Nil(t, err)
}
