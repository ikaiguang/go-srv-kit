# component

泛型懒加载组件容器、命名组件组和生命周期关闭管理，适合 Redis、MySQL、HTTP client、SDK client 等基础设施 manager 的延迟初始化与统一关闭。

> 目录名是 `component`，当前 Go 包名是 `componentpkg`。外部使用时建议显式设置 import alias。

## 单组件用法

`Component[T]` 会在首次 `Get()` 时调用 factory。初始化成功后复用同一个实例；初始化失败不会缓存结果，下次 `Get()` 会重新尝试。

```go
package main

import (
	"fmt"
	"io"

	componentpkg "github.com/ikaiguang/go-kit/component"
)

type Client struct{}

func (c *Client) Close() error { return nil }

func main() {
	lc := &componentpkg.Lifecycle{}
	defer func() {
		if err := lc.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	clientComponent := componentpkg.NewComponent("client", func() (*Client, error) {
		return &Client{}, nil
	}, lc)

	client, err := clientComponent.Get()
	if err != nil {
		panic(err)
	}
	_ = client

	var _ io.Closer = client
}
```

如果 `T` 实现了 `io.Closer`，组件初始化成功后会自动注册到 `Lifecycle`。调用 `Lifecycle.Close()` 时，组件会按注册逆序关闭。

## 命名组件组

`ComponentGroup[T]` 用于同一种组件存在多个命名实例的场景，例如多 Redis 实例、多数据库连接或多租户 SDK client。

```go
lc := &componentpkg.Lifecycle{}
group := componentpkg.NewComponentGroup("client", func(name string) func() (*Client, error) {
	return func() (*Client, error) {
		return &Client{}, nil
	}
}, lc)

defaultClient, err := group.Get("default")
if err != nil {
	return err
}
_ = defaultClient

archiveClient, err := group.Get("archive")
if err != nil {
	return err
}
_ = archiveClient
```

同名实例只会创建一个 `Component`；不同名称会分别懒加载并注册到同一个生命周期管理器。

## 并发语义

- 已初始化的 `Component.Get()` 通过原子读返回，热路径无锁。
- 首次初始化使用互斥锁和双重检查，保证成功的 factory 只发布一次。
- factory 返回错误时不缓存失败，调用方下次 `Get()` 会重新触发初始化。
- `ComponentGroup.Get(name)` 对已存在命名实例使用读锁，多个 goroutine 可并发命中。
- `Lifecycle.Close()` 会清空已注册 closer；通常只在进程退出或服务停止时调用一次。

## 注意事项

- factory 内应完成连接参数校验、超时控制和必要的 ping/check。
- factory 不应依赖未受控的全局状态；错误要原样返回给调用方处理。
- `Lifecycle` 的零值可用，外部包直接使用 `&componentpkg.Lifecycle{}` 创建。
- 如果组件需要关闭，请让返回值实现 `io.Closer`，否则不会自动注册到生命周期。
- `Close()` 会尝试关闭所有已注册组件，并用 `errors.Join` 返回多个关闭错误。

## 验证

```bash
go test ./component
```
