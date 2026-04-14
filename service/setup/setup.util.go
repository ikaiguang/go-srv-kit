package setuputil

import (
	stdlog "log"

	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	debugpkg "github.com/ikaiguang/go-srv-kit/debug"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	authutil "github.com/ikaiguang/go-srv-kit/service/auth"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	consulutil "github.com/ikaiguang/go-srv-kit/service/consul"
	jaegerutil "github.com/ikaiguang/go-srv-kit/service/jaeger"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	mongoutil "github.com/ikaiguang/go-srv-kit/service/mongo"
	mysqlutil "github.com/ikaiguang/go-srv-kit/service/mysql"
	postgresutil "github.com/ikaiguang/go-srv-kit/service/postgres"
	rabbitmqutil "github.com/ikaiguang/go-srv-kit/service/rabbitmq"
	redisutil "github.com/ikaiguang/go-srv-kit/service/redis"
)

// launcherManager 实现 LauncherManager 接口
type launcherManager struct {
	conf *configpb.Bootstrap
	lc   *Lifecycle

	// 单实例组件
	loggerComp     *Component[loggerutil.LoggerManager]
	redisComp      *Component[redisutil.RedisManager]
	mysqlComp      *Component[mysqlutil.MysqlManager]
	postgresComp   *Component[postgresutil.PostgresManager]
	mongoComp      *Component[mongoutil.MongoManager]
	consulComp     *Component[consulutil.ConsulManager]
	jaegerComp     *Component[jaegerutil.JaegerManager]
	rabbitmqComp   *Component[rabbitmqutil.RabbitmqManager]
	authComp       *Component[authutil.AuthInstance]
	serviceAPIComp *Component[clientutil.ServiceAPIManager]

	// 命名实例组
	mysqlGroup    *ComponentGroup[mysqlutil.MysqlManager]
	postgresGroup *ComponentGroup[postgresutil.PostgresManager]
	redisGroup    *ComponentGroup[redisutil.RedisManager]
	mongoGroup    *ComponentGroup[mongoutil.MongoManager]
	consulGroup   *ComponentGroup[consulutil.ConsulManager]
	jaegerGroup   *ComponentGroup[jaegerutil.JaegerManager]
	rabbitmqGroup *ComponentGroup[rabbitmqutil.RabbitmqManager]
}

// New 创建 LauncherManager，纯懒加载模式
func New(conf *configpb.Bootstrap, opts ...Option) (LauncherManager, error) {
	if conf == nil {
		return nil, errorpkg.ErrorBadRequest("bootstrap config is required")
	}

	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	lc := newLifecycle()
	lm := &launcherManager{
		conf: conf,
		lc:   lc,
	}

	// 注册单实例组件 factory
	lm.loggerComp = NewComponent(ComponentLogger, lm.newLoggerManager, lc)
	lm.redisComp = NewComponent(ComponentRedis, lm.newRedisManager, lc)
	lm.mysqlComp = NewComponent(ComponentMysql, lm.newMysqlManager, lc)
	lm.postgresComp = NewComponent(ComponentPostgres, lm.newPostgresManager, lc)
	lm.mongoComp = NewComponent(ComponentMongo, lm.newMongoManager, lc)
	lm.consulComp = NewComponent(ComponentConsul, lm.newConsulManager, lc)
	lm.jaegerComp = NewComponent(ComponentJaeger, lm.newJaegerManager, lc)
	lm.rabbitmqComp = NewComponent(ComponentRabbitmq, lm.newRabbitmqManager, lc)
	lm.authComp = NewComponent(ComponentAuth, lm.newAuthInstance, lc)
	lm.serviceAPIComp = NewComponent(ComponentServiceAPI, lm.newServiceAPIManager, lc)

	// 注册命名实例组 factory
	lm.mysqlGroup = NewComponentGroup(ComponentMysql, lm.newNamedMysqlManager, lc)
	lm.postgresGroup = NewComponentGroup(ComponentPostgres, lm.newNamedPostgresManager, lc)
	lm.redisGroup = NewComponentGroup(ComponentRedis, lm.newNamedRedisManager, lc)
	lm.mongoGroup = NewComponentGroup(ComponentMongo, lm.newNamedMongoManager, lc)
	lm.consulGroup = NewComponentGroup(ComponentConsul, lm.newNamedConsulManager, lc)
	lm.jaegerGroup = NewComponentGroup(ComponentJaeger, lm.newNamedJaegerManager, lc)
	lm.rabbitmqGroup = NewComponentGroup(ComponentRabbitmq, lm.newNamedRabbitmqManager, lc)

	// 日志始终初始化（其他组件依赖日志）
	loggerManager, err := lm.loggerComp.Get()
	if err != nil {
		return nil, err
	}
	loggerForHelper, err := loggerManager.GetLoggerForHelper()
	if err != nil {
		return nil, err
	}
	logpkg.Setup(loggerForHelper)
	debugpkg.Setup(loggerForHelper)

	// 按需急切初始化指定组件
	if err := lm.eagerInit(o.eagerComponents); err != nil {
		return nil, err
	}

	return lm, nil
}

// eagerInit 急切初始化指定的组件
func (lm *launcherManager) eagerInit(components []string) error {
	initMap := map[string]func() error{
		ComponentRedis:    func() error { _, err := lm.redisComp.Get(); return err },
		ComponentMysql:    func() error { _, err := lm.mysqlComp.Get(); return err },
		ComponentPostgres: func() error { _, err := lm.postgresComp.Get(); return err },
		ComponentMongo:    func() error { _, err := lm.mongoComp.Get(); return err },
		ComponentConsul:   func() error { _, err := lm.consulComp.Get(); return err },
		ComponentJaeger:   func() error { _, err := lm.jaegerComp.Get(); return err },
		ComponentRabbitmq: func() error { _, err := lm.rabbitmqComp.Get(); return err },
		ComponentAuth:     func() error { _, err := lm.authComp.Get(); return err },
	}

	for _, name := range components {
		initFn, ok := initMap[name]
		if !ok {
			return errorpkg.ErrorBadRequest("unknown component: %s", name)
		}
		if err := initFn(); err != nil {
			return err
		}
	}
	return nil
}

// NewWithCleanup 便捷函数，加载配置并创建 LauncherManager
func NewWithCleanup(configFilePath string, configOpts ...configutil.Option) (LauncherManager, func(), error) {
	conf, err := configutil.Loading(configFilePath, configOpts...)
	if err != nil {
		return nil, nil, err
	}
	apputil.SetConfig(conf)

	lm, err := New(conf)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		closeErr := lm.Close()
		if closeErr != nil {
			stdlog.Printf("==> launcherManager.Close failed: %+v\n", closeErr)
		}
	}
	return lm, cleanup, nil
}
