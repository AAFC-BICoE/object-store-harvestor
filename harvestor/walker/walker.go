package walker

import (
	"harvestor/config"
	l "harvestor/logger"
)

func Run() {
	// Create new logger
	var logger = l.NewLogger()
	conf := config.GetConf()
	logger.Info("Harvester Walker is about to scan : ", conf.Walker.Path())
}
