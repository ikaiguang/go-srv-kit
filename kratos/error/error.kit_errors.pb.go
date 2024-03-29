// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errorpkg

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// UNKNOWN 常规
func IsUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNKNOWN.String() && e.Code == 500
}

// UNKNOWN 常规
func ErrorUnknown(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_UNKNOWN.String(), fmt.Sprintf(format, args...))
}

func IsRequestFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_FAILED.String() && e.Code == 400
}

func ErrorRequestFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ERROR_REQUEST_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsRecordNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_RECORD_NOT_FOUND.String() && e.Code == 404
}

func ErrorRecordNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ERROR_RECORD_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsRecordAlreadyExists(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_RECORD_ALREADY_EXISTS.String() && e.Code == 400
}

func ErrorRecordAlreadyExists(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ERROR_RECORD_ALREADY_EXISTS.String(), fmt.Sprintf(format, args...))
}

func IsNetworkError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NETWORK_ERROR.String() && e.Code == 500
}

func ErrorNetworkError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_NETWORK_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsNetworkTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NETWORK_TIMEOUT.String() && e.Code == 504
}

func ErrorNetworkTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, ERROR_NETWORK_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

func IsConnection(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CONNECTION.String() && e.Code == 500
}

func ErrorConnection(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_CONNECTION.String(), fmt.Sprintf(format, args...))
}

func IsUninitialized(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNINITIALIZED.String() && e.Code == 500
}

func ErrorUninitialized(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_UNINITIALIZED.String(), fmt.Sprintf(format, args...))
}

func IsUnimplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNIMPLEMENTED.String() && e.Code == 500
}

func ErrorUnimplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_UNIMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

func IsInvalidParameter(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_INVALID_PARAMETER.String() && e.Code == 400
}

func ErrorInvalidParameter(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ERROR_INVALID_PARAMETER.String(), fmt.Sprintf(format, args...))
}

func IsRequestNotSupport(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_NOT_SUPPORT.String() && e.Code == 500
}

func ErrorRequestNotSupport(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_REQUEST_NOT_SUPPORT.String(), fmt.Sprintf(format, args...))
}

// 第三方服务错误
func IsThirdPartyServiceInternalError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_THIRD_PARTY_SERVICE_INTERNAL_ERROR.String() && e.Code == 500
}

// 第三方服务错误
func ErrorThirdPartyServiceInternalError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_THIRD_PARTY_SERVICE_INTERNAL_ERROR.String(), fmt.Sprintf(format, args...))
}

// 第三方服务响应结果有误
func IsThirdPartyServiceInvalidCode(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_THIRD_PARTY_SERVICE_INVALID_CODE.String() && e.Code == 400
}

// 第三方服务响应结果有误
func ErrorThirdPartyServiceInvalidCode(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ERROR_THIRD_PARTY_SERVICE_INVALID_CODE.String(), fmt.Sprintf(format, args...))
}

func IsDb(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_DB.String() && e.Code == 500
}

func ErrorDb(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_DB.String(), fmt.Sprintf(format, args...))
}

func IsMysql(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MYSQL.String() && e.Code == 500
}

func ErrorMysql(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_MYSQL.String(), fmt.Sprintf(format, args...))
}

func IsMongo(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MONGO.String() && e.Code == 500
}

func ErrorMongo(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_MONGO.String(), fmt.Sprintf(format, args...))
}

func IsCache(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CACHE.String() && e.Code == 500
}

func ErrorCache(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_CACHE.String(), fmt.Sprintf(format, args...))
}

func IsRedis(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REDIS.String() && e.Code == 500
}

func ErrorRedis(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_REDIS.String(), fmt.Sprintf(format, args...))
}

func IsMq(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MQ.String() && e.Code == 500
}

func ErrorMq(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_MQ.String(), fmt.Sprintf(format, args...))
}

func IsRabbitMq(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_RABBIT_MQ.String() && e.Code == 500
}

func ErrorRabbitMq(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_RABBIT_MQ.String(), fmt.Sprintf(format, args...))
}

func IsKafka(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_KAFKA.String() && e.Code == 500
}

func ErrorKafka(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_KAFKA.String(), fmt.Sprintf(format, args...))
}

func IsPanic(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PANIC.String() && e.Code == 500
}

func ErrorPanic(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_PANIC.String(), fmt.Sprintf(format, args...))
}

