package redispkg

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/go-redsync/redsync/v4"
	lockerpkg "github.com/ikaiguang/go-srv-kit/kit/v3/locker"
)

// mutexLock ...
type mutexLock struct {
	mutex       *redsync.Mutex
	stopExtend  chan struct{}
	extendDone  chan struct{}
	stopOnce    sync.Once
	extendErrMu sync.Mutex
	extendErr   error
	lockName    string
}

// Unlock ...
func (s *mutexLock) Unlock(ctx context.Context) (ok bool, err error) {
	if err = s.stopExtending(ctx); err != nil {
		return false, err
	}
	if s.mutex == nil {
		return false, errors.New("redis mutex is nil")
	}

	// 取锁的有效期。在获取锁之前，该值将为零值。
	if s.mutex.Until().IsZero() {
		return ok, err
	}
	ok, err = s.mutex.UnlockContext(ctx)
	if extendErr := s.getExtendError(); extendErr != nil {
		if err == nil {
			return ok, extendErr
		}
		return ok, errors.Join(extendErr, err)
	}
	return ok, err
}

func (s *mutexLock) Name() string {
	return s.lockName
}

// extend ...
func (s *mutexLock) extend(ctx context.Context) {
	defer close(s.extendDone)
	ticker := time.NewTicker(_lockExtendDelay)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if ok, err := s.mutex.ExtendContext(ctx); err != nil || !ok {
				if err == nil {
					err = redsync.ErrExtendFailed
				}
				s.setExtendError(lockerpkg.ErrorExtendFailed(s.mutex.Name(), err))
				return
			}
		case <-ctx.Done():
			return
		case <-s.stopExtend:
			return
		}
	}
}

func (s *mutexLock) stopExtending(ctx context.Context) error {
	if ctx == nil {
		return errors.New("redis unlock context is nil")
	}
	if s.stopExtend == nil {
		return nil
	}
	s.stopOnce.Do(func() {
		close(s.stopExtend)
	})
	if s.extendDone == nil {
		return nil
	}
	select {
	case <-s.extendDone:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *mutexLock) setExtendError(err error) {
	s.extendErrMu.Lock()
	defer s.extendErrMu.Unlock()
	s.extendErr = err
}

func (s *mutexLock) getExtendError() error {
	s.extendErrMu.Lock()
	defer s.extendErrMu.Unlock()
	return s.extendErr
}
