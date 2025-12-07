package middlewarepkg

import (
	"context"
	stderrors "errors"

	ratelimitpkg "github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

var _ = ratelimit.Server()

// ErrLimitExceed is service unavailable due to rate limit exceeded.
var ErrLimitExceed = errors.New(429, "RATELIMIT", "service unavailable due to rate limit exceeded")

// Option is ratelimit option.
type Option func(*options)

// WithLimiter set Limiter implementation,
// default is bbr limiter
func WithLimiter(limiter ratelimitpkg.Limiter) Option {
	return func(o *options) {
		o.limiter = limiter
	}
}

type options struct {
	limiter ratelimitpkg.Limiter
}

// RateLimit ratelimiter middleware
func RateLimit(opts ...Option) middleware.Middleware {
	options := &options{
		limiter: bbr.NewLimiter(),
	}
	for _, o := range opts {
		o(options)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			done, e := options.limiter.Allow()
			if e != nil {
				e := errorpkg.ErrorTooManyRequests(err.Error())
				// rejected
				return reply, errorpkg.Wrap(e, stderrors.New("[RATELIMIT] service unavailable due to rate limit exceeded"))
				//return nil, ErrLimitExceed
			}
			// allowed
			reply, err = handler(ctx, req)
			done(ratelimitpkg.DoneInfo{Err: err})
			return
		}
	}
}
