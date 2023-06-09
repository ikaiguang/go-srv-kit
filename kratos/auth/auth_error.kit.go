package authpkg

import (
	"github.com/go-kratos/kratos/v2/errors"

	errorpkg "gitlab.realibox.cn/designhub/app-server/hub-kratos-pkg/error"
)

func ErrMissingToken() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_MISSING.String(), "token is missing")
}
func ErrMissingSignKeyFunc() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_KEY_MISSING.String(), "keyFunc is missing")
}
func ErrUnSupportSigningMethod() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_METHOD_MISSING.String(), "Wrong signing method")
}
func ErrTokenParseFail() *errors.Error {
	return errors.Unauthorized(ERROR_AUTHENTICATION_FAILED.String(), "Fail to parse JWT token ")
}
func ErrTokenExpired() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_EXPIRED.String(), "JWT token has expired")
}
func ErrWrongContext() *errors.Error {
	return errors.Unauthorized(ERROR_UNAUTHORIZED.String(), "Wrong context for middleware")
}
func ErrTokenInvalid() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_INVALID.String(), "Token is invalid")
}
func ErrNeedTokenProvider() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_INVALID.String(), "Token provider is missing")
}
func ErrSignToken() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_INVALID.String(), "Can not sign token.Is the key correct?")
}
func ErrGetKey() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_INVALID.String(), "Can not get key while signing token")
}
func ErrInvalidAuthToken() *errors.Error {
	return errors.Unauthorized(ERROR_VERIFICATION_FAILED.String(), "[validator] invalid auth token")
}
func ErrInvalidClaims() *errors.Error {
	return errors.Unauthorized(ERROR_INVALID_CLAIMS.String(), "[validator] get redis data failed")
}
func ErrBlacklist() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_DEPRECATED.String(), "[validator] deprecated token")
}
func ErrWhitelist() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_INVALID.String(), "[validator] invalid token")
}

// Is ...
func Is(err, target error) bool {
	return errorpkg.Is(err, target)
}
