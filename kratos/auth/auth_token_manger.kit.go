package authpkg

import (
	"context"
	stderrors "errors"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	redispkg "github.com/ikaiguang/go-srv-kit/data/redis"
	lockerpkg "github.com/ikaiguang/go-srv-kit/kit/locker"
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
	DefaultClearTokenKeyPrefix RedisCacheKeyPrefix = "kit:auth_clear:"
)

// AuthCacheKeyPrefix ...
type AuthCacheKeyPrefix struct {
	TokensKeyPrefix     RedisCacheKeyPrefix // 用户令牌
	BlackTokenKeyPrefix RedisCacheKeyPrefix // 黑名单
	LimitTokenKeyPrefix RedisCacheKeyPrefix // 登录限制
	ClearTokenKeyPrefix RedisCacheKeyPrefix // 清除token
}

// CheckAuthCacheKeyPrefix ...
func CheckAuthCacheKeyPrefix(inputKeyPrefix *AuthCacheKeyPrefix) *AuthCacheKeyPrefix {
	keyPrefix := &AuthCacheKeyPrefix{
		TokensKeyPrefix:     DefaultAuthTokenKeyPrefix,  // 用户令牌
		BlackTokenKeyPrefix: DefaultBlackTokenKeyPrefix, // 黑名单
		LimitTokenKeyPrefix: DefaultLoginLimitKeyPrefix, // 登录限制
		ClearTokenKeyPrefix: DefaultClearTokenKeyPrefix, // 清除token
	}
	if inputKeyPrefix == nil {
		return keyPrefix
	}
	if inputKeyPrefix.TokensKeyPrefix == "" {
		keyPrefix.TokensKeyPrefix = inputKeyPrefix.TokensKeyPrefix
	}
	if inputKeyPrefix.BlackTokenKeyPrefix == "" {
		keyPrefix.BlackTokenKeyPrefix = inputKeyPrefix.BlackTokenKeyPrefix
	}
	if inputKeyPrefix.LimitTokenKeyPrefix == "" {
		keyPrefix.LimitTokenKeyPrefix = inputKeyPrefix.LimitTokenKeyPrefix
	}
	if inputKeyPrefix.ClearTokenKeyPrefix == "" {
		keyPrefix.ClearTokenKeyPrefix = inputKeyPrefix.ClearTokenKeyPrefix
	}
	return keyPrefix
}

var _ TokenManger = (*tokenManger)(nil)

