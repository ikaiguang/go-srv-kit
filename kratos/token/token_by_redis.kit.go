package tokenutil

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/proto"
	"time"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
)

// redisToken ...
type redisToken struct {
	redisCC *redis.Client
	opt     options
}

// NewRedisTokenRepo ...
func NewRedisTokenRepo(redisCC *redis.Client, opts ...Option) TokenRepo {
	o := &options{}
	for i := range opts {
		opts[i](o)
	}
	return &redisToken{
		redisCC: redisCC,
		opt:     *o,
	}
}

// CacheKey 缓存key、...
func (s *redisToken) CacheKey(ctx context.Context, authClaims *authutil.Claims) string {
	if s.opt.cacheKeyFunc != nil {
		return s.opt.cacheKeyFunc(ctx, authClaims)
	}
	return AuthKeyForDefault(authClaims.Payload)
}

// JWTKeyFunc KeyFunc == jwt.Keyfunc
func (s *redisToken) JWTKeyFunc(ctx context.Context) (context.Context, jwt.Keyfunc) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		myClaims, ok := token.Claims.(*authutil.Claims)
		if !ok || myClaims.Payload == nil {
			err := errorutil.WithStack(authutil.ErrInvalidAuthInfo)
			return []byte(""), err
		}

		// key
		key := s.CacheKey(ctx, myClaims)
		authInfo, err := s.getCacheData(ctx, key)
		if err != nil {
			return nil, err
		}

		// 密码
		secret := authInfo.Secret
		if s.opt.secretFunc != nil {
			secret = s.opt.secretFunc(ctx, secret)
		}

		// 存储
		ctx = authutil.NewRedisContext(ctx, authInfo)

		// 响应
		return []byte(secret), nil
	}
	return ctx, keyFunc
}

// ValidateAuthInfo 校验：验证信息
func (s *redisToken) ValidateAuthInfo(authClaims *authutil.Claims, authInfo *authv1.Auth) (err error) {
	// 无限制
	switch authInfo.Payload.L {
	case authv1.LimitTypeEnum_UNKNOWN, authv1.LimitTypeEnum_UNLIMITED:
		return nil
	case authv1.LimitTypeEnum_ONLY_ONE:
		// 仅一个
		if authClaims.Payload.T.AsTime().UnixNano() != authInfo.Payload.T.AsTime().UnixNano() {
			err = errorutil.WithStack(authutil.ErrLoginLimit)
			return err
		}
	case authv1.LimitTypeEnum_SAME_PLATFORM:
		// 平台限制一个
		if authClaims.Payload.P == authInfo.Payload.P &&
			authClaims.Payload.T.AsTime().UnixNano() != authInfo.Payload.T.AsTime().UnixNano() {
			err = errorutil.WithStack(authutil.ErrLoginLimit)
			return err
		}
	}
	return err
}

// getCacheData 获取 缓存数据
func (s *redisToken) getCacheData(ctx context.Context, cacheKey string) (*authv1.Auth, error) {
	cacheBytes, cacheErr := s.redisCC.Get(ctx, cacheKey).Bytes()
	if cacheErr != nil {
		err := authutil.ErrGetRedisData
		err.Metadata = map[string]string{"error": cacheErr.Error()}
		return nil, errorutil.WithStack(err)
	}

	// cache
	cacheData := &authv1.Auth{}
	unmarshalErr := proto.Unmarshal(cacheBytes, cacheData)
	if unmarshalErr != nil {
		err := authutil.ErrUnmarshalRedisData
		err.Metadata = map[string]string{"error": unmarshalErr.Error()}
		return cacheData, errorutil.WithStack(err)
	}
	return cacheData, nil
}

// saveCacheData 存储 缓存数据
func (s *redisToken) saveCacheData(ctx context.Context, cacheKey string, authInfo *authv1.Auth, expiration time.Duration) error {
	cacheData, marshalErr := proto.Marshal(authInfo)
	if marshalErr != nil {
		err := authutil.ErrMarshalRedisData
		err.Metadata = map[string]string{"error": marshalErr.Error()}
		return errorutil.WithStack(err)
	}

	cacheErr := s.redisCC.Set(ctx, cacheKey, cacheData, expiration).Err()
	if cacheErr != nil {
		err := authutil.ErrSetRedisData
		err.Metadata = map[string]string{"error": cacheErr.Error()}
		return errorutil.WithStack(err)
	}
	return nil
}
