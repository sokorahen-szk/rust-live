package logger

import (
	"github.com/sirupsen/logrus"
)

type logFormat struct {
	TimestampFormat string
}

func init() {
	formatter := logFormat{}
	formatter.TimestampFormat = "2006-01-02 15:04:05"

	logrus.SetFormatter(&logrus.JSONFormatter{})
}
