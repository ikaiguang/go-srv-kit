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

// NewRedisContext ...
func NewRedisContext(ctx context.Context, info *authv1.Auth) context.Context {
	return context.WithValue(ctx, redisAuthKey{}, info)
}

// FromRedisContext extract auth info from context
func FromRedisContext(ctx context.Context) (info *authv1.Auth, ok bool) {
	info, ok = ctx.Value(redisAuthKey{}).(*authv1.Auth)
	return
}
