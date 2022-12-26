package service

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Can't create logfile")
	}

	log := &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
		},
		ReportCaller: true,
	}

	log.Info("Logger has been started")

	return log
}
