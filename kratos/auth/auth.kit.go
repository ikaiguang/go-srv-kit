package authutil

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
)

const (
	// ExpireDuration 过期时间
	ExpireDuration = time.Hour * 24 * 7

	// AuthorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	AuthorizationKey string = "Authorization"

	// BearerWord the bearer key word for authorization
	BearerWord string = "Bearer"
	// BearerFormat authorization token format
	BearerFormat string = "Bearer %s"
)

// DefaultExpireTime 令牌过期时间
func DefaultExpireTime() *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(ExpireDuration))
}

// Claims jwt.Claims
// 查看更多信息 jwt.RegisteredClaims
type Claims struct {
	jwt.RegisteredClaims

	Payload *authv1.Payload `json:"p,omitempty"`
}

// KeyFunc 自定义 jwt.Keyfunc
type KeyFunc func(ctx context.Context) (context.Context, jwt.Keyfunc)

// Server is a server auth middleware. Check the token and extract the info from token.
func Server(customKeyFunc KeyFunc, opts ...Option) middleware.Middleware {
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				var keyFunc jwt.Keyfunc
				if customKeyFunc == nil {
					return nil, errorutil.WithStack(ErrMissingKeyFunc)
				}
				ctx, keyFunc = customKeyFunc(ctx)
				if keyFunc == nil {
					return nil, errorutil.WithStack(ErrMissingKeyFunc)
				}
				auths := strings.SplitN(header.RequestHeader().Get(AuthorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], BearerWord) {
					return nil, errorutil.WithStack(ErrMissingJwtToken)
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
						return nil, errorutil.WithStack(ErrTokenInvalid)
					}
					if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						return nil, errorutil.WithStack(ErrTokenExpired)
					}
					return nil, errorutil.WithStack(ErrTokenParseFail)
				}
				if !tokenInfo.Valid {
					return nil, errorutil.WithStack(ErrTokenInvalid)
				}
				if tokenInfo.Method != o.signingMethod {
					return nil, errorutil.WithStack(ErrUnSupportSigningMethod)
				}
				if o.validator != nil {
					if err = o.validator(ctx, tokenInfo); err != nil {
						return nil, err
					}
				}
				ctx = NewJWTContext(ctx, tokenInfo.Claims)
				return handler(ctx, req)
			}
			return nil, errorutil.WithStack(ErrWrongContext)
		}
	}
}

// Client is a client jwt middleware.
func Client(customKeyFunc KeyFunc, opts ...Option) middleware.Middleware {
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
			var keyProvider jwt.Keyfunc
			if customKeyFunc == nil {
				return nil, errorutil.WithStack(ErrMissingKeyFunc)
			}
			ctx, keyProvider = customKeyFunc(ctx)
			if keyProvider == nil {
				return nil, errorutil.WithStack(ErrMissingKeyFunc)
			}
			if keyProvider == nil {
				return nil, errorutil.WithStack(ErrNeedTokenProvider)
			}
			token := jwt.NewWithClaims(o.signingMethod, o.claims())
			if o.tokenHeader != nil {
				for k, v := range o.tokenHeader {
					token.Header[k] = v
				}
			}
			key, err := keyProvider(token)
			if err != nil {
				return nil, errorutil.WithStack(ErrGetKey)
			}
			tokenStr, err := token.SignedString(key)
			if err != nil {
				return nil, errorutil.WithStack(ErrSignToken)
			}
			if o.validator != nil {
				if err = o.validator(ctx, token); err != nil {
					return nil, err
				}
			}
			if clientContext, ok := transport.FromClientContext(ctx); ok {
				clientContext.RequestHeader().Set(AuthorizationKey, fmt.Sprintf(BearerFormat, tokenStr))
				return handler(ctx, req)
			}
			return nil, errorutil.WithStack(ErrWrongContext)
		}
	}
}
