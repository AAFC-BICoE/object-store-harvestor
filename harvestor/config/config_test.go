package config

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"reflect"
	"testing"
)

func TestStructConfiguration(t *testing.T) {
	fd := GetConf()
	want := true
	assert.Equal(t, want, reflect.DeepEqual(fd.Database, DatabaseConfiguration{}))
	assert.Equal(t, want, reflect.DeepEqual(fd.Walker, FileWalkerConfiguration{}))
	assert.Equal(t, want, reflect.DeepEqual(fd.HttpClient, HttpClientConfiguration{}))
	assert.Equal(t, want, reflect.DeepEqual(fd.Logger, LoggerConfiguration{}))
	assert.Equal(t, want, reflect.DeepEqual(fd.App, AppConfiguration{}))
}

func TestConfigYmlExtention(t *testing.T) {
	file := "/go/src/harvestor/harvestor_config.yml"
	want := true
	assert.Equal(t, want, isValidConfigFile(file))
}

func TestConfigJsonExtention(t *testing.T) {
	file := "/go/src/harvestor/harvestor_config.json"
	want := false
	assert.Equal(t, want, isValidConfigFile(file))
}

func TestDefaultConfiguration(t *testing.T) {
	file := "/go/src/harvestor/harvestor_config.yml"
	Load(file)
	fd := GetConf()
	want := true
	fDatabase := DatabaseConfiguration{"harvestor.db"}
	fWalker := FileWalkerConfiguration{"/tmp/data-test"}
	fHttpClient := HttpClientConfiguration{300, "http://localhost:8080/api/v1/", "object"}
	fLogger := LoggerConfiguration{"Debug", "/var/logs/AAFC", "harvestor.log"}
	fAppConfiguration := AppConfiguration{"0.01", "harvestor", "dev"}
	assert.Equal(t, want, reflect.DeepEqual(fd.Database, fDatabase))
	assert.Equal(t, want, reflect.DeepEqual(fd.Walker, fWalker))
	assert.Equal(t, want, reflect.DeepEqual(fd.HttpClient, fHttpClient))
	assert.Equal(t, want, reflect.DeepEqual(fd.Logger, fLogger))
	assert.Equal(t, want, reflect.DeepEqual(fd.App, fAppConfiguration))
}

func TestDefaultLoggerLevel(t *testing.T) {
	file := "/go/src/harvestor/harvestor_config.yml"
	Load(file)
	want := "debug"
	assert.Equal(t, want, GetLoggerLevel())
}

func TestDefaultFileExtension(t *testing.T) {
	filename := "/go/src/harvestor/harvestor_config.yml"
	_, file := filepath.Split(filename)
	want := "yml"
	assert.Equal(t, want, getFileExtension(file))
}

func TestDefaultFileName(t *testing.T) {
	filename := "/go/src/harvestor/harvestor_config.yml"
	_, file := filepath.Split(filename)
	want := "harvestor_config"
	assert.Equal(t, want, getFileName(file))
}
