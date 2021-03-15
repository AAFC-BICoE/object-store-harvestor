package config

// place holder for now
type KeycloakConfiguration struct {
	Host           string
	AdminClientID  string
	GrantType      string
	Debug          bool
	NewTokenBefore int
}

// Define all interfaces for this struct
type IKeycloakConfiguration interface {
	GetHost() string
	GetAdminClientID() string
	GetGrantType() string
	IsDebug() bool
	GetNewTokenBefore() int
}

// Implementation
func (k KeycloakConfiguration) GetHost() string {
	return k.Host
}

func (k KeycloakConfiguration) GetAdminClientID() string {
	return k.AdminClientID
}

func (k KeycloakConfiguration) GetGrantType() string {
	return k.GrantType
}

func (k KeycloakConfiguration) IsDebug() bool {
	return k.Debug
}

func (k KeycloakConfiguration) GetNewTokenBefore() int {
	return k.NewTokenBefore
}
