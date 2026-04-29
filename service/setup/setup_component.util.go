package setuputil

import errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"

// componentNotRegisteredError 组件未注册错误。
func componentNotRegisteredError(name string) error {
	e := errorpkg.ErrorBadRequest("component not registered: %s; import corresponding service module and pass its WithSetup() option", name)
	return errorpkg.WithStack(e)
}

// ComponentNotFoundError 返回命名组件配置不存在错误。
func ComponentNotFoundError(component, name string) error {
	e := errorpkg.ErrorNotFound("%s instance not found: %s", component, name)
	return errorpkg.WithStack(e)
}

// GetComponentValue 获取已注册组件的 manager 实例。
func GetComponentValue[T any](launcherManager LauncherManager, name string) (T, error) {
	comp, ok := GetComponent[T](launcherManager.GetRegistry(), name)
	if !ok {
		var zero T
		return zero, componentNotRegisteredError(name)
	}
	return comp.Get()
}

// GetNamedComponentValue 获取已注册命名组件的 manager 实例。
func GetNamedComponentValue[T any](launcherManager LauncherManager, name, instanceName string) (T, error) {
	group, ok := GetComponentGroup[T](launcherManager.GetRegistry(), name)
	if !ok {
		var zero T
		return zero, componentNotRegisteredError(name)
	}
	return group.Get(instanceName)
}
