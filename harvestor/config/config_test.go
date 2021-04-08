package config

import (
	"github.com/stretchr/testify/assert"
	_ "log"
	"path/filepath"
	"reflect"
	"testing"
)

func TestConfigYmlExtention(t *testing.T) {
	file := "../default_harvestor_config.yml"
	want := true

	assert.Equal(t, want, isValidConfigFile(file))
}

func TestConfigJsonExtention(t *testing.T) {
	file := "../harvestor_config.json"
	want := false

	assert.Equal(t, want, isValidConfigFile(file))
}

func TestDefaultYmlExistingConfiguration(t *testing.T) {
	file := "../default_harvestor_config.yml"
	e := readFromFile(file)

	assert.Nil(t, e)
}

func TestDefaultJsonExistingConfiguration(t *testing.T) {
	file := "../harvestor_config.json"
	e := readFromFile(file)

	assert.NotNil(t, e)
}

func TestDefaultNotExistingConfiguration(t *testing.T) {
	file := "../not_existing_default_harvestor_config.yml"
	e := readFromFile(file)

	assert.NotNil(t, e)
}

func TestDefaultConfiguration(t *testing.T) {
	file := "../default_harvestor_config.yml"
	Load(file)
	fd := GetConf()
	want := true
	fDatabase := &DatabaseConfiguration{2, 2, 30, "/tmp/db-test/harvestor.db"}
	fWalker := &FileWalkerConfiguration{"/tmp/data-test", "jpg, png"}
	fHttpClient := &HttpClientConfiguration{300, 2, 3, 10, "http://localhost:8081", "/api/v1/file", "dev-group", "/api/v1/metadata", "/api/v1/managed-attribute-map", "/api/v1/derivative"}
	fLogger := &LoggerConfiguration{"Info", "/var/logs/AAFC", "harvestor.log"}
	fAppConfiguration := &AppConfiguration{"0.01", "harvestor", "dev"}

	assert.Equal(t, want, reflect.DeepEqual(fd.Database, *fDatabase))
	assert.Equal(t, want, reflect.DeepEqual(fd.Walker, *fWalker))
	assert.Equal(t, want, reflect.DeepEqual(fd.HttpClient, *fHttpClient))
	assert.Equal(t, want, reflect.DeepEqual(fd.Logger, *fLogger))
	assert.Equal(t, want, reflect.DeepEqual(fd.App, *fAppConfiguration))
}

func TestDefaultLoggerLevel(t *testing.T) {
	file := "../default_harvestor_config.yml"
	Load(file)
	conf := GetConf()
	l := conf.Logger.GetLevel()
	want := "info"
	assert.Equal(t, want, l)
}

func TestDefaultFileExtension(t *testing.T) {
	filename := "../default_harvestor_config.yml"
	_, file := filepath.Split(filename)
	want := "yml"

	assert.Equal(t, want, getFileExtension(file))
}

func TestDefaultFileName(t *testing.T) {
	filename := "../default_harvestor_config.yml"
	_, file := filepath.Split(filename)
	want := "default_harvestor_config"

	assert.Equal(t, want, getFileName(file))
}
