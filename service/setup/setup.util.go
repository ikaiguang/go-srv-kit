package setuputil

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/go-kratos/kratos/v2/log"
	consulapi "github.com/hashicorp/consul/api"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	debugpkg "github.com/ikaiguang/go-srv-kit/debug"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"gorm.io/gorm"
	stdlog "log"
)

type LauncherManager interface {
	GetConfig() *configpb.Bootstrap

	GetLogger() (log.Logger, error)
	GetLoggerForMiddleware() (log.Logger, error)
	GetLoggerForHelper() (log.Logger, error)

	GetRedisClient() (redis.UniversalClient, error)
	GetMysqlDBConn() (*gorm.DB, error)
	GetPostgresDBConn() (*gorm.DB, error)
	GetMongoClient() (*mongo.Client, error)
	GetConsulClient() (*consulapi.Client, error)
	GetJaegerExporter() (*otlptrace.Exporter, error)
	GetRabbitmqConn() (*amqp.ConnectionWrapper, error)

	GetTokenManager() (authpkg.TokenManger, error)
	GetAuthManager() (authpkg.AuthRepo, error)

	GetServiceApiManager() (clientutil.ServiceAPIManager, error)

	Close() error
}

func LoadingConfig(configFilePath string, configOpts ...configutil.Option) (*configpb.Bootstrap, error) {
	conf, err := configutil.Loading(configFilePath, configOpts...)
	if err != nil {
		return nil, err
	}
	apputil.SetConfig(conf)
	return conf, nil
}

func NewLauncherManagerWithCleanup(configFilePath string, configOpts ...configutil.Option) (LauncherManager, func(), error) {
	launcherManager, err := NewLauncherManager(configFilePath, configOpts...)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		closeErr := launcherManager.Close()
		if closeErr != nil {
			stdlog.Printf("==> launcherManager.Close failed: %+v\n", closeErr)
		}
	}
	return launcherManager, cleanup, nil

}

func NewLauncherManager(configFilePath string, configOpts ...configutil.Option) (LauncherManager, error) {
	// 开始配置
	stdlog.Println("|==================== LOADING PROGRAM : START ====================|")
	defer stdlog.Println("|==================== LOADING PROGRAM : END ====================|")

	// 加载配置文件
	bootstrap, err := LoadingConfig(configFilePath, configOpts...)
	if err != nil {
		return nil, err
	}
	launcher := &launcherManager{
		conf: bootstrap,
	}

	// 初始化日志
	loggerManager, err := launcher.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	loggerForHelper, err := loggerManager.GetLoggerForHelper()
	if err != nil {
		return nil, err
	}
	logpkg.Setup(loggerForHelper)
	debugpkg.Setup(loggerForHelper)

	// redis
	redisConfig := bootstrap.GetRedis()
	if redisConfig.GetEnable() {
		_, err = launcher.GetRedisClient()
		if err != nil {
			return nil, err
		}
	}

	// mysql
	mysqlConfig := bootstrap.GetMysql()
	if mysqlConfig.GetEnable() {
		_, err = launcher.GetMysqlDBConn()
		if err != nil {
			return nil, err
		}
	}

	// postgres
	psqlConfig := bootstrap.GetPsql()
	if psqlConfig.GetEnable() {
		_, err = launcher.GetPostgresDBConn()
		if err != nil {
			return nil, err
		}
	}

	// mongo
	mongoConfig := bootstrap.GetMongo()
	if mongoConfig.GetEnable() {
		_, err = launcher.GetMongoClient()
		if err != nil {
			return nil, err
		}
	}

	// consul
	consulConfig := bootstrap.GetConsul()
	if consulConfig.GetEnable() {
		_, err = launcher.GetConsulClient()
		if err != nil {
			return nil, err
		}
	}

	// jaeger.GetExporter()
	jaegerConfig := bootstrap.GetJaeger()
	if jaegerConfig.GetEnable() {
		_, err = launcher.GetJaegerExporter()
		if err != nil {
			return nil, err
		}
	}

	// rabbitmq
	rabbitmqConfig := bootstrap.GetRabbitmq()
	if rabbitmqConfig.GetEnable() {
		_, err = launcher.GetRabbitmqConn()
		if err != nil {
			return nil, err
		}
	}

	// token
	settingConfig := bootstrap.GetSetting()
	if settingConfig.GetEnableAuthMiddleware() {
		_, err = launcher.getSingletonAuthInstance()
		if err != nil {
			return nil, err
		}
	}
	return launcher, nil
}
