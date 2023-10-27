package config

import "go.uber.org/zap"

// Logger is a wrapper for an logger library
type Logger struct {
	logger *zap.Logger
}

// NewLogger creates a new logger instance
func NewLogger(isEnabled bool) (*Logger, error) {
	if !isEnabled {
		return &Logger{
			logger: zap.NewNop(),
		}, nil
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logger{
		logger: zapLogger,
	}, nil
}

// Info logs an informational message.
func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

// Error logs an error message.
func (l *Logger) Error(message string) {
	l.logger.Error(message)
}

// Sync calls the underlying Core's Sync method, flushing any buffered log entries. Applications should take care to call Sync before exiting.
func (l *Logger) Sync() {
	err := l.logger.Sync()
	if err != nil {
		panic(err)
	}
}
