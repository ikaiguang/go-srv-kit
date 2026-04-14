package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
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

// Ping 请求消息常量
const (
	pingMessageLogger      = "logger"
	pingMessageError       = "error"
	pingMessagePanic       = "panic"
	pingMessageHTTPAndGRPC = "http_and_grpc"
)

func (s *pingService) Ping(ctx context.Context, req *resourcev1.PingReq) (*resourcev1.PingResp, error) {
	// logger
	if req.GetMessage() == pingMessageLogger {
		s.testLogger(ctx, req)
	}

	// error
	if req.GetMessage() == pingMessageError {
		e := errorv1.ErrorContentError(req.GetMessage())
		md := map[string]string{
			"testdata": "testdata",
		}
		return nil, errorpkg.WrapWithMetadata(e, md)
	}

	// panic
	if req.GetMessage() == pingMessagePanic {
		panic("testdata panic")
	}

	// request
	if req.GetMessage() == pingMessageHTTPAndGRPC {
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
	// 统一使用 WithContext(ctx).Infow 结构化日志
	s.log.WithContext(ctx).Infow("msg", "Ping Received", "message", in.GetMessage())
}

func (s *pingService) requestClusterServiceAPI(ctx context.Context, in *resourcev1.PingReq) error {
	err := s.pingBiz.TestingRequest(ctx)
	if err != nil {
		return err
	}
	return nil
}
