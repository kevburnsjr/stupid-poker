package internal

import (
	"time"

	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

func newLogger(levelStr string) *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  time.RFC3339,
		DisableTimestamp: false,
	})

	logger.AddHook(filename.NewHook())

	return logger
}
