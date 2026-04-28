package contextpkg

import (
	"github.com/ikaiguang/go-srv-kit/kit/header"
)

// TrustedPlatform 信任的平台
var (
	defaultTrustedPlatform = headerpkg.RemoteAddr
)

// SetTrustedPlatform 设置信任的平台
func SetTrustedPlatform(platformHeader string) {
	defaultTrustedPlatform = platformHeader
}
