package redis

import (
	"context"
)

// onceLock ...
type onceLock struct {
	mutex    *redsync.Mutex
	lockName string
}

// Unlock 解锁
func (s *onceLock) Unlock(ctx context.Context) (ok bool, err error) {
	return s.mutex.UnlockContext(ctx)
}

func (s *onceLock) Name() string {
	return s.lockName
}
