package websocketutil

import (
	"github.com/gorilla/websocket"

	connectionutil "github.com/ikaiguang/go-srv-kit/kit/connection"
)

// IsCloseError .
func IsCloseError(wsErr error) bool {
	isClose := websocket.IsCloseError(
		wsErr,
		websocket.CloseNormalClosure,
		websocket.CloseGoingAway,
		websocket.CloseProtocolError,
		websocket.CloseUnsupportedData,
		websocket.CloseNoStatusReceived,
		websocket.CloseAbnormalClosure,
	)
	if isClose {
		return isClose
	}

	if connectionutil.IsConnCloseErr(wsErr) {
		return true
	}
	return false
}
