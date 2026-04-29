# registry_etcd - Etcd 服务注册

`registry_etcd/` 提供基于 Etcd 的服务注册与发现。

## 包名

```go
import etcdregistry "github.com/ikaiguang/go-srv-kit/kratos/registry_etcd"
```

## 使用

```go
registry, err := etcdregistry.NewEtcdRegistry(etcdClient, opts...)
```
