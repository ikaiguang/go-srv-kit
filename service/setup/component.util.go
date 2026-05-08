package setuputil

import (
	stderrors "errors"
	stdlog "log"
	"sync"
	"sync/atomic"
)

// closerEntry 关闭器条目
type closerEntry struct {
	name   string
	closer Closer
}

// Lifecycle 管理组件的生命周期，按注册逆序关闭
//
// 说明：Register 与 Close 在进程生命周期内调用次数很少（与基础设施数量同阶），
// 这里保留 sync.Mutex 是因为写路径简单、可读性高，lock-free 改造收益不大。
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

// componentValue 封装 Component 的成功值，用于 atomic.Pointer
// 使用指针包装，确保 Load() == nil 严格表示"未初始化"，避免值类型零值歧义
type componentValue[T any] struct {
	value T
}

// Component 泛型懒加载组件容器
// T 为底层 manager 类型（如 redisutil.RedisManager）
//
// 并发语义：
//   - Get() 热路径完全无锁：通过 atomic.Pointer.Load 读取已初始化的值
//   - 首次初始化通过 mu 互斥 + 双检，保证 factory 只成功执行一次
//   - factory 失败不缓存，下次 Get() 会重新尝试（与原 sync.Once 重置语义一致）
//   - 相较原 sync.Once+重置方案，修复了 factory 失败与并发 Get 之间的 data race
//     及可能返回 (零值, nil) 的正确性问题
type Component[T any] struct {
	state   atomic.Pointer[componentValue[T]] // nil=未初始化；非 nil=已成功
	mu      sync.Mutex                        // 仅 factory 执行期间互斥
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
// 热路径（已初始化）无锁；首次初始化通过 mu 互斥；失败不缓存，下次重试
func (c *Component[T]) Get() (T, error) {
	// 快速路径：原子读，已初始化时无锁返回
	if cv := c.state.Load(); cv != nil {
		return cv.value, nil
	}

	// 慢速路径：互斥 + 双检
	c.mu.Lock()
	defer c.mu.Unlock()

	// 双检：可能在等待锁期间已被其他 goroutine 初始化完成
	if cv := c.state.Load(); cv != nil {
		return cv.value, nil
	}

	v, err := c.factory()
	if err != nil {
		// 失败不发布到 state，下次 Get() 会重新执行 factory
		var zero T
		return zero, err
	}

	// 成功：先发布 value 到 state，再注册 Closer
	// 顺序保证：state 可见时 Closer 已准备好（Register 仍在 mu 保护内）
	c.state.Store(&componentValue[T]{value: v})
	if closer, ok := any(v).(Closer); ok {
		c.lc.Register(c.name, closer)
	}
	return v, nil
}

// ComponentGroup 管理同类型组件的多个命名实例
//
// 并发语义：
//   - Get(name) 热路径使用 RWMutex 的读锁，多个 goroutine 可并发读已存在实例
//   - 首次创建某个命名实例时升级为写锁，通过双检避免重复创建
//   - 底层 Component.Get 仍保持无锁快速路径
type ComponentGroup[T any] struct {
	mu        sync.RWMutex
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
	// 快速路径：读锁查询已创建的 Component
	g.mu.RLock()
	comp, ok := g.instances[name]
	g.mu.RUnlock()
	if ok {
		return comp.Get()
	}

	// 慢速路径：写锁 + 双检，创建新 Component
	g.mu.Lock()
	comp, ok = g.instances[name]
	if !ok {
		compName := g.baseType + ":" + name
		comp = NewComponent(compName, g.factoryFn(name), g.lc)
		g.instances[name] = comp
	}
	g.mu.Unlock()
	return comp.Get()
}
