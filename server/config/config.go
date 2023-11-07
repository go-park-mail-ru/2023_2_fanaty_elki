package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Cfg = zap.Config{
	Encoding:         "json",
	Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
	OutputPaths:      []string{"../logs/logs.txt", "stdout"},
	ErrorOutputPaths: []string{"../logs/errors.txt", "stderr"}, 
	// EncoderConfig: zapcore.EncoderConfig{
	// 	MessageKey: "message",
	// 	LevelKey:   "level",
	// 	TimeKey:    "time",
	// },
}