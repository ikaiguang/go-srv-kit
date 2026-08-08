package redispkg

import (
	"context"
	"errors"

	"github.com/go-redsync/redsync/v4"
)

// onceLock ...
type onceLock struct {
	mutex    *redsync.Mutex
	lockName string
}

// Unlock 解锁
func (s *onceLock) Unlock(ctx context.Context) (ok bool, err error) {
	if ctx == nil {
		return false, errors.New("redis unlock context is nil")
	}
	if s.mutex == nil {
		return false, errors.New("redis mutex is nil")
	}
	return s.mutex.UnlockContext(ctx)
}

func (s *onceLock) Name() string {
	return s.lockName
}
