package authpkg

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

// AccessTokenValidate 自定义验证
type AccessTokenValidate func(context.Context, *Claims) error

// Option is jwt option.
type Option func(*options)

// Parser is a jwt parser
type options struct {
	signingMethod         jwt.SigningMethod
	claims                func() jwt.Claims
	accessTokenHeader     map[string]interface{}
	accessTokenValidators []AccessTokenValidate
}

// WithSigningMethod with signing method option.
func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// WithClaims with customer claim
// If you use it in Server, f needs to return a new jwt.Claims object each time to avoid concurrent write problems
// If you use it in Client, f only needs to return a single object to provide performance
func WithClaims(f func() jwt.Claims) Option {
	return func(o *options) {
		o.claims = f
	}
}

// WithAccessTokenHeader withe customer accessTokenHeader for client side
func WithAccessTokenHeader(header map[string]interface{}) Option {
	return func(o *options) {
		o.accessTokenHeader = header
	}
}

// WithAccessTokenValidator token验证
func WithAccessTokenValidator(tokenValidators ...AccessTokenValidate) Option {
	return func(o *options) {
		o.accessTokenValidators = append(o.accessTokenValidators, tokenValidators...)
	}
}
