package redisutil

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	lockerutil "github.com/ikaiguang/go-srv-kit/kit/locker"
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
func NewLocker(redisCC *redis.Client, opts ...redsync.Option) lockerutil.Lock {
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
func (s *Locker) Once(ctx context.Context, lockName string) (locker lockerutil.Unlock, err error) {
	m := &onceLock{}
	m.mutex = s.rs.NewMutex(lockName, s.opts...)
	if err = m.mutex.LockContext(ctx); err != nil {
		err = lockerutil.NewErrLockerFailed(true, lockName, err)
		return m, err
	}
	return m, err
}

// Mutex ...
func (s *Locker) Mutex(ctx context.Context, lockName string) (locker lockerutil.Unlock, err error) {
	m := &mutexLock{}
	m.mutex = s.rs.NewMutex(lockName, s.opts...)
	if err = m.mutex.LockContext(ctx); err != nil {
		err = lockerutil.NewErrLockerFailed(true, lockName, err)
		return m, err
	}

	// 续期锁，防止锁自动过期
	m.extendChannel = make(chan bool)
	go m.extend(ctx)

	return m, err
}
