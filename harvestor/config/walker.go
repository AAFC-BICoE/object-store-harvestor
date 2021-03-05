package config

// place holder for now
type FileWalkerConfiguration struct {
	EntryPointPath  string
	FilesOfInterest string
}

// Define all interfaces for this struct
type IFileWalkerConfiguration interface {
	Path() string
	Interest() string
}

// Implementation
func (w FileWalkerConfiguration) Path() string {
	return w.EntryPointPath
}

func (w FileWalkerConfiguration) Interest() string {
	return w.FilesOfInterest
}
