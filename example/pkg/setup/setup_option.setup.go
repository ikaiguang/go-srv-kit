package setuppkg

// options 配置可选项
type options struct {
	configPath string
}

// Option is config option.
type Option func(*options)

// WithConfigPath 配置路径
func WithConfigPath(configPath string) Option {
	return func(o *options) {
		o.configPath = configPath
	}
}
