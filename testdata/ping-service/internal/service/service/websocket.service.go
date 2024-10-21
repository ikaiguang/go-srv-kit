package service

import (
	"github.com/go-kratos/kratos/v2/log"
	bizrepo "github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/repo"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// WebsocketService ...
type WebsocketService struct {
	log          *log.Helper
	websocketBiz bizrepo.WebsocketBizRepo
}

// NewWebsocketService ...
func NewWebsocketService(
	logger log.Logger,
	websocketBiz bizrepo.WebsocketBizRepo,
) *WebsocketService {
	logHelper := log.NewHelper(log.With(logger, "module", "ping/service/websocket"))

	return &WebsocketService{
		log:          logHelper,
		websocketBiz: websocketBiz,
	}
}

// TestWebsocket ...
func (s *WebsocketService) TestWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != stdhttp.MethodGet {
		e := errorpkg.ErrorMethodNotAllowed("METHOD_NOT_ALLOWED")
		w.WriteHeader(stdhttp.StatusBadRequest)
		_, _ = w.Write([]byte(e.Error()))
		return
	}

	err := s.websocketBiz.Wss(r.Context(), w, r)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	return
}
