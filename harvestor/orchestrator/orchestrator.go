package orchestrator

import (
	"harvestor/config"
	"harvestor/httpclient"
	l "harvestor/logger"
	"harvestor/walker"
	"time"
)

func Run() {
	// time the start
	start := time.Now()
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()

	// Bio cluster case only
	if conf.App.GetEnvironment() == "cluster" {
		logger.Info("Orchestrator is about to run on the Bio cluster ...")
		walker.ClusterRun()
		httpclient.ClusterRun()
		logger.Info("Orchestrator has finished the run on the Bio cluster !!!")
	}

	// TODO Development
	// Future development for for PC of scientists
	if conf.App.GetEnvironment() == "PC" {
		// TODO
		logger.Fatal("PC Run NOT SUPPORTED YET")
		// ...
		logger.Info("Orchestrator is about to run on the PC ...")
		walker.PcRun()
		httpclient.PcRun()
		logger.Info("Orchestrator has finished the run on the PC !!!")
	}

	// time the end
	t := time.Now()
	elapsed := t.Sub(start)
	logger.Debug("= = = = = = = = D O N E = = = = = = = =")
	logger.Debug("app harvestor is done and it took : ", elapsed)
}
