package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	errorv1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/ping-service/v1/errors"
	resourcev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/ping-service/v1/resources"
	servicev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/ping-service/v1/services"
	bizrepo "github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/repo"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/service/dto"
)

type pingService struct {
	servicev1.UnimplementedSrvPingServer

	log     *log.Helper
	pingBiz bizrepo.PingBizRepo
}

func NewPingService(logger log.Logger, pingBiz bizrepo.PingBizRepo) servicev1.SrvPingServer {
	logHelper := log.NewHelper(log.With(logger, "module", "ping/service/ping"))

	return &pingService{
		log:     logHelper,
		pingBiz: pingBiz,
	}
}

func (s *pingService) Ping(ctx context.Context, req *resourcev1.PingReq) (*resourcev1.PingResp, error) {
	// logger
	if req.GetMessage() == "logger" {
		s.testLogger(ctx, req)
	}

	// error
	if req.GetMessage() == "error" {
		e := errorv1.ErrorContentError(req.GetMessage())
		md := map[string]string{
			"testdata": "testdata",
		}
		return nil, errorpkg.WrapWithMetadata(e, md)
	}

	// panic
	if req.GetMessage() == "panic" {
		panic("testdata panic")
	}

	// request
	if req.GetMessage() == "http_and_grpc" {
		err := s.requestClusterServiceAPI(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	param := dto.PingDTO.ToBoGetPingMessageParam(req)
	reply, err := s.pingBiz.GetPingMessage(ctx, param)
	if err != nil {
		return nil, err
	}
	resp := &resourcev1.PingResp{
		Data: dto.PingDTO.ToPbPingRespData(reply),
	}
	return resp, err
}

func (s *pingService) testLogger(ctx context.Context, in *resourcev1.PingReq) {
	s.log.WithContext(ctx).Infof("==> s.log.WithContext(ctx).Infof : Ping Received: %s", in.GetMessage())
	s.log.Infow("==> s.log.Infow : Ping Received: ", in.GetMessage())
	logpkg.InfoWithContext(ctx, "==> logpkg.InfoWithContext : Ping Received: ", in.GetMessage())
	logpkg.InfowWithContext(ctx, "==> logpkg.InfowWithContext : Ping Received: ", in.GetMessage())
	logpkg.Info("==> logpkg.Info : Ping Received: ", in.GetMessage())
	debugutil.Printw("==> debugutil.Print : Ping Received: ", in.GetMessage())
}

func (s *pingService) requestClusterServiceAPI(ctx context.Context, in *resourcev1.PingReq) error {
	err := s.pingBiz.TestingRequest(ctx)
	if err != nil {
		return err
	}
	return nil
}
