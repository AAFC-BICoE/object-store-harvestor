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
	l := config.GetLoggerLevel()
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

	// TODO Logging to file
	//logFile := config.GetLoggerFile()
	//file, err := OpenFile(logFile, O_RDWR|O_CREATE|O_APPEND, 0666)
	//standardLogger.SetOutput(file)

	return standardLogger
}
