package mysqlutil

import (
	"time"

	pkgerrors "github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

// NewDB .
func NewDB(conf *confv1.Data_MySQL, opts ...Option) (db *gorm.DB, err error) {
	// 可选项
	options := options{
		writers: nil,
	}
	for _, o := range opts {
		o(&options)
	}

	// 拨号
	dialect := mysql.Open(conf.Dsn)

	// 日志
	loggerConfig := &logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	}
	loggerHandler := NewLogger(loggerConfig, options.writers...)

	// 选项
	opt := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 loggerHandler,
	}

	// 数据库链接
	db, err = gorm.Open(dialect, opt)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return
	}

	// 连接池
	connPool, err := db.DB()
	if err != nil {
		err = pkgerrors.WithStack(err)
		return
	}
	// 连接池中空闲连接的最大数量
	if conf.ConnMaxIdle > 0 {
		connPool.SetMaxIdleConns(int(conf.ConnMaxIdle))
	}
	// 设置连接空闲的最长时间
	if conf.ConnMaxIdleTime.GetSeconds() > 0 {
		connPool.SetConnMaxIdleTime(conf.ConnMaxLifetime.AsDuration())
	}
	//  连接的最大数量
	if conf.ConnMaxActive > 0 {
		connPool.SetMaxOpenConns(int(conf.ConnMaxActive))
	}
	// 连接可复用的最大时间
	if conf.ConnMaxLifetime.GetSeconds() > 0 {
		connPool.SetConnMaxLifetime(conf.ConnMaxLifetime.AsDuration())
	}
	return
}
