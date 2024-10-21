package service

import (
	"github.com/go-kratos/kratos/v2/log"
	servicev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/testdata-service/v1/services"
	bizrepos "github.com/ikaiguang/go-srv-kit/testdata/testing-service/internal/biz/repo"
)

type testingV1Service struct {
	servicev1.UnimplementedSrvTestdataServer

	log        *log.Helper
	testingBiz bizrepos.TestingBizRepo
}

func NewTestingV1Service(logger log.Logger, testingBiz bizrepos.TestingBizRepo) servicev1.SrvTestdataServer {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/service/service"))
	return &testingV1Service{
		log:        logHelper,
		testingBiz: testingBiz,
	}
}
