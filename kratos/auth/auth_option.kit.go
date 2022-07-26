package authutil

import "github.com/golang-jwt/jwt/v4"

// Option is jwt option.
type Option func(*options)

// Parser is a jwt parser
type options struct {
	signingMethod jwt.SigningMethod
	claims        func() jwt.Claims
	tokenHeader   map[string]interface{}
	validator     func(*jwt.Token) error
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

// WithTokenHeader withe customer tokenHeader for client side
func WithTokenHeader(header map[string]interface{}) Option {
	return func(o *options) {
		o.tokenHeader = header
	}
}

// WithValidator 自定义验证
func WithValidator(validator func(*jwt.Token) error) Option {
	return func(o *options) {
		o.validator = validator
	}
}
