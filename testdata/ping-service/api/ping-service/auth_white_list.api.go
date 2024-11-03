package pingapi

import (
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/ping-service/v1/services"
)

// GetAuthWhiteList 验证白名单
func GetAuthWhiteList() map[string]middlewareutil.TransportServiceKind {
	// 白名单
	whiteList := make(map[string]middlewareutil.TransportServiceKind)

	// 测试
	whiteList[pingservicev1.OperationSrvPingPing] = middlewareutil.TransportServiceKindALL
	whiteList["/ws/v1/testdata/websocket"] = middlewareutil.TransportServiceKindALL

	return whiteList
}
