package middlewarepkg

import (
	"github.com/go-kratos/kratos/contrib/otel/v3/tracing"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v3/middleware/metadata"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/v3/app"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/log"
)

// DefaultServerMiddlewares 中间件
func DefaultServerMiddlewares(logger logpkg.Logger) []middleware.Middleware {
	slogLogger := logpkg.ToSlog(logger)
	return []middleware.Middleware{
		recovery.Recovery(
			recovery.WithHandler(RecoveryHandler()),
			recovery.WithLogger(slogLogger),
		),
		tracing.Server(),
		RateLimit(), // ratelimit.Server(),
		metadata.Server(),
		RequestAndResponseHeader(),
		apppkg.ServerLog(slogLogger, apppkg.WithDefaultDepth()),
		Validator(), // validate.Validator(),
	}
}

// DefaultClientMiddlewares 中间件
func DefaultClientMiddlewares(logger logpkg.Logger) []middleware.Middleware {
	slogLogger := logpkg.ToSlog(logger)
	return []middleware.Middleware{
		recovery.Recovery(
			recovery.WithHandler(RecoveryHandler()),
			recovery.WithLogger(slogLogger),
		),
		tracing.Client(),
		circuitbreaker.Client(), // error: circuitbreaker.ErrNotAllowed  reason=CIRCUITBREAKER
		metadata.Client(),
		apppkg.ClientLog(slogLogger),
	}
}
