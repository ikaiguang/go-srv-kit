package errorpkg

import (
	"runtime"

	"github.com/go-kratos/kratos/v2/errors"
)

// FromError try to convert an error to *errors.Error
func FromError(err error) *errors.Error {
	return errors.FromError(Cause(err))
}

// Cause returns the underlying cause of the error
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

// Is ...
func Is(err, target error) bool {
	if err == nil {
		return false
	}
	errCause := FromError(err)
	targetCause := FromError(target)

	return errors.Is(errCause, targetCause)
}

// IsCode .
func IsCode(err error, code int) bool {
	return Code(err) == code
}

// IsReason .
func IsReason(err error, reason string) bool {
	return Reason(err) == reason
}

// Code returns the http code for a error.
func Code(err error) int {
	if err == nil {
		return 200
	}
	if se := FromError(err); se != nil {
		return int(se.Code)
	}
	return errors.UnknownCode
}

// Reason returns the reason for a particular error.
func Reason(err error) string {
	if err == nil {
		return errors.UnknownReason
	}
	if se := FromError(err); se != nil {
		return se.Reason
	}
	return errors.UnknownReason
}

// Metadata returns the metadata for a particular error.
func Metadata(err error) (metadata map[string]string, ok bool) {
	if err == nil {
		return metadata, ok
	}
	if se := FromError(err); se != nil && se.Metadata != nil {
		metadata = se.Metadata
		ok = true
		return metadata, ok
	}
	return metadata, ok
}

// Message returns the message for a particular error.
func Message(err error) string {
	if err == nil {
		return ""
	}
	if se := FromError(err); se != nil {
		return se.Message
	}
	return ""
}

// RecoverStack ...
func RecoverStack() string {
	buf := make([]byte, 64<<10) //nolint:gomnd
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
