package jwtutil

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrMissingJwtToken        = errors.Unauthorized(Reason, "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(Reason, "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(Reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(Reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(Reason, "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(Reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(Reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(Reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(Reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(Reason, "Can not get key while signing token")
)
