# transport - Transport 工具

`transport/` 提供从 Context 中获取 Transport 类型的工具。

## 包名

```go
import transportpkg "github.com/ikaiguang/go-srv-kit/kratos/transport"
```

## 使用

```go
kind, ok := transportpkg.TransportKind(ctx)
// kind: "http" 或 "grpc"
```
