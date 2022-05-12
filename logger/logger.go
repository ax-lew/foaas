package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Sugar logger to use on the project
var Logger *zap.SugaredLogger

func Initialize() error {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()
	if err != nil {
		return err
	}

	Logger = logger.Sugar()
	return nil
}
