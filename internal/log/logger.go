package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger() *zap.SugaredLogger {
	encodingConfig := zap.NewProductionEncoderConfig()

	jsonEncoder := zapcore.NewJSONEncoder(encodingConfig)

	encodingConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodingConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encodingConfig)

	fileLevel := zap.DebugLevel
	consoleLevel := zap.DebugLevel

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), fileLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), consoleLevel),
	)

	return zap.New(core).Sugar()
}
