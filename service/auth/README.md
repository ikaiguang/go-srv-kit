# auth - 认证管理

`auth/` 提供认证实例的懒加载管理，封装 Token 和 Auth 的初始化逻辑。

## 包名

```go
import authutil "github.com/ikaiguang/go-srv-kit/service/auth"
```

## 使用

```go
authInstance, err := authutil.NewAuthInstance(encryptConfig, redisClient, loggerManager)

// 获取 Token 管理器
tokenManager, err := authInstance.GetTokenManger()

// 获取 Auth 管理器
authManager, err := authInstance.GetAuthManger()
```

## 依赖

- Redis 客户端（Token 存储和分布式锁）
- 加密配置（`encrypt.token_encrypt`）
- 日志管理器

## 与 kratos/auth 的关系

- `kratos/auth/` 提供 JWT 中间件和 Token 操作的底层实现
- `service/auth/` 负责将配置、Redis、日志组装为可用的认证实例
