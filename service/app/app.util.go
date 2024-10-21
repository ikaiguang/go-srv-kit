package apputil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func ToAppConfig(appConfig *configpb.App) *AppConfig {
	ac := &AppConfig{}
	ac.SetByPbApp(appConfig)
	return ac
}

type AppConfig struct {
	ProjectName   string //
	ServerName    string //
	ServerEnv     string // 开发环境；值：DEVELOP、TESTING、PREVIEW、PRODUCTION
	ServerVersion string //
}

func (a *AppConfig) SetByPbApp(appConfig *configpb.App) {
	if appConfig == nil {
		return
	}
	a.ProjectName = appConfig.ProjectName
	a.ServerName = appConfig.ServerName
	a.ServerEnv = appConfig.ServerEnv
	a.ServerVersion = appConfig.ServerVersion
}

const (
	RedisSep = ":"
	PathSep  = "/"
)

// ID 程序ID
// 例：go-srv-saas:user-service:DEVELOP:v1.0.0
func ID(appConfig *AppConfig) string {
	return Identifier(appConfig, RedisSep)
}

func Path(appConfig *AppConfig) string {
	return Identifier(appConfig, PathSep)
}

func AbsPath(appConfig *AppConfig) string {
	return PathSep + Identifier(appConfig, PathSep)
}

// Identifier app 唯一标准
// @result = app.ProjectName + "/" + app.ServerName + "/" + app.ServerEnv + "/" + app.ServerVersion
func Identifier(appConfig *AppConfig, sep string) string {
	var ss = make([]string, 0, 4)
	if appConfig.ProjectName != "" {
		ss = append(ss, appConfig.ProjectName)
	}
	serviceName := appConfig.ServerName
	if appConfig.ServerName == "" {
		serviceName = "my-service"
	}
	ss = append(ss, serviceName)
	ss = append(ss, strings.ToLower(apppkg.ParseEnv(appConfig.ServerEnv).String()))
	if appConfig.ServerVersion != "" {
		ss = append(ss, appConfig.ServerVersion)
	}
	return strings.Join(ss, sep)
}

// CurrentPath ...
func CurrentPath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}

// RuntimePath ...
func RuntimePath() (string, error) {
	p, err := os.Getwd()
	if err != nil {
		e := errorpkg.ErrorInternalServer("get runtime path failed")
		return "", errorpkg.WithStack(e)
	}
	return p, nil
}

// ServerDecoderEncoder ...
func ServerDecoderEncoder() []http.ServerOption {
	apppkg.RegisterCodec()
	return []http.ServerOption{
		http.RequestDecoder(apppkg.RequestDecoder),
		http.ResponseEncoder(apppkg.SuccessResponseEncoder),
		http.ErrorEncoder(apppkg.ErrorResponseEncoder),
	}
}

// ClientDecoderEncoder ...
func ClientDecoderEncoder() []http.ClientOption {
	return []http.ClientOption{
		http.WithResponseDecoder(apppkg.SuccessResponseDecoder),
	}
}
