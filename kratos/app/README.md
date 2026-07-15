# app

`app` 包的包名为 `apppkg`，提供 Kratos 应用运行环境、HTTP 响应模型、JSON/protobuf 编解码、请求/响应 encoder/decoder、404 handler 和服务端/客户端日志中间件。适合在 Kratos HTTP/gRPC 服务的启动装配、handler 编解码和调用链日志中使用。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/kratos
```

```go
import apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
```

## 核心能力

- `SetRuntimeEnv`、`GetRuntimeEnv`、`ParseEnv`、`IsDebugMode`：维护应用运行环境。
- `RegisterCodec`、`RequestDecoder`、`ContentSubtype`：注册和使用 `multipart/form-data`、`octet-stream` 等请求解码辅助。
- `MarshalJSON`、`UnmarshalJSON`、`SetJSONMarshalOptions`、`SetJSONUnmarshalOptions`：统一标准 JSON 和 protobuf JSON 行为。
- `SuccessResponseEncoder`、`ErrorResponseEncoder`、`SuccessResponseDecoderBody`、`ErrorResponseDecodeBody`：HTTP 响应编码和响应体解码。
- `HTTPResponse`、`ToResponseError`、`IsSuccessCode`：通用响应结构和错误转换。
- `ServerLog`、`ClientLog`、`ServerLoggingKvs`、`ClientLoggingKvs`：Kratos 中间件日志。
- `NotFound404`：生成 Kratos HTTP 404 handler option。

## 快速使用

```go
apppkg.SetRuntimeEnv(apppkg.ParseEnv("develop"))
apppkg.RegisterCodec()

if apppkg.IsDebugMode() {
    // enable debug-only behavior
}
```

```go
srv := http.NewServer(
    http.RequestDecoder(apppkg.RequestDecoder),
    http.ResponseEncoder(apppkg.SuccessResponseEncoder),
    http.ErrorEncoder(apppkg.ErrorResponseEncoder),
    apppkg.NotFound404(),
)
_ = srv
```

## 测试

```bash
go test ./app
```

## 注意事项

- `app.kit.pb.go` 由 `app.kit.proto` 生成，不要手工修改。
- `RequestDecoder` 会读取并重置 `r.Body`，并限制请求体大小为 `DefaultUploadMaxSize`。
- `multipart/form-data` 当前注册为跳过 body 自动反序列化，文件上传应在业务 handler 中按 Kratos/标准库方式处理。
- `app_response_custom.kit.go` 中的旧自定义响应实现均处于注释状态，不应在文档和新代码中当作可用 API。
