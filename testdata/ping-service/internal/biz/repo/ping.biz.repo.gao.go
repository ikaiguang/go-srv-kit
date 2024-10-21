package bizrepo

import (
	"context"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/bo"
)

type PingBizRepo interface {
	GetPingMessage(ctx context.Context, param *bo.GetPingMessageParam) (*bo.GetPingMessageResult, error)
	TestingRequest(ctx context.Context) error
}
