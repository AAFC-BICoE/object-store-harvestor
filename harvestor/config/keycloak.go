// Package config provides a functionality to read from yml config file
// and provides values for each key in the file
// the package is splitted into multiple for simplicity and readability
// Keycloak config
package config

// Keycloak config struct
type KeycloakConfiguration struct {
	Host           string // Keycloak hostname
	AdminClientID  string // Keycloak admin client id is configured by users
	UserName       string // Keycloak User name which is used to connect on behalf of AdminClientID
	UserPassword   string // Keycloak User password which is used to connect on behalf of AdminClientID
	GrantType      string // Keycloak OAuth2-compliant GrantType
	RealmName      string // Keycloak Realm Name for Keycloak UserName
	NewTokenBefore int    // Number of seconds before Access token expires it will try to get new one
	Debug          bool   // Keycloak Logger mode
	Enabled        bool   // Keycloak ON/OFF
}

// Define all interfaces for this struct
type IKeycloakConfiguration interface {
	GetHost() string
	GetAdminClientID() string
	GetUserName() string
	GetUserPassword() string
	GetRealmName() string
	GetGrantType() string
	GetNewTokenBefore() int
	IsDebug() bool
	IsEnabled() bool
}

// Implementation
func (k KeycloakConfiguration) GetHost() string {
	return k.Host
}

func (k KeycloakConfiguration) GetAdminClientID() string {
	return k.AdminClientID
}

func (k KeycloakConfiguration) GetUserName() string {
	return k.UserName
}

func (k KeycloakConfiguration) GetUserPassword() string {
	return k.UserPassword
}

func (k KeycloakConfiguration) GetRealmName() string {
	return k.RealmName
}

func (k KeycloakConfiguration) GetGrantType() string {
	return k.GrantType
}

func (k KeycloakConfiguration) GetNewTokenBefore() int {
	return k.NewTokenBefore
}

func (k KeycloakConfiguration) IsDebug() bool {
	return k.Debug
}

func (k KeycloakConfiguration) IsEnabled() bool {
	return k.Enabled
}
