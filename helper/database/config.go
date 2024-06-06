package database

type DBConfig struct {
	Driver                string      `json:"driver,default=postgres"`
	Host                  string      `json:"host"`
	Port                  int         `json:"port"`
	DBName                string      `json:"name"`
	Username              string      `json:"username"`
	Password              string      `json:"password"`
	SchemaName            string      `json:"schema-name,optional"`
	TimeZone              string      `json:"time-zone,default=Asia/Ho_Chi_Minh"`
	ConnectTimeout        int         `json:"connect-timeout,default=2"`
	MaxIdleConnection     int         `json:"max-idle-connections,default=20"`
	MaxOpenConnection     int         `json:"max-open-connections,default=100"`
	ConnectionMaxLifeTime int         `json:"connection-max-lifetime,default=1200"`
	ConnectionMaxIdleTime int         `json:"connection-max-idle-time,default=1"`
	LogConfig             DBLogConfig `json:"logger"`
}

// DBLogConfig represents log configuration
type DBLogConfig struct {
	Writer         string `json:"writer,default=pmc"`
	Level          string `json:"level,default=warn"`
	SlowThreshold  int    `json:"slow-threshold,default=200"`
	IgnoreNotFound bool   `json:"ignore-not-found,default=true"`
}

// GetDriver returns the database driver
func (c *DBConfig) GetDriver() string {
	return c.Driver
}

// GetHost returns the database host
func (c *DBConfig) GetHost() string {
	return c.Host
}

// GetPort returns the database port
func (c *DBConfig) GetPort() int {
	return c.Port
}

// GetDBName returns the database name
func (c *DBConfig) GetDBName() string {
	return c.DBName
}

// GetUsername returns the database username
func (c *DBConfig) GetUsername() string {
	return c.Username
}

// GetPassword returns the database password
func (c *DBConfig) GetPassword() string {
	return c.Password
}

// GetSchemaName returns the database schema name
func (c *DBConfig) GetSchemaName() string {
	return c.SchemaName
}

// GetTimeZone returns the database time zone
func (c *DBConfig) GetTimeZone() string {
	return c.TimeZone
}

// GetConnectTimeout returns the database connection timeout
func (c *DBConfig) GetConnectTimeout() int {
	return c.ConnectTimeout
}

// GetMaxIdleConnection returns the maximum number of idle connections
func (c *DBConfig) GetMaxIdleConnection() int {
	return c.MaxIdleConnection
}

// GetMaxOpenConnection returns the maximum number of open connections
func (c *DBConfig) GetMaxOpenConnection() int {
	return c.MaxOpenConnection
}

// GetConnectionMaxLifeTime returns the maximum connection lifetime
func (c *DBConfig) GetConnectionMaxLifeTime() int {
	return c.ConnectionMaxLifeTime
}

// GetConnectionMaxIdleTime returns the maximum connection idle time
func (c *DBConfig) GetConnectionMaxIdleTime() int {
	return c.ConnectionMaxIdleTime
}

func (c *DBConfig) GetLogConfig() LogConfig {
	return &c.LogConfig
}

// GetWriter returns the log writer
func (c *DBLogConfig) GetWriter() string {
	return c.Writer
}

// GetLevel returns the log level
func (c *DBLogConfig) GetLevel() string {
	return c.Level
}

// GetSlowThreshold returns the slow query threshold
func (c *DBLogConfig) GetSlowThreshold() int {
	return c.SlowThreshold
}

// GetIgnoreNotFound returns whether to ignore not found errors
func (c *DBLogConfig) GetIgnoreNotFound() bool {
	return c.IgnoreNotFound
}
