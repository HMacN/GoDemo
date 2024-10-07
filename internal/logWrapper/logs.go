package logWrapper

import (
	"fmt"
	"log/slog"
	"os"
)

func New() LogWrapper {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return LogWrapper{logger: logger}
}

func (wrapper *LogWrapper) Info(format string, a ...any) {
	wrapper.logger.Info(fmt.Sprintf(format, a...))
}

func (wrapper *LogWrapper) Error(format string, a ...any) {
	wrapper.logger.Error(fmt.Sprintf(format, a...))
}

type LogWrapper struct {
	logger *slog.Logger
}
