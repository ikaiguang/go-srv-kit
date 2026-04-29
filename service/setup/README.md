# setup - LauncherManager 核心入口

`setup/` 是 go-srv-kit 服务启动的核心包，负责配置、日志、组件注册表、生命周期和资源关闭。

核心包不再直接 import Redis、MySQL、PostgreSQL、MongoDB、Consul、Jaeger、RabbitMQ、Auth 或 cluster service API 等基础设施实现。业务服务需要什么组件，就在入口文件显式 import 对应 `service/<component>` 包，并传入该包提供的 `WithSetup()`。

## 包名

```go
import setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
```

## 核心接口

`LauncherManager` 只包含核心能力：

| Provider | 方法 | 说明 |
|---|---|---|
| `ConfigProvider` | `GetConfig()` | 获取启动配置 |
| `LoggerProvider` | `GetLogger()` / `GetLoggerForMiddleware()` / `GetLoggerForHelper()` | 获取日志 |
| `RegistryProvider` | `GetRegistry()` | 获取组件注册表 |
| `LifecycleProvider` | `GetLifecycle()` | 获取生命周期管理器 |
| `Closer` | `Close()` | 按注册顺序关闭资源 |

具体基础设施由各组件包注册和读取，例如：

- `redisutil.WithSetup()` + `redisutil.GetClient(launcher)`
- `postgresutil.WithSetup()` + `postgresutil.GetDB(launcher)`
- `clientutil.WithSetup()` + `clientutil.GetManager(launcher)`

## 使用方式

### 核心启动

只加载配置和日志，不注册额外基础设施：

```go
lm, cleanup, err := setuputil.NewWithCleanup(configFilePath)
if err != nil {
    return err
}
defer cleanup()
```

### 按需注册组件

```go
import (
    configutil "github.com/ikaiguang/go-srv-kit/service/config"
    postgresutil "github.com/ikaiguang/go-srv-kit/service/postgres"
    redisutil "github.com/ikaiguang/go-srv-kit/service/redis"
    setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

conf, err := configutil.Loading(configFilePath)
if err != nil {
    return err
}

lm, err := setuputil.New(conf,
    postgresutil.WithSetup(),
    redisutil.WithSetup(),
)
if err != nil {
    return err
}
defer lm.Close()

db, err := postgresutil.GetDB(lm)
redisClient, err := redisutil.GetClient(lm)
```

### 加载配置并按需注册组件

```go
lm, cleanup, err := setuputil.NewWithCleanupOptions(
    configFilePath,
    nil,
    postgresutil.WithSetup(),
    redisutil.WithSetup(),
)
if err != nil {
    return err
}
defer cleanup()
```

## 组件注册

组件注册由对应子包提供：

| 组件 | 注册方式 | 获取方式 |
|---|---|---|
| Redis | `redisutil.WithSetup()` | `redisutil.GetClient(lm)` / `redisutil.GetNamedClient(lm, name)` |
| MySQL | `mysqlutil.WithSetup()` | `mysqlutil.GetDB(lm)` / `mysqlutil.GetNamedDB(lm, name)` |
| PostgreSQL | `postgresutil.WithSetup()` | `postgresutil.GetDB(lm)` / `postgresutil.GetNamedDB(lm, name)` |
| MongoDB | `mongoutil.WithSetup()` | `mongoutil.GetClient(lm)` / `mongoutil.GetNamedClient(lm, name)` |
| Consul | `consulutil.WithSetup()` | `consulutil.GetClient(lm)` / `consulutil.GetNamedClient(lm, name)` |
| Jaeger | `jaegerutil.WithSetup()` | `jaegerutil.GetExporter(lm)` / `jaegerutil.GetNamedExporter(lm, name)` |
| RabbitMQ | `rabbitmqutil.WithSetup()` | `rabbitmqutil.GetConn(lm)` / `rabbitmqutil.GetNamedConn(lm, name)` |
| Auth | `authutil.WithSetup()` | `authutil.GetTokenManager(lm)` / `authutil.GetAuthManager(lm)` |
| Cluster service API | `clientutil.WithSetup()` | `clientutil.GetManager(lm)` |

`setuputil.WithComponentRegistrar` 是组件包内部使用的注册扩展点，业务入口优先使用对应组件包的 `WithSetup()`。

## 急切初始化

默认情况下组件懒加载，在首次 `GetXxx()` 时初始化。需要启动时立即验证连接时，可以配合 `WithEagerInit`：

```go
lm, err := setuputil.New(conf,
    postgresutil.WithSetup(),
    redisutil.WithSetup(),
    setuputil.WithEagerInit(
        setuputil.ComponentPostgres,
        setuputil.ComponentRedis,
    ),
)
```

## 设计原则

- **按需依赖**：核心 setup 包不隐式拉入所有基础设施组件。
- **显式注册**：业务入口通过 `service/<component>.WithSetup()` 声明所需组件。
- **懒加载**：组件在首次使用时才初始化。
- **有序关闭**：`Close()` 按注册的逆序关闭资源。
