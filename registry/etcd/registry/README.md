# registry

`registry` 目录提供包 `registrypkg`，用于基于已有 `*clientv3.Client` 创建 Kratos etcd 注册中心。它适合在 Kratos 服务启动或 Wire 装配阶段复用统一的 etcd registry 创建逻辑。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/registry/etcd/v3
```

```go
import registrypkg "github.com/ikaiguang/go-srv-kit/registry/etcd/v3/registry"
```

## 核心能力

- `NewEtcdRegistry`：接收 etcd v3 client 和来自 `github.com/go-kratos/kratos/contrib/registry/etcd/v3` 的可选 registry option，返回 `*etcdregistry.Registry`。
- 默认配置：函数内部默认添加 `etcdregistry.MaxRetry(3)`，再追加调用方传入的自定义 option。

## 快速使用

```go
package main

import (
	"time"

	registrypkg "github.com/ikaiguang/go-srv-kit/registry/etcd/v3/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	registry, err := registrypkg.NewEtcdRegistry(client)
	if err != nil {
		panic(err)
	}

	_ = registry
}
```

调用方也可以传入 Kratos etcd registry 的原生 option：

```go
registry, err := registrypkg.NewEtcdRegistry(
	client,
	// etcdregistry.Namespace("/services"),
)
```

## 测试

```bash
go test ./registry
```

## 注意事项

- 调用前需要由业务服务自行创建并管理 `*clientv3.Client` 的生命周期。
- etcd endpoint、账号、Token、证书路径等配置不要硬编码在公开源码或文档中。
- 当前包不负责启动、关闭 Kratos server，也不负责创建 etcd client。
- 若需要覆盖默认行为，请优先使用 Kratos etcd registry 已提供的 option。
