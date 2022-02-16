package pingsrv

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/ikaiguang/go-srv-kit/api/ping/exception"
	"github.com/ikaiguang/go-srv-kit/api/ping/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
)

// ping .
type ping struct {
	v1.UnimplementedSrvPingServer

	log *log.Helper
}

// NewPingService .
func NewPingService(logger log.Logger) v1.SrvPingServer {
	return &ping{
		log: log.NewHelper(logger),
	}
}

func (s *ping) Ping(ctx context.Context, in *v1.PingReq) (out *v1.PingResp, err error) {
	s.log.WithContext(ctx).Infof("Ping Received: %v", in.GetMessage())

	if in.GetMessage() == "" {
		err = exception.ErrorContentMissing("content missing")
		return out, errorutil.WithStack(err)
	}
	if in.GetMessage() == "error" {
		err = exception.ErrorContentError("testing error")
		return out, errorutil.WithStack(err)
	}
	//s.log.Log(log.LevelInfo, "name", "msyql", "msg", "xxx")
	return &v1.PingResp{Message: "Received Message : " + in.GetMessage()}, err
}
