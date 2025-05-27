package lockerpkg

import (
	"context"
	stderrors "errors"
	"sync"
	"time"
)

var _ LocalLocker = (*local)(nil)

type local struct {
	sm sync.Map
}

// NewLocalLocker ...
func NewLocalLocker() Locker {
	return &local{}
}

func (s *local) Mutex(ctx context.Context, lockName string) (Unlocker, error) {
	locker := newLocalLock(&s.sm, &sync.Mutex{}, lockName)
	lockerInterface, _ := s.sm.LoadOrStore(lockName, locker)
	locker = lockerInterface.(*localLock)
	if !locker.mu.TryLock() {
		err := ErrorLockerFailed(lockName, stderrors.New("try lock failed"))
		return locker, err
	}
	return locker, nil
}

func (s *local) Once(ctx context.Context, lockName string) (Unlocker, error) {
	locker, err := s.Mutex(ctx, lockName)
	if err != nil {
		return locker, err
	}
	defer func() {
		time.AfterFunc(_cacheLockerExpire, func() {
			_, _ = locker.Unlock(ctx)
		})
	}()
	return locker, nil
}

func (s *local) Unlock(ctx context.Context, lockName string) {
	lockerInterface, ok := s.sm.LoadAndDelete(lockName)
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
	sm       *sync.Map
	mu       *sync.Mutex
	lockName string
}

func newLocalLock(sm *sync.Map, mu *sync.Mutex, lockName string) *localLock {
	return &localLock{
		sm:       sm,
		mu:       mu,
		lockName: lockName,
	}
}

// Unlock ...
func (s *localLock) Unlock(ctx context.Context) (bool, error) {
	_ = s.mu.TryLock()
	s.mu.Unlock()
	s.sm.Delete(s.lockName)
	return true, nil
}

func (s *localLock) Name() string {
	return s.lockName
}