func IsFatal(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_FATAL.String() && e.Code == 500
}

func ErrorFatal(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_FATAL.String(), fmt.Sprintf(format, args...))
}

// CONTINUE Continue
func IsContinue(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CONTINUE.String() && e.Code == 100
}

// CONTINUE Continue
func ErrorContinue(format string, args ...interface{}) *errors.Error {
	return errors.New(100, ERROR_CONTINUE.String(), fmt.Sprintf(format, args...))
}

func IsSwitchingProtocols(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_SWITCHING_PROTOCOLS.String() && e.Code == 101
}

func ErrorSwitchingProtocols(format string, args ...interface{}) *errors.Error {
	return errors.New(101, ERROR_SWITCHING_PROTOCOLS.String(), fmt.Sprintf(format, args...))
}

func IsProcessing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PROCESSING.String() && e.Code == 102
}

func ErrorProcessing(format string, args ...interface{}) *errors.Error {
	return errors.New(102, ERROR_PROCESSING.String(), fmt.Sprintf(format, args...))
}

func IsEarlyHints(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_EARLY_HINTS.String() && e.Code == 103
}

func ErrorEarlyHints(format string, args ...interface{}) *errors.Error {
	return errors.New(103, ERROR_EARLY_HINTS.String(), fmt.Sprintf(format, args...))
}

// OK OK
func IsOk(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_OK.String() && e.Code == 200
}

// OK OK
func ErrorOk(format string, args ...interface{}) *errors.Error {
	return errors.New(200, ERROR_OK.String(), fmt.Sprintf(format, args...))
}

func IsCreated(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CREATED.String() && e.Code == 201
}

func ErrorCreated(format string, args ...interface{}) *errors.Error {
	return errors.New(201, ERROR_CREATED.String(), fmt.Sprintf(format, args...))
}

func IsAccepted(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_ACCEPTED.String() && e.Code == 202
}

func ErrorAccepted(format string, args ...interface{}) *errors.Error {
	return errors.New(202, ERROR_ACCEPTED.String(), fmt.Sprintf(format, args...))
}

func IsNonAuthoritativeInfo(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NON_AUTHORITATIVE_INFO.String() && e.Code == 203
}

func ErrorNonAuthoritativeInfo(format string, args ...interface{}) *errors.Error {
	return errors.New(203, ERROR_NON_AUTHORITATIVE_INFO.String(), fmt.Sprintf(format, args...))
}

func IsNoContent(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NO_CONTENT.String() && e.Code == 204
}

func ErrorNoContent(format string, args ...interface{}) *errors.Error {
	return errors.New(204, ERROR_NO_CONTENT.String(), fmt.Sprintf(format, args...))
}

func IsResetContent(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_RESET_CONTENT.String() && e.Code == 205
}

func ErrorResetContent(format string, args ...interface{}) *errors.Error {
	return errors.New(205, ERROR_RESET_CONTENT.String(), fmt.Sprintf(format, args...))
}

func IsPartialContent(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PARTIAL_CONTENT.String() && e.Code == 206
}

func ErrorPartialContent(format string, args ...interface{}) *errors.Error {
	return errors.New(206, ERROR_PARTIAL_CONTENT.String(), fmt.Sprintf(format, args...))
}

func IsMultiStatus(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MULTI_STATUS.String() && e.Code == 207
}

func ErrorMultiStatus(format string, args ...interface{}) *errors.Error {
	return errors.New(207, ERROR_MULTI_STATUS.String(), fmt.Sprintf(format, args...))
}

func IsAlreadyReported(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_ALREADY_REPORTED.String() && e.Code == 208
}

func ErrorAlreadyReported(format string, args ...interface{}) *errors.Error {
	return errors.New(208, ERROR_ALREADY_REPORTED.String(), fmt.Sprintf(format, args...))
}

func IsIMUsed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_I_M_USED.String() && e.Code == 226
}

func ErrorIMUsed(format string, args ...interface{}) *errors.Error {
	return errors.New(226, ERROR_I_M_USED.String(), fmt.Sprintf(format, args...))
}

// MULTIPLE_CHOICES MultipleChoices
func IsMultipleChoices(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MULTIPLE_CHOICES.String() && e.Code == 300
}

// MULTIPLE_CHOICES MultipleChoices
func ErrorMultipleChoices(format string, args ...interface{}) *errors.Error {
	return errors.New(300, ERROR_MULTIPLE_CHOICES.String(), fmt.Sprintf(format, args...))
}

