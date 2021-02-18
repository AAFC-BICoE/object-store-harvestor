package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
httpclient:
  # Number of max Open connections to SQLite DB
  timeOut: 300
  # Number of max Open connections to SQLite DB
  baseApiUrl: "http://localhost:8080"
  # Number of max Open connections to SQLite DB
  uri: "/api/v1/object"
*/

func TestGetTimeOut(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := 300
	assert.Equal(t, want, conf.HttpClient.GetTimeOut())

}

func TestGetBaseApiUrl(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "http://localhost:8080"
	assert.Equal(t, want, conf.HttpClient.GetBaseApiUrl())
}

func TestGetUri(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "/api/v1/object"
	assert.Equal(t, want, conf.HttpClient.GetUri())
}
