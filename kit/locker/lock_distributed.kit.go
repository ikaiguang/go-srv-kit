package lockerpkg

import "context"

// Deprecated: DistributedLocker 仅是 Locker 的方法重命名包装，
// MutexLock 等价于 Mutex，EasyLock 等价于 Once。
// 请直接使用 Locker 接口。
type DistributedLocker interface {
	// MutexLock 互斥锁，一直等待直到解锁
	MutexLock(ctx context.Context, lockName string) (Unlocker, error)
	// EasyLock 简单锁，等待解锁或者锁定时间过期后自动解锁
	EasyLock(ctx context.Context, lockName string) (Unlocker, error)
}

// Deprecated: 请直接使用 Locker 接口。
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
