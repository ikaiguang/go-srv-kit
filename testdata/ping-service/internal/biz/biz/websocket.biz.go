package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	websocketpkg "github.com/ikaiguang/go-srv-kit/kratos/websocket"
	bizrepo "github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/repo"
	stdhttp "net/http"
)

// websocketBiz ...
type websocketBiz struct {
	log *log.Helper
}

// NewWebsocketBiz ...
func NewWebsocketBiz(logger log.Logger) bizrepo.WebsocketBizRepo {
	return &websocketBiz{
		log: log.NewHelper(log.With(logger, "module", "ping/biz/websocket")),
	}
}

// WsMessage ws
type WsMessage struct {
	Type    int
	Content []byte
}

// Wss ws
func (s *websocketBiz) Wss(ctx context.Context, w stdhttp.ResponseWriter, r *stdhttp.Request) error {
	// 升级连接
	cc, err := websocketpkg.UpgradeConn(w, r, w.Header())
	if err != nil {
		e := errorpkg.ErrorInternalError("upgrade conn failed")
		s.log.WithContext(ctx).Error(e.Error())
		return errorpkg.Wrap(e, err)
	}
	defer func() { _ = cc.Close() }()

	// 处理消息
	err = s.ProcessWssMsg(ctx, cc)
	if err != nil {
		return err
	}
	return err
}

func (s *websocketBiz) ProcessWssMsg(ctx context.Context, cc *websocket.Conn) error {
	// 读消息
	for {
		messageType, messageBytes, wsErr := cc.ReadMessage()
		if wsErr != nil {
			if websocketpkg.IsCloseError(wsErr) {
				s.log.WithContext(ctx).Infow("websocket close", wsErr.Error())
				break
			}
			s.log.WithContext(ctx).Error(wsErr)
			e := errorpkg.ErrorInternalError("ws读取信息失败")
			return errorpkg.Wrap(e, wsErr)
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
			s.log.WithContext(ctx).Error(err)
			e := errorpkg.ErrorInternalError("ws处理信息失败")
			return errorpkg.Wrap(e, err)
		}

		// 响应
		err = cc.WriteMessage(messageType, messageBytes)
		if err != nil {
			if websocketpkg.IsCloseError(err) {
				s.log.WithContext(ctx).Infow("websocket close", err.Error())
				break
			}
			s.log.WithContext(ctx).Error(err)
			e := errorpkg.ErrorInternalError("JSON编码错误")
			return errorpkg.Wrap(e, err)
		}

		// 关闭
		if needCloseConn {
			return err
		}
	}
	return nil
}

// processWsMessage 处理ws-message
func (s *websocketBiz) processWsMessage(ctx context.Context, wsMessage *WsMessage) (needCloseConn bool, err error) {
	s.log.WithContext(ctx).Infow("ws-message type", wsMessage.Type)
	s.log.WithContext(ctx).Infow("ws-message message", string(wsMessage.Content))
	if string(wsMessage.Content) == "close" {
		needCloseConn = true
	}
	return needCloseConn, err
}
