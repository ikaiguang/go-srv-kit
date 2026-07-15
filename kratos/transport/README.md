# transport

`transport` 包的包名为 `transportpkg`，提供从 Kratos server context 中读取 transport kind 的轻量工具。适合在日志、审计或分支处理逻辑中判断当前请求来自 HTTP 还是 gRPC。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/kratos
```

```go
import transportpkg "github.com/ikaiguang/go-srv-kit/kratos/transport"
```

## 核心能力

- `TransportKind`：从 `transport.FromServerContext` 读取 `transport.Kind`。

## 快速使用

```go
kind, ok := transportpkg.TransportKind(ctx)
if ok {
    _ = kind.String()
}
```

## 测试

```bash
go test ./transport
```

## 注意事项

- 只有 Kratos server context 中存在 transport 信息时，`ok` 才会为 `true`。
- 如果需要访问 header、operation 或 endpoint 等更完整信息，可使用 `context` 包中的 transport 匹配辅助。
