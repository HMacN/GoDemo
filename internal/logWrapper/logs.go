package logWrapper

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func New() LogWrapper {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Could not start zap logger: %v", err)
	}
	defer logger.Sync()

	return LogWrapper{
		logger: logger.WithOptions(),
	}
}

func (wrapper *LogWrapper) Debug(format string, a ...any) {
	wrapper.logger.WithOptions(zap.AddCallerSkip(1)).Debug(fmt.Sprintf(format, a...))
	defer wrapper.logger.Sync()
}

func (wrapper *LogWrapper) Info(format string, a ...any) {
	wrapper.logger.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, a...))
	defer wrapper.logger.Sync()
}

func (wrapper *LogWrapper) Warn(format string, a ...any) {
	wrapper.logger.WithOptions(zap.AddCallerSkip(1)).Warn(fmt.Sprintf(format, a...))
	defer wrapper.logger.Sync()
}

func (wrapper *LogWrapper) Error(format string, a ...any) {
	wrapper.logger.WithOptions(zap.AddCallerSkip(1)).Error(fmt.Sprintf(format, a...))
	defer wrapper.logger.Sync()
}

type LogWrapper struct {
	logger *zap.Logger
}
