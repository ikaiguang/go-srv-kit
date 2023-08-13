package lockerpkg

import (
	"context"
	stderrors "errors"
	"sync"
)

var _ LocalLocker = (*local)(nil)

type local struct{ sm sync.Map }

// NewLocalLocker ...
func NewLocalLocker() Lock {
	return &local{}
}

func (s *local) Mutex(ctx context.Context, lockName string) (Unlock, error) {
	locker := newLocalLock(&sync.Mutex{})
	lockerInterface, _ := s.sm.LoadOrStore(lockName, locker)
	locker = lockerInterface.(*localLock)
	if !locker.mu.TryLock() {
		err := ErrorLockerFailed(lockName, stderrors.New("try lock failed"))
		return locker, err
	}
	return locker, nil
}

func (s *local) Once(ctx context.Context, lockName string) (Unlock, error) {
	return s.Mutex(ctx, lockName)
}

func (s *local) Unlock(ctx context.Context, lockName string) {
	lockerInterface, ok := s.sm.Load(lockName)
	if !ok {
		return
	}
	locker := lockerInterface.(*localLock)
	_ = locker.mu.TryLock()
	_, _ = locker.Unlock(ctx)
	return
}

// localLock ...
type localLock struct {
	mu *sync.Mutex
}

func newLocalLock(mu *sync.Mutex) *localLock {
	return &localLock{mu: mu}
}

// Unlock ...
func (s *localLock) Unlock(ctx context.Context) (bool, error) {
	_ = s.mu.TryLock()
	s.mu.Unlock()
	return true, nil
}
