package threadpkg

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -v -count 1 ./kratos/thread -run TestGoSafe
func TestGoSafe(t *testing.T) {
	type args struct {
		fn func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "#TestGoSafe",
			args: args{fn: func() { panic(1) }},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoSafe(tt.args.fn)
		})
	}
	time.Sleep(time.Second)
}

func TestGoSafeWithContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")
	done := make(chan string, 1)

	GoSafeWithContext(ctx, func(ctx context.Context) {
		done <- ctx.Value("key").(string)
	})

	select {
	case got := <-done:
		assert.Equal(t, "value", got)
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for goroutine")
	}
}

func TestGoWithContextNilContext(t *testing.T) {
	done := make(chan struct{}, 1)

	GoWithContext(nil, func(ctx context.Context) {
		require.NotNil(t, ctx)
		done <- struct{}{}
	})

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for goroutine")
	}
}
