package redisutil

import (
	"sync"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

var (
	_keyPrefix     string
	_keyPrefixOnce sync.Once
)

// Key redis key
func Key(prefix, key string) string {
	return prefix + ":" + key
}

// KeyPrefix redis key prefix
func KeyPrefix(app *confv1.App) string {
	_keyPrefixOnce.Do(func() {
		_keyPrefix = NewKeyPrefix(app)
	})
	return _keyPrefix
}

// NewKeyPrefix redis key prefix
func NewKeyPrefix(app *confv1.App) string {
	return app.Name + ":" + app.Version + ":" + app.Env
}
