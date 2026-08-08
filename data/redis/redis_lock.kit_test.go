package redispkg

import (
	"context"
	"testing"
	"time"

	lockerpkg "github.com/ikaiguang/go-srv-kit/kit/v3/locker"
)

func TestLockerOnce(t *testing.T) {
	config := testRedisConfig(t)
	db, err := NewDB(config)
	if err != nil {
		t.Fatalf("NewDB() error = %v", err)
	}
	defer db.Close()

	locker := NewLocker(db)
	first, err := locker.Once(context.Background(), "once-lock")
	if err != nil {
		t.Fatalf("Once() error = %v", err)
	}
	if _, err = locker.Once(context.Background(), "once-lock"); !lockerpkg.IsErrorLockFailed(err) {
		t.Fatalf("second Once() error = %v, want lock failed", err)
	}
	if ok, unlockErr := first.Unlock(context.Background()); unlockErr != nil || !ok {
		t.Fatalf("Unlock() = (%v, %v), want (true, nil)", ok, unlockErr)
	}
}

func TestLockerMutexCanBeReacquired(t *testing.T) {
	config := testRedisConfig(t)
	db, err := NewDB(config)
	if err != nil {
		t.Fatalf("NewDB() error = %v", err)
	}
	defer db.Close()

	locker := NewLocker(db)
	first, err := locker.Mutex(context.Background(), "mutex-lock")
	if err != nil {
		t.Fatalf("Mutex() error = %v", err)
	}
	if _, err = locker.Mutex(context.Background(), "mutex-lock"); !lockerpkg.IsErrorLockFailed(err) {
		t.Fatalf("second Mutex() error = %v, want lock failed", err)
	}
	if ok, unlockErr := first.Unlock(context.Background()); unlockErr != nil || !ok {
		t.Fatalf("Unlock() = (%v, %v), want (true, nil)", ok, unlockErr)
	}
	if second, reacquireErr := locker.Mutex(context.Background(), "mutex-lock"); reacquireErr != nil {
		t.Fatalf("reacquire Mutex() error = %v", reacquireErr)
	} else {
		_, _ = second.Unlock(context.Background())
	}
}

func TestMutexLockStopExtendingIsIdempotent(t *testing.T) {
	lock := &mutexLock{stopExtend: make(chan struct{}), extendDone: make(chan struct{})}
	done := make(chan struct{})
	go func() {
		<-lock.stopExtend
		close(lock.extendDone)
		close(done)
	}()

	if err := lock.stopExtending(context.Background()); err != nil {
		t.Fatalf("stopExtending() error = %v", err)
	}
	if err := lock.stopExtending(context.Background()); err != nil {
		t.Fatalf("second stopExtending() error = %v", err)
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("stopExtending() did not close the stop channel")
	}
}

func TestNewLockerRejectsNilClient(t *testing.T) {
	locker := NewLocker(nil)
	if _, err := locker.Once(context.Background(), "nil-client"); !lockerpkg.IsErrorLockFailed(err) {
		t.Fatalf("Once() error = %v, want lock failed", err)
	}
	if _, err := locker.Mutex(nil, "nil-context"); !lockerpkg.IsErrorLockFailed(err) {
		t.Fatalf("Mutex(nil) error = %v, want lock failed", err)
	}
}
