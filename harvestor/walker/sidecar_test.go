package walker

import (
	"github.com/stretchr/testify/assert"
	_ "harvestor/config"
	_ "harvestor/db"
	"testing"
)

func TestGetSidecarFile(t *testing.T) {
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.yml"
	_, err := GetSidecarFile(filePath)
	assert.Nil(t, err)
}
