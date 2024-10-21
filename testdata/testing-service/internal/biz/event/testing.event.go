package events

import (
	"github.com/go-kratos/kratos/v2/log"
	bizrepos "github.com/ikaiguang/go-srv-kit/testdata/testing-service/internal/biz/repo"
)

type testingEvent struct {
	log *log.Helper
}

func NewTestingEvent(
	logger log.Logger,
) bizrepos.TestingEventRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/event"))

	return &testingEvent{
		log: logHelper,
	}
}
