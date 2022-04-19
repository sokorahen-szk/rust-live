package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

type logFormat struct {
	TimestampFormat string
}

func init() {
	log = NewLogger()
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	formatter := logFormat{}
	formatter.TimestampFormat = "2006-01-02 15:04:05"

	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Fatal(args ...interface{}) {
	log.Info(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
