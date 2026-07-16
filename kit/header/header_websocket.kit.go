package headerpkg

import (
	"net/http"
)

const (
	_defaultIsWebsocketValue = "1"
)

// SetIsWebSocket marks a header as a WebSocket request.
func SetIsWebSocket(header http.Header) {
	if header == nil {
		return
	}
	header.Set(IsWebsocket, _defaultIsWebsocketValue)
}

// GetIsWebSocket reports whether the header has the WebSocket marker.
func GetIsWebSocket(header http.Header) bool {
	return header.Get(IsWebsocket) == _defaultIsWebsocketValue
}

// Deprecated: use SetIsWebSocket instead.
func SetIsWebsocket(header http.Header) { SetIsWebSocket(header) }

// Deprecated: use GetIsWebSocket instead.
func GetIsWebsocket(header http.Header) bool { return GetIsWebSocket(header) }
