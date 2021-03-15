package orchestrator

import (
	"harvestor/httpclient"
	l "harvestor/logger"
	"harvestor/walker"
	"time"
)

func Run() {
	// time the start
	start := time.Now()
	// Create new logger
	var logger = l.NewLogger()
	logger.Info("Orchestrator is about to run ...")
	walker.Run()
	httpclient.Run()
	logger.Info("Orchestrator has finished the run !!!")
	t := time.Now()
	elapsed := t.Sub(start)
	logger.Debug("harvestor took : ", elapsed)
}
