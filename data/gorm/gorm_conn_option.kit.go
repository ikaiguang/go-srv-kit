package gormpkg

import (
	"time"

	"gorm.io/gorm/logger"
)

// ConnOption 连接配置
type ConnOption struct {
	LoggerEnable              bool
	LoggerLevel               logger.LogLevel
	LoggerWriters             []logger.Writer
	LoggerColorful            bool
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool

	ConnMaxActive   int
	ConnMaxLifetime time.Duration
	ConnMaxIdle     int
	ConnMaxIdleTime time.Duration
}

// Option is config option.
type Option func(*ConnOption)

// WithWriters with config writers.
func WithWriters(writers ...logger.Writer) Option {
	return func(o *ConnOption) {
		o.LoggerWriters = writers
	}
}
