package apppkg

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClientLogMeasuresHandlerLatency(t *testing.T) {
	var output bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&output, nil))
	middleware := ClientLog(logger)
	handler := middleware(func(context.Context, any) (any, error) {
		time.Sleep(10 * time.Millisecond)
		return nil, nil
	})

	_, err := handler(context.Background(), nil)
	require.NoError(t, err)

	var record map[string]any
	require.NoError(t, json.Unmarshal(output.Bytes(), &record))
	latency, err := time.ParseDuration(record["latency"].(string))
	require.NoError(t, err)
	require.GreaterOrEqual(t, latency, 10*time.Millisecond)
}

func TestMaxRequestArgSizeConcurrentAccess(t *testing.T) {
	defer SetMaxRequestArgSize(defaultMaxRequestArgs)
	var wg sync.WaitGroup
	for range 16 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			SetMaxRequestArgSize(3)
		}()
		go func() {
			defer wg.Done()
			_ = truncateRequestArgs("abcdef")
		}()
	}
	wg.Wait()
	require.Equal(t, "abc", truncateRequestArgs("abcdef"))
}

func TestLoggingMiddlewareAcceptsNilLogger(t *testing.T) {
	require.NotPanics(t, func() {
		ServerLog(nil)
		ClientLog(nil)
	})
}
