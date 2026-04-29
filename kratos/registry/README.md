# registry - Consul 服务注册

`registry/` 提供基于 Consul 的服务注册与发现。

## 包名

```go
import registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
```

## 使用

```go
registry, err := registrypkg.NewConsulRegistry(consulClient, opts...)
```

默认配置：
- 健康检查：启用
- 心跳：启用
- 超时：60 秒
