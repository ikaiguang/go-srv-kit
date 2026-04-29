package setuputil

// Option 构造选项函数
type Option func(*options)

// ComponentRegistrar 组件注册器，由各 service 子模块的 WithXxx() 调用
type ComponentRegistrar func(lm *launcherManager)

type options struct {
	eagerComponents     []string             // 需要急切初始化的组件名称列表
	componentRegistrars []ComponentRegistrar // 组件注册器列表
}

// WithEagerInit 指定需要在构造时立即初始化的组件
// 例如: WithEagerInit(ComponentRedis, ComponentMysql)
func WithEagerInit(components ...string) Option {
	return func(o *options) {
		o.eagerComponents = append(o.eagerComponents, components...)
	}
}

// WithComponentRegistrar 注册组件注册器（内部使用，由各 WithXxx 调用）
func WithComponentRegistrar(registrar ComponentRegistrar) Option {
	return func(o *options) {
		o.componentRegistrars = append(o.componentRegistrars, registrar)
	}
}
