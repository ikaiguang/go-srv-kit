package setup

import (
	strerrors "errors"
	"fmt"
	"io"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
)

var (
	_ Config = &configuration{}
	_ Engine = &engines{}

	// ErrUnimplemented 未实现
	ErrUnimplemented = strerrors.New("unimplemented")
	// ErrUninitialized 未初始化
	ErrUninitialized = strerrors.New("uninitialized")
)

// IsUnimplementedError 未实现
func IsUnimplementedError(err error) bool {
	return strerrors.Is(pkgerrors.Cause(err), ErrUnimplemented)
}

// IsUninitializedError 未初始化
func IsUninitializedError(err error) bool {
	return strerrors.Is(pkgerrors.Cause(err), ErrUninitialized)
}

// LoggerPrefixField with logger fields.
type LoggerPrefixField struct {
	AppName    string `json:"name"`
	AppVersion string `json:"version"`
	AppEnv     string `json:"env"`
	Hostname   string `json:"hostname"`
	ServerIP   string `json:"serverIP"`
}

// String returns the string representation of LoggerPrefixField.
func (s *LoggerPrefixField) String() string {
	strSlice := []string{
		"name:" + fmt.Sprintf("%q", s.AppName),
		"version:" + fmt.Sprintf("%q", s.AppVersion),
		"env:" + fmt.Sprintf("%q", s.AppEnv),
		"hostname:" + fmt.Sprintf("%q", s.Hostname),
		"serverIP:" + fmt.Sprintf("%q", s.ServerIP),
	}
	return strings.Join(strSlice, " ")
}

// Config 配置
type Config interface {
	// AppConfig APP配置
	AppConfig() *confv1.App
	// AppAuthConfig APP验证配置
	AppAuthConfig() *confv1.App_Auth
	// AppSettingConfig APP设置配置
	AppSettingConfig() *confv1.App_Setting

	// Env app环境
	Env() envv1.Env
	// IsDebugMode 是否启用 调试模式
	IsDebugMode() bool
	// EnableLoggingConsole 是否启用 日志输出到控制台
	EnableLoggingConsole() bool
	// EnableLoggingFile 是否启用 日志输出到文件
	EnableLoggingFile() bool

	// LoggerConfigForConsole 日志配置 控制台
	LoggerConfigForConsole() *confv1.Log_Console
	// LoggerConfigForFile 日志配置 文件
	LoggerConfigForFile() *confv1.Log_File
	// MySQLConfig mysql配置
	MySQLConfig() *confv1.Data_MySQL
	// PostgresConfig postgres配置
	PostgresConfig() *confv1.Data_PSQL
	// RedisConfig redis配置
	RedisConfig() *confv1.Data_Redis
	// HTTPConfig http配置
	HTTPConfig() *confv1.Server_HTTP
	// GRPCConfig grpc配置
	GRPCConfig() *confv1.Server_GRPC
}

// Engine 引擎模块、组件、单元
type Engine interface {
	// Config 配置
	Config

	// LoggerPrefixField 日志前缀字段
	LoggerPrefixField() *LoggerPrefixField
	// LoggerFileWriter 文件日志写手柄
	LoggerFileWriter() (io.Writer, error)
	// Logger 日志处理实例 runtime.caller.skip + 1
	// 用于 log.Helper 输出；例子：log.Helper.Info
	Logger() (log.Logger, []io.Closer, error)
	// LoggerHelper 日志处理实例 runtime.caller.skip + 2
	// 用于包含 log.Helper 输出；例子：func Info(){log.Helper.Info()}
	LoggerHelper() (log.Logger, []io.Closer, error)
	// LoggerMiddleware 日志处理实例 runtime.caller.skip - 1
	// 用于包含 http.Middleware(logging.Server)
	LoggerMiddleware() (log.Logger, []io.Closer, error)

	// GetMySQLGormDB mysql gorm 数据库
	GetMySQLGormDB() (*gorm.DB, error)
	// GetPostgresGormDB postgres gorm 数据库
	GetPostgresGormDB() (*gorm.DB, error)
	// GetRedisClient redis 客户端
	GetRedisClient() (*redis.Client, error)

	// Close 关闭
	Close() error
}
