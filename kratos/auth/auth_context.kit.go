package authutil

import (
	"context"
	"github.com/golang-jwt/jwt/v4"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
)

// jwtAuthKey context.Context key
type jwtAuthKey struct{}

// NewJWTContext put auth info into context
func NewJWTContext(ctx context.Context, info jwt.Claims) context.Context {
	return context.WithValue(ctx, jwtAuthKey{}, info)
}

// FromJWTContext extract auth info from context
func FromJWTContext(ctx context.Context) (token jwt.Claims, ok bool) {
	token, ok = ctx.Value(jwtAuthKey{}).(jwt.Claims)
	return
}

// FromAuthContext extract auth info from context
func FromAuthContext(ctx context.Context) (token *Claims, ok bool) {
	token, ok = ctx.Value(jwtAuthKey{}).(*Claims)
	return
}

// redisAuthKey context.Context key
type redisAuthKey struct{}

const (
	// _authInfoKey 存储验证信息到Token.Header中
	_authInfoKey = "kit:auth:v1:auth_info"
)

// SaveAuthInfo 存储 验证信息到Token.Header
func SaveAuthInfo(tokenHeader map[string]interface{}, data interface{}) {
	tokenHeader[_authInfoKey] = data
}

// GetAuthInfo 获取 Token.Header中的验证信息
func GetAuthInfo(tokenHeader map[string]interface{}) (info *authv1.Auth, ok bool) {
	i, ok := tokenHeader[_authInfoKey]
	if !ok {
		return info, ok
	}
	info, ok = i.(*authv1.Auth)
	return info, ok
}

// NewRedisContext ...
func NewRedisContext(ctx context.Context, info *authv1.Auth) context.Context {
	return context.WithValue(ctx, redisAuthKey{}, info)
}

// FromRedisContext extract auth info from context
func FromRedisContext(ctx context.Context) (info *authv1.Auth, ok bool) {
	info, ok = ctx.Value(redisAuthKey{}).(*authv1.Auth)
	return
}
