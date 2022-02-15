// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package exception

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUNKNOWN(err error) bool {
	e := errors.FromError(err)
	return e.Reason == Error_UNKNOWN.String() && e.Code == 404
}

func ErrorUNKNOWN(format string, args ...interface{}) *errors.Error {
	return errors.New(404, Error_UNKNOWN.String(), fmt.Sprintf(format, args...))
}

func IsContentMissing(err error) bool {
	e := errors.FromError(err)
	return e.Reason == Error_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, Error_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}

func IsContentError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == Error_CONTENT_ERROR.String() && e.Code == 400
}

func ErrorContentError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, Error_CONTENT_ERROR.String(), fmt.Sprintf(format, args...))
}
