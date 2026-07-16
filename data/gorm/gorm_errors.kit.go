package gorm

import (
	stderrors "errors"

	"gorm.io/gorm"
)

// IsErrRecordNotFound ...
func IsErrRecordNotFound(err error) bool {
	return stderrors.Is(err, gorm.ErrRecordNotFound)
}

// IsErrDuplicatedKey ...
func IsErrDuplicatedKey(err error) bool {
	if err == nil {
		return false
	}
	if stderrors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	return false
}
