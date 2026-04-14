package setuputil

// Option 构造选项函数
type Option func(*options)

type options struct {
	eagerComponents []string // 需要急切初始化的组件名称列表
}

// WithEagerInit 指定需要在构造时立即初始化的组件
// 例如: WithEagerInit(ComponentRedis, ComponentMysql)
func WithEagerInit(components ...string) Option {
	return func(o *options) {
		o.eagerComponents = append(o.eagerComponents, components...)
	}
}
