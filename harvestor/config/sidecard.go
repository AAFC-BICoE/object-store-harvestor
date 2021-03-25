package config

// place holder for now
type SideCarConfiguration struct {
	Enabled bool
	Debug   bool
}

// Define all interfaces for this struct
type ISideCarConfiguration interface {
	IsEnabled() bool
	IsDebug() bool
}

// Implementation
func (s SideCarConfiguration) IsEnabled() bool {
	return s.Enabled
}

func (s SideCarConfiguration) IsDebug() bool {
	return s.Debug
}
