package dbmigrate

// options ...
type options struct {
	closeEngine bool
}

// Option ...
type Option func(*options)

// WithCloseEngineHandler 关闭engine
func WithCloseEngineHandler() Option {
	return func(o *options) {
		o.closeEngine = true
	}
}
