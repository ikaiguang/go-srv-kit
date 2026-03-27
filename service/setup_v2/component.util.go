package setupv2

import (
	stderrors "errors"
	stdlog "log"
	"sync"
)

// Closer 关闭接口
type Closer interface {
	Close() error
}

// closerEntry 关闭器条目
type closerEntry struct {
	name   string
	closer Closer
}

// Lifecycle 管理组件的生命周期，按注册逆序关闭
type Lifecycle struct {
	mu      sync.Mutex
	closers []closerEntry
}

// newLifecycle 创建生命周期管理器
func newLifecycle() *Lifecycle {
	return &Lifecycle{}
}

// Register 注册一个需要关闭的组件
func (l *Lifecycle) Register(name string, closer Closer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.closers = append(l.closers, closerEntry{name: name, closer: closer})
}

// Close 按注册逆序关闭所有组件
func (l *Lifecycle) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	var errs []error
	for i := len(l.closers) - 1; i >= 0; i-- {
		entry := l.closers[i]
		stdlog.Printf("|*** STOP: close: %s", entry.name)
		if err := entry.closer.Close(); err != nil {
			stdlog.Printf("|*** STOP: close: %s failed: %s", entry.name, err.Error())
			errs = append(errs, err)
		}
	}
	l.closers = nil
	if len(errs) > 0 {
		return stderrors.Join(errs...)
	}
	return nil
}

// Component 泛型懒加载组件容器
// T 为底层 manager 类型（如 redisutil.RedisManager）
type Component[T any] struct {
	mu      sync.Mutex
	once    sync.Once
	value   T
	factory func() (T, error)
	lc      *Lifecycle
	name    string
}

// NewComponent 创建组件容器
func NewComponent[T any](name string, factory func() (T, error), lc *Lifecycle) *Component[T] {
	return &Component[T]{
		factory: factory,
		lc:      lc,
		name:    name,
	}
}

// Get 获取组件实例，首次调用时触发 factory 初始化
func (c *Component[T]) Get() (T, error) {
	var err error
	c.once.Do(func() {
		c.value, err = c.factory()
		if err == nil {
			// 如果 T 实现了 Closer 接口，自动注册到生命周期管理
			if closer, ok := any(c.value).(Closer); ok {
				c.lc.Register(c.name, closer)
			}
		}
	})
	if err != nil {
		c.mu.Lock()
		c.once = sync.Once{}
		c.mu.Unlock()
		var zero T
		return zero, err
	}
	return c.value, nil
}

// ComponentGroup 管理同类型组件的多个命名实例
type ComponentGroup[T any] struct {
	mu        sync.Mutex
	instances map[string]*Component[T]
	lc        *Lifecycle
	baseType  string
	factoryFn func(name string) func() (T, error)
}

// NewComponentGroup 创建命名实例容器
func NewComponentGroup[T any](baseType string, factoryFn func(name string) func() (T, error), lc *Lifecycle) *ComponentGroup[T] {
	return &ComponentGroup[T]{
		instances: make(map[string]*Component[T]),
		lc:        lc,
		baseType:  baseType,
		factoryFn: factoryFn,
	}
}

// Get 获取指定名称的实例，首次调用时创建 Component 并懒加载
func (g *ComponentGroup[T]) Get(name string) (T, error) {
	g.mu.Lock()
	comp, ok := g.instances[name]
	if !ok {
		compName := g.baseType + ":" + name
		comp = NewComponent(compName, g.factoryFn(name), g.lc)
		g.instances[name] = comp
	}
	g.mu.Unlock()
	return comp.Get()
}
