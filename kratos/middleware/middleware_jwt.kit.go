package middlewarepkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/selector"
)

// NewWhiteListMatcher 路由白名单
func NewWhiteListMatcher(whiteList map[string]struct{}) selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
