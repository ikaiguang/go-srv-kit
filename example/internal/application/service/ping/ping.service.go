package pingsrv

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pingerrorv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/errors"
	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// ping .
type ping struct {
	pingservicev1.UnimplementedSrvPingServer

	log *log.Helper
}

// NewPingService .
func NewPingService(logger log.Logger) pingservicev1.SrvPingServer {
	return &ping{
		log: log.NewHelper(logger),
	}
}

func (s *ping) Ping(ctx context.Context, in *pingv1.PingReq) (out *pingv1.PingResp, err error) {
	s.log.WithContext(ctx).Infof("Ping Received: %v", in.GetMessage())

	// 空
	if in.GetMessage() == "" {
		err = pingerrorv1.ErrorContentMissing("content missing")
		return out, errorutil.WithStack(err)
	}
	// logger
	if in.GetMessage() == "logger" {
		s.testLogger(ctx, in)
	}
	// error
	if in.GetMessage() == "error" {
		e := pingerrorv1.ErrorContentError("testing error")
		e.Metadata = map[string]string{
			"testdata": "testdata",
		}
		return out, errorutil.WithStack(e)
	}

	return &pingv1.PingResp{Message: "Received Message : " + in.GetMessage()}, err
}

// testLogger .
// curl http://127.0.0.1:8081/api/v1/ping/logger
func (s *ping) testLogger(ctx context.Context, in *pingv1.PingReq) {
	s.log.WithContext(ctx).Infof("==> s.log.WithContext(ctx).Infof : Ping Received: %s", in.GetMessage())
	s.log.Infow("==> s.log.Infow : Ping Received: ", in.GetMessage())
	logutil.InfoWithContext(ctx, "==> loghelper.ContextInfo : Ping Received: ", in.GetMessage())
	logutil.InfowWithContext(ctx, "==> loghelper.ContextInfow : Ping Received: ", in.GetMessage())
	logutil.Info("==> loghelper.Info : Ping Received: ", in.GetMessage())
	debugutil.Printw("==> debugutil.Print : Ping Received: ", in.GetMessage())
}
