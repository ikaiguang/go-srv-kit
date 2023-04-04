package setuppkg

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"gorm.io/gorm"
	"io"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	tokenutil "github.com/ikaiguang/go-srv-kit/kratos/token"
)

var _ Engine = (*Unimplemented)(nil)

// Unimplemented 未实现
type Unimplemented struct{}

// NotImplementedError ...
func NotImplementedError() error {
	reason := errorv1.ERROR_NOT_IMPLEMENTED
	message := errorv1.ERROR_NOT_IMPLEMENTED.String()
	err := errorutil.NotImplemented(reason.String(), message)
	return err
}

func (*Unimplemented) Watch(string, config.Observer) error {
	err := NotImplementedError()
	return err
}

func (*Unimplemented) AppConfig() *confv1.App {
	return nil
}

func (*Unimplemented) ServerConfig() *confv1.Server {
	return nil
}

func (*Unimplemented) HTTPConfig() *confv1.Server_HTTP {
	return nil
}

func (*Unimplemented) GRPCConfig() *confv1.Server_GRPC {
	return nil
}

func (*Unimplemented) BusinessAuthConfig() *confv1.Business_Auth {
	return nil
}

func (*Unimplemented) BaseSettingConfig() *confv1.Base_Setting {
	return nil
}

func (*Unimplemented) Env() envv1.Env {
	return envv1.Env_PRODUCTION
}

func (*Unimplemented) IsDebugMode() bool {
	return false
}

func (*Unimplemented) EnableLoggingConsole() bool {
	return false
}

func (*Unimplemented) EnableLoggingFile() bool {
	return false
}

func (*Unimplemented) LoggerConfigForConsole() *confv1.Log_Console {
	return nil
}

func (*Unimplemented) LoggerConfigForFile() *confv1.Log_File {
	return nil
}

func (*Unimplemented) DataConfig() *confv1.Data {
	return nil
}

func (*Unimplemented) MySQLConfig() *confv1.Data_MySQL {
	return nil
}

func (*Unimplemented) PostgresConfig() *confv1.Data_PSQL {
	return nil
}

func (*Unimplemented) RedisConfig() *confv1.Data_Redis {
	return nil
}

func (*Unimplemented) ConsulConfig() *confv1.Base_Consul {
	return nil
}

func (*Unimplemented) JaegerTracerConfig() *confv1.Base_JaegerTracer {
	return nil
}

func (*Unimplemented) SnowflakeWorkerConfig() *confv1.Base_SnowflakeWorker {
	return nil
}

func (*Unimplemented) LoggerPrefixField() *LoggerPrefixField {
	return nil
}

func (*Unimplemented) LoggerFileWriter() (io.Writer, error) {
	err := NotImplementedError()
	return nil, err
}

func (*Unimplemented) Logger() (log.Logger, []io.Closer, error) {
	err := NotImplementedError()
	return nil, nil, err
}

func (*Unimplemented) LoggerHelper() (log.Logger, []io.Closer, error) {
	err := NotImplementedError()
	return nil, nil, err
}

func (*Unimplemented) LoggerMiddleware() (log.Logger, []io.Closer, error) {
	err := NotImplementedError()
	return nil, nil, err
}

func (*Unimplemented) GetMySQLGormDB() (*gorm.DB, error) {
	err := NotImplementedError()
	return nil, err
}

func (*Unimplemented) GetPostgresGormDB() (*gorm.DB, error) {
	err := NotImplementedError()
	return nil, err
}

func (*Unimplemented) GetRedisClient() (*redis.Client, error) {
	err := NotImplementedError()
	return nil, err
}

func (*Unimplemented) GetConsulClient() (*api.Client, error) {
	err := NotImplementedError()
	return nil, err
}

func (*Unimplemented) GetJaegerTraceExporter() (*jaeger.Exporter, error) {
	err := NotImplementedError()
	return nil, err
}

func (*Unimplemented) GetAuthTokenRepo(redisCC *redis.Client) tokenutil.AuthTokenRepo {
	return nil
}

func (*Unimplemented) Close() error {
	return nil
}
