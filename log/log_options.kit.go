package logutil

import (
	"io"
)

// options 配置可选项
type options struct {
	writer io.Writer
}

// Option is config option.
type Option func(*options)

// WithWriter with config writer.
func WithWriter(writer io.Writer) Option {
	return func(o *options) {
		o.writer = writer
	}
}
