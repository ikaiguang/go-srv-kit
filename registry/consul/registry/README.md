# registry

`registry` 目录提供包 `registrypkg`，用于为 Kratos 服务创建 Consul registry。它适合已经使用 HashiCorp Consul client，并希望复用统一健康检查、心跳和默认超时配置的服务。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/registry/consul/v3
```

```go
import registrypkg "github.com/ikaiguang/go-srv-kit/registry/consul/v3/registry"
```

## 核心 API

- `DefaultTimeout`：默认 registry 超时时间，当前为 `time.Minute`。
- `NewConsulRegistry(consulClient *api.Client, opts ...consulregistry.Option) (*consulregistry.Registry, error)`：基于调用方传入的 Consul client 创建 Kratos Consul registry，并默认启用健康检查、心跳和 `DefaultTimeout`。

## 快速使用

```go
package main

import (
	"log"
	"time"

	consulregistry "github.com/go-kratos/kratos/contrib/registry/consul/v3"
	registrypkg "github.com/ikaiguang/go-srv-kit/registry/consul/v3/registry"
	"github.com/hashicorp/consul/api"
)

func main() {
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	registry, err := registrypkg.NewConsulRegistry(
		consulClient,
		consulregistry.WithTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	_ = registry
}
```

## 测试

```bash
go test ./registry
```

## 注意事项

- `NewConsulRegistry` 不创建 Consul client，也不读取配置文件；Consul 地址、ACL token、TLS、连接超时和其他认证配置由调用方在 `api.NewClient` 前后处理。
- 默认参数会启用健康检查和心跳。如果业务服务需要不同策略，可通过追加 `consulregistry.Option` 覆盖 Kratos Consul registry 支持的配置。
- 不要在示例、配置或日志中记录未脱敏的 Consul token、账号密码、私钥或内网地址。
