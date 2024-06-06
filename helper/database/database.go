package database

import (
	"errors"
	"github/tronglv_authen_author/helper/cache"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm/logger"
	gormtracing "gorm.io/plugin/opentelemetry/tracing"
)

const (
	MysqlDBDriver    string = "mysql"
	PostgresDBDriver string = "postgres"
	SqliteDBDriver   string = "sqlite"
	MongoDBDriver    string = "mongodb"
)

type (
	DBOption struct {
		Migrate          []any
		Cache            cache.Cache
		Logger           logger.Interface
		TraceProvider    trace.TracerProvider
		TraceGormOptions []gormtracing.Option
		DisableTracing   bool
	}

	Option func(s *DBOption)
)

func WithCache(cache cache.Cache) Option {
	return func(m *DBOption) {
		m.Cache = cache
	}
}

func WithDisableTracing() Option {
	return func(m *DBOption) {
		m.DisableTracing = true
	}
}

func New(c Config, opts ...Option) (Database, error) {
	switch c.GetDriver() {
	case PostgresDBDriver, SqliteDBDriver, MysqlDBDriver:
		return NewGormConnect(c, opts...)
	}
	return nil, errors.New("the database driver does not support")
}

func Must(c Config, opts ...Option) Database {
	conn, err := New(c, opts...)
	if err != nil {
		panic(err)
	}
	return conn
}
