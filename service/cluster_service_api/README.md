# cluster_service_api

`cluster_service_api` 包用于按服务名创建并缓存 HTTP 或 gRPC 客户端连接。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/service
```

```go
import clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
```

## 获取 gRPC 连接

```go
package serviceapi

import (
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	"google.golang.org/grpc"
)

const PingService clientutil.ServiceName = "ping-service"

func NewPingGRPCConn(manager clientutil.ServiceAPIManager) (*grpc.ClientConn, error) {
	conn, err := clientutil.NewSingletonServiceAPIConnection(manager, PingService)
	if err != nil {
		return nil, err
	}
	return conn.GetGRPCConnection()
}
```

HTTP 服务使用相同的连接获取流程，最后调用 `conn.GetHTTPClient()`。
