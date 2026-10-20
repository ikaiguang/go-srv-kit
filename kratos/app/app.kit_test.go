package apppkg

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRuntimeEnvConcurrentAccess(t *testing.T) {
	defer SetRuntimeEnv(RuntimeEnvEnum_PRODUCTION)
	var wg sync.WaitGroup
	for range 16 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			SetRuntimeEnv(RuntimeEnvEnum_TESTING)
		}()
		go func() {
			defer wg.Done()
			_ = GetRuntimeEnv()
			_ = IsDebugMode()
		}()
	}
	wg.Wait()
	require.Contains(t, []RuntimeEnvEnum_RuntimeEnv{
		RuntimeEnvEnum_PRODUCTION,
		RuntimeEnvEnum_TESTING,
	}, GetRuntimeEnv())
}
