package tokenutil

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"strconv"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
)

// TokenRepo ...
type TokenRepo interface {
	// CacheKey 缓存key、...
	CacheKey(context.Context, *authutil.Claims) string
	// JWTKeyFunc 响应 jwt.Keyfunc
	JWTKeyFunc(context.Context) (context.Context, jwt.Keyfunc)
}

const (
	// KeyPrefixDefault 密码前缀、缓存key前缀
	KeyPrefixDefault = "default_"
	KeyPrefixApp     = "app_"
	KeyPrefixService = "service_"
	KeyPrefixAdmin   = "admin_"
	KeyPrefixApi     = "api_"
	KeyPrefixWeb     = "web_"
)

var (
	// DefaultCachePrefix 默认key前缀；防止与其他缓存冲突；
	DefaultCachePrefix = "token:"
)

// AuthKey ...
func AuthKey(prefix string, payload *authv1.Payload) string {
	if payload.Uid != "" {
		return CacheKey(prefix, payload.Uid)
	}
	if payload.Id > 0 {
		return CacheKey(prefix, strconv.FormatUint(payload.Id, 10))
	}
	return CacheKey(prefix, "null")
}

// CacheKey 例子：token:xxx_xxx
func CacheKey(prefix, identifier string) string {
	return DefaultCachePrefix + prefix + identifier
}

// Secret ...
func Secret(prefix, secret string) []byte {
	return []byte(prefix + secret)
}

// CacheKeyForDefault 例子：token:xxx_xxx
func CacheKeyForDefault(identifier string) string {
	return CacheKey(KeyPrefixDefault, identifier)
}

// CacheIDForDefault 例子：token:xxx_xxx
func CacheIDForDefault(id uint64) string {
	return CacheKey(KeyPrefixDefault, strconv.FormatUint(id, 10))
}

// SecretForDefault ...
func SecretForDefault(secret string) []byte {
	return Secret(KeyPrefixDefault, secret)
}

// AuthKeyForDefault ...
func AuthKeyForDefault(payload *authv1.Payload) string {
	return AuthKey(KeyPrefixDefault, payload)
}

// CacheKeyForApp 例子：token:xxx_xxx
func CacheKeyForApp(identifier string) string {
	return CacheKey(KeyPrefixApp, identifier)
}

// CacheIDForApp 例子：token:xxx_xxx
func CacheIDForApp(id uint64) string {
	return CacheKey(KeyPrefixApp, strconv.FormatUint(id, 10))
}

// SecretForApp ...
func SecretForApp(secret string) []byte {
	return Secret(KeyPrefixApp, secret)
}

// AuthKeyForApp ...
func AuthKeyForApp(payload *authv1.Payload) string {
	return AuthKey(KeyPrefixApp, payload)
}

// CacheKeyForService 例子：token:xxx_xxx
func CacheKeyForService(identifier string) string {
	return CacheKey(KeyPrefixService, identifier)
}

// CacheIDForService 例子：token:xxx_xxx
func CacheIDForService(id uint64) string {
	return CacheKey(KeyPrefixService, strconv.FormatUint(id, 10))
}

// SecretForService ...
func SecretForService(secret string) []byte {
	return Secret(KeyPrefixService, secret)
}

// AuthKeyForService ...
func AuthKeyForService(payload *authv1.Payload) string {
	return AuthKey(KeyPrefixService, payload)
}

// CacheKeyForAdmin 例子：token:xxx_xxx
func CacheKeyForAdmin(identifier string) string {
	return CacheKey(KeyPrefixAdmin, identifier)
}

// CacheIDForAdmin 例子：token:xxx_xxx
func CacheIDForAdmin(id uint64) string {
	return CacheKey(KeyPrefixAdmin, strconv.FormatUint(id, 10))
}

// SecretForAdmin ...
func SecretForAdmin(secret string) []byte {
	return Secret(KeyPrefixAdmin, secret)
}

// AuthKeyForAdmin ...
func AuthKeyForAdmin(payload *authv1.Payload) string {
	return AuthKey(KeyPrefixAdmin, payload)
}

// CacheKeyForApi 例子：token:xxx_xxx
func CacheKeyForApi(identifier string) string {
	return CacheKey(KeyPrefixApi, identifier)
}

// CacheIDForApi 例子：token:xxx_xxx
func CacheIDForApi(id uint64) string {
	return CacheKey(KeyPrefixApi, strconv.FormatUint(id, 10))
}

// SecretForApi ...
func SecretForApi(secret string) []byte {
	return Secret(KeyPrefixApi, secret)
}

// AuthKeyForApi ...
func AuthKeyForApi(payload *authv1.Payload) string {
	return AuthKey(KeyPrefixApi, payload)
}

// CacheKeyForWeb 例子：token:xxx_xxx
func CacheKeyForWeb(identifier string) string {
	return CacheKey(KeyPrefixWeb, identifier)
}

// CacheIDForWeb 例子：token:xxx_xxx
func CacheIDForWeb(id uint64) string {
	return CacheKey(KeyPrefixWeb, strconv.FormatUint(id, 10))
}

// SecretForWeb ...
func SecretForWeb(secret string) []byte {
	return Secret(KeyPrefixWeb, secret)
}

// AuthKeyForWeb ...
func AuthKeyForWeb(payload *authv1.Payload) string {
	return AuthKey(KeyPrefixWeb, payload)
}
