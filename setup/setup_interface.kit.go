package setuputil

import (
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
)

// Config 配置
type Config interface {
	// Env app环境
	Env() envv1.Env
	// IsDebugMode 是否启用 调试模式
	IsDebugMode() bool
	// EnableLoggingConsole 是否启用 日志输出到控制台
	EnableLoggingConsole() bool
	// EnableLoggingFile 是否启用 日志输出到文件
	EnableLoggingFile() bool

	// AppConfig APP配置
	AppConfig() *confv1.App
	// LoggerConfig 日志配置
	LoggerConfig() *confv1.Log
	// DataConfig 数据配置
	DataConfig() *confv1.Data
	// MySQLConfig mysql配置
	MySQLConfig() *confv1.Data_MySQL
	// RedisConfig redis配置
	RedisConfig() *confv1.Data_Redis
}
