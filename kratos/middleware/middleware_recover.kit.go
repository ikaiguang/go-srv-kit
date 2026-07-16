package middlewarepkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/error"
)

var _ = recovery.ErrUnknownRequest

// RecoveryHandler ...
func RecoveryHandler() recovery.HandlerFunc {
	return func(ctx context.Context, req, err any) error {
		e := errorpkg.ErrorPanic(errorpkg.ERROR_INTERNAL_SERVER.String())
		return errorpkg.Wrap(e)
	}
}
