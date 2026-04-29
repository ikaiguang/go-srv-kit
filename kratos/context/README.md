# context - Context 工具

`context/` 提供从 Context 中提取 HTTP/gRPC Transport 信息的工具函数。

## 包名

```go
import contextpkg "github.com/ikaiguang/go-srv-kit/kratos/context"
```

## 使用

```go
// 判断是 HTTP 还是 gRPC 请求
httpTr, ok := contextpkg.MatchHTTPServerContext(ctx)
grpcTr, ok := contextpkg.MatchGRPCServerContext(ctx)

// 获取 Transport
tr, ok := contextpkg.FromServerContext(ctx)

// 创建新 Context（保留 Span 信息）
newCtx := contextpkg.NewContext(ctx)

// 获取 SpanContext
spanCtx := contextpkg.SpanContext(ctx)
```
