package serverutil

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func GetApp(serverManager ServerManager) (*kratos.App, error) {
	return serverManager.GetApp()
}

func GetHTTPServer(serverManager ServerManager) (*http.Server, error) {
	return serverManager.GetHTTPServer()
}

func GetGRPCServer(serverManager ServerManager) (*grpc.Server, error) {
	return serverManager.GetGRPCServer()
}
