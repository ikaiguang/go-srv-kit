package redisutil

import (
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

// KeyPrefix is the prefix of redis key
func KeyPrefix(app *confv1.App) string {
	return app.Name + ":" + app.Version + ":" + app.Env + ":"
}
