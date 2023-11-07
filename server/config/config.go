package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Cfg = zap.Config{
	Encoding:         "json",
	Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
	OutputPaths:      []string{"../logs/logs.json", "stdout"},
	ErrorOutputPaths: []string{"../logs/errors.json", "stderr"}, 
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,
	},
}