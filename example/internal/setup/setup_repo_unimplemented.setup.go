package setup

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"gorm.io/gorm"
	"io"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	tokenutil "github.com/ikaiguang/go-srv-kit/kratos/token"
)

var _ Engine = (*UnimplementedEngine)(nil)

// UnimplementedEngine ...
type UnimplementedEngine struct{}

func (UnimplementedEngine) Watch(key string, o config.Observer) error {
	return errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) AppConfig() *confv1.App {
	return nil
}

func (UnimplementedEngine) ServerConfig() *confv1.Server {
	return nil
}

func (UnimplementedEngine) ServerAuthConfig() *confv1.Server_Auth {
	return nil
}

func (UnimplementedEngine) ServerSettingConfig() *confv1.Server_Setting {
	return nil
}

func (UnimplementedEngine) Env() envv1.Env {
	return envv1.Env_UNKNOWN
}

func (UnimplementedEngine) IsDebugMode() bool {
	return false
}

func (UnimplementedEngine) EnableLoggingConsole() bool {
	return false
}

func (UnimplementedEngine) EnableLoggingFile() bool {
	return false
}

func (UnimplementedEngine) LoggerConfigForConsole() *confv1.Log_Console {
	return nil
}

func (UnimplementedEngine) LoggerConfigForFile() *confv1.Log_File {
	return nil
}

func (UnimplementedEngine) DataConfig() *confv1.Data {
	return nil
}

func (UnimplementedEngine) MySQLConfig() *confv1.Data_MySQL {
	return nil
}

func (UnimplementedEngine) PostgresConfig() *confv1.Data_PSQL {
	return nil
}

func (UnimplementedEngine) RedisConfig() *confv1.Data_Redis {
	return nil
}

func (UnimplementedEngine) ConsulConfig() *confv1.Data_Consul {
	return nil
}

func (UnimplementedEngine) JaegerTraceConfig() *confv1.Data_JaegerTrace {
	return nil
}

func (UnimplementedEngine) HTTPConfig() *confv1.Server_HTTP {
	return nil
}

func (UnimplementedEngine) GRPCConfig() *confv1.Server_GRPC {
	return nil
}

func (UnimplementedEngine) LoggerPrefixField() *LoggerPrefixField {
	return nil
}

func (UnimplementedEngine) LoggerFileWriter() (io.Writer, error) {
	return nil, errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) Logger() (log.Logger, []io.Closer, error) {
	return nil, nil, errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) LoggerHelper() (log.Logger, []io.Closer, error) {
	return nil, nil, errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) LoggerMiddleware() (log.Logger, []io.Closer, error) {
	return nil, nil, errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) GetMySQLGormDB() (*gorm.DB, error) {
	return nil, errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) GetPostgresGormDB() (*gorm.DB, error) {
	return nil, errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) GetRedisClient() (*redis.Client, error) {
	return nil, errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) GetConsulClient() (*api.Client, error) {
	return nil, errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) GetJaegerTraceExporter() (*jaeger.Exporter, error) {
	return nil, errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}

func (UnimplementedEngine) GetAuthTokenRepo(redisCC *redis.Client) tokenutil.AuthTokenRepo {
	return nil
}

func (UnimplementedEngine) Close() error {
	return errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "")
}
