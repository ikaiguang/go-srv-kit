package middlewarepkg

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
)

// DefaultServerMiddlewares 中间件
func DefaultServerMiddlewares(logHelper *log.Helper) []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(recovery.WithHandler(RecoveryHandler())),
		metadata.Server(),
		tracing.Server(),
		RequestAndResponseHeader(),
		apppkg.ServerLog(logHelper),
		Validator(), // validate.Validator(),
	}
}

// DefaultClientMiddlewares 中间件
func DefaultClientMiddlewares(logHelper *log.Helper) []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(recovery.WithHandler(RecoveryHandler())),
		metadata.Client(),
		tracing.Client(),
		apppkg.ClientLog(logHelper),
	}
}
