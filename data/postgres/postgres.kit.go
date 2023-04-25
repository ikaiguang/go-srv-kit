package psqlutil

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
)

// Config postgres conig
type Config struct {
	Dsn string
	// SlowThreshold 慢查询
	SlowThreshold  *durationpb.Duration
	LoggerEnable   bool
	LoggerColorful bool
	// LoggerLevel 日志级别；值：DEBUG、INFO、WARN、ERROR、FATAL
	LoggerLevel string
	// ConnMaxActive 连接可复用的最大时间
	ConnMaxActive uint32
	// ConnMaxLifetime 可复用的最大时间
	ConnMaxLifetime *durationpb.Duration
	// ConnMaxIdle 连接池中空闲连接的最大数量
	ConnMaxIdle uint32
	// ConnMaxIdleTime 设置连接空闲的最长时间
	ConnMaxIdleTime *durationpb.Duration
}

// NewPostgresDB .
func NewPostgresDB(conf *Config, opts ...gormutil.Option) (db *gorm.DB, err error) {
	return NewDB(conf, opts...)
}

// NewDB 初始化
func NewDB(conf *Config, opts ...gormutil.Option) (db *gorm.DB, err error) {
	// 链接选项
	connOption := &gormutil.ConnOption{
		LoggerEnable:              conf.LoggerEnable,
		LoggerLevel:               gormutil.ParseLoggerLevel(conf.LoggerLevel),
		LoggerWriters:             nil,
		LoggerColorful:            conf.LoggerColorful,
		SlowThreshold:             conf.SlowThreshold.AsDuration(),
		IgnoreRecordNotFoundError: false,

		ConnMaxActive:   int(conf.ConnMaxActive),
		ConnMaxLifetime: conf.ConnMaxLifetime.AsDuration(),
		ConnMaxIdle:     int(conf.ConnMaxIdle),
		ConnMaxIdleTime: conf.ConnMaxIdleTime.AsDuration(),
	}
	for _, o := range opts {
		o(connOption)
	}

	// 拨号
	dialect := postgres.Open(conf.Dsn)

	return gormutil.NewDB(dialect, connOption)
}
