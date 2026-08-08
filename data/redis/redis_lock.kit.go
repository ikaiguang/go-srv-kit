package redispkg

import (
	"context"
	"errors"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	lockerpkg "github.com/ikaiguang/go-srv-kit/kit/v3/locker"
	threadpkg "github.com/ikaiguang/go-srv-kit/kit/v3/thread"
	"github.com/redis/go-redis/v9"
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
	err  error
}

// NewLocker ..
func NewLocker(redisCC redis.UniversalClient, opts ...redsync.Option) lockerpkg.Locker {
	// 锁选项
	lockerOpts := []redsync.Option{
		redsync.WithExpiry(_lockExpire),
		redsync.WithTries(_lockTries),
	}
	lockerOpts = append(lockerOpts, opts...)

	if redisCC == nil {
		return &Locker{opts: lockerOpts, err: errors.New("redis client is nil")}
	}

	// 锁
	return &Locker{
		rs:   redsync.New(goredis.NewPool(redisCC)),
		opts: lockerOpts,
	}
}

// Once ...
func (s *Locker) Once(ctx context.Context, lockName string) (locker lockerpkg.Unlocker, err error) {
	if ctx == nil {
		return nil, lockerpkg.ErrorLockerFailed(lockName, errors.New("redis lock context is nil"))
	}
	if s == nil || s.err != nil || s.rs == nil {
		if s != nil && s.err != nil {
			err = s.err
		} else {
			err = errors.New("redis locker is not initialized")
		}
		return nil, lockerpkg.ErrorLockerFailed(lockName, err)
	}
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
	if ctx == nil {
		return nil, lockerpkg.ErrorLockerFailed(lockName, errors.New("redis lock context is nil"))
	}
	if s == nil || s.err != nil || s.rs == nil {
		if s != nil && s.err != nil {
			err = s.err
		} else {
			err = errors.New("redis locker is not initialized")
		}
		return nil, lockerpkg.ErrorLockerFailed(lockName, err)
	}
	m := &mutexLock{lockName: lockName}
	m.mutex = s.rs.NewMutex(lockName, s.opts...)
	if err = m.mutex.LockContext(ctx); err != nil {
		err = lockerpkg.ErrorLockerFailed(lockName, err)
		return m, err
	}

	// 续期锁，防止锁自动过期
	m.stopExtend = make(chan struct{})
	m.extendDone = make(chan struct{})
	threadpkg.GoSafe(func() {
		m.extend(ctx)
	})

	return m, err
}
