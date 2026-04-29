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

// componentNotRegisteredError 组件未注册错误
func componentNotRegisteredError(name string) error {
	e := errorpkg.ErrorBadRequest("component not registered: %s; use corresponding WithXxx() option", name)
	return errorpkg.WithStack(e)
}

// WithAllComponents 注册所有组件（向后兼容）
func WithAllComponents() Option {
	return WithComponentRegistrar(func(lm *launcherManager) {
		registerRedis(lm)
		registerMysql(lm)
		registerPostgres(lm)
		registerMongo(lm)
		registerConsul(lm)
		registerJaeger(lm)
		registerRabbitmq(lm)
		registerAuth(lm)
		registerServiceAPI(lm)
	})
}

// ==================== 组件注册函数（供 WithAllComponents 和各 WithXxx 使用） ====================

func registerRedis(lm *launcherManager) {
	RegisterComponent(lm.registry, ComponentRedis, lm.newRedisManager, lm.lc)
	RegisterComponentGroup(lm.registry, ComponentRedis, lm.newNamedRedisManager, lm.lc)
}

func registerMysql(lm *launcherManager) {
	RegisterComponent(lm.registry, ComponentMysql, lm.newMysqlManager, lm.lc)
	RegisterComponentGroup(lm.registry, ComponentMysql, lm.newNamedMysqlManager, lm.lc)
}

func registerPostgres(lm *launcherManager) {
	RegisterComponent(lm.registry, ComponentPostgres, lm.newPostgresManager, lm.lc)
	RegisterComponentGroup(lm.registry, ComponentPostgres, lm.newNamedPostgresManager, lm.lc)
}

func registerMongo(lm *launcherManager) {
	RegisterComponent(lm.registry, ComponentMongo, lm.newMongoManager, lm.lc)
	RegisterComponentGroup(lm.registry, ComponentMongo, lm.newNamedMongoManager, lm.lc)
}

func registerConsul(lm *launcherManager) {
	RegisterComponent(lm.registry, ComponentConsul, lm.newConsulManager, lm.lc)
	RegisterComponentGroup(lm.registry, ComponentConsul, lm.newNamedConsulManager, lm.lc)
}

func registerJaeger(lm *launcherManager) {
	RegisterComponent(lm.registry, ComponentJaeger, lm.newJaegerManager, lm.lc)
	RegisterComponentGroup(lm.registry, ComponentJaeger, lm.newNamedJaegerManager, lm.lc)
}

func registerRabbitmq(lm *launcherManager) {
	RegisterComponent(lm.registry, ComponentRabbitmq, lm.newRabbitmqManager, lm.lc)
	RegisterComponentGroup(lm.registry, ComponentRabbitmq, lm.newNamedRabbitmqManager, lm.lc)
}

func registerAuth(lm *launcherManager) {
	RegisterComponent(lm.registry, ComponentAuth, lm.newAuthInstance, lm.lc)
}

func registerServiceAPI(lm *launcherManager) {
	RegisterComponent(lm.registry, ComponentServiceAPI, lm.newServiceAPIManager, lm.lc)
}

// ==================== 单实例 factory 方法 ====================

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
	var opts = []clientutil.Option{
		clientutil.WithLogger(loggerForMiddleware),
		clientutil.WithSkipRegistryCheck(),
	}
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

func (lm *launcherManager) GetConfig() *configpb.Bootstrap {
	return lm.conf
}

// ==================== Provider 方法：LoggerProvider ====================

func (lm *launcherManager) GetLogger() (log.Logger, error) {
	mgr, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetLogger()
}

func (lm *launcherManager) GetLoggerForMiddleware() (log.Logger, error) {
	mgr, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetLoggerForMiddleware()
}

func (lm *launcherManager) GetLoggerForHelper() (log.Logger, error) {
	mgr, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetLoggerForHelper()
}

// ==================== Provider 方法：DatabaseProvider ====================

