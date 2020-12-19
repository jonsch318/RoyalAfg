package log

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func RegisterService() *zap.SugaredLogger {
	logger = NewLogger()
	logger.Warn("Application started")
	return logger
}

func CleanLogger() {
	logger.Warn("Logger is exiting")
	logger.Desugar().Sync()
}
