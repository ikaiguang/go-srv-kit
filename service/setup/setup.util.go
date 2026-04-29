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
	registry *componentRegistry
}

// New 创建 LauncherManager，纯懒加载模式
// 通过 WithXxx() Option 按需注册组件
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
		registry: newComponentRegistry(),
	}

	// 日志始终注册（其他组件依赖日志）
	lm.loggerComp = NewComponent(ComponentLogger, lm.newLoggerManager, lc)

	// 通过 WithXxx() Option 注册的组件
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
	initMap := map[string]func() error{
		ComponentRedis:    func() error { _, err := lm.GetRedisClient(); return err },
		ComponentMysql:    func() error { _, err := lm.GetMysqlDBConn(); return err },
		ComponentPostgres: func() error { _, err := lm.GetPostgresDBConn(); return err },
		ComponentMongo:    func() error { _, err := lm.GetMongoClient(); return err },
		ComponentConsul:   func() error { _, err := lm.GetConsulClient(); return err },
		ComponentJaeger:   func() error { _, err := lm.GetJaegerExporter(); return err },
		ComponentRabbitmq: func() error { _, err := lm.GetRabbitmqConn(); return err },
		ComponentAuth:     func() error { _, err := lm.GetTokenManager(); return err },
	}

	for _, name := range components {
		initFn, ok := initMap[name]
		if !ok {
			return errorpkg.ErrorBadRequest("unknown component: %s", name)
		}
		if err := initFn(); err != nil {
			return err
		}
	}
	return nil
}

// GetRegistry 获取组件注册表（供 WithXxx Option 使用）
func (lm *launcherManager) GetRegistry() *componentRegistry {
	return lm.registry
}

// GetLifecycle 获取生命周期管理器（供 WithXxx Option 使用）
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

// NewWithCleanup 便捷函数，加载配置并创建 LauncherManager
// 向后兼容：自动注入 Consul 配置加载器和所有组件
func NewWithCleanup(configFilePath string, configOpts ...configutil.Option) (LauncherManager, func(), error) {
	// 向后兼容：自动注入 Consul 配置加载器
	configOpts = append([]configutil.Option{configutil.WithConsulConfigLoader(configutil.NewConsulConfigLoader())}, configOpts...)

	conf, err := configutil.Loading(configFilePath, configOpts...)
	if err != nil {
		return nil, nil, err
	}
	apputil.SetConfig(conf)

	// 向后兼容：注册所有组件
	lm, err := New(conf, WithAllComponents())
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
