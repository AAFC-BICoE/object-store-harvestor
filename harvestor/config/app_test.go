package config

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGetRelease(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}$`), conf.App.GetRelease())

}

func TestGetName(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "harvestor"
	assert.Equal(t, want, conf.App.GetName())
}

func TestGetEnvironment(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "cluster"
	assert.Equal(t, want, conf.App.GetEnvironment())
}

func TestGetObjectTimezone(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "EST"
	assert.Equal(t, want, conf.App.GetObjectTimezone())
}
