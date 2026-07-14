# client

`client` 包的包名为 `clientpkg`，封装 Kratos HTTP/gRPC 客户端创建，并提供适配 `apppkg.Response` 的响应解码器。适合在服务间调用或 SDK 代码中复用基础客户端配置。

## 安装

```bash
go get github.com/ikaiguang/go-kratos-kit
```

```go
import clientpkg "github.com/ikaiguang/go-kratos-kit/client"
```

## 核心能力

- `NewHTTPClient`：直接调用 Kratos `http.NewClient`。
- `NewGRPCClient`：按 `insecure` 参数选择 `grpc.DialInsecure` 或 `grpc.Dial`。
- `NewSampleHTTPClient`：创建带 recovery、endpoint 和自定义 `ResponseDecoder` 的 HTTP client。
- `ResponseDecoder`：解码 `apppkg.Response.Data`，目标可以是 protobuf message 或普通 JSON 结构。

## 快速使用

```go
ctx := context.Background()

cli, err := clientpkg.NewSampleHTTPClient(ctx, "http://127.0.0.1:8000")
if err != nil {
    return err
}
_ = cli
```

```go
conn, err := clientpkg.NewGRPCClient(ctx, true)
if err != nil {
    return err
}
defer conn.Close()
```

## 测试

```bash
go test ./client
```

## 注意事项

- `ResponseDecoder` 会关闭 `res.Body`。
- 普通结构体解码依赖 `apppkg.ResponseData.Data` 中保存的 JSON 字符串。
- gRPC 使用安全连接还是非安全连接由调用方通过 `insecure` 参数明确选择。
