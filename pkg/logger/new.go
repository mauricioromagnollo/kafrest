package logger

import (
	"go.uber.org/zap"
)

type logger struct {
	Logger
	log *zap.Logger
}

// NewLogger creates and returns a new Logger instance configured with the given application name, environment, and log level.
func NewLogger(appName, appEnvironment, logLevel string) Logger {
	zapLogger := buildZapLogger(appName, appEnvironment, logLevel)

	return &logger{
		log: zapLogger,
	}
}
