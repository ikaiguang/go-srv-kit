package debugpkg

import (
	"sync"
	"testing"

	logpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/log"
)

func TestSetupConcurrentLogging(_ *testing.T) {
	defer CloseDebug()
	var wg sync.WaitGroup
	for range 8 {
		wg.Add(3)
		go func() {
			defer wg.Done()
			for range 100 {
				Setup(logpkg.NewNopLogger())
			}
		}()
		go func() {
			defer wg.Done()
			for range 100 {
				CloseDebug()
			}
		}()
		go func() {
			defer wg.Done()
			for range 100 {
				Info("concurrent")
			}
		}()
	}
	wg.Wait()
}
