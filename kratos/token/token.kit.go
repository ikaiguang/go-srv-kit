package tokenutil

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"strconv"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
)

const (
	// KeyPrefixDefault 密码前缀、缓存key前缀
	KeyPrefixDefault = "default_"
	KeyPrefixService = "service_"
	KeyPrefixAdmin   = "admin_"
	KeyPrefixApi     = "api_"
	KeyPrefixWeb     = "web_"
	KeyPrefixApp     = "app_"
	KeyPrefixH5      = "h5_"
)

// AuthTokenRepo ...
type AuthTokenRepo interface {
	// SigningSecret 签名密码
	SigningSecret(ctx context.Context, authClaims *authutil.Claims, passwordHash string) string
	// JWTSigningMethod jwt 签名方法
	JWTSigningMethod() *jwt.SigningMethodHMAC
	// SignedToken 签证Token
	SignedToken(authClaims *authutil.Claims, secret string) (string, error)
	// CacheKey 缓存key、...
	CacheKey(context.Context, *authutil.Claims) string
	// SaveCacheData 存储缓存
	SaveCacheData(ctx context.Context, authClaims *authutil.Claims, authInfo *authv1.Auth) error
	// JWTKeyFunc 响应 jwt.Keyfunc
	JWTKeyFunc(context.Context) (context.Context, jwt.Keyfunc)
}

var (
	// DefaultCachePrefix 默认key前缀；防止与其他缓存冲突；
	DefaultCachePrefix = "token:"
)

// NewAuthKey ...
func NewAuthKey(prefix string, payload *authv1.Payload) string {
	if payload.Uid != "" {
		return NewCacheKey(prefix, payload.Uid)
	}
	if payload.Id > 0 {
		return NewCacheKey(prefix, strconv.FormatUint(payload.Id, 10))
	}
	return NewCacheKey(prefix, "null")
}

// NewCacheKey 例子：token:xxx_xxx
func NewCacheKey(prefix, identifier string) string {
	return DefaultCachePrefix + prefix + identifier
}

// NewSecret ...
func NewSecret(prefix, secret string) string {
	return prefix + secret
}

// CacheKeyForDefault 例子：token:xxx_xxx
func CacheKeyForDefault(identifier string) string {
	return NewCacheKey(KeyPrefixDefault, identifier)
}

// CacheIDForDefault 例子：token:xxx_xxx
func CacheIDForDefault(id uint64) string {
	return NewCacheKey(KeyPrefixDefault, strconv.FormatUint(id, 10))
}

// AuthKeyForDefault ...
func AuthKeyForDefault(payload *authv1.Payload) string {
	return NewAuthKey(KeyPrefixDefault, payload)
}

// CacheKeyForService 例子：token:xxx_xxx
func CacheKeyForService(identifier string) string {
	return NewCacheKey(KeyPrefixService, identifier)
}

// CacheIDForService 例子：token:xxx_xxx
func CacheIDForService(id uint64) string {
	return NewCacheKey(KeyPrefixService, strconv.FormatUint(id, 10))
}

// AuthKeyForService ...
func AuthKeyForService(payload *authv1.Payload) string {
	return NewAuthKey(KeyPrefixService, payload)
}

// CacheKeyForAdmin 例子：token:xxx_xxx
func CacheKeyForAdmin(identifier string) string {
	return NewCacheKey(KeyPrefixAdmin, identifier)
}

// CacheIDForAdmin 例子：token:xxx_xxx
func CacheIDForAdmin(id uint64) string {
	return NewCacheKey(KeyPrefixAdmin, strconv.FormatUint(id, 10))
}

// AuthKeyForAdmin ...
func AuthKeyForAdmin(payload *authv1.Payload) string {
	return NewAuthKey(KeyPrefixAdmin, payload)
}

// CacheKeyForApi 例子：token:xxx_xxx
func CacheKeyForApi(identifier string) string {
	return NewCacheKey(KeyPrefixApi, identifier)
}

// CacheIDForApi 例子：token:xxx_xxx
func CacheIDForApi(id uint64) string {
	return NewCacheKey(KeyPrefixApi, strconv.FormatUint(id, 10))
}

// AuthKeyForApi ...
func AuthKeyForApi(payload *authv1.Payload) string {
	return NewAuthKey(KeyPrefixApi, payload)
}

// CacheKeyForWeb 例子：token:xxx_xxx
func CacheKeyForWeb(identifier string) string {
	return NewCacheKey(KeyPrefixWeb, identifier)
}

// CacheIDForWeb 例子：token:xxx_xxx
func CacheIDForWeb(id uint64) string {
	return NewCacheKey(KeyPrefixWeb, strconv.FormatUint(id, 10))
}

// AuthKeyForWeb ...
func AuthKeyForWeb(payload *authv1.Payload) string {
	return NewAuthKey(KeyPrefixWeb, payload)
}

// CacheKeyForApp 例子：token:xxx_xxx
func CacheKeyForApp(identifier string) string {
	return NewCacheKey(KeyPrefixApp, identifier)
}

// CacheIDForApp 例子：token:xxx_xxx
func CacheIDForApp(id uint64) string {
	return NewCacheKey(KeyPrefixApp, strconv.FormatUint(id, 10))
}

// AuthKeyForApp ...
func AuthKeyForApp(payload *authv1.Payload) string {
	return NewAuthKey(KeyPrefixApp, payload)
}

// CacheKeyForH5 例子：token:xxx_xxx
func CacheKeyForH5(identifier string) string {
	return NewCacheKey(KeyPrefixH5, identifier)
}

// CacheIDForH5 例子：token:xxx_xxx
func CacheIDForH5(id uint64) string {
	return NewCacheKey(KeyPrefixH5, strconv.FormatUint(id, 10))
}

// AuthKeyForH5 ...
func AuthKeyForH5(payload *authv1.Payload) string {
	return NewAuthKey(KeyPrefixH5, payload)
}
