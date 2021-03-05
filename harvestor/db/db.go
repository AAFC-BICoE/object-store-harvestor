package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"harvestor/config"
	l "harvestor/logger"
	"time"
)

var dbHarvester *gorm.DB

func Init() {
	initHarvester()
	migrateHarvester()
}

func initHarvester() {
	// Create new logger
	var logger = l.NewLogger()
	// getting config for our db
	conf := config.GetConf()
	// trying to open sqlite DB
	d, err := gorm.Open(sqlite.Open(conf.Database.DBFile()), &gorm.Config{})
	//dbHarvester, err := gorm.Open("sqlite3", conf.Database.DBFile())
	if err != nil {
		logger.Fatal("Can NOT open SQLite DB:", err)
	}
	// assign to global
	dbHarvester = d
	// details
	db, err := d.DB()
	if err != nil {
		logger.Fatal("Can NOT get Stats for SQLite DB:", err)
	}
	// Controlling open connections
	db.SetMaxOpenConns(conf.Database.MaxOpenConnections())
	// Max Idle connections
	db.SetMaxIdleConns(conf.Database.MaxIdleConnections())
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Duration(conf.Database.MaxConnectionLifeTime()) * time.Minute)
	// Logging DB Stats
	logger.Info("Harvester Database connected ||| dbHarvester Stats ... ", db.Stats())

}

func migrateHarvester() {
	// Create new logger
	var logger = l.NewLogger()
	// just a plcae holder
	logger.Info("Harvester Database checking for migration ... ")
	// Migrate the schema
	db := GetHarvesterDB()
	err := db.AutoMigrate(&File{})
	if err != nil {
		logger.Fatal("Can NOT AutoMigrate for SQLite DB:", err)
	}
	logger.Info("Harvester Database AutoMigrate is Done !!!")
}

// Get our db for sql operations
func GetHarvesterDB() *gorm.DB {
	return dbHarvester
}
