package middlewareutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
)

// DefaultServerMiddlewares 中间件
func DefaultServerMiddlewares(logHelper *log.Helper) []middleware.Middleware {
	return middlewarepkg.DefaultServerMiddlewares(logHelper)
}

// DefaultClientMiddlewares 中间件
func DefaultClientMiddlewares(logHelper *log.Helper) []middleware.Middleware {
	return middlewarepkg.DefaultClientMiddlewares(logHelper)
}
