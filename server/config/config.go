package config

import (
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

//Cfg config of logger
var Cfg = zap.Config{
  Encoding:         "json",
  Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
//  OutputPaths:      []string{"../logs/logs.json", "stdout"},
  OutputPaths:      []string{"stdout"},
  ErrorOutputPaths: []string{"stderr"},
  //ErrorOutputPaths: []string{"../logs/errors.json", "stderr"}, 
  EncoderConfig: zapcore.EncoderConfig{
    MessageKey: "message",
    LevelKey:   "level",
    TimeKey:    "time",
    EncodeTime: zapcore.ISO8601TimeEncoder,
  },
}

//ErrorCfg config of error logger
var ErrorCfg = zap.Config{
  Encoding:         "json",
  Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
  OutputPaths:      []string{"stderr"},
  //OutputPaths:      []string{"../logs/errors.json", "stderr"},
  EncoderConfig: zapcore.EncoderConfig{
    MessageKey: "message",
    LevelKey:   "level",
    TimeKey:    "time",
    EncodeTime: zapcore.ISO8601TimeEncoder,
  },
}