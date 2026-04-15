package setuputil

import (
	"github.com/go-kratos/kratos/v2/log"
	consulapi "github.com/hashicorp/consul/api"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	rabbitmqpkg "github.com/ikaiguang/go-srv-kit/data/rabbitmq"
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
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"gorm.io/gorm"
)

// ==================== 单实例 factory 方法 ====================

// newLoggerManager 创建日志管理器
func (lm *launcherManager) newLoggerManager() (loggerutil.LoggerManager, error) {
	return loggerutil.NewLoggerManager(lm.conf.GetLog(), lm.conf.GetApp())
}

// newRedisManager 创建 Redis 管理器
func (lm *launcherManager) newRedisManager() (redisutil.RedisManager, error) {
	return redisutil.NewRedisManager(lm.conf.GetRedis())
}

// newMysqlManager 创建 MySQL 管理器
func (lm *launcherManager) newMysqlManager() (mysqlutil.MysqlManager, error) {
	loggerManager, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mysqlutil.NewMysqlManager(lm.conf.GetMysql(), loggerManager)
}

// newPostgresManager 创建 PostgreSQL 管理器
func (lm *launcherManager) newPostgresManager() (postgresutil.PostgresManager, error) {
	loggerManager, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return postgresutil.NewPostgresManager(lm.conf.GetPsql(), loggerManager)
}

// newMongoManager 创建 MongoDB 管理器
func (lm *launcherManager) newMongoManager() (mongoutil.MongoManager, error) {
	loggerManager, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mongoutil.NewMongoManager(lm.conf.GetMongo(), loggerManager)
}

// newConsulManager 创建 Consul 管理器
func (lm *launcherManager) newConsulManager() (consulutil.ConsulManager, error) {
	return consulutil.NewConsulManager(lm.conf.GetConsul())
}

// newJaegerManager 创建 Jaeger 管理器
func (lm *launcherManager) newJaegerManager() (jaegerutil.JaegerManager, error) {
	return jaegerutil.NewJaegerManager(lm.conf.GetJaeger())
}

// newRabbitmqManager 创建 RabbitMQ 管理器
func (lm *launcherManager) newRabbitmqManager() (rabbitmqutil.RabbitmqManager, error) {
	loggerManager, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return rabbitmqutil.NewRabbitmqManager(lm.conf.GetRabbitmq(), loggerManager)
}

// newAuthInstance 创建认证实例
func (lm *launcherManager) newAuthInstance() (authutil.AuthInstance, error) {
	loggerManager, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	redisClient, err := lm.GetRedisClient()
	if err != nil {
		return nil, err
	}
	return authutil.NewAuthInstance(lm.conf.GetEncrypt().GetTokenEncrypt(), redisClient, loggerManager)
}

// newServiceAPIManager 创建集群服务 API 管理器
func (lm *launcherManager) newServiceAPIManager() (clientutil.ServiceAPIManager, error) {
	apiConfigs, diffRT, err := clientutil.ToConfig(lm.conf.GetClusterServiceApi())
	if err != nil {
		return nil, err
	}
	loggerForMiddleware, err := lm.GetLoggerForMiddleware()
	if err != nil {
		return nil, err
	}
	var opts = []clientutil.Option{clientutil.WithLogger(loggerForMiddleware)}
	for rt := range diffRT {
		switch rt {
		case configpb.RegistryTypeEnum_CONSUL:
			consulClient, err := lm.GetConsulClient()
			if err != nil {
				return nil, err
			}
			opts = append(opts, clientutil.WithConsulClient(consulClient))
		case configpb.RegistryTypeEnum_ETCD:
			e := errorpkg.ErrorUnimplemented("uninitialized setup etcd")
			return nil, errorpkg.WithStack(e)
		}
	}
	return clientutil.NewServiceAPIManager(apiConfigs, opts...)
}

// ==================== 命名实例 factory 方法 ====================

// newNamedMysqlManager 根据实例名称生成 MySQL factory 函数
func (lm *launcherManager) newNamedMysqlManager(name string) func() (mysqlutil.MysqlManager, error) {
	return func() (mysqlutil.MysqlManager, error) {
		instances := lm.conf.GetMysqlInstances()
		mysqlConfig, ok := instances[name]
		if !ok {
			return nil, errorpkg.ErrorNotFound("mysql instance not found: %s", name)
		}
		loggerManager, err := lm.loggerComp.Get()
		if err != nil {
			return nil, err
		}
		return mysqlutil.NewMysqlManager(mysqlConfig, loggerManager)
	}
}

