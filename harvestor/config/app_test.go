package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	want := "America/New_York"
	assert.Equal(t, want, conf.App.GetObjectTimezone())
}
