package database

import (
	"github/tronglv_authen_author/helper/cache"

	"gorm.io/gorm"
)

type Database interface {
	GetGormClient() *gorm.DB
	GetCacheClient() cache.Cache
}

type Config interface {
	GetDriver() string
	GetHost() string
	GetPort() int
	GetDBName() string
	GetUsername() string
	GetPassword() string
	GetSchemaName() string
	GetTimeZone() string
	GetConnectTimeout() int
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
	GetConnectionMaxLifeTime() int
	GetConnectionMaxIdleTime() int
	GetLogConfig() LogConfig
}

type LogConfig interface {
	GetWriter() string
	GetLevel() string
	GetSlowThreshold() int
	GetIgnoreNotFound() bool
}
