package serviceapi

import (
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
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
