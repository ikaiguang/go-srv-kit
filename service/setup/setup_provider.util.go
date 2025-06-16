package setuputil

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
	"sync"
)

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

// GetRecommendDBConn 获取数据库连接
func GetRecommendDBConn(launcherManager LauncherManager) (*gorm.DB, error) {
	return GetDBConn(launcherManager)
}

// GetDBConn recommend Postgres
var GetDBConn = func(launcherManager LauncherManager) (*gorm.DB, error) {
	dbConn, err := launcherManager.GetPostgresDBConn()
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func GetConfig(launcherManager LauncherManager) *configpb.Bootstrap {
	return launcherManager.GetConfig()
}
func GetLogger(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLogger()
}
func GetLoggerForMiddleware(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLoggerForMiddleware()
}
func GetLoggerForHelper(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLoggerForHelper()
}
func GetRedisClient(launcherManager LauncherManager) (redis.UniversalClient, error) {
	return launcherManager.GetRedisClient()
}
func GetMysqlDBConn(launcherManager LauncherManager) (*gorm.DB, error) {
	return launcherManager.GetMysqlDBConn()
}
func GetPostgresDBConn(launcherManager LauncherManager) (*gorm.DB, error) {
	return launcherManager.GetPostgresDBConn()
}
func GetMongoClient(launcherManager LauncherManager) (*mongo.Client, error) {
	return launcherManager.GetMongoClient()
}
func GetConsulClient(launcherManager LauncherManager) (*consulapi.Client, error) {
	return launcherManager.GetConsulClient()
}
func GetJaegerExporter(launcherManager LauncherManager) (*otlptrace.Exporter, error) {
	return launcherManager.GetJaegerExporter()
}
func GetRabbitmqConn(launcherManager LauncherManager) (*rabbitmqpkg.ConnectionWrapper, error) {
	return launcherManager.GetRabbitmqConn()
}
func GetTokenManager(launcherManager LauncherManager) (authpkg.TokenManger, error) {
	return launcherManager.GetTokenManager()
}
func GetAuthManager(launcherManager LauncherManager) (authpkg.AuthRepo, error) {
	return launcherManager.GetAuthManager()
}
func GetServiceAPIManager(launcherManager LauncherManager) (clientutil.ServiceAPIManager, error) {
	return launcherManager.GetServiceApiManager()
}
func Close(launcherManager LauncherManager) error {
	return launcherManager.Close()
}
