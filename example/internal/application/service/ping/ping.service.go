package pingsrv

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pingerror "github.com/ikaiguang/go-srv-kit/api/ping/error"
	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
)

// ping .
type ping struct {
	pingv1.UnimplementedSrvPingServer

	log *log.Helper
}

// NewPingService .
func NewPingService(logger log.Logger) pingv1.SrvPingServer {
	return &ping{
		log: log.NewHelper(logger),
	}
}

func (s *ping) Ping(ctx context.Context, in *pingv1.PingReq) (out *pingv1.PingResp, err error) {
	s.log.WithContext(ctx).Infof("Ping Received: %v", in.GetMessage())

	if in.GetMessage() == "" {
		err = pingerror.ErrorContentMissing("content missing")
		return out, errorutil.WithStack(err)
	}
	if in.GetMessage() == "error" {
		e := pingerror.ErrorContentError("testing error")
		e.Metadata = map[string]string{
			"testdata": "testdata",
		}
		return out, errorutil.WithStack(e)
	}
	//s.log.Log(log.LevelInfo, "name", "msyql", "msg", "xxx")
	return &pingv1.PingResp{Message: "Received Message : " + in.GetMessage()}, err
}