func (lm *launcherManager) GetMysqlDBConn() (*gorm.DB, error) {
	comp, ok := GetComponent[mysqlutil.MysqlManager](lm.registry, ComponentMysql)
	if !ok {
		return nil, componentNotRegisteredError(ComponentMysql)
	}
	mgr, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

func (lm *launcherManager) GetPostgresDBConn() (*gorm.DB, error) {
	comp, ok := GetComponent[postgresutil.PostgresManager](lm.registry, ComponentPostgres)
	if !ok {
		return nil, componentNotRegisteredError(ComponentPostgres)
	}
	mgr, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

func (lm *launcherManager) GetNamedMysqlDBConn(name string) (*gorm.DB, error) {
	group, ok := GetComponentGroup[mysqlutil.MysqlManager](lm.registry, ComponentMysql)
	if !ok {
		return nil, componentNotRegisteredError(ComponentMysql)
	}
	mgr, err := group.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

func (lm *launcherManager) GetNamedPostgresDBConn(name string) (*gorm.DB, error) {
	group, ok := GetComponentGroup[postgresutil.PostgresManager](lm.registry, ComponentPostgres)
	if !ok {
		return nil, componentNotRegisteredError(ComponentPostgres)
	}
	mgr, err := group.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

// ==================== Provider 方法：RedisProvider ====================

func (lm *launcherManager) GetRedisClient() (redis.UniversalClient, error) {
	comp, ok := GetComponent[redisutil.RedisManager](lm.registry, ComponentRedis)
	if !ok {
		return nil, componentNotRegisteredError(ComponentRedis)
	}
	mgr, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

func (lm *launcherManager) GetNamedRedisClient(name string) (redis.UniversalClient, error) {
	group, ok := GetComponentGroup[redisutil.RedisManager](lm.registry, ComponentRedis)
	if !ok {
		return nil, componentNotRegisteredError(ComponentRedis)
	}
	mgr, err := group.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// ==================== Provider 方法：MongoProvider ====================

func (lm *launcherManager) GetMongoClient() (*mongo.Client, error) {
	comp, ok := GetComponent[mongoutil.MongoManager](lm.registry, ComponentMongo)
	if !ok {
		return nil, componentNotRegisteredError(ComponentMongo)
	}
	mgr, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetMongoClient()
}

func (lm *launcherManager) GetNamedMongoClient(name string) (*mongo.Client, error) {
	group, ok := GetComponentGroup[mongoutil.MongoManager](lm.registry, ComponentMongo)
	if !ok {
		return nil, componentNotRegisteredError(ComponentMongo)
	}
	mgr, err := group.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetMongoClient()
}

// ==================== Provider 方法：ConsulProvider ====================

func (lm *launcherManager) GetConsulClient() (*consulapi.Client, error) {
	comp, ok := GetComponent[consulutil.ConsulManager](lm.registry, ComponentConsul)
	if !ok {
		return nil, componentNotRegisteredError(ComponentConsul)
	}
	mgr, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

func (lm *launcherManager) GetNamedConsulClient(name string) (*consulapi.Client, error) {
	group, ok := GetComponentGroup[consulutil.ConsulManager](lm.registry, ComponentConsul)
	if !ok {
		return nil, componentNotRegisteredError(ComponentConsul)
	}
	mgr, err := group.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// ==================== Provider 方法：TracerProvider ====================

func (lm *launcherManager) GetJaegerExporter() (*otlptrace.Exporter, error) {
	comp, ok := GetComponent[jaegerutil.JaegerManager](lm.registry, ComponentJaeger)
	if !ok {
		return nil, componentNotRegisteredError(ComponentJaeger)
	}
	mgr, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetExporter()
}

func (lm *launcherManager) GetNamedJaegerExporter(name string) (*otlptrace.Exporter, error) {
	group, ok := GetComponentGroup[jaegerutil.JaegerManager](lm.registry, ComponentJaeger)
	if !ok {
		return nil, componentNotRegisteredError(ComponentJaeger)
	}
	mgr, err := group.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetExporter()
}

// ==================== Provider 方法：MessageQueueProvider ====================

func (lm *launcherManager) GetRabbitmqConn() (*rabbitmqpkg.ConnectionWrapper, error) {
	comp, ok := GetComponent[rabbitmqutil.RabbitmqManager](lm.registry, ComponentRabbitmq)
	if !ok {
		return nil, componentNotRegisteredError(ComponentRabbitmq)
	}
	mgr, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

func (lm *launcherManager) GetNamedRabbitmqConn(name string) (*rabbitmqpkg.ConnectionWrapper, error) {
	group, ok := GetComponentGroup[rabbitmqutil.RabbitmqManager](lm.registry, ComponentRabbitmq)
	if !ok {
		return nil, componentNotRegisteredError(ComponentRabbitmq)
	}
	mgr, err := group.Get(name)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// ==================== Provider 方法：AuthProvider ====================

func (lm *launcherManager) GetTokenManager() (authpkg.TokenManager, error) {
	comp, ok := GetComponent[authutil.AuthInstance](lm.registry, ComponentAuth)
	if !ok {
		return nil, componentNotRegisteredError(ComponentAuth)
	}
	authInstance, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return authInstance.GetTokenManger()
}

func (lm *launcherManager) GetAuthManager() (authpkg.AuthRepo, error) {
	comp, ok := GetComponent[authutil.AuthInstance](lm.registry, ComponentAuth)
	if !ok {
		return nil, componentNotRegisteredError(ComponentAuth)
	}
	authInstance, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return authInstance.GetAuthManger()
}

// ==================== Provider 方法：ServiceAPIProvider ====================

func (lm *launcherManager) GetServiceApiManager() (clientutil.ServiceAPIManager, error) {
	comp, ok := GetComponent[clientutil.ServiceAPIManager](lm.registry, ComponentServiceAPI)
	if !ok {
		return nil, componentNotRegisteredError(ComponentServiceAPI)
	}
	mgr, err := comp.Get()
	if err != nil {
		return nil, err
	}
	return mgr, nil
}

// ==================== Closer ====================

func (lm *launcherManager) Close() error {
	return lm.lc.Close()
}

// ==================== 辅助方法（供 loggerutil 使用） ====================

func (lm *launcherManager) getLoggerManager() (loggerutil.LoggerManager, error) {
	return lm.loggerComp.Get()
}
