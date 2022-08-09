package tokenutil

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	uuidutil "github.com/ikaiguang/go-srv-kit/kit/uuid"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// redisToken ...
type redisToken struct {
	redisCC        *redis.Client
	tokenTypeMutex sync.Mutex
	opt            options
}

// NewRedisTokenRepo ...
func NewRedisTokenRepo(redisCC *redis.Client, opts ...Option) AuthTokenRepo {
	o := &options{}
	for i := range opts {
		opts[i](o)
	}
	if o.authConfig == nil {
		o.authConfig = &confv1.App_Auth{}
	}
	if o.tokenTypeMap == nil {
		o.tokenTypeMap = newTokenTypeMap()
	}
	return &redisToken{
		redisCC: redisCC,
		opt:     *o,
	}
}

// SigningSecret 签名密码
func (s *redisToken) SigningSecret(ctx context.Context, tokenType authv1.TokenTypeEnum_TokenType, passwordHash string) string {
	if passwordHash == "" {
		passwordHash = uuidutil.New()
		logutil.Warn("SignedString secret is empty; new secret = " + passwordHash)
	}
	if s.opt.signingSecretFunc != nil {
		return s.opt.signingSecretFunc(ctx, tokenType, passwordHash)
	}
	return NewSecret(s.opt.authConfig, tokenType, passwordHash)
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
	return NewCacheKey(authClaims.Payload)
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

// DeleteCacheData 删除缓存
func (s *redisToken) DeleteCacheData(ctx context.Context, authClaims *authutil.Claims) error {
	return s.deleteCacheData(ctx, s.CacheKey(ctx, authClaims))
}

// SetTokenType 设置令牌类型
func (s *redisToken) SetTokenType(operation string, tokenType authv1.TokenTypeEnum_TokenType) {
	// 写锁
	s.tokenTypeMutex.Lock()
	defer s.tokenTypeMutex.Unlock()

	s.opt.tokenTypeMap[operation] = tokenType
}

// GetTokenType 获取令牌类型
func (s *redisToken) GetTokenType(operation string) authv1.TokenTypeEnum_TokenType {
	// 暂不加锁：不考虑动态增加令牌，并发读
	//s.tokenTypeRWMutex.RLock()
	//defer s.tokenTypeRWMutex.RUnlock()

	return s.opt.tokenTypeMap[operation]
}

// JWTKeyFunc 验证工具： authutil.KeyFunc，提供最终的 jwt.Keyfunc
func (s *redisToken) JWTKeyFunc() authutil.KeyFunc {
	// authutil.KeyFunc
	authKeyFunc := func(ctx context.Context) jwt.Keyfunc {
		// 令牌类型
		tokenType := authv1.TokenTypeEnum_DEFAULT
		if s.opt.tokenTypeFunc != nil {
			tokenType = s.opt.tokenTypeFunc(ctx, s.opt.tokenTypeMap)
		} else {
			tokenType = s.GetTokenType(GetRequestOperation(ctx))
		}

		// jwt key func
		jwtKeyFunc := func(token *jwt.Token) (interface{}, error) {
			myClaims, ok := token.Claims.(*authutil.Claims)
			if !ok || myClaims.Payload == nil {
				return []byte(""), authutil.ErrInvalidAuthInfo
			}

			// key
			key := s.CacheKey(ctx, myClaims)
			authInfo, err := s.getCacheData(ctx, key)
			if err != nil {
				return nil, err
			}

			// 存储信息
			authutil.SaveAuthInfo(token.Header, authInfo)

			// 密码
			signingSecret := s.SigningSecret(ctx, tokenType, authInfo.Secret)
			return []byte(signingSecret), nil
		}
		return jwtKeyFunc
	}
	return authKeyFunc
}

// ValidateFunc 自定义验证
func (s *redisToken) ValidateFunc() authutil.ValidateFunc {
	validator := func(ctx context.Context, jwtToken *jwt.Token) error {
		authClaims, ok := jwtToken.Claims.(*authutil.Claims)
		if !ok {
			return authutil.ErrInvalidAuthInfo
		}
		authInfo, ok := authutil.GetAuthInfo(jwtToken.Header)
		if !ok {
			return authutil.ErrInvalidAuthInfo
		}
		return s.validateAuthInfo(authClaims, authInfo)
	}
	return validator
}

// validateAuthInfo 校验：验证信息
func (s *redisToken) validateAuthInfo(authClaims *authutil.Claims, authInfo *authv1.Auth) (err error) {
	// 无限制
	switch authInfo.Payload.Lt {
	case authv1.LimitTypeEnum_UNLIMITED:
		return err
	case authv1.LimitTypeEnum_ONLY_ONE:
		// 仅一个
		if authClaims.Payload.St.AsTime().UnixNano() != authInfo.Payload.St.AsTime().UnixNano() {
			return authutil.ErrLoginLimit
		}
	case authv1.LimitTypeEnum_PLATFORM_ONE:
		// 平台限制一个
		if authClaims.Payload.Lp == authInfo.Payload.Lp &&
			authClaims.Payload.St.AsTime().UnixNano() != authInfo.Payload.St.AsTime().UnixNano() {
			return authutil.ErrLoginLimit
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
		return nil, err
	}

	// cache
	cacheData := &authv1.Auth{}
	unmarshalErr := proto.Unmarshal(cacheBytes, cacheData)
	if unmarshalErr != nil {
		err := authutil.ErrUnmarshalRedisData
		err.Metadata = map[string]string{"error": unmarshalErr.Error()}
		return cacheData, err
	}
	return cacheData, nil
}

// saveCacheData 存储 缓存数据
func (s *redisToken) saveCacheData(ctx context.Context, cacheKey string, authInfo *authv1.Auth, expiration time.Duration) error {
	cacheData, marshalErr := proto.Marshal(authInfo)
	if marshalErr != nil {
		err := authutil.ErrMarshalRedisData
		err.Metadata = map[string]string{"error": marshalErr.Error()}
		return err
	}

	cacheErr := s.redisCC.Set(ctx, cacheKey, cacheData, expiration).Err()
	if cacheErr != nil {
		err := authutil.ErrSetRedisData
		err.Metadata = map[string]string{"error": cacheErr.Error()}
		return err
	}
	return nil
}

// deleteCacheData 删除 缓存数据
func (s *redisToken) deleteCacheData(ctx context.Context, cacheKey string) error {
	cacheErr := s.redisCC.Del(ctx, cacheKey).Err()
	if cacheErr != nil {
		err := authutil.ErrSetRedisData
		err.Metadata = map[string]string{"error": cacheErr.Error()}
		return err
	}
	return nil
}
