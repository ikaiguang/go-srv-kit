package authpkg

import (
	"github.com/go-kratos/kratos/v2/errors"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// Is ...
func Is(err, target error) bool {
	return errorpkg.Is(errorpkg.Cause(err), target)
}

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
	return errors.Unauthorized(ERROR_VERIFICATION_FAILED.String(), "[validator] Invalid auth token")
}
func ErrInvalidClaims() *errors.Error {
	return errors.Unauthorized(ERROR_INVALID_CLAIMS.String(), "[validator] Invalid auth claims")
}
func ErrBlacklist() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_DEPRECATED.String(), "[validator] Deprecated token")
}
func ErrWhitelist() *errors.Error {
	return errors.Unauthorized(ERROR_TOKEN_NOT_IN_WHITELIST.String(), "[validator] The token is not in the valid list")
}
