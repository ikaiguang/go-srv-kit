# setup - LauncherManager 核心入口

`setup/` 是 go-srv-kit 服务层的核心，提供 `LauncherManager` 接口，统一管理所有基础设施组件的生命周期。

## 包名

```go
import setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
```

## 核心接口

`LauncherManager` 组合了所有 Provider 接口：

| Provider | 方法 | 说明 |
|----------|------|------|
| `ConfigProvider` | `GetConfig()` | 获取配置 |
| `LoggerProvider` | `GetLogger()` | 获取日志 |
| `DatabaseProvider` | `GetMysqlDBConn()` / `GetPostgresDBConn()` | 获取数据库连接 |
| `RedisProvider` | `GetRedisClient()` | 获取 Redis 客户端 |
| `MongoProvider` | `GetMongoClient()` | 获取 MongoDB 客户端 |
| `ConsulProvider` | `GetConsulClient()` | 获取 Consul 客户端 |
| `TracerProvider` | `GetJaegerExporter()` | 获取 Jaeger Exporter |
| `MessageQueueProvider` | `GetRabbitmqConn()` | 获取 RabbitMQ 连接 |
| `AuthProvider` | `GetTokenManager()` / `GetAuthManager()` | 获取认证管理器 |
| `ServiceAPIProvider` | `GetServiceApiManager()` | 获取集群服务 API 管理器 |

所有 `GetNamed*` 方法支持多实例（通过名称区分）。

## 使用方式

```go
// 方式一：完整创建（推荐）
lm, cleanup, err := setuputil.NewWithCleanup(configFilePath)
defer cleanup()

// 方式二：按需注册组件
lm, err := setuputil.New(conf,
    setuputil.WithRedis(),
    setuputil.WithMysql(),
)
defer lm.Close()
```

## 组件注册

通过 `WithXxx()` Option 按需注册组件，未注册的组件调用时会返回错误：

- `WithAllComponents()` - 注册所有组件（向后兼容）
- `WithRedis()` - 注册 Redis
- `WithMysql()` - 注册 MySQL
- `WithPostgres()` - 注册 PostgreSQL
- `WithMongo()` - 注册 MongoDB
- `WithConsul()` - 注册 Consul
- `WithJaeger()` - 注册 Jaeger
- `WithRabbitmq()` - 注册 RabbitMQ
- `WithAuth()` - 注册认证

## 设计原则

- **懒加载**：组件在首次调用 `Get*` 时才初始化
- **单例**：同一组件只初始化一次（`sync.Once`）
- **有序关闭**：`Close()` 按注册的逆序关闭所有组件
