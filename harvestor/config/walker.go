package config

// place holder for now
type FileWalkerConfiguration struct {
	EntryPointPath string
}

// Define all interfaces for this struct
type IFileWalkerConfiguration interface {
	Path() string
}

// Implementation
func (w FileWalkerConfiguration) Path() string {
	return w.EntryPointPath
}
