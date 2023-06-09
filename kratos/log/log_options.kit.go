package logpkg

import (
	"io"
)

// options 配置可选项
type options struct {
	writer         io.Writer
	filenameSuffix string
	loggerKeys     map[LoggerKey]string
	timeFormat     string
}

// Option is config option.
type Option func(*options)

// WithWriter with config writer.
func WithWriter(writer io.Writer) Option {
	return func(o *options) {
		o.writer = writer
	}
}

// WithLoggerKey with config writer.
func WithLoggerKey(keys map[LoggerKey]string) Option {
	return func(o *options) {
		if o.loggerKeys == nil {
			o.loggerKeys = DefaultLoggerKey()
		}
		for k := range keys {
			o.loggerKeys[k] = keys[k]
		}
	}
}

// WithFilenameSuffix 文件名后缀
func WithFilenameSuffix(suffix string) Option {
	return func(o *options) {
		o.filenameSuffix = suffix
	}
}

// WithTimeFormat 时间格式
// DefaultTimeFormat "2006-01-02T15:04:05.999"
func WithTimeFormat(timeFormat string) Option {
	return func(o *options) {
		o.timeFormat = timeFormat
	}
}
