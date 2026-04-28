package headerpkg

import (
	"net/http"
)

const (
	_defaultIsWebsocketValue = "1"
)

// SetIsWebsocket 设置是否websocket
func SetIsWebsocket(header http.Header) {
	header.Set(IsWebsocket, _defaultIsWebsocketValue)
}

// GetIsWebsocket 获取是否websocket
func GetIsWebsocket(header http.Header) bool {
	return header.Get(IsWebsocket) == _defaultIsWebsocketValue
}
