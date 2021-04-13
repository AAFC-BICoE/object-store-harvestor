// Package config provides a functionality to read from yml config file
// and provides values for each key in the file
// the package is splitted into multiple for simplicity and readability
// Walker config
package config

// Walker config struct
type FileWalkerConfiguration struct {
	EntryPointPath  string // File walker root path for files, walker traverse all folders under the root path
	FilesOfInterest string // slice of file extensions we are interested in
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
