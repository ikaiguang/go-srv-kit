package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
	resourcev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/ping-service/v1/resources"
)

type HomeService struct {
	log *log.Helper
}

func NewHomeService(logger log.Logger) *HomeService {
	logHelper := log.NewHelper(log.With(logger, "module", "ping/service/home"))

	return &HomeService{
		log: logHelper,
	}
}

func (s *HomeService) Homepage(w http.ResponseWriter, r *http.Request) {
	data := &resourcev1.PingResp{
		Data: &resourcev1.PingRespData{
			Message: "Hello World!",
		},
	}
	err := apppkg.SuccessResponseEncoder(w, r, data)
	if err != nil {
		apppkg.ErrorResponseEncoder(w, r, err)
	}
}
