package contextpkg

import (
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// TrustedPlatform 信任的平台
var (
	defaultTrustedPlatform = headerpkg.RemoteAddr
)

// SetTrustedPlatform 设置信任的平台
func SetTrustedPlatform(platformHeader string) {
	defaultTrustedPlatform = platformHeader
}
