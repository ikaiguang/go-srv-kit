package lockerpkg

import (
	"context"
	"fmt"
)

type LocalLocker interface {
	Lock

	Unlock(ctx context.Context, lockName string)
}

// Unlock 解锁
type Unlock interface {
	Unlock(ctx context.Context) (ok bool, err error)
}

// Lock 加锁
type Lock interface {
	// Mutex 互斥锁，一直等待直到解锁
	Mutex(ctx context.Context, lockName string) (Unlock, error)
	// Once 简单锁，等待解锁或者锁定时间过期后自动解锁
	Once(ctx context.Context, lockName string) (Unlock, error)
}

// ErrorLockerFailed .
func ErrorLockerFailed(name string, err error) error {
	return &errLockerFailed{
		name: name,
		err:  err,
	}
}

// errLockerFailed ...
type errLockerFailed struct {
	name string
	err  error
}

// Error error
func (e *errLockerFailed) Error() string {
	return fmt.Sprintf("Lock(%s) failed : %s", e.name, e.err.Error())
}

// IsErrorLockFailed 锁失败
func IsErrorLockFailed(err error) bool {
	_, ok := err.(*errLockerFailed)
	return ok
}

// ErrorExtendFailed .
func ErrorExtendFailed(name string, err error) error {
	return &errExtendFailed{
		name: name,
		err:  err,
	}
}

// errExtendFailed ...
type errExtendFailed struct {
	name string
	err  error
}

// Error error
func (e *errExtendFailed) Error() string {
	return fmt.Sprintf("Lock(%s) failed : %s", e.name, e.err.Error())
}

// IsErrorExtendFailed 延长锁失败
func IsErrorExtendFailed(err error) bool {
	_, ok := err.(*errExtendFailed)

	return ok
}
