package lockerpkg

import "context"

// DistributedLocker 分布式锁
type DistributedLocker interface {
	// MutexLock 互斥锁，一直等待直到解锁
	MutexLock(ctx context.Context, lockName string) (Unlocker, error)
	// EasyLock 简单锁，等待解锁或者锁定时间过期后自动解锁
	EasyLock(ctx context.Context, lockName string) (Unlocker, error)
}

// NewDistributedLocker 请确保传递进来的锁支持分布式锁
func NewDistributedLocker(locker Locker) DistributedLocker {
	return &distributedLock{locker: locker}
}

type distributedLock struct {
	locker Locker
}

func (s *distributedLock) MutexLock(ctx context.Context, lockName string) (Unlocker, error) {
	return s.locker.Mutex(ctx, lockName)
}

func (s *distributedLock) EasyLock(ctx context.Context, lockName string) (Unlocker, error) {
	return s.locker.Once(ctx, lockName)
}
