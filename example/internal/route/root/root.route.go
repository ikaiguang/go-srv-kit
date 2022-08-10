package rootroute

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	stdlog "log"

	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
)

// RegisterRoutes 注册路由
func RegisterRoutes(hs *http.Server, gs *grpc.Server, logger log.Logger) (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &pingv1.PingResp{
			Message: "Hello World!",
		}
		err := apputil.ResponseEncoder(w, r, data)
		if err != nil {
			apputil.ErrorEncoder(w, r, err)
		}
	})

	stdlog.Println("|*** 注册路由：Root(/)")
	hs.Handle("/", router)
	return err
}
