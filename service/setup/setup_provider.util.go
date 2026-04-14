package setuputil

import (
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"

	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	rabbitmqpkg "github.com/ikaiguang/go-srv-kit/data/rabbitmq"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"

	"gorm.io/gorm"
)

// GetRecommendDBConn 获取推荐的数据库连接（默认 Postgres）
func GetRecommendDBConn(launcherManager LauncherManager) (*gorm.DB, error) {
	return GetDBConn(launcherManager)
}

// GetDBConn 获取数据库连接（默认 Postgres，可通过赋值切换）
var GetDBConn = func(launcherManager LauncherManager) (*gorm.DB, error) {
	return launcherManager.GetPostgresDBConn()
}

// GetConfig 获取配置
func GetConfig(launcherManager LauncherManager) *configpb.Bootstrap {
	return launcherManager.GetConfig()
}

// GetLogger 获取日志
func GetLogger(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLogger()
}

// GetLoggerForMiddleware 获取中间件日志
func GetLoggerForMiddleware(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLoggerForMiddleware()
}

// GetLoggerForHelper 获取辅助工具日志
func GetLoggerForHelper(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLoggerForHelper()
}

// GetRedisClient 获取 Redis 客户端
func GetRedisClient(launcherManager LauncherManager) (redis.UniversalClient, error) {
	return launcherManager.GetRedisClient()
}

// GetMysqlDBConn 获取 MySQL 数据库连接
func GetMysqlDBConn(launcherManager LauncherManager) (*gorm.DB, error) {
	return launcherManager.GetMysqlDBConn()
}

// GetPostgresDBConn 获取 PostgreSQL 数据库连接
func GetPostgresDBConn(launcherManager LauncherManager) (*gorm.DB, error) {
	return launcherManager.GetPostgresDBConn()
}

// GetMongoClient 获取 MongoDB 客户端
func GetMongoClient(launcherManager LauncherManager) (*mongo.Client, error) {
	return launcherManager.GetMongoClient()
}

// GetConsulClient 获取 Consul 客户端
func GetConsulClient(launcherManager LauncherManager) (*consulapi.Client, error) {
	return launcherManager.GetConsulClient()
}

// GetJaegerExporter 获取 Jaeger 导出器
func GetJaegerExporter(launcherManager LauncherManager) (*otlptrace.Exporter, error) {
	return launcherManager.GetJaegerExporter()
}

// GetRabbitmqConn 获取 RabbitMQ 连接
func GetRabbitmqConn(launcherManager LauncherManager) (*rabbitmqpkg.ConnectionWrapper, error) {
	return launcherManager.GetRabbitmqConn()
}

// GetTokenManager 获取 Token 管理器
func GetTokenManager(launcherManager LauncherManager) (authpkg.TokenManger, error) {
	return launcherManager.GetTokenManager()
}

// GetAuthManager 获取认证管理器
func GetAuthManager(launcherManager LauncherManager) (authpkg.AuthRepo, error) {
	return launcherManager.GetAuthManager()
}

// GetServiceAPIManager 获取集群服务 API 管理器
func GetServiceAPIManager(launcherManager LauncherManager) (clientutil.ServiceAPIManager, error) {
	return launcherManager.GetServiceApiManager()
}

// Close 关闭所有组件
func Close(launcherManager LauncherManager) error {
	return launcherManager.Close()
}

// NewLauncherManagerWithCleanup 向后兼容：等价于 NewWithCleanup
func NewLauncherManagerWithCleanup(configFilePath string, configOpts ...configutil.Option) (LauncherManager, func(), error) {
	return NewWithCleanup(configFilePath, configOpts...)
}

// NewLauncherManager 向后兼容：调用 NewWithCleanup 并丢弃 cleanup
func NewLauncherManager(configFilePath string, configOpts ...configutil.Option) (LauncherManager, error) {
	lm, _, err := NewWithCleanup(configFilePath, configOpts...)
	return lm, err
}

// NewSingletonLauncherManager 单例模式构造函数
var (
	singletonMutex           sync.Once
	singletonLauncherManager LauncherManager
)

func NewSingletonLauncherManager(configFilePath string) (LauncherManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonLauncherManager, err = NewLauncherManager(configFilePath)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonLauncherManager, err
}

// LoadingConfig 加载配置文件并设置全局配置
func LoadingConfig(configFilePath string, configOpts ...configutil.Option) (*configpb.Bootstrap, error) {
	conf, err := configutil.Loading(configFilePath, configOpts...)
	if err != nil {
		return nil, err
	}
	apputil.SetConfig(conf)
	return conf, nil
}
