package logger

import (
	"time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006.01.02 15:04:05.000"))
}

func CreateLogger(logName string) *zap.Logger {
	var infoLogSetting = lumberjack.Logger{
		Filename: "data/log/" + logName + "_info.log",
		MaxSize: 100,
		Compress: false,
		MaxAge: 365,
		MaxBackups: 2,
		LocalTime: true, 
	}

	var errorLogSetting = infoLogSetting
	errorLogSetting.Filename = "data/log/" + logName + "_error.log" 

	var encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder 

	var infoCore = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(&infoLogSetting),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)

	var errorCore = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(&errorLogSetting),
		zap.NewAtomicLevelAt(zapcore.ErrorLevel),
	)

	var l = zap.New(
		zapcore.NewTee([]zapcore.Core{infoCore, errorCore}...),
		zap.AddCaller(),
	)

	return l 
}
