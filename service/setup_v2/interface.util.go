package setupv2

import (
	"github.com/go-kratos/kratos/v2/log"
	consulapi "github.com/hashicorp/consul/api"

	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	rabbitmqpkg "github.com/ikaiguang/go-srv-kit/data/rabbitmq"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"gorm.io/gorm"
)

// ConfigProvider 配置提供者
type ConfigProvider interface {
	GetConfig() *configpb.Bootstrap
}

// LoggerProvider 日志提供者
type LoggerProvider interface {
	GetLogger() (log.Logger, error)
	GetLoggerForMiddleware() (log.Logger, error)
	GetLoggerForHelper() (log.Logger, error)
}

// DatabaseProvider 数据库提供者
type DatabaseProvider interface {
	GetMysqlDBConn() (*gorm.DB, error)
	GetPostgresDBConn() (*gorm.DB, error)
	GetNamedMysqlDBConn(name string) (*gorm.DB, error)
	GetNamedPostgresDBConn(name string) (*gorm.DB, error)
}

// RedisProvider Redis 提供者
type RedisProvider interface {
	GetRedisClient() (redis.UniversalClient, error)
	GetNamedRedisClient(name string) (redis.UniversalClient, error)
}

// MongoProvider MongoDB 提供者
type MongoProvider interface {
	GetMongoClient() (*mongo.Client, error)
	GetNamedMongoClient(name string) (*mongo.Client, error)
}

// ConsulProvider 服务发现与配置中心提供者
type ConsulProvider interface {
	GetConsulClient() (*consulapi.Client, error)
	GetNamedConsulClient(name string) (*consulapi.Client, error)
}

// TracerProvider 链路追踪提供者
type TracerProvider interface {
	GetJaegerExporter() (*otlptrace.Exporter, error)
	GetNamedJaegerExporter(name string) (*otlptrace.Exporter, error)
}

// MessageQueueProvider 消息队列提供者
type MessageQueueProvider interface {
	GetRabbitmqConn() (*rabbitmqpkg.ConnectionWrapper, error)
	GetNamedRabbitmqConn(name string) (*rabbitmqpkg.ConnectionWrapper, error)
}

// AuthProvider 认证提供者
type AuthProvider interface {
	GetTokenManager() (authpkg.TokenManger, error)
	GetAuthManager() (authpkg.AuthRepo, error)
}

// ServiceAPIProvider 集群服务 API 提供者
type ServiceAPIProvider interface {
	GetServiceApiManager() (clientutil.ServiceAPIManager, error)
}

// LauncherManager 组合接口，保持向后兼容
type LauncherManager interface {
	ConfigProvider
	LoggerProvider
	DatabaseProvider
	RedisProvider
	MongoProvider
	ConsulProvider
	TracerProvider
	MessageQueueProvider
	AuthProvider
	ServiceAPIProvider
	Closer
}
