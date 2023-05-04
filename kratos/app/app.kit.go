package apputil

import (
	"strings"
)

const (
	AppIDSep         = ":"
	AppConfigPathSep = "/"
)

// App application
type App struct {
	// BelongTo 属于哪个项目
	BelongTo string
	// Name app名字
	Name string
	// Version app版本
	Version string
	// Env app 环境
	Env string
	// EnvBranch 环境分支
	EnvBranch string
	// Endpoints app站点
	Endpoints []string
	// Metadata 元数据
	Metadata map[string]string
}

// ID 程序ID
// 例：go-srv-services/DEVELOP/main/v1.0.0/user-service
func ID(appConfig *App) string {
	return appIdentifier(appConfig, AppIDSep)
}

// ConfigPath 配置路径；用于配置中心，如：consul、etcd、...
// @result = app.BelongTo + "/" + app.RuntimeEnv + "/" + app.Branch + "/" + app.Version + "/" + app.Name
// 例：go-srv-services/DEVELOP/main/v1.0.0/user-service
func ConfigPath(appConfig *App) string {
	return appIdentifier(appConfig, AppConfigPathSep)
}

// appIdentifier app 唯一标准
// @result = app.BelongTo + "/" + app.RuntimeEnv + "/" + app.Branch + "/" + app.Version + "/" + app.Name
// 例：go-srv-services/DEVELOP/main/v1.0.0/user-service
func appIdentifier(appConfig *App, sep string) string {
	var ss = make([]string, 0, 5)
	if appConfig.BelongTo != "" {
		ss = append(ss, appConfig.BelongTo)
	}
	if appConfig.Env != "" {
		ss = append(ss, appConfig.Env)
	}
	if appConfig.EnvBranch != "" {
		branchString := strings.Replace(appConfig.EnvBranch, " ", ":", -1)
		ss = append(ss, branchString)
	}
	if appConfig.Version != "" {
		ss = append(ss, appConfig.Version)
	}
	if appConfig.Name != "" {
		ss = append(ss, appConfig.Name)
	}
	return strings.Join(ss, sep)
}
