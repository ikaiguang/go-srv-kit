package redispkg

import (
	"context"

	"github.com/go-redsync/redsync/v4"
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
