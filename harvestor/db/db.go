package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	dbHarvester, err := gorm.Open("sqlite3", conf.Database.DBFile())
	if err != nil {
		logger.Fatal("Can NOT open SQLite DB:", err)
	}
	// Controlling open connections
	dbHarvester.DB().SetMaxOpenConns(conf.Database.MaxOpenConnections())
	// Max Idle connections
	dbHarvester.DB().SetMaxIdleConns(conf.Database.MaxIdleConnections())
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dbHarvester.DB().SetConnMaxLifetime(time.Duration(conf.Database.MaxConnectionLifeTime()) * time.Minute)
	// Logging DB Stats
	logger.Info("Harvester Database connected ||| dbHarvester Stats ... ", dbHarvester.DB().Stats())
	// deferring db close
	defer dbHarvester.Close()
}

// Get our db for sql operations
func GetarvesterDB() *gorm.DB {
	return dbHarvester
}

func migrateHarvester() {
	// Create new logger
	var logger = l.NewLogger()
	// just a plcae holder
	logger.Info("Harvester Database checking for migration ... ")
	/*
		if !dbHarvester.HasTable(&WalkFiles{}) {
			err := dbHarvester.HasTable.CreateTable(&WalkFiles{})
			if err != nil {
				logger.Fatal("Can not migrate walk_files table :", err)

			}
		}
	*/
}
