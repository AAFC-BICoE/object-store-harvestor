package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
app:
  release: "0.01"
  name: "harvestor"
  env: "dev"
*/

func TestGetRelease(t *testing.T) {
	var file = "../default_harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "0.01"
	assert.Equal(t, want, conf.App.GetRelease())

}

func TestGetName(t *testing.T) {
	var file = "../default_harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "harvestor"
	assert.Equal(t, want, conf.App.GetName())
}

func TestGetEnvironment(t *testing.T) {
	var file = "../default_harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "dev"
	assert.Equal(t, want, conf.App.GetEnvironment())
}
