package serverutil

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"sync"
)

type ServerManager interface {
	GetApp() (*kratos.App, error)
	GetHTTPServer() (*http.Server, error)
	GetGRPCServer() (*grpc.Server, error)
}

type serverManager struct {
	launcherManager setuputil.LauncherManager
	authWhiteList   map[string]middlewareutil.TransportServiceKind

	appOnce        sync.Once
	appInstance    *kratos.App
	httpServerOnce sync.Once
	httpServer     *http.Server
	grpcServerOnce sync.Once
	grpcServer     *grpc.Server
}

func NewServerManager(
	launcherManager setuputil.LauncherManager,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
) (ServerManager, error) {
	return &serverManager{
		launcherManager: launcherManager,
		authWhiteList:   authWhiteList,
	}, nil
}

func (s *serverManager) GetApp() (*kratos.App, error) {
	app, err := s.getSingletonApp()
	if err != nil {
		return nil, err
	}
	return app, nil
}
func (s *serverManager) getApp() (*kratos.App, error) {
	hs, err := s.getSingletonHTTPServer()
	if err != nil {
		return nil, err
	}
	gs, err := s.getSingletonGRPCServer()
	if err != nil {
		return nil, err
	}
	app, err := NewApp(s.launcherManager, hs, gs)
	if err != nil {
		return nil, err
	}
	s.appInstance = app
	return app, nil
}
func (s *serverManager) getSingletonApp() (*kratos.App, error) {
	var err error
	s.appOnce.Do(func() {
		s.appInstance, err = s.getApp()
	})
	if err != nil {
		s.appOnce = sync.Once{}
	}
	return s.appInstance, err
}

func (s *serverManager) GetHTTPServer() (*http.Server, error) {
	hs, err := s.getSingletonHTTPServer()
	if err != nil {
		return nil, err
	}
	return hs, nil
}
func (s *serverManager) getHTTPServer() (*http.Server, error) {
	hs, err := NewHTTPServer(s.launcherManager, s.authWhiteList)
	if err != nil {
		return nil, err
	}
	s.httpServer = hs
	return hs, nil
}
func (s *serverManager) getSingletonHTTPServer() (*http.Server, error) {
	var err error
	s.httpServerOnce.Do(func() {
		s.httpServer, err = s.getHTTPServer()
	})
	if err != nil {
		s.httpServerOnce = sync.Once{}
	}
	return s.httpServer, err
}

func (s *serverManager) GetGRPCServer() (*grpc.Server, error) {
	gs, err := s.getSingletonGRPCServer()
	if err != nil {
		return nil, err
	}
	return gs, nil
}
func (s *serverManager) getGRPCServer() (*grpc.Server, error) {
	gs, err := NewGRPCServer(s.launcherManager, s.authWhiteList)
	if err != nil {
		return nil, err
	}
	s.grpcServer = gs
	return gs, nil
}
func (s *serverManager) getSingletonGRPCServer() (*grpc.Server, error) {
	var err error
	s.grpcServerOnce.Do(func() {
		s.grpcServer, err = s.getGRPCServer()
	})
	if err != nil {
		s.grpcServerOnce = sync.Once{}
	}
	return s.grpcServer, err
}
