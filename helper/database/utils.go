package database

import (
	"time"

	"gorm.io/gorm/logger"
)

func GetLogConfig(logConfig LogConfig) logger.Config {
	return logger.Config{
		SlowThreshold:             time.Duration(logConfig.GetSlowThreshold()) * time.Millisecond,
		LogLevel:                  GetLogLevel(logConfig.GetLevel()),
		IgnoreRecordNotFoundError: logConfig.GetIgnoreNotFound(),
	}
}

func GetLogLevel(level string) logger.LogLevel {
	var logLevel logger.LogLevel
	switch level {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Warn
	}
	return logLevel
}
