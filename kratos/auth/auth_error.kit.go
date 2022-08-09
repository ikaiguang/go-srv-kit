package authutil

import (
	"github.com/go-kratos/kratos/v2/errors"

	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "Can not get key while signing token")

	ErrLoginLimit         = errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), "ValidateFunc : Token is invalid")
	ErrInvalidRedisKey    = errors.BadRequest(errorv1.ERROR_BAD_REQUEST.String(), "RedisSecretFunc : invalid redis key")
	ErrInvalidAuthInfo    = errors.BadRequest(errorv1.ERROR_BAD_REQUEST.String(), "ValidateFunc : invalid auth info")
	ErrGetRedisData       = errors.BadRequest(errorv1.ERROR_EXPECTATION_FAILED.String(), "RedisSecretFunc : get redis data failed")
	ErrSetRedisData       = errors.BadRequest(errorv1.ERROR_EXPECTATION_FAILED.String(), "RedisSecretFunc : set redis data failed")
	ErrUnmarshalRedisData = errors.BadRequest(errorv1.ERROR_EXPECTATION_FAILED.String(), "RedisSecretFunc : unmarshal redis data failed")
	ErrMarshalRedisData   = errors.BadRequest(errorv1.ERROR_EXPECTATION_FAILED.String(), "RedisSecretFunc : marshal redis data failed")
)

// Is ...
func Is(err, target error) bool {
	return errorutil.Is(err, target)
}
