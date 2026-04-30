package jaegerutil

import (
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
)

// WithSetup 注册 Jaeger 组件。
func WithSetup() setuputil.Option {
	return setuputil.WithComponentRegistrar(func(ctx setuputil.ComponentRegistrarContext) {
		setuputil.RegisterComponent(ctx.GetRegistry(), setuputil.ComponentJaeger, func() (JaegerManager, error) {
			return NewJaegerManager(ctx.GetConfig().GetJaeger())
		}, ctx.GetLifecycle())
		setuputil.RegisterComponentGroup(ctx.GetRegistry(), setuputil.ComponentJaeger, func(name string) func() (JaegerManager, error) {
			return func() (JaegerManager, error) {
				jaegerConfig, ok := ctx.GetConfig().GetJaegerInstances()[name]
				if !ok {
					return nil, setuputil.ComponentNotFoundError(setuputil.ComponentJaeger, name)
				}
				return NewJaegerManager(jaegerConfig)
			}
		}, ctx.GetLifecycle())
	})
}

// GetExporter 从 LauncherManager 获取默认 Jaeger exporter。
func GetExporter(launcherManager setuputil.LauncherManager) (*otlptrace.Exporter, error) {
	mgr, err := setuputil.GetComponentValue[JaegerManager](launcherManager, setuputil.ComponentJaeger)
	if err != nil {
		return nil, err
	}
	return mgr.GetExporter()
}

// GetNamedExporter 从 LauncherManager 获取命名 Jaeger exporter。
func GetNamedExporter(launcherManager setuputil.LauncherManager, name string) (*otlptrace.Exporter, error) {
	mgr, err := setuputil.GetNamedComponentValue[JaegerManager](launcherManager, setuputil.ComponentJaeger, name)
	if err != nil {
		return nil, err
	}
	return mgr.GetExporter()
}

// TracerOptions 为 all-in-one 启动注入 Jaeger tracer。
func TracerOptions(launcherManager setuputil.LauncherManager) ([]middlewarepkg.TracerOption, error) {
	exp, err := GetExporter(launcherManager)
	if err != nil {
		return nil, err
	}
	return []middlewarepkg.TracerOption{
		middlewarepkg.WithTracerJaegerExporter(exp),
	}, nil
}

// WithTracerOptions 将 Jaeger tracer 接入 all-in-one 启动。
func WithTracerOptions() serverutil.Option {
	return serverutil.WithTracerOptionProvider(TracerOptions)
}
