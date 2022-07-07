package contextutil

import (
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// TrustedPlatform 信任的平台
var (
	defaultTrustedPlatform = headerutil.RemoteAddr
)

// SetTrustedPlatform 设置信任的平台
func SetTrustedPlatform(platformHeader string) {
	defaultTrustedPlatform = platformHeader
}
