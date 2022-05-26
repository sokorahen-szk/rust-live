package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

type logFormat struct {
	TimestampFormat string
}

func init() {
	log = NewLogger(os.Getenv("APP_LOG_LEVEL"))
}

func NewLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()

	formatter := logFormat{}
	formatter.TimestampFormat = "2006-01-02 15:04:05"

	logger.SetFormatter(&logrus.JSONFormatter{})

	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	logger.SetLevel(lvl)

	return logger
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Info(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
