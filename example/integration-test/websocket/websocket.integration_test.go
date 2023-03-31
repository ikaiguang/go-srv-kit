package testwebsocket

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
	"time"

	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	websocketutil "github.com/ikaiguang/go-srv-kit/kratos/websocket"
	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// go test -v -count=1 ./example/integration-test/websocket -test.run=TestApi_Websocket
func TestApi_Websocket(t *testing.T) {
	urlPath := "/api/v1/testdata/websocket"

	// 开启ws
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8081", Path: urlPath}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	require.Nil(t, err)
	defer func() { _ = c.Close() }()

	runTestWebsocket(c)
}

// go test -v -count=1 ./example/integration-test/websocket -test.run=TestWs_Websocket
func TestWs_Websocket(t *testing.T) {
	urlPath := "/ws/v1/websocket"

	// 开启ws
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8081", Path: urlPath}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	require.Nil(t, err)
	defer func() { _ = c.Close() }()

	runTestWebsocket(c)
}

// runTestWebsocket ...
func runTestWebsocket(c *websocket.Conn) {
	// 读取信息
	var (
		err  error
		done = make(chan struct{})
	)
	go func() {
		defer close(done)
		for {
			messageType, messageBytes, wsErr := c.ReadMessage()
			if wsErr != nil {
				if websocketutil.IsCloseError(wsErr) {
					logutil.Infow("websocket close", wsErr.Error())
					break
				}
				err = errorutil.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "ws读取信息失败", wsErr)
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
				err = errorutil.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "ws读取信息失败", err)
				logutil.Error(err)
				return
			}
		}
		counter++
	}
}
