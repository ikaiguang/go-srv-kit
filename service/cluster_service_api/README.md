# cluster_service_api - 集群服务间 API 调用

`cluster_service_api/` 提供微服务间的 HTTP/gRPC 客户端管理，支持直连和服务发现（Consul/Etcd）两种模式。

## 包名

```go
import clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
```

## 核心概念

### ServiceAPIManager

管理所有服务的连接配置和客户端创建：

```go
manager, err := clientutil.NewServiceAPIManager(apiConfigs, opts...)

// 创建服务连接
conn, err := manager.NewAPIConnection(serviceName)

// 获取 gRPC 连接
grpcConn, err := conn.GetGRPCConnection()

// 获取 HTTP 客户端
httpClient, err := conn.GetHTTPClient()
```

### 连接模式

| 模式 | 说明 |
|------|------|
| 直连 | 通过 `host:port` 直接连接目标服务 |
| Consul 注册发现 | 通过 Consul 注册中心发现服务 |
| Etcd 注册发现 | 通过 Etcd 注册中心发现服务 |

### 单例连接

同一服务名的连接默认复用（单例模式），避免重复创建：

```go
conn, err := clientutil.NewSingletonServiceAPIConnection(manager, serviceName)
```

## 使用示例

```go
// 创建 Ping 服务的 gRPC 客户端
conn, err := clientutil.NewSingletonServiceAPIConnection(manager, "ping-service")
grpcConn, err := conn.GetGRPCConnection()
client := pingservicev1.NewSrvPingClient(grpcConn)
```

## 参考

- 使用示例：[testdata/service-api/](../../testdata/service-api/)
