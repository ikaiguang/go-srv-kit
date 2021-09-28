package setuphandler

import (
	pkgerrors "github.com/pkg/errors"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
)

// setupDebugUtil 设置调试工具
func (s *setup) setupDebugUtil() (err error) {
	if s.isInit {
		return pkgerrors.New("处理手柄未初始化")
	}
	// 调试工具
	if !s.enableDebug {
		return err
	}
	return debugutil.Setup()
}
