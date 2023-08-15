package middlewarepkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// validator ...
type validator interface {
	Validate() error
}

// Validator is a validator middleware.
func Validator() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if v, ok := req.(validator); ok {
				if err := v.Validate(); err != nil {
					e := errorpkg.ErrorInvalidParameter(err.Error())
					return reply, errorpkg.Wrap(e, err)
				}
			}
			return handler(ctx, req)
		}
	}
}
