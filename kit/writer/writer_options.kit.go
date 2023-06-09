package writerpkg

// options 配置可选项
type options struct {
	filenameSuffix string
}

// Option is config option.
type Option func(*options)

// WithFilenameSuffix 文件名后缀
func WithFilenameSuffix(suffix string) Option {
	return func(o *options) {
		o.filenameSuffix = suffix
	}
}
