package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
walker:
  # File walker root path for media files
  EntryPointPath: "/tmp/data-test"
*/

func TestPath(t *testing.T) {
	var file = "../default_harvestor_config.yml"
	Load(file)
	conf := GetConf()
	want := "/tmp/data-test"
	assert.Equal(t, want, conf.Walker.Path())

}
