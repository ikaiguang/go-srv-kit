# consul

`consulpkg` 包负责把项目内的 protobuf 配置 `Config` 转换为 HashiCorp Consul API client 配置，并创建 `*api.Client`。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/data/consul/v3
```

```go
import "github.com/ikaiguang/go-srv-kit/data/consul/v3"
```

## 核心能力

- `Config`：由 `config.proto` 生成的 Consul 连接配置，覆盖地址、ACL token、命名空间、分区、HTTP Basic Auth 和 TLS PEM 等字段。
- `NewClient(conf *Config, opts ...Option)`：根据 `Config` 创建 `*api.Client`，并通过只读 GET `ping` 验证连接。
- `NewConsulClient(conf *Config, opts ...Option)`：`NewClient` 的兼容封装。
- `WithWriter(writer io.Writer)`：设置 Consul API logger 的输出目标。

## 快速使用

```go
package main

import (
	"log"

	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul/v3"
)

func main() {
	client, err := consulpkg.NewConsulClient(&consulpkg.Config{
		Scheme:     "http",
		Address:    "127.0.0.1:8500",
		Datacenter: "dc1",
	})
	if err != nil {
		log.Fatal(err)
	}

	_ = client
}
```

## 配置字段

| 字段 | 说明 |
| --- | --- |
| `Enable` | 是否启用 Consul 配置；当前 `NewClient` 不读取该字段。 |
| `Scheme` | Consul scheme，例如 `http` 或 `https`。 |
| `Address` | Consul 地址，例如 `127.0.0.1:8500`。 |
| `PathPrefix` | Consul API path prefix。 |
| `Datacenter` | Consul datacenter。 |
| `WaitTime` | Consul blocking query wait time。 |
| `Token` | Consul ACL token。 |
| `Namespace` | Consul Enterprise namespace。 |
| `Partition` | Consul Enterprise partition。 |
| `WithHttpBasicAuth` | 是否设置 HTTP Basic Auth。 |
| `AuthUsername` / `AuthPassword` | HTTP Basic Auth 用户名和密码。 |
| `InsecureSkipVerify` | 是否跳过 TLS 证书校验。 |
| `TlsAddress` | TLS server name/address。 |
| `TlsCaPem` | CA PEM 内容。 |
| `TlsCertPem` | client cert PEM 内容。 |
| `TlsKeyPem` | client key PEM 内容。 |

## 生成

`Config` 来自 `config.proto`。修改 protobuf 后在仓库根目录运行：

```bash
protoc --proto_path=. --proto_path="$(go env GOPATH)/src" --proto_path=./third_party \
  --go_out=paths=source_relative:. \
  --validate_out=paths=source_relative,lang=go:. \
  data/consul/config.proto
```

不要手改 `*.pb.go`、`*.validate.go` 等生成文件。

## 测试

```bash
go test ./...
```

当前单元测试使用内存 HTTP transport 验证连接探测为 GET，不依赖真实 Consul 服务。

## 注意事项

- `NewClient` 会读取 Consul KV `ping` 验证连接，不会创建或修改该 key。
- `Token`、`AuthPassword`、`TlsKeyPem` 属于敏感配置，不应输出到日志或提交真实值。
- `InsecureSkipVerify` 会降低 TLS 安全性，只应在明确风险的测试环境使用。
