package config

// place holder for now
type KeycloakConfiguration struct {
	Host           string
	AdminClientID  string
	UserName       string
	UserPassword   string
	GrantType      string
	RealmName      string
	NewTokenBefore int
	Debug          bool
	Enabled        bool
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