// newNamedPostgresManager 根据实例名称生成 PostgreSQL factory 函数
func (lm *launcherManager) newNamedPostgresManager(name string) func() (postgresutil.PostgresManager, error) {
	return func() (postgresutil.PostgresManager, error) {
		instances := lm.conf.GetPsqlInstances()
		psqlConfig, ok := instances[name]
		if !ok {
			return nil, errorpkg.ErrorNotFound("postgres instance not found: %s", name)
		}
		loggerManager, err := lm.loggerComp.Get()
		if err != nil {
			return nil, err
		}
		return postgresutil.NewPostgresManager(psqlConfig, loggerManager)
	}
}

// newNamedRedisManager 根据实例名称生成 Redis factory 函数
func (lm *launcherManager) newNamedRedisManager(name string) func() (redisutil.RedisManager, error) {
	return func() (redisutil.RedisManager, error) {
		instances := lm.conf.GetRedisInstances()
		redisConfig, ok := instances[name]
		if !ok {
			return nil, errorpkg.ErrorNotFound("redis instance not found: %s", name)
		}
		return redisutil.NewRedisManager(redisConfig)
	}
}

// newNamedMongoManager 根据实例名称生成 MongoDB factory 函数
func (lm *launcherManager) newNamedMongoManager(name string) func() (mongoutil.MongoManager, error) {
	return func() (mongoutil.MongoManager, error) {
		instances := lm.conf.GetMongoInstances()
		mongoConfig, ok := instances[name]
		if !ok {
			return nil, errorpkg.ErrorNotFound("mongo instance not found: %s", name)
		}
		loggerManager, err := lm.loggerComp.Get()
		if err != nil {
			return nil, err
		}
		return mongoutil.NewMongoManager(mongoConfig, loggerManager)
	}
}

// newNamedConsulManager 根据实例名称生成 Consul factory 函数
func (lm *launcherManager) newNamedConsulManager(name string) func() (consulutil.ConsulManager, error) {
	return func() (consulutil.ConsulManager, error) {
		instances := lm.conf.GetConsulInstances()
		consulConfig, ok := instances[name]
		if !ok {
			return nil, errorpkg.ErrorNotFound("consul instance not found: %s", name)
		}
		return consulutil.NewConsulManager(consulConfig)
	}
}

// newNamedJaegerManager 根据实例名称生成 Jaeger factory 函数
func (lm *launcherManager) newNamedJaegerManager(name string) func() (jaegerutil.JaegerManager, error) {
	return func() (jaegerutil.JaegerManager, error) {
		instances := lm.conf.GetJaegerInstances()
		jaegerConfig, ok := instances[name]
		if !ok {
			return nil, errorpkg.ErrorNotFound("jaeger instance not found: %s", name)
		}
		return jaegerutil.NewJaegerManager(jaegerConfig)
	}
}

// newNamedRabbitmqManager 根据实例名称生成 RabbitMQ factory 函数
func (lm *launcherManager) newNamedRabbitmqManager(name string) func() (rabbitmqutil.RabbitmqManager, error) {
	return func() (rabbitmqutil.RabbitmqManager, error) {
		instances := lm.conf.GetRabbitmqInstances()
		rabbitmqConfig, ok := instances[name]
		if !ok {
			return nil, errorpkg.ErrorNotFound("rabbitmq instance not found: %s", name)
		}
		loggerManager, err := lm.loggerComp.Get()
		if err != nil {
			return nil, err
		}
		return rabbitmqutil.NewRabbitmqManager(rabbitmqConfig, loggerManager)
	}
}

// ==================== Provider 方法：ConfigProvider ====================

// GetConfig 获取 Bootstrap 配置
func (lm *launcherManager) GetConfig() *configpb.Bootstrap {
	return lm.conf
}

// ==================== Provider 方法：LoggerProvider ====================

// GetLogger 获取日志实例
func (lm *launcherManager) GetLogger() (log.Logger, error) {
	mgr, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetLogger()
}

// GetLoggerForMiddleware 获取中间件日志实例
func (lm *launcherManager) GetLoggerForMiddleware() (log.Logger, error) {
	mgr, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetLoggerForMiddleware()
}

// GetLoggerForHelper 获取辅助工具日志实例
func (lm *launcherManager) GetLoggerForHelper() (log.Logger, error) {
	mgr, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetLoggerForHelper()
}

// ==================== Provider 方法：DatabaseProvider ====================

