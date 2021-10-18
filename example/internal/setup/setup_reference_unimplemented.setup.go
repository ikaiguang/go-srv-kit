package setup

import (
	"io"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
)

var (
	_ Packages = &UnimplementedPackages{}
	_ Config   = &UnimplementedConfig{}
)

// UnimplementedPackages
type UnimplementedPackages struct{}

// LoggerFileWriter 文件日志写手柄
func (s *UnimplementedPackages) LoggerFileWriter() (io.Writer, error) {
	return nil, pkgerrors.WithStack(ErrUnimplemented)
}

// Logger 日志处理示例
func (s *UnimplementedPackages) Logger() (log.Logger, error) {
	return nil, pkgerrors.WithStack(ErrUnimplemented)
}

// MysqlGormDB mysql gorm 数据库
func (s *UnimplementedPackages) MysqlGormDB() (*gorm.DB, error) {
	return nil, pkgerrors.WithStack(ErrUnimplemented)
}

// RedisClient redis 客户端
func (s *UnimplementedPackages) RedisClient() (*redis.Client, error) {
	return nil, pkgerrors.WithStack(ErrUnimplemented)
}

// UnimplementedConfig
type UnimplementedConfig struct{}

// Env app环境
func (s *UnimplementedConfig) Env() envv1.Env {
	return envv1.Env_UNKNOWN
}

// IsDebugMode 是否启用 调试模式
func (s *UnimplementedConfig) IsDebugMode() bool {
	return false
}

// EnableLoggingConsole 是否启用 日志输出到控制台
func (s *UnimplementedConfig) EnableLoggingConsole() bool {
	return false
}

// EnableLoggingFile 是否启用 日志输出到文件
func (s *UnimplementedConfig) EnableLoggingFile() bool {
	return false
}

// AppConfig APP配置
func (s *UnimplementedConfig) AppConfig() *confv1.App {
	return nil
}

// LoggerConfig 日志配置
func (s *UnimplementedConfig) LoggerConfig() *confv1.Log {
	return nil
}

// DataConfig 数据配置
func (s *UnimplementedConfig) DataConfig() *confv1.Data {
	return nil
}

// MySQLConfig mysql配置
func (s *UnimplementedConfig) MySQLConfig() *confv1.Data_MySQL {
	return nil
}

// RedisConfig redis配置
func (s *UnimplementedConfig) RedisConfig() *confv1.Data_Redis {
	return nil
}
