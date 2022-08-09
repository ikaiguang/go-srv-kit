package tokenutil

import (
	"context"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
)

// Option token option.
type Option func(*options)

// CacheKeyFunc 自定义缓存key
type CacheKeyFunc func(context.Context, *authutil.Claims) string

// SigningSecretFunc 自定义密码
type SigningSecretFunc func(ctx context.Context, tokenType authv1.TokenTypeEnum_TokenType, passwordHash string) string

// TokenTypeFunc 令牌类型
type TokenTypeFunc func(context.Context, TokenTypeMap) authv1.TokenTypeEnum_TokenType

// options Parser is a jwt parser
type options struct {
	authConfig        *confv1.App_Auth
	tokenTypeMap      TokenTypeMap
	cacheKeyFunc      CacheKeyFunc
	signingSecretFunc SigningSecretFunc
	tokenTypeFunc     TokenTypeFunc
}

// WithAuthConfig 签名密码
func WithAuthConfig(authConfig *confv1.App_Auth) Option {
	return func(o *options) {
		o.authConfig = authConfig
	}
}

// WithTokenTypeMap 令牌类型映射
func WithTokenTypeMap(tokenTypeMap TokenTypeMap) Option {
	return func(o *options) {
		o.tokenTypeMap = tokenTypeMap
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

// WithTokenTypeFunc 令牌类型
func WithTokenTypeFunc(tokenTypeFunc TokenTypeFunc) Option {
	return func(o *options) {
		o.tokenTypeFunc = tokenTypeFunc
	}
}
