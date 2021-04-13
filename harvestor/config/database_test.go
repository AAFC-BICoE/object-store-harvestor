package config

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestMaxOpenConnections(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[0-9]{1,2}$`), conf.Database.MaxOpenConnections())
}

func TestMaxIdleConnections(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[0-9]{1,2}$`), conf.Database.MaxIdleConnections())
}

func TestMaxConnectionLifeTime(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[0-9]{1,2}$`), conf.Database.MaxConnectionLifeTime())
}

func TestDBFile(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`((?:[^/]*/)*)(.db)`), conf.Database.DBFile())
}
