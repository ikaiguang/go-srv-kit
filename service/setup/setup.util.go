package setuputil

import (
	stdlog "log"

	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	debugpkg "github.com/ikaiguang/go-srv-kit/kratos/debug"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
)

// launcherManager 实现 LauncherManager 接口
type launcherManager struct {
	conf *configpb.Bootstrap
	lc   *Lifecycle

	// 日志组件（始终注册）
	loggerComp *Component[loggerutil.LoggerManager]

	// 类型擦除的组件注册表
	registry *ComponentRegistry
}

// New 创建 LauncherManager，纯懒加载模式。
// 通过各 service 子模块的 WithSetup() Option 按需注册组件。
func New(conf *configpb.Bootstrap, opts ...Option) (LauncherManager, error) {
	if conf == nil {
		return nil, errorpkg.ErrorBadRequest("bootstrap config is required")
	}

	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	lc := newLifecycle()
	lm := &launcherManager{
		conf:     conf,
		lc:       lc,
		registry: NewComponentRegistry(),
	}

	// 日志始终注册（其他组件依赖日志）
	lm.loggerComp = NewComponent(ComponentLogger, lm.newLoggerManager, lc)

	// 通过各 service 子模块的 WithSetup() Option 注册组件。
	for _, registrar := range o.componentRegistrars {
		registrar(lm)
	}

	// 日志始终初始化（其他组件依赖日志）
	loggerManager, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	loggerForHelper, err := loggerManager.GetLoggerForHelper()
	if err != nil {
		return nil, err
	}
	logpkg.Setup(loggerForHelper)
	debugpkg.Setup(loggerForHelper)

	// 按需急切初始化指定组件
	if err := lm.eagerInit(o.eagerComponents); err != nil {
		return nil, err
	}

	return lm, nil
}

// eagerInit 急切初始化指定的组件
func (lm *launcherManager) eagerInit(components []string) error {
	for _, name := range components {
		comp, ok := lm.registry.Get(RegistryKeyComp(name))
		if !ok {
			return componentNotRegisteredError(name)
		}
		initializer, ok := comp.(interface{ Init() error })
		if !ok {
			return errorpkg.ErrorBadRequest("component cannot be eagerly initialized: %s", name)
		}
		if err := initializer.Init(); err != nil {
			return err
		}
	}
	return nil
}

// GetRegistry 获取组件注册表（供组件 WithSetup Option 使用）
func (lm *launcherManager) GetRegistry() *ComponentRegistry {
	return lm.registry
}

// GetLifecycle 获取生命周期管理器（供组件 WithSetup Option 使用）
func (lm *launcherManager) GetLifecycle() *Lifecycle {
	return lm.lc
}

// GetLoggerComp 获取日志组件（供其他组件 factory 使用）
func (lm *launcherManager) GetLoggerComp() *Component[loggerutil.LoggerManager] {
	return lm.loggerComp
}

// newLoggerManager 创建日志管理器
func (lm *launcherManager) newLoggerManager() (loggerutil.LoggerManager, error) {
	return loggerutil.NewLoggerManager(lm.conf.GetLog(), lm.conf.GetApp())
}

// NewWithCleanup 便捷函数，加载配置并创建 LauncherManager。
func NewWithCleanup(configFilePath string, configOpts ...configutil.Option) (LauncherManager, func(), error) {
	return NewWithCleanupOptions(configFilePath, configOpts)
}

// NewWithCleanupOptions 加载配置并创建 LauncherManager，同时允许调用方显式注册组件。
func NewWithCleanupOptions(configFilePath string, configOpts []configutil.Option, setupOpts ...Option) (LauncherManager, func(), error) {
	conf, err := configutil.Loading(configFilePath, configOpts...)
	if err != nil {
		return nil, nil, err
	}
	apputil.SetConfig(conf)

	lm, err := New(conf, setupOpts...)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		closeErr := lm.Close()
		if closeErr != nil {
			stdlog.Printf("==> launcherManager.Close failed: %+v\n", closeErr)
		}
	}
	return lm, cleanup, nil
}