func IsMovedPermanently(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MOVED_PERMANENTLY.String() && e.Code == 301
}

func ErrorMovedPermanently(format string, args ...interface{}) *errors.Error {
	return errors.New(301, ERROR_MOVED_PERMANENTLY.String(), fmt.Sprintf(format, args...))
}

func IsFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_FOUND.String() && e.Code == 302
}

func ErrorFound(format string, args ...interface{}) *errors.Error {
	return errors.New(302, ERROR_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsSeeOther(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_SEE_OTHER.String() && e.Code == 303
}

func ErrorSeeOther(format string, args ...interface{}) *errors.Error {
	return errors.New(303, ERROR_SEE_OTHER.String(), fmt.Sprintf(format, args...))
}

func IsNotModified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_MODIFIED.String() && e.Code == 304
}

func ErrorNotModified(format string, args ...interface{}) *errors.Error {
	return errors.New(304, ERROR_NOT_MODIFIED.String(), fmt.Sprintf(format, args...))
}

func IsUseProxy(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_USE_PROXY.String() && e.Code == 305
}

func ErrorUseProxy(format string, args ...interface{}) *errors.Error {
	return errors.New(305, ERROR_USE_PROXY.String(), fmt.Sprintf(format, args...))
}

func IsEmpty306(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_EMPTY306.String() && e.Code == 306
}

func ErrorEmpty306(format string, args ...interface{}) *errors.Error {
	return errors.New(306, ERROR_EMPTY306.String(), fmt.Sprintf(format, args...))
}

func IsTemporaryRedirect(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_TEMPORARY_REDIRECT.String() && e.Code == 307
}

func ErrorTemporaryRedirect(format string, args ...interface{}) *errors.Error {
	return errors.New(307, ERROR_TEMPORARY_REDIRECT.String(), fmt.Sprintf(format, args...))
}

func IsPermanentRedirect(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PERMANENT_REDIRECT.String() && e.Code == 308
}

func ErrorPermanentRedirect(format string, args ...interface{}) *errors.Error {
	return errors.New(308, ERROR_PERMANENT_REDIRECT.String(), fmt.Sprintf(format, args...))
}

// BAD_REQUEST Bad Request
func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_BAD_REQUEST.String() && e.Code == 400
}

// BAD_REQUEST Bad Request
func ErrorBadRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ERROR_BAD_REQUEST.String(), fmt.Sprintf(format, args...))
}

func IsUnauthorized(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNAUTHORIZED.String() && e.Code == 401
}

func ErrorUnauthorized(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ERROR_UNAUTHORIZED.String(), fmt.Sprintf(format, args...))
}

func IsPaymentRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PAYMENT_REQUIRED.String() && e.Code == 402
}

func ErrorPaymentRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(402, ERROR_PAYMENT_REQUIRED.String(), fmt.Sprintf(format, args...))
}

func IsForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_FORBIDDEN.String() && e.Code == 403
}

func ErrorForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ERROR_FORBIDDEN.String(), fmt.Sprintf(format, args...))
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_FOUND.String() && e.Code == 404
}

func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ERROR_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_METHOD_NOT_ALLOWED.String() && e.Code == 405
}

func ErrorMethodNotAllowed(format string, args ...interface{}) *errors.Error {
	return errors.New(405, ERROR_METHOD_NOT_ALLOWED.String(), fmt.Sprintf(format, args...))
}

func IsNotAcceptable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_ACCEPTABLE.String() && e.Code == 406
}

func ErrorNotAcceptable(format string, args ...interface{}) *errors.Error {
	return errors.New(406, ERROR_NOT_ACCEPTABLE.String(), fmt.Sprintf(format, args...))
}

func IsProxyAuthRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PROXY_AUTH_REQUIRED.String() && e.Code == 407
}

func ErrorProxyAuthRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(407, ERROR_PROXY_AUTH_REQUIRED.String(), fmt.Sprintf(format, args...))
}

func IsRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_TIMEOUT.String() && e.Code == 408
}

func ErrorRequestTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(408, ERROR_REQUEST_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

func IsConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CONFLICT.String() && e.Code == 409
}

func ErrorConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, ERROR_CONFLICT.String(), fmt.Sprintf(format, args...))
}

