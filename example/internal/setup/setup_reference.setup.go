package setup

import (
	strerrors "errors"
	"io"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
)

var (
	// Config implement
	_ Config = &configuration{}
	// Packages implement
	_ Packages = &up{}

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

// Args 参数
type Args interface {
	// AppConfig APP配置
	AppConfig() *confv1.App

	// Env app环境
	Env() envv1.Env
	// IsDebugMode 是否启用 调试模式
	IsDebugMode() bool
	// EnableLoggingConsole 是否启用 日志输出到控制台
	EnableLoggingConsole() bool
	// EnableLoggingFile 是否启用 日志输出到文件
	EnableLoggingFile() bool
}

// Config 配置
type Config interface {
	// Args 参数
	Args

	// LoggerConfig 日志配置
	LoggerConfig() *confv1.Log
	// DataConfig 数据配置
	DataConfig() *confv1.Data
	// MySQLConfig mysql配置
	MySQLConfig() *confv1.Data_MySQL
	// RedisConfig redis配置
	RedisConfig() *confv1.Data_Redis
}

// Packages 包/依赖
type Packages interface {
	// Args 参数
	Args

	// LoggerFileWriter 文件日志写手柄
	LoggerFileWriter() (io.Writer, error)
	// Logger 日志处理实例
	Logger() (log.Logger, error)

	// MysqlGormDB mysql gorm 数据库
	MysqlGormDB() (*gorm.DB, error)

	// RedisClient redis 客户端
	RedisClient() (*redis.Client, error)
}
