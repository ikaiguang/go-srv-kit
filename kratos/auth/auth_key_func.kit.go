package authutil

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/proto"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// RedisKeyFuncRepo ...
type RedisKeyFuncRepo interface {
	KeyFunc(ctx context.Context) (context.Context, jwt.Keyfunc)
}

// NewRedisKeyFunc ...
func NewRedisKeyFunc(redisCC *redis.Client) RedisKeyFuncRepo {
	return &redisKeyFunc{
		redisCC: redisCC,
	}
}

// redisKeyFunc ...
type redisKeyFunc struct {
	redisCC *redis.Client
}

// KeyFunc 默认 KeyFunc == jwt.Keyfunc
func (s *redisKeyFunc) KeyFunc(ctx context.Context) (context.Context, jwt.Keyfunc) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		myClaims, ok := token.Claims.(*Claims)
		if !ok || myClaims.AuthPayload == nil || myClaims.AuthPayload.Key == "" {
			logutil.WarnwWithContext(ctx,
				"error", ErrInvalidRedisKey,
				"token.Claims.(*Claims):OK", ok,
				"token.Claims.(*Claims):Content", fmt.Sprintf("%#v", myClaims),
			)
			err := errorutil.WithStack(ErrInvalidKeyFunc)
			return []byte(""), err
		}

		// key
		authInfo, err := s.getCacheData(ctx, myClaims)
		if err != nil {
			return nil, err
		}
		if authInfo.Payload == nil {
			logutil.WarnWithContext(ctx, "authInfo.Payload == nil")
			err := errorutil.WithStack(ErrInvalidKeyFunc)
			return []byte(""), err
		}
		ctx = NewRedisContext(ctx, authInfo)
		return []byte(authInfo.Payload.Key), nil
	}
	return ctx, keyFunc
}

// getCacheData 获取缓存数据
func (s *redisKeyFunc) getCacheData(ctx context.Context, claims *Claims) (*authv1.Auth, error) {
	cacheBytes, cacheErr := s.redisCC.Get(ctx, claims.AuthPayload.Key).Bytes()
	if cacheErr != nil {
		err := ErrGetRedisData
		err.Metadata = map[string]string{"error": cacheErr.Error()}
		return nil, err
	}

	// cache
	cacheData := &authv1.Auth{}
	unmarshalErr := proto.Unmarshal(cacheBytes, cacheData)
	if unmarshalErr != nil {
		err := ErrUnmarshalRedisData
		err.Metadata = map[string]string{"error": unmarshalErr.Error()}
		return cacheData, errorutil.WithStack(err)
	}
	return cacheData, nil
}
