# middleware

`middleware` 包的包名为 `middlewarepkg`，提供 Kratos 服务端和客户端常用中间件组合，以及限流、恢复、请求 ID、参数校验、CORS、白名单匹配和 tracing provider 设置能力。适合在服务启动装配阶段统一复用。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/kratos/v3
```

```go
import middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/v3/middleware"
```

## 核心能力

- `DefaultServerMiddlewares`：组合 recovery、tracing、rate limit、metadata、request id、server log 和 validator。
- `DefaultClientMiddlewares`：组合 recovery、tracing、circuit breaker、metadata 和 client log。
- `RateLimit`、`WithLimiter`：默认使用 BBR limiter，也可传入自定义 limiter。
- `RecoveryHandler`：将 panic 转换为 `errorpkg.ErrorPanic`。
- `RequestAndResponseHeader`：保证 request/reply header 中存在 request id。
- `Validator`：调用请求对象的 `Validate() error`。
- `NewCORS`：返回 gorilla handlers CORS filter。
- `NewWhiteListMatcher`：为 Kratos selector 构造 operation 白名单 matcher。
- `SetTracer`、`SetTracerProvider`：设置 OpenTelemetry tracer provider。

## 快速使用

```go
logger, err := logpkg.NewStdLogger(&logpkg.ConfigStd{Level: log.LevelInfo})
if err != nil {
    return err
}
defer logger.Close()

middlewares := middlewarepkg.DefaultServerMiddlewares(log.NewHelper(logger))
_ = middlewares
```

```go
whiteList := map[string]struct{}{
    "/healthz": {},
}
matcher := middlewarepkg.NewWhiteListMatcher(whiteList)
_ = matcher
```

## 测试

```bash
go test ./middleware
```

## 注意事项

- `NewCORS` 当前允许所有 origin 且启用 credentials，公开服务上线前应按实际域名收敛。
- `RateLimit` 默认使用 BBR limiter；如果服务需要固定速率或分租户限流，应传入自定义 limiter。
- `RecoveryHandler` 不暴露 panic 详情，排查需要结合日志和 trace。
- `Validator` 只处理实现了 `Validate() error` 的请求对象。
