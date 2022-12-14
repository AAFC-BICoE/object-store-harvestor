package config

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTimeOut(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[0-9]{1,3}$`), conf.HttpClient.GetTimeOut())
}

func TestGetRetryMax(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[1-9]{1,1}$`), conf.HttpClient.GetRetryMax())
}

func TestGetRetryWaitMin(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[0-9]{1,2}$`), conf.HttpClient.GetRetryWaitMin())
}

func TestGetMaxIdleConnections(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[0-9]{1,2}$`), conf.HttpClient.GetMaxIdleConnections())
}

func TestGetLocalCertFile(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	// this is optional
	assert.Equal(t, "", conf.HttpClient.GetLocalCertFile())
}

func TestGetBaseApiUrl(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`(http[s]?://.*):(\d*)\/?(.*)`), conf.HttpClient.GetBaseApiUrl())
}

func TestGetUploadUri(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "/api/v1/file"
	assert.Equal(t, want, conf.HttpClient.GetUploadUri())
}

func TestGetUploadGroup(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[a-z]{1,9}$`), conf.HttpClient.GetUploadGroup())
}

func TestGetMetaUri(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "/api/v1/metadata"
	assert.Equal(t, want, conf.HttpClient.GetMetaUri())
}

func TestGetDerivativeUri(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "/api/v1/derivative"
	assert.Equal(t, want, conf.HttpClient.GetDerivativeUri())
}
