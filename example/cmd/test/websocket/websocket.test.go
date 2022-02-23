package websockettest

import (
	stdlog "log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"

	baseexception "github.com/ikaiguang/go-srv-kit/api/base/exception"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
	websocketutil "github.com/ikaiguang/go-srv-kit/kratos/websocket"
	loghelper "github.com/ikaiguang/go-srv-kit/log/helper"
)

// RunTestWebsocket 测试websocket
func RunTestWebsocket() {
	processResp, err := TestWebsocket()
	if err != nil {
		stdlog.Printf("%+v\n", err)
		return
	}
	jsonContent, _ := apputil.JSON(processResp)
	stdlog.Println(string(jsonContent))
}

// TestWebsocket 测试websocket
func TestWebsocket() (processResp interface{}, err error) {
	urlPath := "/api/v1/testdata/websocket"

	// 开启ws
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: urlPath}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		err = errors.WithStack(err)
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
					loghelper.Infow("websocket close", wsErr.Error())
					break
				}
				err = errorutil.InternalServer(baseexception.ERROR_INTERNAL_SERVER.String(), "ws读取信息失败", wsErr)
				loghelper.Error(err)
				return
			}

			loghelper.Infow("ws-message type", messageType)
			loghelper.Infow("ws-message message", string(messageBytes))
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
					loghelper.Infow("websocket close", wsErr.Error())
					break
				}
				err = errorutil.InternalServer(baseexception.ERROR_INTERNAL_SERVER.String(), "ws读取信息失败", err)
				loghelper.Error(err)
				return
			}
		}
		counter++
	}

	return processResp, err
}