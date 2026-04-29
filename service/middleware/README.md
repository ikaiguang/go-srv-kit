# middleware - 中间件配置

`middleware/` 提供服务中间件的配置工具，包括 JWT 认证白名单管理。

## 包名

```go
import middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
```

## JWT 白名单

定义无需认证的路径：

```go
func ExportAuthWhitelist() []map[string]middlewareutil.TransportServiceKind {
    return []map[string]middlewareutil.TransportServiceKind{
        {"/health": middlewareutil.TransportServiceKindAll},
        {"/api/v1/ping/say_hello": middlewareutil.TransportServiceKindHTTP},
    }
}
```

### TransportServiceKind

| 值 | 说明 |
|----|------|
| `TransportServiceKindHTTP` | 仅 HTTP 协议白名单 |
| `TransportServiceKindGRPC` | 仅 gRPC 协议白名单 |
| `TransportServiceKindAll` | HTTP + gRPC 双协议白名单 |

## 白名单合并

多个服务的白名单可以合并：

```go
whitelist := middlewareutil.MergeWhitelist(service1Whitelist, service2Whitelist)
```
