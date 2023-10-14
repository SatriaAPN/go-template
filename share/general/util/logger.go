package util

import (
	dtohttp "go-template/dto/http"

	log "github.com/sirupsen/logrus"
)

var logUnit Logger

func GetLogger() Logger {
	if logUnit == nil {
		logUnit = newLogger()
	}

	return logUnit
}

type Logger interface {
	Errorf(ld dtohttp.LoggingData)
	// Fatalf(format string, args ...interface{})
	// Fatal(args ...interface{})
	Infof(ld dtohttp.LoggingData)
	// Info( args ...interface{})
	// Warnf(format string, args ...interface{})
	// Debugf(format string, args ...interface{})
	// Debug(args ...interface{})
}

func newLogger() Logger {
	return &loggerWrapper{}
}

type loggerWrapper struct {
}

func (l *loggerWrapper) Errorf(ld dtohttp.LoggingData) {
	log.WithFields(ld.GetParam()).Error(ld.GetInfo())
}

func (l *loggerWrapper) Infof(ld dtohttp.LoggingData) {
	log.WithFields(ld.GetParam()).Info(ld.GetInfo())
}
