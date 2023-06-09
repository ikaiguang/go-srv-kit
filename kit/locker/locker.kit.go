package lockerpkg

import (
	"context"
	"fmt"
)

// Unlock 解锁
type Unlock interface {
	Unlock(ctx context.Context) (ok bool, err error)
}

// Lock 加锁
type Lock interface {
	Mutex(ctx context.Context, lockName string) (locker Unlock, err error)
	Once(ctx context.Context, lockName string) (locker Unlock, err error)
}

// NewErrLockerFailed .
func NewErrLockerFailed(isLockFailed bool, name string, err error) error {
	return &ErrLockerFailed{
		isLockFailed: isLockFailed,
		name:         name,
		err:          err,
	}
}

// ErrLockerFailed ...
type ErrLockerFailed struct {
	isLockFailed bool
	name         string
	err          error
}

// Error error
func (e *ErrLockerFailed) Error() string {
	return fmt.Sprintf("Lock(%s) failed : %s", e.name, e.err.Error())
}

// NewErrExtendFailed .
func NewErrExtendFailed(isExtendFailed bool, name string, err error) error {
	return &ErrExtendFailed{
		isExtendFailed: isExtendFailed,
		name:           name,
		err:            err,
	}
}

// ErrExtendFailed ...
type ErrExtendFailed struct {
	isExtendFailed bool
	name           string
	err            error
}

// Error error
func (e *ErrExtendFailed) Error() string {
	return fmt.Sprintf("Lock(%s) failed : %s", e.name, e.err.Error())
}

// IsErrLockFailed 锁失败
func IsErrLockFailed(err error) bool {
	if e, ok := err.(*ErrLockerFailed); ok {
		return e.isLockFailed
	}
	return false
}

// IsErrExtendFailed 延长锁失败
func IsErrExtendFailed(err error) bool {
	if e, ok := err.(*ErrExtendFailed); ok {
		return e.isExtendFailed
	}
	return false
}
