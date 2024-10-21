package datarepo

import (
	"context"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/po"
)

type PingDataRepo interface {
	GetMockPingMessage(ctx context.Context, param *po.MockPingMessageParam) (*po.MockPingMessageReply, error)
}
