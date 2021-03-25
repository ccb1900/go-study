package logger

import (
	log "github.com/sirupsen/logrus"
)

type FileLogger struct {
}

func (f FileLogger) Error(msg string) {
	log.Error(msg)
}
