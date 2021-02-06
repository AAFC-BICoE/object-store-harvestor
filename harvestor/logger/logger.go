package logger

import (
	"github.com/sirupsen/logrus"
	"harvestor/config"
	"os"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

var loglevels = map[string]logrus.Level{
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
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	level := config.GetLoggerLevel()
	loglevel, ok := loglevels[level]
	if ok {
		standardLogger.SetLevel(loglevel)
	} else {
		standardLogger.SetLevel(logrus.InfoLevel)
	}

	standardLogger.SetOutput(os.Stdout)

	// TODO Logging to file
	//logFile := config.GetLoggerFile()
	//file, err := OpenFile(logFile, O_RDWR|O_CREATE|O_APPEND, 0666)
	//standardLogger.SetOutput(file)

	return standardLogger
}
