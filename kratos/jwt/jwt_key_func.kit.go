package jwtutil

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"

	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// DefaultKeyFunc 默认 KeyFunc == jwt.Keyfunc
func DefaultKeyFunc(secret string) KeyFunc {
	return func(ctx context.Context) func(token *jwt.Token) (interface{}, error) {
		return func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		}
	}
}

// RedisSecretFunc 缓存密码
type RedisSecretFunc func(*Claims) ([]byte, error)

// KeyFuncFromRedis redis缓存  KeyFunc == jwt.Keyfunc
func KeyFuncFromRedis(redisSecretFunc RedisSecretFunc) KeyFunc {
	return func(ctx context.Context) func(token *jwt.Token) (interface{}, error) {
		return func(token *jwt.Token) (interface{}, error) {
			myClaims, ok := token.Claims.(*Claims)
			if !ok {
				return []byte(""), nil
			}
			return redisSecretFunc(myClaims)
		}
	}
}

// DefaultRedisSecretFunc 默认 缓存密码
func DefaultRedisSecretFunc(ctx context.Context, redisCC *redis.Client) func(*Claims) ([]byte, error) {
	return func(claims *Claims) ([]byte, error) {
		if claims == nil || claims.AuthPayload == nil || claims.AuthPayload.Key == "" {
			logutil.WarnWithContext(ctx, "RedisSecretFunc : invalid redis key")
			return []byte(""), nil
		}
		getResp := redisCC.Get(ctx, claims.AuthPayload.Key)
		if getResp.Err() != nil {
			logutil.WarnWithContext(ctx, "RedisSecretFunc : get secret failed : "+getResp.Err().Error())
			return []byte(""), nil
		}
		return []byte(getResp.String()), nil
	}
}
