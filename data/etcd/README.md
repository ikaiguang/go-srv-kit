# etcd - Etcd 客户端

`etcd/` 提供 Etcd 客户端的创建，支持 TLS。

## 包名

```go
import etcdpkg "github.com/ikaiguang/go-srv-kit/data/etcd"
```

## 使用

```go
client, err := etcdpkg.NewEtcdClient(config)
```

创建时会自动执行 Ping 验证连接可用性。
