package logutil

import (
	"io"
)

// options 配置可选项
type options struct {
	writer         io.Writer
	filenameSuffix string
}

// Option is config option.
type Option func(*options)

// WithWriter with config writer.
func WithWriter(writer io.Writer) Option {
	return func(o *options) {
		o.writer = writer
	}
}

// WithFilenameSuffix 文件名后缀
func WithFilenameSuffix(suffix string) Option {
	return func(o *options) {
		o.filenameSuffix = suffix
	}
}
