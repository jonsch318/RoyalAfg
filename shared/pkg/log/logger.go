package log

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() *zap.SugaredLogger {
	encodingConfig := zap.NewProductionEncoderConfig()

	jsonEncoder := zapcore.NewJSONEncoder(encodingConfig)

	encodingConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodingConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encodingConfig.EncodeCaller = zapcore.ShortCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encodingConfig)

	fileLevel := zap.DebugLevel
	consoleLevel := zap.DebugLevel

	homedir, err := homedir.Dir()

	if err != nil {
		return nil
	}

	logSync := &lumberjack.Logger{
		Filename:   homedir + "/logs/RoyalAfgInGo/log.log",
		MaxSize:    50,
		MaxBackups: 5,
		MaxAge:     14,
		Compress:   false,
	}

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(logSync), fileLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), consoleLevel),
	)

	return zap.New(core, zap.AddCaller()).Sugar()
}