func IsGone(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_GONE.String() && e.Code == 410
}

func ErrorGone(format string, args ...interface{}) *errors.Error {
	return errors.New(410, ERROR_GONE.String(), fmt.Sprintf(format, args...))
}

func IsLengthRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_LENGTH_REQUIRED.String() && e.Code == 411
}

func ErrorLengthRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(411, ERROR_LENGTH_REQUIRED.String(), fmt.Sprintf(format, args...))
}

func IsPreconditionFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PRECONDITION_FAILED.String() && e.Code == 412
}

func ErrorPreconditionFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(412, ERROR_PRECONDITION_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsRequestEntityTooLarge(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_ENTITY_TOO_LARGE.String() && e.Code == 413
}

func ErrorRequestEntityTooLarge(format string, args ...interface{}) *errors.Error {
	return errors.New(413, ERROR_REQUEST_ENTITY_TOO_LARGE.String(), fmt.Sprintf(format, args...))
}

func IsRequestUriTooLong(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_URI_TOO_LONG.String() && e.Code == 414
}

func ErrorRequestUriTooLong(format string, args ...interface{}) *errors.Error {
	return errors.New(414, ERROR_REQUEST_URI_TOO_LONG.String(), fmt.Sprintf(format, args...))
}

func IsUnsupportedMediaType(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNSUPPORTED_MEDIA_TYPE.String() && e.Code == 415
}

func ErrorUnsupportedMediaType(format string, args ...interface{}) *errors.Error {
	return errors.New(415, ERROR_UNSUPPORTED_MEDIA_TYPE.String(), fmt.Sprintf(format, args...))
}

func IsRequestedRangeNotSatisfiable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUESTED_RANGE_NOT_SATISFIABLE.String() && e.Code == 416
}

func ErrorRequestedRangeNotSatisfiable(format string, args ...interface{}) *errors.Error {
	return errors.New(416, ERROR_REQUESTED_RANGE_NOT_SATISFIABLE.String(), fmt.Sprintf(format, args...))
}

func IsExpectationFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_EXPECTATION_FAILED.String() && e.Code == 417
}

func ErrorExpectationFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(417, ERROR_EXPECTATION_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsTeapot(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_TEAPOT.String() && e.Code == 418
}

func ErrorTeapot(format string, args ...interface{}) *errors.Error {
	return errors.New(418, ERROR_TEAPOT.String(), fmt.Sprintf(format, args...))
}

func IsMisdirectedRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MISDIRECTED_REQUEST.String() && e.Code == 421
}

func ErrorMisdirectedRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(421, ERROR_MISDIRECTED_REQUEST.String(), fmt.Sprintf(format, args...))
}

func IsUnprocessableEntity(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNPROCESSABLE_ENTITY.String() && e.Code == 422
}

func ErrorUnprocessableEntity(format string, args ...interface{}) *errors.Error {
	return errors.New(422, ERROR_UNPROCESSABLE_ENTITY.String(), fmt.Sprintf(format, args...))
}

func IsLocked(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_LOCKED.String() && e.Code == 423
}

func ErrorLocked(format string, args ...interface{}) *errors.Error {
	return errors.New(423, ERROR_LOCKED.String(), fmt.Sprintf(format, args...))
}

func IsFailedDependency(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_FAILED_DEPENDENCY.String() && e.Code == 424
}

func ErrorFailedDependency(format string, args ...interface{}) *errors.Error {
	return errors.New(424, ERROR_FAILED_DEPENDENCY.String(), fmt.Sprintf(format, args...))
}

func IsTooEarly(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_TOO_EARLY.String() && e.Code == 425
}

func ErrorTooEarly(format string, args ...interface{}) *errors.Error {
	return errors.New(425, ERROR_TOO_EARLY.String(), fmt.Sprintf(format, args...))
}

func IsUpgradeRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UPGRADE_REQUIRED.String() && e.Code == 426
}

func ErrorUpgradeRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(426, ERROR_UPGRADE_REQUIRED.String(), fmt.Sprintf(format, args...))
}

func IsPreconditionRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PRECONDITION_REQUIRED.String() && e.Code == 428
}

func ErrorPreconditionRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(428, ERROR_PRECONDITION_REQUIRED.String(), fmt.Sprintf(format, args...))
}

func IsTooManyRequests(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_TOO_MANY_REQUESTS.String() && e.Code == 429
}

