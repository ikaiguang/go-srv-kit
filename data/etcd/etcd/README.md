# etcd

`etcd` 包用于把项目内的 protobuf 配置转换为 `go.etcd.io/etcd/client/v3` 客户端配置，并创建可用的 etcd v3 客户端。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/data/etcd/v3
```

```go
import etcdpkg "github.com/ikaiguang/go-srv-kit/data/etcd/v3/etcd"
```

## 核心能力

- `Config`：由 `config.proto` 生成的 etcd 配置结构，包含 endpoints、账号密码、拨号超时、CA 证书和 `insecure_skip_verify`。
- `NewEtcdClient`：接收 `*Config`，转换为 `clientv3.Config`，并在存在 `CaCert` 时配置 TLS Root CA。
- `NewClient`：接收原生 `*clientv3.Config`，创建客户端并通过写入 `/ping=pong` 验证连接。
- `Validate` / `ValidateAll`：由 `protoc-gen-validate` 生成的配置校验方法。

## 快速使用

```go
package main

import (
	"log"
	"time"

	etcdpkg "github.com/ikaiguang/go-srv-kit/data/etcd/v3/etcd"
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

不要手工修改 `*.pb.go` 或 `*.validate.go`。修改 `config.proto` 后，优先使用仓库 Makefile 重新生成：

```bash
make protoc-config-protobuf
```

## 测试

```bash
go test ./etcd
```

当前目录没有 `*_test.go`。`NewClient` 会连接真实 etcd 并写入 `/ping`，后续补测试时建议使用可控的集成测试环境，或对构造逻辑拆出可单测的部分。

## 注意事项

- `NewEtcdClient` 假设 `Config.DialTimeout` 非空；传入 nil 可能导致运行时 panic。调用前应提供有效的 `google.protobuf.Duration`。
- `NewClient` 使用 `context.Background()` 执行探测写入，不接收外部 `context.Context`。
- `NewClient` 创建客户端后，如果探测写入失败，会同时返回非 nil client 和 error，调用方需要自行决定是否关闭该 client。
- 生产环境不要默认开启 `InsecureSkipVerify`；确需开启时应记录风险和边界。
- `config.proto` 的 `go_package` 与当前 module 和包目录保持一致。
