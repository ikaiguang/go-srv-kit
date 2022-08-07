package tokenutil

import (
	"context"

	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
)

// Option token option.
type Option func(*options)

// CacheKeyFunc 自定义缓存key
type CacheKeyFunc func(ctx context.Context, authClaims *authutil.Claims) string

// SecretFunc 自定义密码
type SecretFunc func(ctx context.Context, secret string) string

// options Parser is a jwt parser
type options struct {
	cacheKeyFunc CacheKeyFunc
	secretFunc   SecretFunc
}

// WithCacheKeyFunc 用途：自定义缓存key、...
func WithCacheKeyFunc(cacheKeyFunc CacheKeyFunc) Option {
	return func(o *options) {
		o.cacheKeyFunc = cacheKeyFunc
	}
}

// WithSecretFunc 用途：密码前缀；用于区分特定的环境；例如：admin/user
func WithSecretFunc(customSecretFunc SecretFunc) Option {
	return func(o *options) {
		o.secretFunc = customSecretFunc
	}
}
