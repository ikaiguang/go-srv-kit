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
	return ErrorTokenMissing("token is missing")
}
func ErrMissingSignKeyFunc() *errors.Error {
	return ErrorTokenKeyMissing("keyFunc is missing")
}
func ErrUnSupportSigningMethod() *errors.Error {
	return ErrorTokenMethodMissing("Wrong signing method")
}
func ErrTokenParseFail() *errors.Error {
	return ErrorAuthenticationFailed("Fail to parse JWT token ")
}
func ErrTokenExpired() *errors.Error {
	return ErrorTokenExpired("JWT token has expired")
}
func ErrWrongContext() *errors.Error {
	return ErrorUnauthorized("Wrong context for middleware")
}
func ErrTokenInvalid() *errors.Error {
	return ErrorTokenInvalid("Token is invalid")
}
func ErrNeedTokenProvider() *errors.Error {
	return ErrorTokenInvalid("Token provider is missing")
}
func ErrSignToken() *errors.Error {
	return ErrorTokenInvalid("Can not sign token.Is the key correct?")
}
func ErrGetKey() *errors.Error {
	return ErrorTokenInvalid("Can not get key while signing token")
}
func ErrInvalidAuthToken() *errors.Error {
	return ErrorVerificationFailed("[validator] Invalid auth token")
}
func ErrInvalidClaims() *errors.Error {
	return ErrorInvalidClaims("[validator] Invalid auth claims")
}
func ErrBlacklist() *errors.Error {
	return ErrorTokenDeprecated("[validator] Deprecated token")
}
func ErrWhitelist() *errors.Error {
	return ErrorTokenNotInWhitelist("[validator] The token is not in the valid list")
}
