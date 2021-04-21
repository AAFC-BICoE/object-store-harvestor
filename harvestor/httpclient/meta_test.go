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

func TestMeta(t *testing.T) {
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
	// getting meta struct to return from mock server
	m := db.Meta{
		ID:        5,
		MetaID:    "c7df3da1-8288-462f-a513-a7cda9254da3",
		FileID:    1,
		UploadID:  5,
		CreatedAt: now,
		UpdatedAt: now,
	}
	// struct to json
	jData, err := json.Marshal(m)
	// create mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
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
	// getting uplad
	u, err := db.GetUploadByFile(f)
	// checking errors
	assert.Nil(t, err)

	// init http client
	InitHttpClient()
	// testing upload against mock server
	_, err = postMeta(u)
	// asserting
	assert.Nil(t, err)
}
