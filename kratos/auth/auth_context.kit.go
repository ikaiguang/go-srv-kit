package authutil

import (
	"context"
	"github.com/golang-jwt/jwt/v4"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
)

// jwtAuthKey context.Context key
type jwtAuthKey struct{}

// redisAuthKey context.Context key
type redisAuthKey struct{}

// NewJWTContext put auth info into context
func NewJWTContext(ctx context.Context, info jwt.Claims) context.Context {
	return context.WithValue(ctx, jwtAuthKey{}, info)
}

// FromJWTContext extract auth info from context
func FromJWTContext(ctx context.Context) (token jwt.Claims, ok bool) {
	token, ok = ctx.Value(jwtAuthKey{}).(jwt.Claims)
	return
}

// NewRedisContext ...
func NewRedisContext(ctx context.Context, info *authv1.Auth) context.Context {
	return context.WithValue(ctx, redisAuthKey{}, info)
}

// FromRedisContext extract auth info from context
func FromRedisContext(ctx context.Context) (info *authv1.Auth, ok bool) {
	info, ok = ctx.Value(jwtAuthKey{}).(*authv1.Auth)
	return
}
