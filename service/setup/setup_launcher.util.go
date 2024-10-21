package setuputil

import (
	stderrors "errors"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/go-kratos/kratos/v2/log"
	consulapi "github.com/hashicorp/consul/api"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	authutil "github.com/ikaiguang/go-srv-kit/service/auth"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	consulutil "github.com/ikaiguang/go-srv-kit/service/consul"
	jaegerutil "github.com/ikaiguang/go-srv-kit/service/jaeger"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	mongoutil "github.com/ikaiguang/go-srv-kit/service/mongo"
	mysqlutil "github.com/ikaiguang/go-srv-kit/service/mysql"
	postgresutil "github.com/ikaiguang/go-srv-kit/service/postgres"
	rabbitmqutil "github.com/ikaiguang/go-srv-kit/service/rabbitmq"
	redisutil "github.com/ikaiguang/go-srv-kit/service/redis"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"gorm.io/gorm"
	stdlog "log"
	"sync"
)

type launcherManager struct {
	conf *configpb.Bootstrap

	loggerManagerOnce     sync.Once
	loggerManager         loggerutil.LoggerManager
	redisManagerOnce      sync.Once
	redisManager          redisutil.RedisManager
	mysqlManagerOnce      sync.Once
	mysqlManager          mysqlutil.MysqlManager
	postgresManagerOnce   sync.Once
	postgresManager       postgresutil.PostgresManager
	mongoManagerOnce      sync.Once
	mongoManager          mongoutil.MongoManager
	consulManagerOnce     sync.Once
	consulManager         consulutil.ConsulManager
	jaegerManagerOnce     sync.Once
	jaegerManager         jaegerutil.JaegerManager
	rabbitmqManagerOnce   sync.Once
	rabbitmqManager       rabbitmqutil.RabbitmqManager
	authInstanceOnce      sync.Once
	authInstance          authutil.AuthInstance
	serviceAPIManagerOnce sync.Once
	serviceAPIManager     clientutil.ServiceAPIManager
}

func (s *launcherManager) GetConfig() *configpb.Bootstrap {
	return s.conf
}

func (s *launcherManager) getLoggerManager() (loggerutil.LoggerManager, error) {
	logConfig := s.conf.GetLog()
	appConfig := s.conf.GetApp()
	loggerManager, err := loggerutil.NewLoggerManager(logConfig, appConfig)
	if err != nil {
		return nil, err
	}
	s.loggerManager = loggerManager
	return loggerManager, nil
}

func (s *launcherManager) getSingletonLoggerManager() (loggerutil.LoggerManager, error) {
	var err error
	s.loggerManagerOnce.Do(func() {
		s.loggerManager, err = s.getLoggerManager()
	})
	if err != nil {
		s.loggerManagerOnce = sync.Once{}
	}
	return s.loggerManager, err
}

func (s *launcherManager) GetLogger() (log.Logger, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	return loggerManager.GetLogger()
}

func (s *launcherManager) GetLoggerForMiddleware() (log.Logger, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	return loggerManager.GetLoggerForMiddleware()
}

func (s *launcherManager) GetLoggerForHelper() (log.Logger, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	return loggerManager.GetLoggerForHelper()
}

func (s *launcherManager) getRedisManager() (redisutil.RedisManager, error) {
	redisConfig := s.conf.GetRedis()
	redisManager, err := redisutil.NewRedisManager(redisConfig)
	if err != nil {
		return nil, err
	}
	s.redisManager = redisManager
	return redisManager, nil
}

func (s *launcherManager) getSingletonRedisManager() (redisutil.RedisManager, error) {
	var err error
	s.redisManagerOnce.Do(func() {
		s.redisManager, err = s.getRedisManager()
	})
	if err != nil {
		s.redisManagerOnce = sync.Once{}
	}
	return s.redisManager, err
}

func (s *launcherManager) GetRedisClient() (redis.UniversalClient, error) {
	redisManager, err := s.getSingletonRedisManager()
	if err != nil {
		return nil, err
	}
	return redisManager.GetClient()
}

