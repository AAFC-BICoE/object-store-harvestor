package walker

import (
	"github.com/stretchr/testify/assert"
	_ "harvestor/config"
	_ "harvestor/db"
	"testing"
)

func TestGetFileName(t *testing.T) {
	assert.Equal(t, "harvestor_config", getFileName("harvestor_config.yml"))
}

func TestGetFileExtension(t *testing.T) {
	assert.Equal(t, "yml", getFileExtension("harvestor_config.yml"))
}
