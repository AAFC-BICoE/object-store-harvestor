// Package config provides a functionality to read from yml config file
// and provides values for each key in the file
// the package is splitted into multiple for simplicity and readability
// SQLite config
package config

// SQLite config struct
type DatabaseConfiguration struct {
	MaxOpenConns    int    // Number of max Open connections to SQLite DB
	MaxIdleConns    int    // Number of max Idle connections with SQLite DB
	ConnMaxLifetime int    // Number of minutes connection will live
	DbFile          string // Path to SQLite DB
}

// Define all interfaces for this struct
type IDatabaseConfiguration interface {
	MaxOpenConnections() int
	MaxIdleConnections() int
	MaxConnectionLifeTime() int
	DBFile() string
}

// Implementation
func (db DatabaseConfiguration) MaxOpenConnections() int {
	return db.MaxOpenConns
}

func (db DatabaseConfiguration) MaxIdleConnections() int {
	return db.MaxIdleConns
}

func (db DatabaseConfiguration) MaxConnectionLifeTime() int {
	return db.ConnMaxLifetime
}

func (db DatabaseConfiguration) DBFile() string {
	return db.DbFile
}
