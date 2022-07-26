// Package jwtutil 摘自kratos子项目
package jwtutil

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
)

// KeyFunc 自定义 jwt.Keyfunc
type KeyFunc func(context.Context) jwt.Keyfunc

// DefaultKeyFunc 默认 KeyFunc == jwt.Keyfunc
func DefaultKeyFunc(secret string) KeyFunc {
	return func(ctx context.Context) jwt.Keyfunc {
		return func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		}
	}
}

// Server is a server auth middleware. Check the token and extract the info from token.
func Server(jwtKeyFunc KeyFunc, opts ...Option) middleware.Middleware {
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				if jwtKeyFunc == nil {
					return nil, authutil.ErrMissingKeyFunc
				}
				keyFunc := jwtKeyFunc(ctx)
				if keyFunc == nil {
					return nil, authutil.ErrMissingKeyFunc
				}
				auths := strings.SplitN(header.RequestHeader().Get(authutil.AuthorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], authutil.BearerWord) {
					return nil, authutil.ErrMissingJwtToken
				}
				jwtToken := auths[1]
				var (
					tokenInfo *jwt.Token
					err       error
				)
				if o.claims != nil {
					tokenInfo, err = jwt.ParseWithClaims(jwtToken, o.claims(), keyFunc)
				} else {
					tokenInfo, err = jwt.Parse(jwtToken, keyFunc)
				}
				if err != nil {
					ve, ok := err.(*jwt.ValidationError)
					if !ok {
						return nil, errors.Unauthorized(errorv1.ERROR_UNAUTHORIZED.String(), err.Error())
					}
					if ve.Errors&jwt.ValidationErrorMalformed != 0 {
						return nil, authutil.ErrTokenInvalid
					}
					if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						return nil, authutil.ErrTokenExpired
					}
					return nil, authutil.ErrTokenParseFail
				}
				if !tokenInfo.Valid {
					return nil, authutil.ErrTokenInvalid
				}
				if tokenInfo.Method != o.signingMethod {
					return nil, authutil.ErrUnSupportSigningMethod
				}
				if o.validator != nil {
					if err = o.validator(tokenInfo); err != nil {
						return nil, err
					}
				}
				ctx = authutil.NewJWTContext(ctx, tokenInfo.Claims)
				return handler(ctx, req)
			}
			return nil, authutil.ErrWrongContext
		}
	}
}

// FromContext ...
func FromContext(ctx context.Context) (token jwt.Claims, ok bool) {
	return authutil.FromJWTContext(ctx)
}

// Client is a client jwt middleware.
func Client(jwtKeyFunc KeyFunc, opts ...Option) middleware.Middleware {
	claims := jwt.RegisteredClaims{}
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
		claims:        func() jwt.Claims { return claims },
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if jwtKeyFunc == nil {
				return nil, authutil.ErrMissingKeyFunc
			}
			keyProvider := jwtKeyFunc(ctx)
			if keyProvider == nil {
				return nil, authutil.ErrNeedTokenProvider
			}
			token := jwt.NewWithClaims(o.signingMethod, o.claims())
			if o.tokenHeader != nil {
				for k, v := range o.tokenHeader {
					token.Header[k] = v
				}
			}
			key, err := keyProvider(token)
			if err != nil {
				return nil, authutil.ErrGetKey
			}
			tokenStr, err := token.SignedString(key)
			if err != nil {
				return nil, authutil.ErrSignToken
			}
			if o.validator != nil {
				if err = o.validator(token); err != nil {
					return nil, err
				}
			}
			if clientContext, ok := transport.FromClientContext(ctx); ok {
				clientContext.RequestHeader().Set(authutil.AuthorizationKey, fmt.Sprintf(authutil.BearerFormat, tokenStr))
				return handler(ctx, req)
			}
			return nil, authutil.ErrWrongContext
		}
	}
}
