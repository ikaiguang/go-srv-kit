package tokenutil

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/proto"
	"time"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	uuidutil "github.com/ikaiguang/go-srv-kit/kit/uuid"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// redisToken ...
type redisToken struct {
	redisCC *redis.Client
	opt     options
}

// NewRedisTokenRepo ...
// 1. 生产签名密码 SigningSecret
// 2. 确定签名方法 JWTSigningMethod
// 3. 签证令牌 SignedToken
// 4. 生产缓存key CacheKey
// 5. 存储令牌 SaveCacheData
func NewRedisTokenRepo(redisCC *redis.Client, opts ...Option) AuthTokenRepo {
	o := &options{}
	for i := range opts {
		opts[i](o)
	}
	if o.authConfig == nil {
		o.authConfig = &confv1.App_Auth{}
	}
	return &redisToken{
		redisCC: redisCC,
		opt:     *o,
	}
}

// SigningSecret 签名密码
func (s *redisToken) SigningSecret(ctx context.Context, authClaims *authutil.Claims, passwordHash string) string {
	if passwordHash == "" {
		passwordHash = uuidutil.New()
		logutil.Warn("SignedString secret is empty; new secret = " + passwordHash)
	}
	if s.opt.signingSecretFunc != nil {
		return s.opt.signingSecretFunc(ctx, authClaims, passwordHash)
	}
	switch authClaims.Payload.Tt {
	case authv1.TokenTypeEnum_DEFAULT:
		return NewSecret(s.opt.authConfig.DefaultKey, passwordHash)
	case authv1.TokenTypeEnum_SERVICE:
		return NewSecret(s.opt.authConfig.ServiceKey, passwordHash)
	case authv1.TokenTypeEnum_ADMIN:
		return NewSecret(s.opt.authConfig.AdminKey, passwordHash)
	case authv1.TokenTypeEnum_API:
		return NewSecret(s.opt.authConfig.ApiKey, passwordHash)
	case authv1.TokenTypeEnum_WEB:
		return NewSecret(s.opt.authConfig.WebKey, passwordHash)
	case authv1.TokenTypeEnum_APP:
		return NewSecret(s.opt.authConfig.AppKey, passwordHash)
	case authv1.TokenTypeEnum_H5:
		return NewSecret(s.opt.authConfig.H5Key, passwordHash)
	default:
		return NewSecret(s.opt.authConfig.DefaultKey, passwordHash)
	}
}

// JWTSigningMethod jwt 签名方法
func (s *redisToken) JWTSigningMethod() *jwt.SigningMethodHMAC {
	return jwt.SigningMethodHS256
}

// SignedToken 签证Token
func (s *redisToken) SignedToken(authClaims *authutil.Claims, signingSecret string) (string, error) {
	return jwt.NewWithClaims(s.JWTSigningMethod(), authClaims).SignedString([]byte(signingSecret))
}

// CacheKey 缓存key、...
func (s *redisToken) CacheKey(ctx context.Context, authClaims *authutil.Claims) string {
	if s.opt.cacheKeyFunc != nil {
		return s.opt.cacheKeyFunc(ctx, authClaims)
	}
	switch authClaims.Payload.Tt {
	case authv1.TokenTypeEnum_DEFAULT:
		return AuthKeyForDefault(authClaims.Payload)
	case authv1.TokenTypeEnum_SERVICE:
		return AuthKeyForService(authClaims.Payload)
	case authv1.TokenTypeEnum_ADMIN:
		return AuthKeyForAdmin(authClaims.Payload)
	case authv1.TokenTypeEnum_API:
		return AuthKeyForApi(authClaims.Payload)
	case authv1.TokenTypeEnum_WEB:
		return AuthKeyForWeb(authClaims.Payload)
	case authv1.TokenTypeEnum_APP:
		return AuthKeyForApp(authClaims.Payload)
	case authv1.TokenTypeEnum_H5:
		return AuthKeyForH5(authClaims.Payload)
	default:
		return AuthKeyForDefault(authClaims.Payload)
	}
}

// SaveCacheData 存储缓存
func (s *redisToken) SaveCacheData(ctx context.Context, authClaims *authutil.Claims, authCache *authv1.Auth) error {
	var (
		cacheKey   = s.CacheKey(ctx, authClaims)
		expiration = authutil.ExpireDuration
	)
	if ex := authClaims.ExpiresAt.Sub(time.Now()); ex > 0 {
		expiration = ex
	}
	return s.saveCacheData(ctx, cacheKey, authCache, expiration)
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

		// 存储
		ctx = authutil.NewRedisContext(ctx, authInfo)

		// 密码
		return []byte(authInfo.Secret), nil
	}
	return ctx, keyFunc
}

// ValidateAuthInfo 校验：验证信息
func (s *redisToken) ValidateAuthInfo(authClaims *authutil.Claims, authInfo *authv1.Auth) (err error) {
	// 无限制
	switch authInfo.Payload.Lt {
	case authv1.LimitTypeEnum_UNKNOWN, authv1.LimitTypeEnum_UNLIMITED:
		return err
	case authv1.LimitTypeEnum_ONLY_ONE:
		// 仅一个
		if authClaims.Payload.St.AsTime().UnixNano() != authInfo.Payload.St.AsTime().UnixNano() {
			err = errorutil.WithStack(authutil.ErrLoginLimit)
			return err
		}
	case authv1.LimitTypeEnum_SAME_PLATFORM:
		// 平台限制一个
		if authClaims.Payload.Lp == authInfo.Payload.Lp &&
			authClaims.Payload.St.AsTime().UnixNano() != authInfo.Payload.St.AsTime().UnixNano() {
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
