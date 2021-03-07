// This is out of the box common logger from logrus
// All logs will be stored as JSON for better searching
// TODO Need to check later on if this app needs
// error handling common interfaces or it's an over kill
// For now just keeping it simple

package logger

import (
	"github.com/sirupsen/logrus"
	"harvestor/config"
	"os"
)

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// Map log level from config
var Levels = map[string]logrus.Level{
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"trace": logrus.TraceLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	// new base logrus
	var baseLogger = logrus.New()
	// introducing our standard logger from base logrus logger
	var standardLogger = &StandardLogger{baseLogger}
	// define log formatter as JSON
	standardLogger.Formatter = &logrus.JSONFormatter{}
	// getting log level from our config
	conf := config.GetConf()
	l := conf.Logger.GetLevel()
	// if level is set correctly then setting ...
	level, ok := Levels[l]
	if ok {
		standardLogger.SetLevel(level)
	} else {
		// fallback in case of an error to Info
		standardLogger.SetLevel(logrus.InfoLevel)
	}

	// Define Output for logs
	standardLogger.SetOutput(os.Stdout)
	// add a caller info to the logs
	if standardLogger.GetLevel() == logrus.DebugLevel {
		standardLogger.SetReportCaller(true)
	}

	// Define Output for logs as a file
	//standardLogger.SetOutputFile("/tmp/harverstor.log")

	return standardLogger
}

// Extra log output
func (l *StandardLogger) SetOutputFile(filename string) {
	// TODO
	// need to open file in orchestrator
	// with defer file.Close()
	// can't close the file here
	// the logger is global
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 640)
	if err != nil {
		l.Fatal(err)
	}
	l.SetOutput(file)
}
