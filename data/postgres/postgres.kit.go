package psqlpkg

import (
	stderrors "errors"
	"time"

	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm/v3"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDB .
func NewPostgresDB(conf *Config, opts ...gormpkg.Option) (db *gorm.DB, err error) {
	return NewDB(conf, opts...)
}

// NewDB 初始化
func NewDB(conf *Config, opts ...gormpkg.Option) (db *gorm.DB, err error) {
	connOption, err := buildConnOption(conf, opts...)
	if err != nil {
		return nil, err
	}

	return gormpkg.NewDB(postgres.Open(conf.Dsn), connOption)
}

func buildConnOption(conf *Config, opts ...gormpkg.Option) (*gormpkg.ConnOption, error) {
	if conf == nil {
		return nil, stderrors.New("postgres config is nil")
	}

	connOption := &gormpkg.ConnOption{
		LoggerEnable:              conf.LoggerEnable,
		LoggerLevel:               gormpkg.ParseLoggerLevel(conf.LoggerLevel),
		LoggerWriters:             nil,
		LoggerColorful:            conf.LoggerColorful,
		SlowThreshold:             durationOrZero(conf.SlowThreshold),
		IgnoreRecordNotFoundError: false,

		ConnMaxActive:   int(conf.ConnMaxActive),
		ConnMaxLifetime: durationOrZero(conf.ConnMaxLifetime),
		ConnMaxIdle:     int(conf.ConnMaxIdle),
		ConnMaxIdleTime: durationOrZero(conf.ConnMaxIdleTime),
	}
	for _, o := range opts {
		if o != nil {
			o(connOption)
		}
	}
	return connOption, nil
}

func durationOrZero(value *durationpb.Duration) time.Duration {
	if value == nil {
		return 0
	}
	return value.AsDuration()
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
