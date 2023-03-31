package testdatasrv

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/rs/xid"

	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	testdatav1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/resources"
	testdataservicev1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/services"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	contextutil "github.com/ikaiguang/go-srv-kit/kratos/context"
	websocketutil "github.com/ikaiguang/go-srv-kit/kratos/websocket"
)

// testdata .
type testdata struct {
	testdataservicev1.UnimplementedSrvTestdataServer

	log *log.Helper
}

// NewTestdataService .
func NewTestdataService(logger log.Logger) testdataservicev1.SrvTestdataServer {
	return &testdata{
		log: log.NewHelper(logger),
	}
}

// Websocket websocket
func (s *testdata) Websocket(ctx context.Context, in *testdatav1.TestReq) (resp *testdatav1.TestResp, err error) {
	//err = errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "未实现")
	//return &v1.TestResp{}, err

	// http
	httpContext, isHTTPContext := contextutil.MatchHTTPContext(ctx)
	if isHTTPContext {
		//return s.exportApp(httpContext, req)
		s.ws(httpContext, in)
		resp = &testdatav1.TestResp{
			Message: xid.New().String(),
		}
		err = errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "未实现")
		return resp, err
	}

	err = errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "未实现")
	return &testdatav1.TestResp{}, err
}

// WsMessage ws
type WsMessage struct {
	Type    int
	Content []byte
}

// ws ws
func (s *testdata) ws(ctx http.Context, in *testdatav1.TestReq) {
	// 升级连接
	c, err := websocketutil.UpgradeConn(ctx.Response(), ctx.Request(), ctx.Response().Header())
	if err != nil {
		err = errorutil.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "ws开启失败", err)
		s.log.Error(err)
		return
	}
	defer func() { _ = c.Close() }()

	// 读消息
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
func (s *testdata) processWsMessage(ctx context.Context, wsMessage *WsMessage) (needCloseConn bool, err error) {
	s.log.Infow("ws-message type", wsMessage.Type)
	s.log.Infow("ws-message message", string(wsMessage.Content))
	if string(wsMessage.Content) == "close" {
		needCloseConn = true
	}
	return needCloseConn, err
}
