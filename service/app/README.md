# app - 应用工具

`app/` 提供应用标识生成、环境信息和 HTTP 编解码器配置。

## 包名

```go
import apputil "github.com/ikaiguang/go-srv-kit/service/app"
```

## 核心功能

### 应用标识

```go
appConfig := apputil.ToAppConfig(conf.GetApp())

// 生成应用 ID（Redis 风格）
// 例：go-srv-saas:ping-service:develop:v1.0.0
id := apputil.ID(appConfig)

// 生成路径风格标识
// 例：go-srv-saas/ping-service/develop/v1.0.0
path := apputil.Path(appConfig)
```

### HTTP 编解码器

配置统一的请求解码和响应编码：

```go
opts := apputil.ServerDecoderEncoder()  // 服务端
opts := apputil.ClientDecoderEncoder()  // 客户端
```
