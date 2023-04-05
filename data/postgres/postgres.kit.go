package psqlutil

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
)

// NewPostgresDB .
func NewPostgresDB(conf *confv1.Data_PSQL, opts ...gormutil.Option) (db *gorm.DB, err error) {
	return NewDB(conf, opts...)
}

// NewDB 初始化
func NewDB(conf *confv1.Data_PSQL, opts ...gormutil.Option) (db *gorm.DB, err error) {
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
