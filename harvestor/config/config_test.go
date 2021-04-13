package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigYmlExtention(t *testing.T) {
	file := "../harvestor_config.yml"
	want := true

	assert.Equal(t, want, isValidConfigFile(file))
}

func TestConfigJsonExtention(t *testing.T) {
	file := "../harvestor_config.json"
	want := false

	assert.Equal(t, want, isValidConfigFile(file))
}
