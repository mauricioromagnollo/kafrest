package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func buildZapLogger(appName, appEnvironment, logLevel string) *zap.Logger {
	var level zap.AtomicLevel

	level, err := zap.ParseAtomicLevel(logLevel)
	if err != nil {
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "severity",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "time",
		EncodeTime:  zapcore.ISO8601TimeEncoder,
	}

	hostname, _ := os.Hostname()

	config := zap.Config{
		Encoding:         "json",
		Level:            level,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"hostname": hostname,
			"language": "golang",
			"app_name": appName,
			"env":      appEnvironment,
		},
		EncoderConfig: encoderConfig,
	}

	zapLogger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return zapLogger
}
