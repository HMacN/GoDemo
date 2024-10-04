package logWrapper

import (
	"log/slog"
	"os"
)

var LogHandler LogWrapper

func New() LogWrapper {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return LogWrapper{logger: logger}
}

func (wrapper *LogWrapper) Info(format string, a ...any) {
	wrapper.logger.Info(format, a...)
}

func (wrapper *LogWrapper) Error(format string, a ...any) {
	LogHandler.logger.Error(format, a...)
}

type LogWrapper struct {
	logger *slog.Logger
}
