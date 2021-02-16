package config

/*
Example :
  MaxOpenConns : 2 # Number of max Open connections to SQLite DB
  MaxIdleConns : 2 # Number of max Idle connections with SQLite DB
  ConnMaxLifetime : 2 # If there is no activity in 2 minutes Idle connections will be destroyed
  DbName: "/tmp/harvestor.db" # for now this is only for Debug
*/

// place holder for now
type DatabaseConfiguration struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	DbFile          string
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