// GetMysqlDBConn 获取 MySQL 数据库连接
func (lm *launcherManager) GetMysqlDBConn() (*gorm.DB, error) {
	mgr, err := lm.mysqlComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

// GetPostgresDBConn 获取 PostgreSQL 数据库连接
func (lm *launcherManager) GetPostgresDBConn() (*gorm.DB, error) {
	mgr, err := lm.postgresComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

// GetNamedMysqlDBConn 获取命名 MySQL 实例的数据库连接
func (lm *launcherManager) GetNamedMysqlDBConn(name string) (*gorm.DB, error) {
	mgr, err := lm.mysqlGroup.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

// GetNamedPostgresDBConn 获取命名 PostgreSQL 实例的数据库连接
func (lm *launcherManager) GetNamedPostgresDBConn(name string) (*gorm.DB, error) {
	mgr, err := lm.postgresGroup.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

// ==================== Provider 方法：RedisProvider ====================

// GetRedisClient 获取 Redis 客户端
func (lm *launcherManager) GetRedisClient() (redis.UniversalClient, error) {
	mgr, err := lm.redisComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// GetNamedRedisClient 获取命名 Redis 实例的客户端
func (lm *launcherManager) GetNamedRedisClient(name string) (redis.UniversalClient, error) {
	mgr, err := lm.redisGroup.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// ==================== Provider 方法：MongoProvider ====================

// GetMongoClient 获取 MongoDB 客户端
func (lm *launcherManager) GetMongoClient() (*mongo.Client, error) {
	mgr, err := lm.mongoComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetMongoClient()
}

// GetNamedMongoClient 获取命名 MongoDB 实例的客户端
func (lm *launcherManager) GetNamedMongoClient(name string) (*mongo.Client, error) {
	mgr, err := lm.mongoGroup.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetMongoClient()
}

// ==================== Provider 方法：ConsulProvider ====================

// GetConsulClient 获取 Consul 客户端
func (lm *launcherManager) GetConsulClient() (*consulapi.Client, error) {
	mgr, err := lm.consulComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// GetNamedConsulClient 获取命名 Consul 实例的客户端
func (lm *launcherManager) GetNamedConsulClient(name string) (*consulapi.Client, error) {
	mgr, err := lm.consulGroup.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// ==================== Provider 方法：TracerProvider ====================

// GetJaegerExporter 获取 Jaeger 导出器
func (lm *launcherManager) GetJaegerExporter() (*otlptrace.Exporter, error) {
	mgr, err := lm.jaegerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetExporter()
}

// GetNamedJaegerExporter 获取命名 Jaeger 实例的导出器
func (lm *launcherManager) GetNamedJaegerExporter(name string) (*otlptrace.Exporter, error) {
	mgr, err := lm.jaegerGroup.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetExporter()
}

// ==================== Provider 方法：MessageQueueProvider ====================

// GetRabbitmqConn 获取 RabbitMQ 连接
func (lm *launcherManager) GetRabbitmqConn() (*rabbitmqpkg.ConnectionWrapper, error) {
	mgr, err := lm.rabbitmqComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// GetNamedRabbitmqConn 获取命名 RabbitMQ 实例的连接
func (lm *launcherManager) GetNamedRabbitmqConn(name string) (*rabbitmqpkg.ConnectionWrapper, error) {
	mgr, err := lm.rabbitmqGroup.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// ==================== Provider 方法：AuthProvider ====================

// GetTokenManager 获取 Token 管理器
func (lm *launcherManager) GetTokenManager() (authpkg.TokenManager, error) {
	authInstance, err := lm.authComp.Get()
	if err != nil {
		return nil, err
	}
	return authInstance.GetTokenManger()
}

// GetAuthManager 获取认证管理器
func (lm *launcherManager) GetAuthManager() (authpkg.AuthRepo, error) {
	authInstance, err := lm.authComp.Get()
	if err != nil {
		return nil, err
	}
	return authInstance.GetAuthManger()
}

// ==================== Provider 方法：ServiceAPIProvider ====================

// GetServiceApiManager 获取集群服务 API 管理器
func (lm *launcherManager) GetServiceApiManager() (clientutil.ServiceAPIManager, error) {
	mgr, err := lm.serviceAPIComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr, nil
}

// ==================== Closer ====================

// Close 关闭所有已初始化的组件，委托给 Lifecycle
func (lm *launcherManager) Close() error {
	return lm.lc.Close()
}
