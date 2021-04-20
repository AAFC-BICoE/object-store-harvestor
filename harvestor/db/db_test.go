package db

import (
	"github.com/stretchr/testify/assert"
	"harvestor/config"
	"testing"
)

func TestDbInit(t *testing.T) {
	file := "../harvestor_config.yml"
	config.Load(file)
	Init()
	assert.NotNil(t, GetHarvesterDB())
}
