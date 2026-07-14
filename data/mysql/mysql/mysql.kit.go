package mysqlpkg

import (
	stderrors "errors"

	mysqldriver "github.com/go-sql-driver/mysql"
	gormpkg "github.com/ikaiguang/go-gorm-kit/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

// IsErrDuplicatedKey ...
func IsErrDuplicatedKey(err error) bool {
	if err == nil {
		return false
	}
	if stderrors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	if mysqlErr, ok := stderrors.AsType[*mysqldriver.MySQLError](err); ok {
		return mysqlErr.Number == 1062
	}
	return false
}
