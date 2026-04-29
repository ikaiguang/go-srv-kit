package setuputil

import (
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
)

// GetConfig 获取配置。
func GetConfig(launcherManager LauncherManager) *configpb.Bootstrap {
	return launcherManager.GetConfig()
}

// GetLogger 获取日志。
func GetLogger(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLogger()
}

// GetLoggerForMiddleware 获取中间件日志。
func GetLoggerForMiddleware(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLoggerForMiddleware()
}

// GetLoggerForHelper 获取辅助工具日志。
func GetLoggerForHelper(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLoggerForHelper()
}

// Close 关闭所有已初始化组件。
func Close(launcherManager LauncherManager) error {
	return launcherManager.Close()
}

// NewLauncherManagerWithCleanup 向后兼容：等价于 NewWithCleanup。
func NewLauncherManagerWithCleanup(configFilePath string, configOpts ...configutil.Option) (LauncherManager, func(), error) {
	return NewWithCleanup(configFilePath, configOpts...)
}

// NewLauncherManager 向后兼容：调用 NewWithCleanup 并丢弃 cleanup。
func NewLauncherManager(configFilePath string, configOpts ...configutil.Option) (LauncherManager, error) {
	lm, _, err := NewWithCleanup(configFilePath, configOpts...)
	return lm, err
}

// NewSingletonLauncherManager 单例模式构造函数。
var (
	singletonMutex           sync.Once
	singletonLauncherManager LauncherManager
)

func NewSingletonLauncherManager(configFilePath string) (LauncherManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonLauncherManager, err = NewLauncherManager(configFilePath)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonLauncherManager, err
}

// LoadingConfig 加载配置文件并设置全局配置。
func LoadingConfig(configFilePath string, configOpts ...configutil.Option) (*configpb.Bootstrap, error) {
	conf, err := configutil.Loading(configFilePath, configOpts...)
	if err != nil {
		return nil, err
	}
	apputil.SetConfig(conf)
	return conf, nil
}
