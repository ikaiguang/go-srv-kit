package authpkg

import (
	"context"
	"strconv"
	"time"

	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	threadpkg "github.com/ikaiguang/go-srv-kit/kratos/thread"
	"github.com/redis/go-redis/v9"
)

// RedisCacheKeyPrefix ...
type RedisCacheKeyPrefix string

func (s RedisCacheKeyPrefix) String() string {
	return string(s)
}

const (
	DefaultBlackTokenKeyPrefix RedisCacheKeyPrefix = "kit:auth_black:"
	DefaultLoginLimitKeyPrefix RedisCacheKeyPrefix = "kit:auth_limit:"
	DefaultAuthTokenKeyPrefix  RedisCacheKeyPrefix = "kit:auth_token:"
)

// AuthCacheKeyPrefix ...
type AuthCacheKeyPrefix struct {
	TokensKeyPrefix     RedisCacheKeyPrefix // 用户令牌
	BlackTokenKeyPrefix RedisCacheKeyPrefix // 黑名单
	LimitTokenKeyPrefix RedisCacheKeyPrefix // 登录限制
}

// CheckAuthCacheKeyPrefix ...
func CheckAuthCacheKeyPrefix(keyPrefix *AuthCacheKeyPrefix) *AuthCacheKeyPrefix {
	if keyPrefix == nil {
		return &AuthCacheKeyPrefix{
			TokensKeyPrefix:     DefaultAuthTokenKeyPrefix,
			BlackTokenKeyPrefix: DefaultBlackTokenKeyPrefix,
		}
	}
	if keyPrefix.TokensKeyPrefix == "" {
		keyPrefix.TokensKeyPrefix = DefaultAuthTokenKeyPrefix
	}
	if keyPrefix.BlackTokenKeyPrefix == "" {
		keyPrefix.BlackTokenKeyPrefix = DefaultBlackTokenKeyPrefix
	}
	if keyPrefix.LimitTokenKeyPrefix == "" {
		keyPrefix.LimitTokenKeyPrefix = DefaultLoginLimitKeyPrefix
	}
	return keyPrefix
}

var _ TokenManger = (*tokenManger)(nil)

