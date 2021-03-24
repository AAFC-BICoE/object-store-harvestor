package config

// place holder for now
type SideCardConfiguration struct {
	Enabled bool
	Debug   bool
}

// Define all interfaces for this struct
type ISideCardConfiguration interface {
	IsEnabled() bool
	IsDebug() bool
}

// Implementation
func (s SideCardConfiguration) IsEnabled() bool {
	return s.Enabled
}

func (s SideCardConfiguration) IsDebug() bool {
	return s.Debug
}
