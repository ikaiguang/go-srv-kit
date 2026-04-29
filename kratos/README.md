# kratos - 框架扩展

`kratos/` 基于 [go-kratos v2](https://github.com/go-kratos/kratos) 框架，提供认证、中间件、错误处理、日志等扩展能力。

这是一个独立的 Go 模块（有自己的 `go.mod`），业务服务通过 `service/` 层间接使用，也可直接引入。

## 模块导入

```go
import "github.com/ikaiguang/go-srv-kit/kratos"
```

## 目录结构

| 目录 | 包名 | 说明 |
|------|------|------|
| `auth/` | `authpkg` | JWT Token 管理、签发、验证、刷新，支持多种 Token 类型 |
| `middleware/` | `middlewarepkg` | 中间件集合：CORS、Recovery、Tracer、Auth、限流、Validator 等 |
| `error/` | `errorpkg` | 统一错误处理，带堆栈信息和元数据 |
| `log/` | `logpkg` | 结构化日志（基于 Zap），支持控制台和文件输出 |
| `app/` | `apppkg` | 应用运行时环境、HTTP 请求/响应编解码、统一响应格式 |
| `client/` | `clientpkg` | gRPC/HTTP 客户端封装 |
| `context/` | `contextpkg` | Context 工具，提取 HTTP/gRPC Transport 信息 |
| `debug/` | `debugpkg` | 调试日志工具，支持多级别输出 |
| `pprof/` | `pprofpkg` | 性能分析路由注册 |
| `registry/` | `registrypkg` | Consul 服务注册与发现 |
| `registry_etcd/` | `etcdregistry` | Etcd 服务注册与发现 |
| `transport/` | `transportpkg` | Transport 类型判断工具 |
| `websocket/` | `websocketpkg` | WebSocket 连接升级 |

## 核心组件

### 认证 (auth/)

JWT 认证中间件，支持多种 Token 类型（USER、ADMIN、EMPLOYEE）：

```go
import authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"

// 服务端中间件
authMiddleware := authpkg.Server(signKeyFunc, authpkg.WithClaims(claimsFunc))

// 从 Context 获取用户信息
claims, ok := authpkg.GetAuthClaimsFromContext(ctx)
userID := claims.Payload.UserID
```

### 错误处理 (error/)

统一错误格式，带堆栈信息和元数据：

```go
import errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"

errorpkg.ErrorBadRequest("invalid parameter")      // 400
errorpkg.ErrorNotFound("user not found")            // 404
errorpkg.ErrorInternal("database connection failed") // 500

// 包装错误（附加堆栈）
errorpkg.FormatError(err)
errorpkg.Wrap(kratosErr, originalErr)
```

### 日志 (log/)

基于 Zap 的结构化日志，支持 Context 传递 TraceID：

```go
import logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"

logpkg.WithContext(ctx).Infow("user created", "user_id", 123)
logpkg.WithContext(ctx).Errorw("operation failed", "error", err.Error())
```

### 中间件 (middleware/)

默认中间件链（按顺序）：

1. **Recovery** - panic 恢复
2. **Tracing** - OpenTelemetry 链路追踪
3. **RateLimit** - 限流
4. **Metadata** - 元数据传递
5. **Header** - 请求/响应头处理
6. **Log** - 请求日志
7. **Validator** - 参数验证

## 参考

- go-kratos 文档：https://go-kratos.dev/
- Uber Go Style Guide：https://github.com/uber-go/guide
