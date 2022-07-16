package websockettest

import (
	stdlog "log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"

	baseerror "github.com/ikaiguang/go-srv-kit/api/base/error"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
	websocketutil "github.com/ikaiguang/go-srv-kit/kratos/websocket"
	logutil "github.com/ikaiguang/go-srv-kit/log"

	pkgerrors "github.com/pkg/errors"
)

// RunTestWebsocket 测试websocket
func RunTestWebsocket() {
	processResp, err := TestWebsocket()
	if err != nil {
		stdlog.Printf("%+v\n", err)
		return
	}
	jsonContent, _ := apputil.JSON(processResp)
	stdlog.Println("==> RunTestWebsocket result :", string(jsonContent))
}

// TestWebsocket 测试websocket
func TestWebsocket() (processResp interface{}, err error) {
	urlPath := "/api/v1/testdata/websocket"

	// 开启ws
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8081", Path: urlPath}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return processResp, err
	}
	defer func() { _ = c.Close() }()

	// 读取信息
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			messageType, messageBytes, wsErr := c.ReadMessage()
			if wsErr != nil {
				if websocketutil.IsCloseError(wsErr) {
					logutil.Infow("websocket close", wsErr.Error())
					break
				}
				err = errorutil.InternalServer(baseerror.ERROR_STATUS_INTERNAL_SERVER.String(), "ws读取信息失败", wsErr)
				logutil.Error(err)
				return
			}

			logutil.Infow("ws-message type", messageType)
			logutil.Infow("ws-message message", string(messageBytes))
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var (
		messageSlice = []string{
			"testdata_1",
			"testdata_2",
			"close",
		}
		counter int
		maxNum  = len(messageSlice)
	)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			if counter >= maxNum {
				continue
			}
			wsErr := c.WriteMessage(websocket.TextMessage, []byte(messageSlice[counter]))
			if wsErr != nil {
				if websocketutil.IsCloseError(wsErr) {
					logutil.Infow("websocket close", wsErr.Error())
					break
				}
				err = errorutil.InternalServer(baseerror.ERROR_STATUS_INTERNAL_SERVER.String(), "ws读取信息失败", err)
				logutil.Error(err)
				return
			}
		}
		counter++
	}
	//return processResp, err
}
