package config

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGetHost(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`(http[s]?://.*):(\d*)\/?(.*)`), conf.Keycloak.GetHost())
}

func TestGetAdminClientID(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[a-z]{1,15}$`), conf.Keycloak.GetAdminClientID())
}

func TestGetUserName(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Equal(t, true, len(conf.Keycloak.GetUserName()) > 0)
}

func TestGetUserPassword(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Equal(t, true, len(conf.Keycloak.GetUserPassword()) > 0)
}

func TestGetRealmName(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Equal(t, true, len(conf.Keycloak.GetRealmName()) > 0)
}

func TestGetGrantType(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Equal(t, "password", conf.Keycloak.GetGrantType())
}

func TestGetGetNewTokenBefore(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.Regexp(t, regexp.MustCompile(`^[0-9]{1,2}$`), conf.Keycloak.GetNewTokenBefore())
}

func TestIsDebug(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.IsType(t, true, conf.Keycloak.IsDebug())
}

func TestIsEnabled(t *testing.T) {
	var file = "../harvestor_config.yml"
	Load(file)
	conf := GetConf()
	assert.IsType(t, true, conf.Keycloak.IsEnabled())
}
