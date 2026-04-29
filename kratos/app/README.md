# app - 应用运行时

`app/` 提供应用运行时环境管理、HTTP 请求/响应编解码和统一响应格式。

## 包名

```go
import apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
```

## 运行时环境

```go
apppkg.SetRuntimeEnv(apppkg.RuntimeEnvEnum_DEVELOP)

if apppkg.IsDebugMode() {
    // LOCAL、DEVELOP、TESTING 环境
}
```

| 环境 | 说明 |
|------|------|
| LOCAL | 本地开发 |
| DEVELOP | 开发环境 |
| TESTING | 测试环境 |
| PREVIEW | 预发布环境 |
| PRODUCTION | 生产环境 |

## HTTP 编解码

自定义请求解码和统一响应格式：

```go
// 注册自定义编解码器
apppkg.RegisterCodec()

// 请求解码器
http.RequestDecoder(apppkg.RequestDecoder)

// 成功响应编码器
http.ResponseEncoder(apppkg.SuccessResponseEncoder)

// 错误响应编码器
http.ErrorEncoder(apppkg.ErrorResponseEncoder)
```

## Proto 定义

```bash
kratos proto client ./kratos/app/app.kit.proto
```
