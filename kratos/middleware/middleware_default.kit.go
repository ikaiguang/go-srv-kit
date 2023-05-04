package middlewareutil

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
)

// DefaultMiddlewares 中间件
func DefaultMiddlewares() []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(),
		metadata.Server(),
		//tracing.Server(),
		//apppkg.ServiceLog(、logger),
	}
}
