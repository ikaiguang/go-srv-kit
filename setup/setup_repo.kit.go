package setuppkg

//import (
//	"github.com/go-kratos/kratos/v2/config"
//	"github.com/go-kratos/kratos/v2/log"
//	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
//	"github.com/redis/go-redis/v9"
//	"go.mongodb.org/mongo-driver/mongo"
//	"io"
//)
//
//var (
//	//_ Config = &configuration{}
//	//_ Engine = &engines{}
//
//	ErrUnimplemented = errorpkg.ErrorUnimplemented("UNIMPLEMENTED")
//	ErrUninitialized = errorpkg.ErrorUninitialized("UNINITIALIZED")
//)
//
//// Config 配置
//type Config interface {
//	Close() error
//	Watch(key string, o config.Observer) error
//
//	// RuntimeEnv app环境
//	RuntimeEnv() envv1.RuntimeEnvEnum_RuntimeEnv
//	// IsDebugMode 是否启用 调试模式
//	IsDebugMode() bool
//	// EnableLoggingConsole 是否启用 日志输出到控制台
//	EnableLoggingConsole() bool
//	EnableLoggingFile() bool
//	EnableLoggingGraylog() bool
//
//	// LoggerConfigForConsole 日志配置 控制台
//	LoggerConfigForConsole() *configv1.Log_Console
//	LoggerConfigForFile() *configv1.Log_File
//	LoggerConfigForGraylog() *configv1.Log_Graylog
//
//	// AppConfig APP配置
//	AppConfig() *configv1.App
//	HubConfig() *configv1.Hub
//	HTTPConfig() *configv1.Server_HTTP
//	GRPCConfig() *configv1.Server_GRPC
//	RegistryConfig() *configv1.Registry
//	RedisConfig() *configv1.Data_Redis
//	MongoConfig() *configv1.Data_Mongo
//	TransferEncryptConfig() *configv1.Secret_TransferEncrypt
//	ServiceEncryptConfig() *configv1.Secret_ServiceEncrypt
//	JwtEncryptConfig() *configv1.Secret_TokenEncrypt
//	RefreshEncryptConfig() *configv1.Secret_RefreshEncrypt
//}
//
//// Engine ...
//type Engine interface {
//	Close() error
//
//	Config
//
//	// Logger 日志处理实例 runtime.caller.skip + 1
//	// 用于 log.Helper 输出；例子：log.Helper.Info
//	Logger() (log.Logger, []io.Closer, error)
//	// LoggerHelper 日志处理实例 runtime.caller.skip + 2
//	// 用于包含 log.Helper 输出；例子：func Info(){log.Helper.Info()}
//	LoggerHelper() (log.Logger, []io.Closer, error)
//	// LoggerMiddleware 日志处理实例 runtime.caller.skip - 1
//	// 用于包含 http.Middleware(logging.Server)
//	LoggerMiddleware() (log.Logger, []io.Closer, error)
//
//	// GetHubAPIEndpoint hub
//	//GetHubAPIEndpoint() (string, error)
//
//	// GetHubServerApiEndpoint hub
//	GetHubServerApiEndpoint(serveName string) (string, error)
//	GetHubPubSubTopic(pubSubName string) (string, error)
//
//	// GetRedisClient 客户端
//	GetRedisClient() (redis.UniversalClient, error)
//	GetMongoDB() (*mongo.Database, error)
//}
//
//// initEngine ...
//func initEngine(conf Config) *engines {
//	return &engines{
//		Config: conf,
//	}
//}
