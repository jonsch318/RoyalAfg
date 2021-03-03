package log

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func RegisterService() *zap.SugaredLogger {
	Logger = NewLogger()
	Logger.Warn("Application started")
	return Logger
}

func CleanLogger() {
	Logger.Warn("Logger is exiting")
	Logger.Desugar().Sync()
}
