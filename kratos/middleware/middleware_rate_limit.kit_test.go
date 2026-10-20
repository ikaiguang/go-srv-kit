package middlewarepkg

import (
	"context"
	"testing"

	ratelimitpkg "github.com/go-kratos/aegis/ratelimit"
	"github.com/stretchr/testify/require"
)

type nilDoneLimiter struct{}

func (nilDoneLimiter) Allow() (ratelimitpkg.DoneFunc, error) {
	return nil, nil
}

func TestRateLimitAcceptsNilOptionsAndDone(t *testing.T) {
	middleware := RateLimit(nil, WithLimiter(nil), WithLimiter(nilDoneLimiter{}))
	handler := middleware(func(context.Context, any) (any, error) {
		return "ok", nil
	})

	require.NotPanics(t, func() {
		reply, err := handler(context.Background(), nil)
		require.NoError(t, err)
		require.Equal(t, "ok", reply)
	})
}
