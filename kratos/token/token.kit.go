package tokenutil

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
	contextutil "github.com/ikaiguang/go-srv-kit/kratos/context"
)

// TokenTypeMap 令牌类型映射
type TokenTypeMap map[string]authv1.TokenTypeEnum_TokenType

const (
	// KeyPrefixDefault 密码前缀、缓存key前缀
	KeyPrefixDefault = "default_"
	KeyPrefixService = "service_"
	KeyPrefixAdmin   = "admin_"
	KeyPrefixApi     = "api_"
	KeyPrefixWeb     = "web_"
	KeyPrefixApp     = "app_"
	KeyPrefixH5      = "h5_"
	KeyPrefixManager = "manager_"
)

var (
	// DefaultCachePrefix 默认key前缀；防止与其他缓存冲突；
	DefaultCachePrefix = "token:"

	// _tokenTypeMutex token类型
	_tokenTypeMap = TokenTypeMap{
		"":            authv1.TokenTypeEnum_DEFAULT,
		"/service/v1": authv1.TokenTypeEnum_SERVICE,
		"/admin/v1":   authv1.TokenTypeEnum_ADMIN,
		"/api/v1":     authv1.TokenTypeEnum_API,
		"/web/v1":     authv1.TokenTypeEnum_WEB,
		"/app/v1":     authv1.TokenTypeEnum_APP,
		"/h5/v1":      authv1.TokenTypeEnum_H5,
		"/manager/v1": authv1.TokenTypeEnum_MANAGER,
	}
)

// AuthTokenRepo 验证令牌
//
// =====
// 生产令牌步骤
// =====
// 1. 生产签名密码 SigningSecret
// 2. 确定签名方法 JWTSigningMethod
// 3. 签证令牌 SignedToken
// 4. 生产缓存key CacheKey
// 5. 存储令牌 SaveCacheData
//
// =====
// 验证令牌步骤
// =====
// 1. 设置令牌类型 SetTokenType
// 2. 获取令牌类型 GetTokenType
// 3. 获取解密密码 JWTKeyFunc
// 4. 额外验证方法 ValidateFunc
// =====
// 删除令牌：退出登录、修改密码
// =====
// 1. 删除令牌类型 DeleteCacheData
type AuthTokenRepo interface {
	// SigningSecret 签名密码
	SigningSecret(ctx context.Context, tokenType authv1.TokenTypeEnum_TokenType, passwordHash string) string
	// JWTSigningMethod jwt 签名方法
	JWTSigningMethod() *jwt.SigningMethodHMAC
	// SignedToken 签证Token
	SignedToken(authClaims *authutil.Claims, secret string) (string, error)
	// CacheKey 缓存key、...
	CacheKey(context.Context, *authutil.Claims) string
	// SaveCacheData 存储缓存
	SaveCacheData(ctx context.Context, authClaims *authutil.Claims, authInfo *authv1.Auth) error
	// DeleteCacheData 删除缓存
	DeleteCacheData(ctx context.Context, authClaims *authutil.Claims) error
	// SetTokenType 设置令牌类型
	SetTokenType(operation string, tokenType authv1.TokenTypeEnum_TokenType)
	// GetTokenType 获取令牌类型
	GetTokenType(operation string) authv1.TokenTypeEnum_TokenType
	// JWTKeyFunc 验证工具： authutil.KeyFunc，提供最终的 jwt.Keyfunc
	JWTKeyFunc() authutil.KeyFunc
	// ValidateFunc 自定义验证
	ValidateFunc() authutil.ValidateFunc
}

// NewCacheKey ...
func NewCacheKey(authPayload *authv1.Payload) string {
	var (
		prefix     = ""
		identifier = "null"
	)

	// identifier
	if authPayload.Uid != "" {
		identifier = authPayload.Uid
	} else if authPayload.Id > 0 {
		identifier = strconv.FormatUint(authPayload.Id, 10)
	}

	// prefix
	switch authPayload.Tt {
	case authv1.TokenTypeEnum_DEFAULT:
		prefix = KeyPrefixDefault
	case authv1.TokenTypeEnum_SERVICE:
		prefix = KeyPrefixService
	case authv1.TokenTypeEnum_ADMIN:
		prefix = KeyPrefixAdmin
	case authv1.TokenTypeEnum_API:
		prefix = KeyPrefixApi
	case authv1.TokenTypeEnum_WEB:
		prefix = KeyPrefixWeb
	case authv1.TokenTypeEnum_APP:
		prefix = KeyPrefixApp
	case authv1.TokenTypeEnum_H5:
		prefix = KeyPrefixH5
	case authv1.TokenTypeEnum_MANAGER:
		prefix = KeyPrefixManager
	default:
		prefix = KeyPrefixDefault
	}
	return DefaultCachePrefix + prefix + identifier
}

// NewSecret ...
func NewSecret(authConfig *confv1.App_Auth, tokenType authv1.TokenTypeEnum_TokenType, passwordHash string) string {
	var (
		prefix = ""
	)
	switch tokenType {
	case authv1.TokenTypeEnum_DEFAULT:
		prefix = authConfig.DefaultKey
	case authv1.TokenTypeEnum_SERVICE:
		prefix = authConfig.ServiceKey
	case authv1.TokenTypeEnum_ADMIN:
		prefix = authConfig.AdminKey
	case authv1.TokenTypeEnum_API:
		prefix = authConfig.ApiKey
	case authv1.TokenTypeEnum_WEB:
		prefix = authConfig.WebKey
	case authv1.TokenTypeEnum_APP:
		prefix = authConfig.AppKey
	case authv1.TokenTypeEnum_H5:
		prefix = authConfig.H5Key
	case authv1.TokenTypeEnum_MANAGER:
		prefix = authConfig.ManagerKey
	default:
		prefix = authConfig.DefaultKey
	}
	return prefix + passwordHash
}

// newTokenTypeMap 令牌类型映射
func newTokenTypeMap() TokenTypeMap {
	m := TokenTypeMap{}
	for key, value := range _tokenTypeMap {
		m[key] = value
	}
	return m
}

// GetRequestOperation 请求路径
func GetRequestOperation(ctx context.Context) (operation string) {
	kratosTr, ok := contextutil.FromServerContext(ctx)
	if ok {
		operation = kratosTr.Operation()
	}
	if httpTr, ok := contextutil.IsHTTPTransporter(kratosTr); ok {
		var (
			pathSeparator = "/"
			splitN        = 4
			urlPathSlice  = strings.SplitN(httpTr.Request().URL.Path, pathSeparator, splitN)
		)
		if len(urlPathSlice) >= splitN {
			operation = strings.Join(urlPathSlice[:splitN-1], "/")
		} else {
			operation = strings.Join(urlPathSlice, "/")
		}
	}
	return operation
}
