# 服务API

* [参考示例](./../testdata/service-api/service_api.util.go)

```go
package serviceapi

import (
	"github.com/go-micro-saas/service-kit/cluster_service_api"
	pingservicev1 "github.com/go-micro-saas/service-kit/testdata/ping-service/api/ping-service/v1/services"
)

// 示例：仅供参考
const (
	PingService   clientutil.ServiceName = "ping-service"
	NodeidService clientutil.ServiceName = "nodeid-service"

	FeishuApi      clientutil.ServiceName = "feishu-openapi"
	DingtalkApi    clientutil.ServiceName = "dingtalk-openapi"
	DingtalkApiOld clientutil.ServiceName = "dingtalk-openapi-old"
)

// NewPingGRPCClient ...
func NewPingGRPCClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (pingservicev1.SrvPingClient, error) {
	serviceName := PingService
	for i := range rewriteServiceName {
		serviceName = rewriteServiceName[i]
	}
	conn, err := clientutil.NewSingletonServiceAPIConnection(serviceAPIManager, serviceName)
	//conn, err := NewServiceAPIConnection(serviceAPIManager, serviceName)
	if err != nil {
		return nil, err
	}
	grpcConn, err := conn.GetGRPCConnection()
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingClient(grpcConn), nil
}

// NewPingHTTPClient ...
func NewPingHTTPClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
	serviceName := PingService
	for i := range rewriteServiceName {
		serviceName = rewriteServiceName[i]
	}
	conn, err := clientutil.NewSingletonServiceAPIConnection(serviceAPIManager, serviceName)
	//conn, err := NewServiceAPIConnection(serviceAPIManager, serviceName)
	if err != nil {
		return nil, err
	}
	httpClient, err := conn.GetHTTPClient()
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingHTTPClient(httpClient), nil
}

```
