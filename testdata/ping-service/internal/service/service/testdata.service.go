package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	resourcev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/testdata-service/v1/resources"
	servicev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/testdata-service/v1/services"
	bizrepo "github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/repo"
)

// testdata .
type testdata struct {
	servicev1.UnimplementedSrvTestdataServer

	log          *log.Helper
	websocketBiz bizrepo.WebsocketBizRepo
}

// NewTestdataService .
func NewTestdataService(
	logger log.Logger,
	websocketBiz bizrepo.WebsocketBizRepo,
) servicev1.SrvTestdataServer {
	logHelper := log.NewHelper(log.With(logger, "module", "ping/service/testdata"))

	return &testdata{
		log:          logHelper,
		websocketBiz: websocketBiz,
	}
}

// Websocket websocket
func (s *testdata) Websocket(ctx context.Context, in *resourcev1.TestReq) (*resourcev1.TestResp, error) {
	e := errorpkg.ErrorUnimplemented("unimplemented")
	return nil, errorpkg.WithStack(e)
}

// wss ws
func (s *testdata) wss(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
	return s.websocketBiz.Wss(ctx, w, r)
}
