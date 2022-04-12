package mysqlutil

import (
	"gorm.io/gorm/logger"
)

// options 配置可选项
type options struct {
	writers []logger.Writer
}

// Option is config option.
type Option func(*options)

// WithWriters with config writers.
func WithWriters(writers ...logger.Writer) Option {
	return func(o *options) {
		o.writers = writers
	}
}
