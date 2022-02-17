package connectionutil

import (
	"net/http"

	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// IsWebSocketConn 是否websocket
func IsWebSocketConn(r *http.Request) bool {
	if r.Method == http.MethodGet &&
		headerutil.ContainsValue(r.Header, headerutil.WebsocketConnection, "upgrade") &&
		headerutil.ContainsValue(r.Header, headerutil.WebsocketUpgrade, "websocket") &&
		r.Header.Get(headerutil.WebsocketSecVersion) != "" &&
		r.Header.Get(headerutil.WebsocketSecKey) != "" {
		return true
	}
	return false
}
