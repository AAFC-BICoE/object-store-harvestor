package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPath(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Equal(t, true, len(conf.Walker.Path()) > 0)
}

func TestInterest(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Equal(t, "yml", conf.Walker.Interest())
}
