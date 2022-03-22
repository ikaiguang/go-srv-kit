package setup

import (
	stdlog "log"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
)

// loadingDebugUtil 加载调试工具
func (s *modules) loadingDebugUtil() error {
	if !s.Config.IsDebugMode() {
		return nil
	}
	stdlog.Printf("|*** 加载调试工具debugutil")
	syncFn, err := debugutil.Setup()
	if err != nil {
		return err
	}
	s.debugHelperCloseFnSlice = append(s.debugHelperCloseFnSlice, syncFn)
	return err
}
