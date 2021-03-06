package orchestrator

import (
	"harvestor/httpclient"
	l "harvestor/logger"
	"harvestor/walker"
)

func Run() {
	// Create new logger
	var logger = l.NewLogger()
	logger.Info("Orchestrator is about to run ...")
	walker.Run()
	httpclient.Run()
	logger.Info("Orchestrator has finished the run !!!")
}