// TokenManger ...
type TokenManger interface {
	SaveAccessTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	ResetPreviousTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	AddBlacklist(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	AddLoginLimit(ctx context.Context, tokenItems []*TokenItem) error

	GetToken(ctx context.Context, userIdentifier string, tokenID string) (item *TokenItem, isNotFound bool, err error)
	GetAllTokens(ctx context.Context, userIdentifier string) (map[string]*TokenItem, error)

	DeleteTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	DeleteExpireTokens(ctx context.Context, userIdentifier string) error

	IsLoginLimit(ctx context.Context, tokenID string) (bool, LoginLimitEnum_LoginLimit, error)
	IsExistToken(ctx context.Context, userIdentifier string, tokenID string) (bool, error)
	IsBlacklist(ctx context.Context, tokenID string) (bool, error)

	// EasyLock 简单锁，等待解锁或者锁定时间过期后自动解锁
	EasyLock(ctx context.Context, lockName string) (lockerpkg.Unlocker, error)
	// MutexLock 互斥锁，一直等待直到解锁
	// MutexLock(ctx context.Context, lockName string) (lockerpkg.Unlocker, error)
}

// tokenManger ...
type tokenManger struct {
	log                *log.Helper
	redisCC            redis.UniversalClient
	authCacheKeyPrefix *AuthCacheKeyPrefix
	locker             lockerpkg.Locker
}

// NewTokenManger ...
func NewTokenManger(
	logger log.Logger,
	redisCC redis.UniversalClient,
	authCacheKeyPrefix *AuthCacheKeyPrefix,
) TokenManger {
	authCacheKeyPrefix = CheckAuthCacheKeyPrefix(authCacheKeyPrefix)
	return &tokenManger{
		log:                log.NewHelper(log.With(logger, "module", "kit.auth.token.manger")),
		redisCC:            redisCC,
		authCacheKeyPrefix: authCacheKeyPrefix,
		locker:             redispkg.NewLocker(redisCC),
	}
}

// SaveAccessTokens ...
func (s *tokenManger) SaveAccessTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error {
	if len(tokenItems) == 0 {
		return nil
	}

	var (
		kvs     = make([]interface{}, 0, 2*len(tokenItems))
		nowUnix = time.Now().Unix()
		expire  = time.Duration(0)
	)
	for i := range tokenItems {
		kvs = append(kvs, tokenItems[i].ID())
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
		// ctx many be cancel
		_ = s.redisCC.Expire(context.Background(), key, expire)
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

// ResetPreviousTokens ...
func (s *tokenManger) ResetPreviousTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error {
	if len(tokenItems) == 0 {
		return nil
	}

	var (
		kvs = make([]interface{}, 0, 2*len(tokenItems))
	)
	for i := range tokenItems {
		kvs = append(kvs, tokenItems[i].ID())
		itemStr, err := tokenItems[i].EncodeToString()
		if err != nil {
			e := errorpkg.ErrorBadRequest("encode token item failed")
			err = errorpkg.Wrap(e, err)
			return err
		}
		kvs = append(kvs, itemStr)
	}

	key := s.genTokensKey(userIdentifier)
	if err := s.redisCC.HSet(ctx, key, kvs...).Err(); err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return nil
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
		hashKeys = append(hashKeys, tokenItems[i].ID())
		blackKey := s.genBlacklistTokenKey(tokenItems[i].ID())
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

// DeleteExpireTokens 删除过期的token
func (s *tokenManger) DeleteExpireTokens(ctx context.Context, userIdentifier string) error {
	var (
		nowUnix    = time.Now().Unix()
		expireList []*TokenItem
	)

	// 防止重复删除
	unlocker, err := s.EasyLock(ctx, s.genClearTokenKey(userIdentifier))
	if err != nil {
		return err
	}
	defer func() { _, _ = unlocker.Unlock(ctx) }()

	allTokens, err := s.GetAllTokens(ctx, userIdentifier)
	if err != nil {
		return err
	}
	for i := range allTokens {
		if allTokens[i].ExpiredAt > nowUnix {
			continue
		}
		expireList = append(expireList, allTokens[i])
	}

	// 删除过期
	if err = s.DeleteTokens(ctx, userIdentifier, expireList); err != nil {
		e := errorpkg.ErrorBadRequest("DeleteTokens failed")
		err = errorpkg.Wrap(e, err)
		return err
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
		hashKeys = append(hashKeys, tokenItems[i].ID())
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
	blackKey := s.genBlacklistTokenKey(tokenID)
	i, err := s.redisCC.Exists(ctx, blackKey).Result()
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return false, err
	}
	return i > 0, nil
}

// EasyLock 简单锁，等待解锁或者锁定时间过期后自动解锁
func (s *tokenManger) EasyLock(ctx context.Context, lockName string) (lockerpkg.Unlocker, error) {
	return s.locker.Once(ctx, lockName)
}

// MutexLock 互斥锁，一直等待直到解锁
func (s *tokenManger) MutexLock(ctx context.Context, lockName string) (lockerpkg.Unlocker, error) {
	return s.locker.Mutex(ctx, lockName)
}

// IsLoginLimit ...
func (s *tokenManger) IsLoginLimit(ctx context.Context, tokenID string) (bool, LoginLimitEnum_LoginLimit, error) {
	limitKey := s.genLimitTokenKey(tokenID)
	loginLimitStr, err := s.redisCC.Get(ctx, limitKey).Result()
	if err != nil {
		if stderrors.Is(err, redis.Nil) {
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
		if stderrors.Is(err, redis.Nil) {
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
	if item.TokenID == "" {
		isNotFound = true
	}
	return item, isNotFound, err
}

// IsExistToken ...
func (s *tokenManger) IsExistToken(ctx context.Context, userIdentifier string, tokenID string) (bool, error) {
	//key := s.genTokensKey(userIdentifier)
	//exist, err := s.redisCC.HExists(ctx, key, tokenID).Result()
	//if err != nil {
	//	e := errorpkg.ErrorInternalServer("")
	//	err = errorpkg.Wrap(e, err)
	//	return false, err
	//}

	tokenItem, isNotFound, err := s.GetToken(ctx, userIdentifier, tokenID)
	if err != nil {
		return false, err
	}
	if isNotFound {
		return false, nil
	}
	if time.Unix(tokenItem.ExpiredAt, 0).Before(time.Now()) {
		threadpkg.GoSafe(func() {
			// ctx many be cancel
			deleteErr := s.DeleteTokens(context.Background(), userIdentifier, []*TokenItem{tokenItem})
			if deleteErr != nil {
				s.log.WithContext(ctx).Warnw("msg", "DeleteTokens failed", "tokenItem", tokenItem)
			}
		})
		return false, nil
	}
	return true, nil
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

// genBlacklistTokenKey ...
func (s *tokenManger) genBlacklistTokenKey(tokenID string) string {
	return s.authCacheKeyPrefix.BlackTokenKeyPrefix.String() + tokenID
}

// genLimitTokenKey ...
func (s *tokenManger) genLimitTokenKey(tokenID string) string {
	return s.authCacheKeyPrefix.LimitTokenKeyPrefix.String() + tokenID
}

// genClearTokenKey ...
func (s *tokenManger) genClearTokenKey(tokenID string) string {
	return s.authCacheKeyPrefix.ClearTokenKeyPrefix.String() + tokenID
}
