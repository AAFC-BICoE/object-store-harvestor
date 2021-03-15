package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
logger:
  level: "Debug"
  path: "/var/logs/AAFC"
  file: "harvestor.log"
*/

func TestGetLevel(t *testing.T) {
	var file = "../default_harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "info"
	assert.Equal(t, want, conf.Logger.GetLevel())

}

func TestGetPath(t *testing.T) {
	var file = "../default_harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "/var/logs/AAFC"
	assert.Equal(t, want, conf.Logger.GetPath())
}

func TestGetFile(t *testing.T) {
	var file = "../default_harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "harvestor.log"
	assert.Equal(t, want, conf.Logger.GetFile())
}
