package setuputil

import (
	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
)

// GetConfig 获取启动配置。
func (lm *launcherManager) GetConfig() *configpb.Bootstrap {
	return lm.conf
}

// GetLogger 获取业务日志。
func (lm *launcherManager) GetLogger() (log.Logger, error) {
	mgr, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetLogger()
}

// GetLoggerForMiddleware 获取中间件日志。
func (lm *launcherManager) GetLoggerForMiddleware() (log.Logger, error) {
	mgr, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetLoggerForMiddleware()
}

// GetLoggerForHelper 获取辅助工具日志。
func (lm *launcherManager) GetLoggerForHelper() (log.Logger, error) {
	mgr, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetLoggerForHelper()
}

// Close 关闭所有已初始化组件。
func (lm *launcherManager) Close() error {
	return lm.lc.Close()
}

// getLoggerManager 供旧内部测试和局部工具使用。
func (lm *launcherManager) getLoggerManager() (loggerutil.LoggerManager, error) {
	return lm.loggerComp.Get()
}
