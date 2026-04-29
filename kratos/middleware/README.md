# middleware - 中间件集合

`middleware/` 提供 HTTP/gRPC 服务的中间件集合。

## 包名

```go
import middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
```

## 默认中间件链

### 服务端

```go
middlewares := middlewarepkg.DefaultServerMiddlewares(logHelper)
```

按顺序：Recovery → Tracing → RateLimit → Metadata → Header → Log → Validator

### 客户端

```go
middlewares := middlewarepkg.DefaultClientMiddlewares(logHelper)
```

按顺序：Recovery → Tracing → CircuitBreaker → Metadata → Log

## 中间件列表

| 文件 | 中间件 | 说明 |
|------|--------|------|
| `middleware_recover.kit.go` | Recovery | panic 恢复，返回 500 错误 |
| `middleware_tracer.kit.go` | Tracer | OpenTelemetry 链路追踪 |
| `middleware_rate_limit.kit.go` | RateLimit | 限流 |
| `middleware_header.kit.go` | Header | 请求/响应头处理 |
| `middleware_jwt.kit.go` | JWT Auth | JWT 认证（白名单模式） |
| `middleware_validator.kit.go` | Validator | 参数验证 |
| `middleware_cors.kit.go` | CORS | 跨域资源共享 |
