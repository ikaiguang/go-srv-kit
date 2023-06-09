package websocketpkg

import (
	"github.com/gorilla/websocket"

	connectionpkg "github.com/ikaiguang/go-srv-kit/kit/connection"
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

	if connectionpkg.IsConnCloseErr(wsErr) {
		return true
	}
	return false
}
