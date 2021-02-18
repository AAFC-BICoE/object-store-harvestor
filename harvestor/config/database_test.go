package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
database:
  # Number of max Open connections to SQLite DB
  maxOpenConns: 2
  # Number of max Idle connections to SQLite DB
  maxIdleConns: 2
  # Number of minutes connection will live
  connMaxLifetime: 30
  # Path to SQLite DB
  dbFile: "/tmp/data-test/harvestor.db"
*/

func TestMaxOpenConnections(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := 2
	assert.Equal(t, want, conf.Database.MaxOpenConnections())

}

func TestMaxIdleConnections(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := 2
	assert.Equal(t, want, conf.Database.MaxIdleConnections())
}

func TestMaxConnectionLifeTime(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := 30
	assert.Equal(t, want, conf.Database.MaxConnectionLifeTime())
}

func TestDBFile(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "/tmp/db-test/harvestor.db"
	assert.Equal(t, want, conf.Database.DBFile())
}