func ErrorTooManyRequests(format string, args ...interface{}) *errors.Error {
	return errors.New(429, ERROR_TOO_MANY_REQUESTS.String(), fmt.Sprintf(format, args...))
}

func IsRequestHeaderFieldsTooLarge(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_HEADER_FIELDS_TOO_LARGE.String() && e.Code == 431
}

func ErrorRequestHeaderFieldsTooLarge(format string, args ...interface{}) *errors.Error {
	return errors.New(431, ERROR_REQUEST_HEADER_FIELDS_TOO_LARGE.String(), fmt.Sprintf(format, args...))
}

func IsUnavailableForLegalReasons(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNAVAILABLE_FOR_LEGAL_REASONS.String() && e.Code == 451
}

func ErrorUnavailableForLegalReasons(format string, args ...interface{}) *errors.Error {
	return errors.New(451, ERROR_UNAVAILABLE_FOR_LEGAL_REASONS.String(), fmt.Sprintf(format, args...))
}

func IsClientClose(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CLIENT_CLOSE.String() && e.Code == 499
}

func ErrorClientClose(format string, args ...interface{}) *errors.Error {
	return errors.New(499, ERROR_CLIENT_CLOSE.String(), fmt.Sprintf(format, args...))
}

// INTERNAL_SERVER Internal Server Error
func IsInternalServer(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_INTERNAL_SERVER.String() && e.Code == 500
}

// INTERNAL_SERVER Internal Server Error
func ErrorInternalServer(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ERROR_INTERNAL_SERVER.String(), fmt.Sprintf(format, args...))
}

func IsNotImplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_IMPLEMENTED.String() && e.Code == 501
}

func ErrorNotImplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(501, ERROR_NOT_IMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

func IsBadGateway(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_BAD_GATEWAY.String() && e.Code == 502
}

func ErrorBadGateway(format string, args ...interface{}) *errors.Error {
	return errors.New(502, ERROR_BAD_GATEWAY.String(), fmt.Sprintf(format, args...))
}

func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_SERVICE_UNAVAILABLE.String() && e.Code == 503
}

func ErrorServiceUnavailable(format string, args ...interface{}) *errors.Error {
	return errors.New(503, ERROR_SERVICE_UNAVAILABLE.String(), fmt.Sprintf(format, args...))
}

func IsGatewayTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_GATEWAY_TIMEOUT.String() && e.Code == 504
}

func ErrorGatewayTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, ERROR_GATEWAY_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

func IsHttpVersionNotSupported(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_HTTP_VERSION_NOT_SUPPORTED.String() && e.Code == 505
}

func ErrorHttpVersionNotSupported(format string, args ...interface{}) *errors.Error {
	return errors.New(505, ERROR_HTTP_VERSION_NOT_SUPPORTED.String(), fmt.Sprintf(format, args...))
}

func IsVariantAlsoNegotiates(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_VARIANT_ALSO_NEGOTIATES.String() && e.Code == 506
}

func ErrorVariantAlsoNegotiates(format string, args ...interface{}) *errors.Error {
	return errors.New(506, ERROR_VARIANT_ALSO_NEGOTIATES.String(), fmt.Sprintf(format, args...))
}

func IsInsufficientStorage(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_INSUFFICIENT_STORAGE.String() && e.Code == 507
}

func ErrorInsufficientStorage(format string, args ...interface{}) *errors.Error {
	return errors.New(507, ERROR_INSUFFICIENT_STORAGE.String(), fmt.Sprintf(format, args...))
}

func IsLoopDetected(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_LOOP_DETECTED.String() && e.Code == 508
}

func ErrorLoopDetected(format string, args ...interface{}) *errors.Error {
	return errors.New(508, ERROR_LOOP_DETECTED.String(), fmt.Sprintf(format, args...))
}

func IsNotExtended(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_EXTENDED.String() && e.Code == 510
}

func ErrorNotExtended(format string, args ...interface{}) *errors.Error {
	return errors.New(510, ERROR_NOT_EXTENDED.String(), fmt.Sprintf(format, args...))
}

func IsNetworkAuthenticationRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NETWORK_AUTHENTICATION_REQUIRED.String() && e.Code == 511
}

func ErrorNetworkAuthenticationRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(511, ERROR_NETWORK_AUTHENTICATION_REQUIRED.String(), fmt.Sprintf(format, args...))
}
