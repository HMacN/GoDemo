package plog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

const stdOutFilePath string = "stdout"

func New(filePath string, filename string, fileExtension string) LogWrapper {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{
		stdOutFilePath,
		filePath + "\\" + filename + "." + fileExtension,
	}

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Could not start zap logger: %v", err)
	}
	defer logger.Sync()

	return LogWrapper{
		logger: logger.WithOptions(),
	}
}

func (wrapper *LogWrapper) Debug(message string, a ...KV) {
	wrapper.logger.WithOptions(zap.AddCallerSkip(1)).Debug(message, unwind(a)...)
	defer wrapper.logger.Sync()
}

func (wrapper *LogWrapper) Info(message string, a ...KV) {
	wrapper.logger.WithOptions(zap.AddCallerSkip(1)).Info(message, unwind(a)...)
	defer wrapper.logger.Sync()
}

func (wrapper *LogWrapper) Warn(message string, a ...KV) {
	wrapper.logger.WithOptions(zap.AddCallerSkip(1)).Warn(message, unwind(a)...)
	defer wrapper.logger.Sync()
}

func (wrapper *LogWrapper) Error(message string, a ...KV) {
	wrapper.logger.WithOptions(zap.AddCallerSkip(1)).Error(message, unwind(a)...)
	defer wrapper.logger.Sync()
}

type LogWrapper struct {
	logger *zap.Logger
}

type KV struct {
	Key   string
	Value any
}

func unwind(keyVals []KV) []zap.Field {
	var output []zap.Field
	for i := range keyVals {
		output = append(output, zap.Any(keyVals[i].Key, keyVals[i].Value))
	}

	return output
}
