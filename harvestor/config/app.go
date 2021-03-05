package config

// place holder for now
type AppConfiguration struct {
	Release string
	Name    string
	Env     string
}

// Define all interfaces for this struct
type IAppConfiguration interface {
	GetRelease() string
	GetName() string
	GetEnvironment() string
}

// Implementation
func (a AppConfiguration) GetRelease() string {
	return a.Release
}

func (a AppConfiguration) GetName() string {
	return a.Name
}

func (a AppConfiguration) GetEnvironment() string {
	return a.Env
}