// TokenManger ...
type TokenManger interface {
	SaveTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	DeleteTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	GetToken(ctx context.Context, userIdentifier string, tokenID string) (item *TokenItem, isNotFound bool, err error)
	GetAllTokens(ctx context.Context, userIdentifier string) (map[string]*TokenItem, error)
	IsExistToken(ctx context.Context, userIdentifier string, tokenID string) (bool, error)
	AddBlacklist(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	IsBlacklist(ctx context.Context, tokenID string) (bool, error)
	AddLoginLimit(ctx context.Context, tokenItems []*TokenItem) error
	IsLoginLimit(ctx context.Context, tokenID string) (bool, LoginLimitEnum_LoginLimit, error)
}

// tokenManger ...
type tokenManger struct {
	redisCC            redis.UniversalClient
	authCacheKeyPrefix *AuthCacheKeyPrefix
}

// NewTokenManger ...
func NewTokenManger(redisCC redis.UniversalClient, authCacheKeyPrefix *AuthCacheKeyPrefix) TokenManger {
	authCacheKeyPrefix = CheckAuthCacheKeyPrefix(authCacheKeyPrefix)
	return &tokenManger{
		redisCC:            redisCC,
		authCacheKeyPrefix: authCacheKeyPrefix,
	}
}

// SaveTokens ...
func (s *tokenManger) SaveTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error {
	if len(tokenItems) == 0 {
		return nil
	}

	var (
		kvs     = make([]interface{}, 0, 2*len(tokenItems))
		nowUnix = time.Now().Unix()
		expire  = time.Duration(0)
	)
	for i := range tokenItems {
		if tokenItems[i].IsRefreshToken {
			kvs = append(kvs, tokenItems[i].RefreshTokenID)
		} else {
			kvs = append(kvs, tokenItems[i].TokenID)
		}
		itemStr, err := tokenItems[i].EncodeToString()
		if err != nil {
			e := errorpkg.ErrorBadRequest("encode token item failed")
			err = errorpkg.Wrap(e, err)
			return err
		}
		kvs = append(kvs, itemStr)

		// 过期时间
		if ex := s.calcExpireTime(tokenItems[i].ExpiredAt, nowUnix); ex > expire {
			expire = ex
		}
	}

	key := s.genTokensKey(userIdentifier)
	if err := s.redisCC.HSet(ctx, key, kvs...).Err(); err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	threadpkg.GoSafe(func() {
		_ = s.redisCC.Expire(ctx, key, expire)
	})
	return nil
}

// calcExpireTime ...
func (s *tokenManger) calcExpireTime(expireAt, nowUnix int64) time.Duration {
	if expireAt == 0 {
		return 0
	}
	t := time.Duration(expireAt-nowUnix) * time.Second
	if t > 0 {
		return t
	}
	return time.Second
}

// AddBlacklist ...
func (s *tokenManger) AddBlacklist(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error {
	if len(tokenItems) == 0 {
		return nil
	}

	var (
		tokensKey = s.genTokensKey(userIdentifier)
		hashKeys  = make([]string, 0, len(tokenItems))
		nowUnix   = time.Now().Unix()
	)

	pipe := s.redisCC.Pipeline()
	for i := range tokenItems {
		var blackKey = ""
		if tokenItems[i].IsRefreshToken {
			hashKeys = append(hashKeys, tokenItems[i].RefreshTokenID)
			blackKey = s.genBlackTokenKey(tokenItems[i].RefreshTokenID)
		} else {
			hashKeys = append(hashKeys, tokenItems[i].TokenID)
			blackKey = s.genBlackTokenKey(tokenItems[i].TokenID)
		}
		// 加入黑名单
		d := s.calcExpireTime(tokenItems[i].ExpiredAt, nowUnix)
		if err := pipe.Set(ctx, blackKey, 0, d).Err(); err != nil {
			e := errorpkg.ErrorInternalServer("")
			return errorpkg.Wrap(e, err)
		}
	}
	if err := pipe.HDel(ctx, tokensKey, hashKeys...).Err(); err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return nil
}

// DeleteTokens ...
func (s *tokenManger) DeleteTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error {
	if len(tokenItems) == 0 {
		return nil
	}

	var (
		tokensKey = s.genTokensKey(userIdentifier)
		hashKeys  = make([]string, 0, len(tokenItems))
	)

	for i := range tokenItems {
		if tokenItems[i].IsRefreshToken {
			hashKeys = append(hashKeys, tokenItems[i].RefreshTokenID)
		} else {
			hashKeys = append(hashKeys, tokenItems[i].TokenID)
		}
	}
	if err := s.redisCC.HDel(ctx, tokensKey, hashKeys...).Err(); err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return nil
}

// AddLoginLimit ...
func (s *tokenManger) AddLoginLimit(ctx context.Context, tokenItems []*TokenItem) error {
	if len(tokenItems) == 0 {
		return nil
	}
	var (
		nowUnix = time.Now().Unix()
	)

	pipe := s.redisCC.Pipeline()
	for i := range tokenItems {
		limitKey := s.genLimitTokenKey(tokenItems[i].TokenID)
		d := s.calcExpireTime(tokenItems[i].ExpiredAt, nowUnix)
		if err := pipe.Set(ctx, limitKey, tokenItems[i].Payload.LoginLimit, d).Err(); err != nil {
			e := errorpkg.ErrorInternalServer("")
			return errorpkg.Wrap(e, err)
		}
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return nil
}

// IsBlacklist ...
func (s *tokenManger) IsBlacklist(ctx context.Context, tokenID string) (bool, error) {
	blackKey := s.genBlackTokenKey(tokenID)
	i, err := s.redisCC.Exists(ctx, blackKey).Result()
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return false, err
	}
	return i > 0, nil
}

// IsLoginLimit ...
func (s *tokenManger) IsLoginLimit(ctx context.Context, tokenID string) (bool, LoginLimitEnum_LoginLimit, error) {
	limitKey := s.genLimitTokenKey(tokenID)
	loginLimitStr, err := s.redisCC.Get(ctx, limitKey).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return false, LoginLimitEnum_UNLIMITED, err
	}
	ll, _ := strconv.Atoi(loginLimitStr)
	return true, LoginLimitEnum_LoginLimit(int32(ll)), nil
}

// GetToken ...
func (s *tokenManger) GetToken(ctx context.Context, userIdentifier string, tokenID string) (item *TokenItem, isNotFound bool, err error) {
	key := s.genTokensKey(userIdentifier)
	res, err := s.redisCC.HGet(ctx, key, tokenID).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return item, isNotFound, err
	}
	item = &TokenItem{}
	err = item.DecodeString(res)
	if err != nil {
		return item, isNotFound, err
	}
	return item, isNotFound, err
}

// IsExistToken ...
func (s *tokenManger) IsExistToken(ctx context.Context, userIdentifier string, tokenID string) (bool, error) {
	key := s.genTokensKey(userIdentifier)

	exist, err := s.redisCC.HExists(ctx, key, tokenID).Result()
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return false, err
	}
	return exist, nil
}

// GetAllTokens ...
func (s *tokenManger) GetAllTokens(ctx context.Context, userIdentifier string) (map[string]*TokenItem, error) {
	key := s.genTokensKey(userIdentifier)
	tokens, err := s.redisCC.HGetAll(ctx, key).Result()
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return nil, err
	}

	var items = make(map[string]*TokenItem)
	for iKey := range tokens {
		item := &TokenItem{}
		err = item.DecodeString(tokens[iKey])
		if err != nil {
			return nil, err
		}
		items[iKey] = item
	}
	return items, nil
}

// genTokensKey ...
func (s *tokenManger) genTokensKey(userIdentifier string) string {
	return s.authCacheKeyPrefix.TokensKeyPrefix.String() + userIdentifier
}

// genBlackTokenKey ...
func (s *tokenManger) genBlackTokenKey(tokenID string) string {
	return s.authCacheKeyPrefix.BlackTokenKeyPrefix.String() + tokenID
}

// genLimitTokenKey ...
func (s *tokenManger) genLimitTokenKey(tokenID string) string {
	return s.authCacheKeyPrefix.LimitTokenKeyPrefix.String() + tokenID
}
