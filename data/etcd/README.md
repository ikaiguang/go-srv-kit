# etcd

`etcdpkg` 包用于把项目内的 protobuf 配置转换为 `go.etcd.io/etcd/client/v3` 客户端配置，并创建可用的 etcd v3 客户端。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/data/etcd/v3
```

```go
import "github.com/ikaiguang/go-srv-kit/data/etcd/v3"
```

## 核心能力

- `Config`：由 `config.proto` 生成的 etcd 配置结构，包含 endpoints、账号密码、拨号超时、CA 证书和 `insecure_skip_verify`。
- `NewEtcdClient`：接收 `*Config`，转换为 `clientv3.Config`，并在存在 `CaCert` 时配置 TLS Root CA。
- `NewClient`：接收原生 `*clientv3.Config`，创建客户端并通过只读 GET `/ping` 验证连接。
- `Validate` / `ValidateAll`：由 `protoc-gen-validate` 生成的配置校验方法。

## 快速使用

```go
package main

import (
	"log"
	"time"

	etcdpkg "github.com/ikaiguang/go-srv-kit/data/etcd/v3"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	client, err := etcdpkg.NewEtcdClient(&etcdpkg.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: durationpb.New(5 * time.Second),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
```

使用 TLS CA 证书：

```go
client, err := etcdpkg.NewEtcdClient(&etcdpkg.Config{
	Endpoints:          []string{"127.0.0.1:2379"},
	Username:           "user",
	Password:           "password",
	DialTimeout:        durationpb.New(5 * time.Second),
	CaCert:             caCertPEM,
	InsecureSkipVerify: false,
})
```

如果调用方已经构造好 etcd 原生配置，可以直接调用：

```go
client, err := etcdpkg.NewClient(&clientv3.Config{
	Endpoints:   []string{"127.0.0.1:2379"},
	DialTimeout: 5 * time.Second,
})
```

## 生成文件

本目录包含以下 protobuf 相关文件：

- `config.proto`：配置定义源文件。
- `config.pb.go`：由 `protoc-gen-go` 生成的 Go 类型。
- `config.pb.validate.go`：由 `protoc-gen-validate` 生成的校验代码。
- `config.swagger.json`：由 Makefile 中的 proto 生成流程产出。

不要手工修改 `*.pb.go` 或 `*.validate.go`。修改 `config.proto` 后在仓库根目录运行：

```bash
protoc --proto_path=. --proto_path="$(go env GOPATH)/src" --proto_path=./third_party \
  --go_out=paths=source_relative:. \
  --validate_out=paths=source_relative,lang=go:. \
  data/etcd/config.proto
```

## 测试

```bash
go test ./...
```

当前单元测试覆盖 nil 配置、无效 CA PEM 和 TLS 配置转换，不依赖真实 etcd 服务。真实连接探测需要在集成环境另行验证。

## 注意事项

- `NewClient` 使用只读 GET `/ping` 探测连接，超时采用 `clientv3.Config.DialTimeout`，未配置时默认为 5 秒。
- 探测失败时会关闭已创建的 client，并返回 nil client 和带上下文的 error。
- CA PEM 无法解析时会在创建 client 前返回错误。
- 生产环境不要默认开启 `InsecureSkipVerify`；确需开启时应记录风险和边界。
- `config.proto` 的 `go_package` 与当前 module 和包目录保持一致。
