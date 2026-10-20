package contextpkg

import (
	"sync/atomic"

	headerpkg "github.com/ikaiguang/go-srv-kit/kit/v3/header"
)

// TrustedPlatform 信任的平台
var defaultTrustedPlatform atomic.Value

func init() {
	defaultTrustedPlatform.Store(headerpkg.RemoteAddr)
}

// SetTrustedPlatform 设置信任的平台
func SetTrustedPlatform(platformHeader string) {
	defaultTrustedPlatform.Store(platformHeader)
}

func trustedPlatform() string {
	return defaultTrustedPlatform.Load().(string)
}
