package walker

import (
	"github.com/stretchr/testify/assert"
	_ "harvestor/config"
	_ "harvestor/db"
	"testing"
)

func TestReadDirNames(t *testing.T) {
	path := "/tmp/data-test/images"
	_, err := readDirNames(path)
	assert.Nil(t, err)
}
