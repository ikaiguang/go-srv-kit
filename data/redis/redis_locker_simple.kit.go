package redisutil

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	lockerutil "github.com/ikaiguang/go-srv-kit/kit/locker"
)

// NewSimpleLocker 获取锁
// 此加锁，V4锁的默认时间：8s；8秒后锁自动过期
func NewSimpleLocker(client *redis.Client) (locker lockerutil.Locker, err error) {
	handler := new(redisSimpleLocker)
	handler.Init(client)

	return handler, err
}

// redisSimpleLocker 简单锁，加锁时间默认8秒；8秒后锁自动过期
type redisSimpleLocker struct {
	rs    *redsync.Redsync
	mutex *redsync.Mutex
}

// Init .
func (s *redisSimpleLocker) Init(client *redis.Client) {
	s.rs = redsync.New(goredis.NewPool(client))
}

// Lock .
func (s *redisSimpleLocker) Lock(ctx context.Context, name string) (err error) {
	s.mutex = s.rs.NewMutex(name, redsync.WithExpiry(lockExpire))
	if err = s.mutex.LockContext(ctx); err != nil {
		err = lockerutil.NewErrLockerFailed(true, name, err)
		return err
	}
	return err
}

// Unlock 解锁
func (s *redisSimpleLocker) Unlock(ctx context.Context) (ok bool, err error) {
	return s.mutex.UnlockContext(ctx)
}
