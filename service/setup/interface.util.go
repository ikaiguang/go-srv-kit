package setuputil

import (
	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
)

// ConfigProvider 配置提供者。
type ConfigProvider interface {
	GetConfig() *configpb.Bootstrap
}

// LoggerProvider 日志提供者。
type LoggerProvider interface {
	GetLogger() (log.Logger, error)
	GetLoggerForMiddleware() (log.Logger, error)
	GetLoggerForHelper() (log.Logger, error)
}

// RegistryProvider 组件注册表提供者。
type RegistryProvider interface {
	GetRegistry() *ComponentRegistry
}

// LifecycleProvider 生命周期提供者。
type LifecycleProvider interface {
	GetLifecycle() *Lifecycle
}

// Closer 关闭接口。
type Closer interface {
	Close() error
}

// LauncherManager 是服务启动核心，只包含配置、日志、组件注册表和生命周期。
// 具体基础设施由各 service 子模块通过 WithSetup() 注册，并通过对应子模块的 GetXxx() 取用。
type LauncherManager interface {
	ConfigProvider
	LoggerProvider
	RegistryProvider
	LifecycleProvider
	Closer
}
