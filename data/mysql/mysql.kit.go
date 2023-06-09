package mysqlpkg

import (
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config MySQL config
type Config struct {
	Dsn string
	// SlowThreshold 慢查询
	SlowThreshold  *durationpb.Duration
	LoggerEnable   bool
	LoggerColorful bool
	// LoggerLevel 日志级别；值：DEBUG、INFO、WARN、ERROR、FATAL
	LoggerLevel string
	// conn_max_active 连接可复用的最大时间
	ConnMaxActive uint32
	// conn_max_lifetime 可复用的最大时间
	ConnMaxLifetime *durationpb.Duration
	// conn_max_idle 连接池中空闲连接的最大数量
	ConnMaxIdle uint32
	// conn_max_idle_time 设置连接空闲的最长时间
	ConnMaxIdleTime *durationpb.Duration
}

// NewMysqlDB .
func NewMysqlDB(conf *Config, opts ...gormpkg.Option) (db *gorm.DB, err error) {
	return NewDB(conf, opts...)
}

// NewDB 初始化
func NewDB(conf *Config, opts ...gormpkg.Option) (db *gorm.DB, err error) {
	// 链接选项
	connOption := &gormpkg.ConnOption{
		LoggerEnable:              conf.LoggerEnable,
		LoggerLevel:               gormpkg.ParseLoggerLevel(conf.LoggerLevel),
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
	dialect := mysql.Open(conf.Dsn)

	return gormpkg.NewDB(dialect, connOption)
}
