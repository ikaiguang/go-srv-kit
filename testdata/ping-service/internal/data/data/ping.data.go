package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/po"
	datarepo "github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/repo"
)

type pingData struct {
	log *log.Helper
}

func NewPingData(logger log.Logger) datarepo.PingDataRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "ping/data/ping"))

	return &pingData{
		log: logHelper,
	}
}

func (p *pingData) GetMockPingMessage(ctx context.Context, param *po.MockPingMessageParam) (*po.MockPingMessageReply, error) {
	return &po.MockPingMessageReply{Message: "mock request message: " + param.Message}, nil
}
