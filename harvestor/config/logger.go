package config

import (
	"strings"
)

type LoggerConfiguration struct {
	Level string // Log level. Example : Info, Debug, Error and etc
	Path  string // Path to the log file. Example : /var/log/aafc_bcoe/
	File  string // Log file name. Example : harvestor.log
}

// Define all interfaces for this struct
type ILoggerConfiguration interface {
	GetLevel() string
	GetPath() string
	GetFile() string
}

// Implementation
func (lc LoggerConfiguration) GetLevel() string {
	return strings.ToLower(lc.Level)
}

func (lc LoggerConfiguration) GetPath() string {
	return lc.Path
}

func (lc LoggerConfiguration) GetFile() string {
	return lc.File
}
