package redispkg

import (
	"context"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"

	lockerpkg "github.com/ikaiguang/go-srv-kit/kit/locker"
)

const (
	_lockExpire      = 8 * time.Second // 锁过期时间
	_lockExtendDelay = 3 * time.Second // 重置锁时间，防止锁自动过期
	_lockTries       = 1               // 尝试次数；修改尝试过大，会导致加锁成功；
)

// Locker 锁
type Locker struct {
	rs   *redsync.Redsync
	opts []redsync.Option
}

// NewLocker ..
func NewLocker(redisCC redis.UniversalClient, opts ...redsync.Option) lockerpkg.Locker {
	// 锁选项
	lockerOpts := []redsync.Option{
		redsync.WithExpiry(_lockExpire),
		redsync.WithTries(_lockTries),
	}
	lockerOpts = append(lockerOpts, opts...)

	// 锁
	return &Locker{
		rs:   redsync.New(goredis.NewPool(redisCC)),
		opts: lockerOpts,
	}
}

// Once ...
func (s *Locker) Once(ctx context.Context, lockName string) (locker lockerpkg.Unlocker, err error) {
	m := &onceLock{lockName: lockName}
	m.mutex = s.rs.NewMutex(lockName, s.opts...)
	if err = m.mutex.LockContext(ctx); err != nil {
		err = lockerpkg.ErrorLockerFailed(lockName, err)
		return m, err
	}
	return m, err
}

// Mutex ...
func (s *Locker) Mutex(ctx context.Context, lockName string) (locker lockerpkg.Unlocker, err error) {
	m := &mutexLock{lockName: lockName}
	m.mutex = s.rs.NewMutex(lockName, s.opts...)
	if err = m.mutex.LockContext(ctx); err != nil {
		err = lockerpkg.ErrorLockerFailed(lockName, err)
		return m, err
	}

	// 续期锁，防止锁自动过期
	m.stopExtend = make(chan bool)
	go m.extend(ctx)

	return m, err
}
