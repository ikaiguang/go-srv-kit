# context

`context` 包的包名为 `contextpkg`，提供 Kratos transport context 匹配、客户端 IP 提取、trace span context 传递和自定义 IP 写入读取能力。适合在 middleware、service handler 和日志/审计代码中使用。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/kratos
```

```go
import contextpkg "github.com/ikaiguang/go-srv-kit/kratos/context"
```

## 核心能力

- `FromServerContext`、`FromClientContext`：转发 Kratos transport context 读取能力。
- `MatchHTTPServerContext`、`MatchGRPCServerContext`：从 server context 匹配 HTTP/gRPC transport。
- `ToHTTPTransporter`、`ToGRPCTransporter`：将 Kratos transport 转为具体 HTTP/gRPC 类型。
- `ClientIP`、`ClientIPFromHTTP`、`ClientIPFromGRPC`：从 HTTP header、gRPC metadata 或 peer 信息提取客户端 IP。
- `SetTrustedPlatform`：设置优先读取的可信平台 header。
- `SetClientIpToContext`、`GetClientIpFromContext`：在 context 中保存和读取客户端 IP。
- `NewContext`、`SpanContext`：保留 trace span 信息创建新 context 或读取 span context。

## 快速使用

```go
clientIP := contextpkg.ClientIP(ctx)
ctx = contextpkg.SetClientIpToContext(ctx, clientIP)
```

```go
if tr, ok := contextpkg.MatchHTTPServerContext(ctx); ok {
    path := tr.Request().URL.Path
    _ = path
}
```

## 测试

```bash
go test ./context
```

## 注意事项

- 只有在网关或反向代理可信时，才应信任 `X-Forwarded-For`、`X-Real-Ip` 或 `SetTrustedPlatform` 设置的 header。
- `MatchHTTPContext` 已标记 deprecated，新代码优先使用 `MatchHTTPServerContext` 或 `MatchGRPCServerContext`。
- `NewContext` 会从原 context 保留 trace span，但以 `context.Background()` 为基础创建新 context，其他 value 和 cancel/deadline 不会自动继承。
