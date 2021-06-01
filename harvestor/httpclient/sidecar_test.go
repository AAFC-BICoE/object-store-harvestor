package httpclient

import (
	"github.com/liamylian/jsontime"
	"github.com/stretchr/testify/assert"
	"harvestor/config"
	"harvestor/db"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Fake struct {
	Data string `json:"data"`
}

func TestSidecarManagedMeta(t *testing.T) {
	// custom json for all time formats
	var json = jsontime.ConfigWithCustomTimeFormat
	// config file
	file := "../harvestor_config.yml"
	// load confif
	config.Load(file)
	// get conf
	conf := config.GetConf()
	// getting fake struct to return from mock server
	d := Fake{
		Data: "Ping-Pong",
	}
	// struct to json
	jData, err := json.Marshal(d)
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
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.yml"
	// init an empty slice of stucked files
	_, err = db.GetSidecarByPath(filePath)
	// checking errors
	assert.Nil(t, err)
}

func TestSidecarDerivative(t *testing.T) {
	// custom json for all time formats
	var json = jsontime.ConfigWithCustomTimeFormat
	// config file
	file := "../harvestor_config.yml"
	config.Load(file)
	// init conf
	conf := config.GetConf()
	// getting meta struct to return from mock server
	d := Fake{
		Data: "Ping-Pong",
	}

	// struct to json
	jData, err := json.Marshal(d)
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
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.yml"
	// init an empty slice of stucked files
	s, err := db.GetSidecarByPath(filePath)
	// checking errors
	assert.Nil(t, err)
	// init http client
	InitHttpClient()
	// testing upload against mock server
	err = postSideCarDerivative(s)
	// asserting
	assert.Nil(t, err)
}
