package snowflakepkg

import (
	"context"
	"time"

	redispkg "github.com/ikaiguang/go-srv-kit/data/redis"
	lockerpkg "github.com/ikaiguang/go-srv-kit/kit/locker"
	"github.com/redis/go-redis/v9"
)

const (
	DefaultMaxNodeId      = 1023             // 最大节点ID
	DefaultIdleDuration   = 60 * time.Second // 空闲ID时间：超过16s不续期，节点ID变为空闲的ID
	DefaultExtentDuration = 5 * time.Second  // 续期间隔时间
)

// Locker ...
type Locker interface {
	Lock(ctx context.Context, lockName string) (locker lockerpkg.Unlocker, err error)
}

var (
	_ Locker = (*cacheRepo)(nil)
	_ Locker = (*redisRepo)(nil)
)

// NewLockerFromCache ...
func NewLockerFromCache() Locker {
	return &cacheRepo{
		locker: lockerpkg.NewCacheLocker(),
	}
}

// NewLockerFromRedis ...
func NewLockerFromRedis(redisCC redis.UniversalClient) Locker {
	return &redisRepo{
		locker: redispkg.NewLocker(redisCC),
	}
}

// Lock ...
func (s *cacheRepo) Lock(ctx context.Context, lockName string) (lockerpkg.Unlocker, error) {
	locker, err := s.locker.Mutex(ctx, lockName)
	if err != nil {
		if lockerpkg.IsErrorLockFailed(err) {
			err = nil
			time.Sleep(time.Millisecond * 30)
			return s.Lock(ctx, lockName)
		}
		return locker, err
	}
	return locker, nil
}

// Lock ...
func (s *redisRepo) Lock(ctx context.Context, lockName string) (lockerpkg.Unlocker, error) {
	locker, err := s.locker.Mutex(ctx, lockName)
	if err != nil {
		if lockerpkg.IsErrorLockFailed(err) {
			err = nil
			time.Sleep(time.Millisecond * 30)
			return s.Lock(ctx, lockName)
		}
		return locker, err
	}
	return locker, nil
}

// cacheRepo ...
type cacheRepo struct {
	locker lockerpkg.Locker
}

// redisRepo ...
type redisRepo struct {
	locker lockerpkg.Locker
}
