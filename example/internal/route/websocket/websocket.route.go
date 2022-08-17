package websocketroute

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	stdlog "log"
	stdhttp "net/http"

	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	websocketutil "github.com/ikaiguang/go-srv-kit/kratos/websocket"
)

// RegisterRoutes 注册路由
// ref https://github.com/go-kratos/examples/blob/main/ws/main.go
func RegisterRoutes(hs *http.Server, gs *grpc.Server, logger log.Logger) (err error) {
	wsHandler := NewWebsocketService(logger)
	router := mux.NewRouter()
	router.HandleFunc("/ws/v1/websocket", wsHandler.TestWebsocket)

	stdlog.Println("|*** 注册路由：Websocket")
	hs.Handle("/ws/v1/websocket", router)
	return err
}

// ws ...
type ws struct {
	log *log.Helper
}

// WsMessage ws
type WsMessage struct {
	Type    int
	Content []byte
}

// NewWebsocketService ...
func NewWebsocketService(logger log.Logger) *ws {
	return &ws{
		log: log.NewHelper(logger),
	}
}

// TestWebsocket ...
func (s *ws) TestWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != stdhttp.MethodGet {
		err := errorutil.InternalServer(errorv1.ERROR_METHOD_NOT_ALLOWED.String(), "ERROR_METHOD_NOT_ALLOWED")
		s.log.Error(err)
		return
	}
	// 升级连接
	c, err := websocketutil.UpgradeConn(w, r, w.Header())
	if err != nil {
		err = errorutil.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "ws: upgrade conn failed", err)
		s.log.Error(err)
		return
	}
	defer func() { _ = c.Close() }()

	// 读消息
	ctx := context.Background()
	for {
		messageType, messageBytes, wsErr := c.ReadMessage()
		if wsErr != nil {
			if websocketutil.IsCloseError(wsErr) {
				s.log.Infow("websocket close", wsErr.Error())
				break
			}
			err = errorutil.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "ws读取信息失败", err)
			s.log.Error(err)
			return
		}

		// 消息
		wsMessage := &WsMessage{
			Type:    messageType,
			Content: messageBytes,
		}
		//messageChan <- wsMessage

		// 处理
		needCloseConn, err := s.processWsMessage(ctx, wsMessage)
		if err != nil {
			err = errorutil.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "ws处理信息失败", err)
			s.log.Error(err)
			return
		}

		// 响应
		err = c.WriteMessage(messageType, messageBytes)
		if err != nil {
			if websocketutil.IsCloseError(wsErr) {
				s.log.Infow("websocket close", wsErr.Error())
				break
			}
			err = errorutil.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "JSON编码错误", err)
			s.log.Error(err)
			return
		}

		// 关闭
		if needCloseConn {
			return
		}
	}
}

// processWsMessage 处理ws-message
func (s *ws) processWsMessage(ctx context.Context, wsMessage *WsMessage) (needCloseConn bool, err error) {
	s.log.Infow("ws-message type", wsMessage.Type)
	s.log.Infow("ws-message message", string(wsMessage.Content))
	if string(wsMessage.Content) == "close" {
		needCloseConn = true
	}
	return needCloseConn, err
}
