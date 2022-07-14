package lockerutil

import (
	"context"
	"fmt"
)

// Locker .
type Locker interface {
	Lock(ctx context.Context, name string) (err error)
	Unlock(ctx context.Context) (ok bool, err error)
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

// IsLockFailedError 锁失败
func IsLockFailedError(err error) bool {
	if e, ok := err.(*ErrLockerFailed); ok {
		return e.isLockFailed
	}
	return false
}

// IsExtendFailedError 延长锁失败
func IsExtendFailedError(err error) bool {
	if e, ok := err.(*ErrExtendFailed); ok {
		return e.isExtendFailed
	}
	return false
}
