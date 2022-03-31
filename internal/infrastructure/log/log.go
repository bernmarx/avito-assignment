package log

import (
	"github.com/orandin/sentrus"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// Logger Provides access to a singleton logger
func Logger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
		logger.Hooks.Add(sentrus.NewHook(
			[]logrus.Level{logrus.WarnLevel, logrus.ErrorLevel},
		))
		logger.Info("logger initialized")
	}

	return logger
}
