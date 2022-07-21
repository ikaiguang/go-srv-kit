// Package jwtutil 摘自kratos子项目
package jwtutil

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
)

// authKey context.Context key
type authKey struct{}

// KeyFunc 自定义 jwt.Keyfunc
type KeyFunc func(context.Context) func(*jwt.Token) (interface{}, error)

const (
	// ExpireDuration 过期时间
	ExpireDuration = time.Hour * 24 * 7

	// BearerWord the bearer key word for authorization
	BearerWord string = "Bearer"
	// BearerFormat authorization token format
	BearerFormat string = "Bearer %s"

	// AuthorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	AuthorizationKey string = "Authorization"

	// Reason holds the error reason.
	Reason string = "UNAUTHORIZED"
)

// Claims jwt.Claims
// 查看更多信息 jwt.RegisteredClaims
type Claims struct {
	jwt.RegisteredClaims

	AuthPayload *authv1.Payload `json:"a_p,omitempty"`
}

// DefaultExpireTime 令牌过期时间
func DefaultExpireTime() *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(ExpireDuration))
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
					return nil, ErrMissingKeyFunc
				}
				keyFunc := jwtKeyFunc(ctx)
				if keyFunc == nil {
					return nil, ErrMissingKeyFunc
				}
				auths := strings.SplitN(header.RequestHeader().Get(AuthorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], BearerWord) {
					return nil, ErrMissingJwtToken
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
						return nil, errors.Unauthorized(Reason, err.Error())
					}
					if ve.Errors&jwt.ValidationErrorMalformed != 0 {
						return nil, ErrTokenInvalid
					}
					if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						return nil, ErrTokenExpired
					}
					return nil, ErrTokenParseFail
				}
				if !tokenInfo.Valid {
					return nil, ErrTokenInvalid
				}
				if tokenInfo.Method != o.signingMethod {
					return nil, ErrUnSupportSigningMethod
				}
				if o.validator != nil {
					if err = o.validator(tokenInfo); err != nil {
						return nil, err
					}
				}
				ctx = NewContext(ctx, tokenInfo.Claims)
				return handler(ctx, req)
			}
			return nil, ErrWrongContext
		}
	}
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
				return nil, ErrMissingKeyFunc
			}
			keyProvider := jwtKeyFunc(ctx)
			if keyProvider == nil {
				return nil, ErrNeedTokenProvider
			}
			token := jwt.NewWithClaims(o.signingMethod, o.claims())
			if o.tokenHeader != nil {
				for k, v := range o.tokenHeader {
					token.Header[k] = v
				}
			}
			key, err := keyProvider(token)
			if err != nil {
				return nil, ErrGetKey
			}
			tokenStr, err := token.SignedString(key)
			if err != nil {
				return nil, ErrSignToken
			}
			if o.validator != nil {
				if err = o.validator(token); err != nil {
					return nil, err
				}
			}
			if clientContext, ok := transport.FromClientContext(ctx); ok {
				clientContext.RequestHeader().Set(AuthorizationKey, fmt.Sprintf(BearerFormat, tokenStr))
				return handler(ctx, req)
			}
			return nil, ErrWrongContext
		}
	}
}

// NewContext put auth info into context
func NewContext(ctx context.Context, info jwt.Claims) context.Context {
	return context.WithValue(ctx, authKey{}, info)
}

// FromContext extract auth info from context
func FromContext(ctx context.Context) (token jwt.Claims, ok bool) {
	token, ok = ctx.Value(authKey{}).(jwt.Claims)
	return
}
