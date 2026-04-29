# client - 客户端封装

`client/` 提供 gRPC 和 HTTP 客户端的便捷创建方法。

## 包名

```go
import clientpkg "github.com/ikaiguang/go-srv-kit/kratos/client"
```

## 使用

```go
// HTTP 客户端
httpClient, err := clientpkg.NewHTTPClient(ctx, opts...)

// gRPC 客户端
grpcConn, err := clientpkg.NewGRPCClient(ctx, insecure, opts...)

// 简易 HTTP 客户端（内置 Recovery 中间件和自定义响应解码）
httpClient, err := clientpkg.NewSampleHTTPClient(ctx, "http://localhost:10101")
```

`ResponseDecoder` 会自动解析统一响应格式（`apppkg.Response`），支持 Proto 和 JSON 两种数据格式。
