package postgres

import (
	stderrors "errors"

	gormpkg "github.com/ikaiguang/go-gorm-kit/gorm"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDB .
func NewPostgresDB(conf *Config, opts ...gormpkg.Option) (db *gorm.DB, err error) {
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
	dialect := postgres.Open(conf.Dsn)

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
	var pgErr *pgconn.PgError
	if stderrors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}
