package database

import (
	"fmt"
	"github/tronglv_authen_author/helper/cache"
	"github/tronglv_authen_author/helper/logify"

	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	gormtracing "gorm.io/plugin/opentelemetry/tracing"
)

func WithGormMigrate(items []any) Option {
	return func(m *DBOption) {
		m.Migrate = items
	}
}

func WithDefaultLog(logConfig LogConfig) Option {
	return func(m *DBOption) {
		m.Logger = gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			GetLogConfig(logConfig),
		)
	}
}

func WithTraceProvider(provider trace.TracerProvider, opts ...gormtracing.Option) Option {
	return func(m *DBOption) {
		m.TraceProvider = provider
		m.TraceGormOptions = opts
	}
}

type gormSvc struct {
	dbConfig Config
	dbConn   *gorm.DB
	*DBOption
}

func NewGormConnect(dbConfig Config, opts ...Option) (Database, error) {
	s := &gormSvc{
		dbConfig: dbConfig,
		DBOption: &DBOption{
			TraceProvider: otel.GetTracerProvider(),
			Logger: logger.New(
				logify.New(),
				GetLogConfig(dbConfig.GetLogConfig()),
			),
		},
	}
	for _, opt := range opts {
		opt(s.DBOption)
	}

	if err := s.getConn(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *gormSvc) GetGormClient() *gorm.DB {
	return s.dbConn
}

func (s *gormSvc) GetCacheClient() cache.Cache {
	return s.Cache
}

func (s *gormSvc) getConn() error {
	db, err := gorm.Open(s.getDialect(), s.options())
	if err != nil {
		return err
	}

	if len(s.Migrate) > 0 {
		if e := db.AutoMigrate(s.Migrate...); e != nil {
			return e
		}
	}

	s.dbConn = db
	sqlDB, e := s.dbConn.DB()
	if e != nil {
		return e
	}

	sqlDB.SetMaxIdleConns(s.dbConfig.GetMaxIdleConnection())
	sqlDB.SetMaxOpenConns(s.dbConfig.GetMaxOpenConnection())
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(s.dbConfig.GetConnectionMaxIdleTime()))
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(s.dbConfig.GetConnectionMaxIdleTime()))

	//apply tracing
	s.applyTracing()
	return nil
}

func (s *gormSvc) options() *gorm.Config {
	schemaName := ""
	if s.dbConfig.GetSchemaName() != "" {
		schemaName = s.dbConfig.GetSchemaName() + "."
	}
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   schemaName,
		},
		Logger: s.DBOption.Logger,
	}
}

func (s *gormSvc) getDialect() gorm.Dialector {
	var dialect gorm.Dialector
	switch s.dbConfig.GetDriver() {
	case PostgresDBDriver:
		dialect = s.postgresOpen()
	case MysqlDBDriver:
		dialect = s.mysqlOpen()
	case SqliteDBDriver:
		dialect = s.sqliteOpen()
	}
	return dialect
}

func (s *gormSvc) postgresOpen() gorm.Dialector {
	return postgres.Open(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s TimeZone=%s",
		s.dbConfig.GetHost(),
		s.dbConfig.GetPort(),
		s.dbConfig.GetUsername(),
		s.dbConfig.GetPassword(),
		s.dbConfig.GetDBName(),
		"disable",
		s.dbConfig.GetSchemaName(),
		s.dbConfig.GetTimeZone(),
	))
}

func (s *gormSvc) mysqlOpen() gorm.Dialector {
	return mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		s.dbConfig.GetUsername(),
		s.dbConfig.GetPassword(),
		s.dbConfig.GetHost(),
		s.dbConfig.GetPort(),
		s.dbConfig.GetDBName(),
	))
}

func (s *gormSvc) sqliteOpen() gorm.Dialector {
	return sqlite.Open(s.dbConfig.GetDBName())
}

func (s *gormSvc) applyTracing() {
	if s.DisableTracing {
		return
	}

	opts := append(
		[]gormtracing.Option{
			gormtracing.WithTracerProvider(s.TraceProvider),
			gormtracing.WithoutMetrics(),
		},
		s.TraceGormOptions...,
	)
	_ = s.dbConn.Use(gormtracing.NewPlugin(opts...))
}
