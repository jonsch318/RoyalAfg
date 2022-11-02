package logging

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger
var PerfLogger *zap.Logger

// RegisterService creates a new Logger and populates the global Logger instance
func RegisterService() *zap.SugaredLogger {
	Logger = NewLogger()
	PerfLogger = Logger.Desugar()
	Logger.Warn("Application started")
	return Logger
}

// CleanLogger flushes and sync the logger.
func CleanLogger() {
	Logger.Warn("Logger is exiting")
	_ = Logger.Desugar().Sync()
}
