package tokenutil

import (
	"context"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
)

// Option token option.
type Option func(*options)

// CacheKeyFunc 自定义缓存key
type CacheKeyFunc func(context.Context, *authutil.Claims) string

// SigningSecretFunc 自定义密码
type SigningSecretFunc func(ctx context.Context, authClaims *authutil.Claims, passwordHash string) string

// options Parser is a jwt parser
type options struct {
	authConfig        *confv1.App_Auth
	cacheKeyFunc      CacheKeyFunc
	signingSecretFunc SigningSecretFunc
}

// WithAuthConfig 签名密码
func WithAuthConfig(authConfig *confv1.App_Auth) Option {
	return func(o *options) {
		o.authConfig = authConfig
	}
}

// WithCacheKeyFunc 用途：自定义缓存key、...
func WithCacheKeyFunc(cacheKeyFunc CacheKeyFunc) Option {
	return func(o *options) {
		o.cacheKeyFunc = cacheKeyFunc
	}
}

// WithSecretFunc 用途：密码前缀；用于区分特定的环境；例如：admin/user
func WithSecretFunc(customSecretFunc SigningSecretFunc) Option {
	return func(o *options) {
		o.signingSecretFunc = customSecretFunc
	}
}
