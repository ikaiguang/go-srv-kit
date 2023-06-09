package snowflakepkg

import (
	"context"
	"fmt"
	"sync"
	"time"

	redispkg "github.com/ikaiguang/go-srv-kit/data/redis"
	lockerpkg "github.com/ikaiguang/go-srv-kit/kit/locker"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

const (
	DefaultMaxNodeId      = 1023             // 最大节点ID
	DefaultIdleDuration   = 60 * time.Second // 空闲ID时间：超过16s不续期，节点ID变为空闲的ID
	DefaultExtentDuration = 5 * time.Second  // 续期间隔时间
)

// Locker ...
type Locker interface {
	Lock(ctx context.Context, lockName string) (locker lockerpkg.Unlock, err error)
}

var (
	_ Locker = (*cacheRepo)(nil)
	_ Locker = (*redisRepo)(nil)
)

// NewLockerFromCache ...
func NewLockerFromCache(cacheHandler *cache.Cache) Locker {
	return &cacheRepo{
		cacheHandler: cacheHandler,
	}
}

// NewLockerFromRedis ...
func NewLockerFromRedis(redisCC redis.UniversalClient) Locker {
	return &redisRepo{
		locker: redispkg.NewLocker(redisCC),
	}
}

// Lock ...
func (s *cacheRepo) Lock(ctx context.Context, lockName string) (lockerpkg.Unlock, error) {
	// 读取锁
	muInterface, ok := s.cacheHandler.Get(lockName)
	if ok {
		mu, ok := muInterface.(*sync.Mutex)
		if ok {
			mu.Lock()
			unlocker := &cacheLocker{
				mu: mu,
			}
			return unlocker, nil
		}
		err := fmt.Errorf("[nodeID] cannot convert sync.Mutex : %s:%v", lockName, muInterface)
		return nil, err
	}

	// 添加锁
	mu := &sync.Mutex{}
	s.cacheHandler.Set(lockName, mu, time.Minute*5)
	mu.Lock()
	unlocker := &cacheLocker{
		mu: mu,
	}
	return unlocker, nil
}

// Lock ...
func (s *redisRepo) Lock(ctx context.Context, lockName string) (locker lockerpkg.Unlock, err error) {
	locker, err = s.locker.Mutex(ctx, lockName)
	if err != nil {
		if lockerpkg.IsErrLockFailed(err) {
			err = nil
			time.Sleep(time.Millisecond * 30)
			return s.Lock(ctx, lockName)
		}
		return locker, err
	}
	return locker, err
}

// cacheRepo ...
type cacheRepo struct {
	cacheHandler *cache.Cache
}

// redisRepo ...
type redisRepo struct {
	locker lockerpkg.Lock
}

// cacheLocker ...
type cacheLocker struct {
	mu *sync.Mutex
}

// Unlock ...
func (s *cacheLocker) Unlock(ctx context.Context) (bool, error) {
	s.mu.Unlock()
	return true, nil
}
