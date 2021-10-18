package mysqlutil

import (
	"strings"

	pkgerrors "github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

// NewMysqlDB .
func NewMysqlDB(conf *confv1.Data_MySQL, opts ...Option) (db *gorm.DB, err error) {
	handler := &mySQL{}
	return handler.New(conf, opts...)
}

// mySQL
type mySQL struct{}

// init 初始化
func (s *mySQL) New(conf *confv1.Data_MySQL, opts ...Option) (db *gorm.DB, err error) {
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
	loggerHandler := s.newLogger(conf, &options)

	// 选项
	opt := &gorm.Config{
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   loggerHandler,
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
	//  连接的最大数量
	if conf.ConnMaxActive > 0 {
		connPool.SetMaxOpenConns(int(conf.ConnMaxActive))
	}
	// 连接可复用的最大时间
	if conf.ConnMaxLifetime.GetSeconds() > 0 {
		connPool.SetConnMaxLifetime(conf.ConnMaxLifetime.AsDuration())
	}
	// 连接池中空闲连接的最大数量
	if conf.ConnMaxIdle > 0 {
		connPool.SetMaxIdleConns(int(conf.ConnMaxIdle))
	}
	// 设置连接空闲的最长时间
	if conf.ConnMaxIdleTime.GetSeconds() > 0 {
		connPool.SetConnMaxIdleTime(conf.ConnMaxIdleTime.AsDuration())
	}
	return db, err
}

// mysqlLogger 日志
func (s *mySQL) newLogger(conf *confv1.Data_MySQL, opt *options) logger.Interface {
	loggerConfig := &logger.Config{
		LogLevel:                  s.parseLevel(conf.LoggerLevel),
		SlowThreshold:             conf.SlowThreshold.AsDuration(),
		Colorful:                  false,
		IgnoreRecordNotFoundError: false,
	}
	if !conf.LoggerEnable {
		return logger.New(NewDummyWriter(), *loggerConfig)
	}
	if len(opt.writers) == 0 {
		return logger.New(NewStdWriter(), *loggerConfig)
	}
	return NewLogger(loggerConfig, opt.writers...)
}

// mysqlLogger 日志
func (s *mySQL) parseLevel(lv string) logger.LogLevel {
	switch strings.ToUpper(lv) {
	case "DEBUG":
		return logger.Info
	case "INFO":
		return logger.Info
	case "WARN":
		return logger.Warn
	case "ERROR":
		return logger.Error
	case "FATAL":
		return logger.Info
	default:
		return logger.Info
	}
}
