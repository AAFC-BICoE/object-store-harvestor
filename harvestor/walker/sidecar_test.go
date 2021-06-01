package walker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSidecarFile(t *testing.T) {
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.yml"
	_, err := GetSidecarFile(filePath)
	// checking errors
	assert.Nil(t, err)
}

func TestGetUploadMediaOrientation(t *testing.T) {
	filePath := "/tmp/data-test/images/2021/03/07/screen_meta.yml"
	sc, err := GetSidecarFile(filePath)
	// checking errors
	assert.Nil(t, err)
	// checking orientation value
	assert.Equal(t, 8, sc.GetOrientation())
}
