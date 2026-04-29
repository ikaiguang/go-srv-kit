# consul - Consul 客户端

`consul/` 提供 Consul 客户端的创建，支持 TLS 和 HTTP Basic Auth。

## 包名

```go
import consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
```

## 使用

```go
client, err := consulpkg.NewConsulClient(config, opts...)
```

创建时会自动执行 Ping 验证连接可用性。

通常通过 `service/consul` 管理器间接使用。
