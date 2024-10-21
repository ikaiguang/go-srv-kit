package dto

import (
	"github.com/ikaiguang/go-srv-kit/testdata/testing-service/internal/biz/bo"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	TestingDto testingDto
)

type testingDto struct{}

func (s *testingDto) ToBoXxx() *bo.Testdata {
	res := &bo.Testdata{}
	return res
}

func (s *testingDto) ToPbXxx() *emptypb.Empty {
	res := &emptypb.Empty{}

	return res
}
