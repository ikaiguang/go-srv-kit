package setuputil

import (
	"sync"
)

// ComponentRegistry 类型擦除的组件注册表
// 用于存储各种类型的 Component 和 ComponentGroup，避免核心模块依赖具体组件包
type ComponentRegistry struct {
	mu         sync.RWMutex
	components map[string]any // name -> *Component[T] or *ComponentGroup[T]
}

// NewComponentRegistry 创建组件注册表
func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{
		components: make(map[string]any),
	}
}

// Register 注册组件
func (r *ComponentRegistry) Register(key string, comp any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.components[key] = comp
}

// Get 获取组件
func (r *ComponentRegistry) Get(key string) (any, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	comp, ok := r.components[key]
	return comp, ok
}

// 注册表键名常量（区分单实例和命名实例组）
const (
	registryKeySuffix      = ""
	registryKeyGroupSuffix = ":group"
)

// RegistryKeyComp 单实例组件的注册表键
func RegistryKeyComp(name string) string {
	return name + registryKeySuffix
}

// RegistryKeyGroup 命名实例组的注册表键
func RegistryKeyGroup(name string) string {
	return name + registryKeyGroupSuffix
}

// RegisterComponent 注册单实例组件到注册表（泛型辅助函数）
func RegisterComponent[T any](registry *ComponentRegistry, name string, factory func() (T, error), lc *Lifecycle) *Component[T] {
	comp := NewComponent(name, factory, lc)
	registry.Register(RegistryKeyComp(name), comp)
	return comp
}

// RegisterComponentGroup 注册命名实例组到注册表（泛型辅助函数）
func RegisterComponentGroup[T any](registry *ComponentRegistry, name string, factoryFn func(name string) func() (T, error), lc *Lifecycle) *ComponentGroup[T] {
	group := NewComponentGroup(name, factoryFn, lc)
	registry.Register(RegistryKeyGroup(name), group)
	return group
}

// GetComponent 从注册表获取单实例组件（泛型辅助函数）
func GetComponent[T any](registry *ComponentRegistry, name string) (*Component[T], bool) {
	comp, ok := registry.Get(RegistryKeyComp(name))
	if !ok {
		return nil, false
	}
	typed, ok := comp.(*Component[T])
	return typed, ok
}

// GetComponentGroup 从注册表获取命名实例组（泛型辅助函数）
func GetComponentGroup[T any](registry *ComponentRegistry, name string) (*ComponentGroup[T], bool) {
	group, ok := registry.Get(RegistryKeyGroup(name))
	if !ok {
		return nil, false
	}
	typed, ok := group.(*ComponentGroup[T])
	return typed, ok
}
