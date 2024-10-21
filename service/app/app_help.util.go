package apputil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

var (
	_bootstrap *configpb.Bootstrap
)

func SetConfig(bootstrap *configpb.Bootstrap) {
	_bootstrap = bootstrap
}

func GetConfig() (*configpb.Bootstrap, error) {
	if _bootstrap == nil {
		e := errorpkg.ErrorUninitialized("bootstrap is uninitialized")
		return nil, errorpkg.WithStack(e)
	}
	return _bootstrap, nil
}

func GetID() string {
	return ID(ToAppConfig(_bootstrap.GetApp()))
}

func GetEnv() apppkg.RuntimeEnvEnum_RuntimeEnv {
	return apppkg.ParseEnv(_bootstrap.GetApp().GetServerEnv())
}

func IsDebugMode() bool {
	switch GetEnv() {
	default:
		return false
	case apppkg.RuntimeEnvEnum_LOCAL, apppkg.RuntimeEnvEnum_DEVELOP, apppkg.RuntimeEnvEnum_TESTING:
		return true
	}
}

func IsLocalMode() bool {
	switch GetEnv() {
	default:
		return false
	case apppkg.RuntimeEnvEnum_LOCAL:
		return true
	}
}