func (s *launcherManager) getMysqlManager() (mysqlutil.MysqlManager, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	mysqlConfig := s.conf.GetMysql()
	mysqlManager, err := mysqlutil.NewMysqlManager(mysqlConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.mysqlManager = mysqlManager
	return mysqlManager, nil
}

func (s *launcherManager) getSingletonMysqlManager() (mysqlutil.MysqlManager, error) {
	var err error
	s.mysqlManagerOnce.Do(func() {
		s.mysqlManager, err = s.getMysqlManager()
	})
	if err != nil {
		s.mysqlManagerOnce = sync.Once{}
	}
	return s.mysqlManager, err
}

func (s *launcherManager) GetMysqlDBConn() (*gorm.DB, error) {
	mysqlManager, err := s.getSingletonMysqlManager()
	if err != nil {
		return nil, err
	}
	return mysqlManager.GetDB()
}

func (s *launcherManager) getPostgresManager() (postgresutil.PostgresManager, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	psqlConfig := s.conf.GetPsql()
	postgresManager, err := postgresutil.NewPostgresManager(psqlConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.postgresManager = postgresManager
	return postgresManager, nil
}

func (s *launcherManager) getSingletonPostgresManager() (postgresutil.PostgresManager, error) {
	var err error
	s.postgresManagerOnce.Do(func() {
		s.postgresManager, err = s.getPostgresManager()
	})
	if err != nil {
		s.postgresManagerOnce = sync.Once{}
	}
	return s.postgresManager, err
}

func (s *launcherManager) GetPostgresDBConn() (*gorm.DB, error) {
	postgresManager, err := s.getSingletonPostgresManager()
	if err != nil {
		return nil, err
	}
	return postgresManager.GetDB()
}

func (s *launcherManager) getMongoManager() (mongoutil.MongoManager, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	mongoConfig := s.conf.GetMongo()
	mongoManager, err := mongoutil.NewMongoManager(mongoConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.mongoManager = mongoManager
	return mongoManager, nil
}

func (s *launcherManager) getSingletonMongoManager() (mongoutil.MongoManager, error) {
	var err error
	s.mongoManagerOnce.Do(func() {
		s.mongoManager, err = s.getMongoManager()
	})
	if err != nil {
		s.mongoManagerOnce = sync.Once{}
	}
	return s.mongoManager, err
}

func (s *launcherManager) GetMongoClient() (*mongo.Client, error) {
	mongoManager, err := s.getSingletonMongoManager()
	if err != nil {
		return nil, err
	}
	return mongoManager.GetMongoClient()
}

func (s *launcherManager) getConsulManager() (consulutil.ConsulManager, error) {
	consulConfig := s.conf.GetConsul()
	consulManager, err := consulutil.NewConsulManager(consulConfig)
	if err != nil {
		return nil, err
	}
	s.consulManager = consulManager
	return consulManager, nil
}

func (s *launcherManager) getSingletonConsulManager() (consulutil.ConsulManager, error) {
	var err error
	s.consulManagerOnce.Do(func() {
		s.consulManager, err = s.getConsulManager()
	})
	if err != nil {
		s.consulManagerOnce = sync.Once{}
	}
	return s.consulManager, err
}

func (s *launcherManager) GetConsulClient() (*consulapi.Client, error) {
	consulManager, err := s.getSingletonConsulManager()
	if err != nil {
		return nil, err
	}
	return consulManager.GetClient()
}

func (s *launcherManager) getJaegerManager() (jaegerutil.JaegerManager, error) {
	jaegerConfig := s.conf.GetJaeger()
	jaegerManager, err := jaegerutil.NewJaegerManager(jaegerConfig)
	if err != nil {
		return nil, err
	}
	s.jaegerManager = jaegerManager
	return jaegerManager, nil
}

func (s *launcherManager) getSingletonJaegerManager() (jaegerutil.JaegerManager, error) {
	var err error
	s.jaegerManagerOnce.Do(func() {
		s.jaegerManager, err = s.getJaegerManager()
	})
	if err != nil {
		s.jaegerManagerOnce = sync.Once{}
	}
	return s.jaegerManager, err
}

func (s *launcherManager) GetJaegerExporter() (*otlptrace.Exporter, error) {
	jaegerManager, err := s.getSingletonJaegerManager()
	if err != nil {
		return nil, err
	}
	return jaegerManager.GetExporter()
}

func (s *launcherManager) getRabbitmqManager() (rabbitmqutil.RabbitmqManager, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	rabbitmqConfig := s.conf.GetRabbitmq()
	rabbitmqManager, err := rabbitmqutil.NewRabbitmqManager(rabbitmqConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.rabbitmqManager = rabbitmqManager
	return rabbitmqManager, nil
}

func (s *launcherManager) getSingletonRabbitmqManager() (rabbitmqutil.RabbitmqManager, error) {
	var err error
	s.rabbitmqManagerOnce.Do(func() {
		s.rabbitmqManager, err = s.getRabbitmqManager()
	})
	if err != nil {
		s.rabbitmqManagerOnce = sync.Once{}
	}
	return s.rabbitmqManager, err
}

func (s *launcherManager) GetRabbitmqConn() (*amqp.ConnectionWrapper, error) {
	rabbitmqManager, err := s.getSingletonRabbitmqManager()
	if err != nil {
		return nil, err
	}
	return rabbitmqManager.GetClient()
}

func (s *launcherManager) getAuthInstance() (authutil.AuthInstance, error) {
	// logger
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	// redis
	universalClient, err := s.GetRedisClient()
	if err != nil {
		return nil, err
	}
	// auth
	encryptTokenEncrypt := s.conf.GetEncrypt().GetTokenEncrypt()
	authInstance, err := authutil.NewAuthInstance(encryptTokenEncrypt, universalClient, loggerManager)
	if err != nil {
		return nil, err
	}
	s.authInstance = authInstance
	return authInstance, err
}

func (s *launcherManager) getSingletonAuthInstance() (authutil.AuthInstance, error) {
	var err error
	s.authInstanceOnce.Do(func() {
		s.authInstance, err = s.getAuthInstance()
	})
	if err != nil {
		s.authInstanceOnce = sync.Once{}
	}
	return s.authInstance, err
}

func (s *launcherManager) GetTokenManager() (authpkg.TokenManger, error) {
	authInstance, err := s.getSingletonAuthInstance()
	if err != nil {
		return nil, err
	}
	return authInstance.GetTokenManger()
}

func (s *launcherManager) GetAuthManager() (authpkg.AuthRepo, error) {
	authInstance, err := s.getSingletonAuthInstance()
	if err != nil {
		return nil, err
	}
	return authInstance.GetAuthManger()
}

func (s *launcherManager) getServiceApiManager() (clientutil.ServiceAPIManager, error) {
	apiConfigs, diffRT, err := clientutil.ToConfig(s.conf.GetClusterServiceApi())
	if err != nil {
		return nil, err
	}
	loggerForMiddleware, err := s.GetLoggerForMiddleware()
	if err != nil {
		return nil, err
	}
	var opts = []clientutil.Option{clientutil.WithLogger(loggerForMiddleware)}
	for rt := range diffRT {
		if rt == configpb.RegistryTypeEnum_CONSUL {
			consulClient, err := s.GetConsulClient()
			if err != nil {
				return nil, err
			}
			opts = append(opts, clientutil.WithConsulClient(consulClient))
		} else if rt == configpb.RegistryTypeEnum_ETCD {
			e := errorpkg.ErrorUnimplemented("uninitialized setup etcd")
			return nil, errorpkg.WithStack(e)
			//etcdClient, err := s.GetEtcdClient()
			//if err != nil {
			//	return nil, err
			//}
			//opts = append(opts, clientutil.WithEtcdClient(etcdClient))
		}
	}
	serviceAPIManager, err := clientutil.NewServiceAPIManager(apiConfigs, opts...)
	if err != nil {
		return nil, err
	}
	s.serviceAPIManager = serviceAPIManager
	return serviceAPIManager, nil
}

func (s *launcherManager) getSingletonServiceApiManager() (clientutil.ServiceAPIManager, error) {
	var err error
	s.serviceAPIManagerOnce.Do(func() {
		s.serviceAPIManager, err = s.getServiceApiManager()
	})
	if err != nil {
		s.serviceAPIManagerOnce = sync.Once{}
	}
	return s.serviceAPIManager, err
}

func (s *launcherManager) GetServiceApiManager() (clientutil.ServiceAPIManager, error) {
	serviceAPIManager, err := s.getSingletonServiceApiManager()
	if err != nil {
		return nil, err
	}
	return serviceAPIManager, nil
}

func (s *launcherManager) Close() error {
	// 退出程序
	stdlog.Println("|==================== EXIT PROGRAM : START ====================|")
	defer stdlog.Println("|==================== EXIT PROGRAM : END ====================|")
	var errs []error

	// redis
	if s.redisManager != nil {
		if err := s.redisManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// mysql
	if s.mysqlManager != nil {
		if err := s.mysqlManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// postgres
	if s.postgresManager != nil {
		if err := s.postgresManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// mongo
	if s.mongoManager != nil {
		if err := s.mongoManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// consul
	if s.consulManager != nil {
		if err := s.consulManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// jaeger
	if s.jaegerManager != nil {
		if err := s.jaegerManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// rabbitmq
	if s.rabbitmqManager != nil {
		if err := s.rabbitmqManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// logger
	if s.loggerManager != nil {
		if err := s.loggerManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return stderrors.Join(errs...)
	}
	return nil
}
