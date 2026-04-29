# consul - Consul 客户端管理

`consul/` 提供 Consul 客户端的懒加载管理。

## 包名

```go
import consulutil "github.com/ikaiguang/go-srv-kit/service/consul"
```

## 使用

```go
manager, err := consulutil.NewConsulManager(consulConfig)

if manager.Enable() {
    client, err := manager.GetClient()
    defer manager.Close()
}
```

底层使用 `data/consul` 包创建客户端连接。
