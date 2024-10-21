package bo

import "github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/po"

type GetPingMessageParam struct {
	Message string
}

func (s *GetPingMessageParam) ToPoMockPingMessageParam() *po.MockPingMessageParam {
	res := &po.MockPingMessageParam{
		Message: s.Message,
	}
	return res
}

type GetPingMessageResult struct {
	Message string
}

func (s *GetPingMessageResult) SetByPoMockPingMessageReply(dataModel *po.MockPingMessageReply) {
	s.Message = dataModel.Message
}
