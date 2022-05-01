package connectionutil

import (
	"net/http"
	"strings"
)

// IsWebSocketConn 是否websocket
func IsWebSocketConn(r *http.Request) bool {
	//if r.Method == http.MethodGet &&
	//	headerutil.ContainsValue(r.Header, headerutil.WebsocketConnection, "upgrade") &&
	//	headerutil.ContainsValue(r.Header, headerutil.WebsocketUpgrade, "websocket") &&
	//	r.Header.Get(headerutil.WebsocketSecVersion) != "" &&
	//	r.Header.Get(headerutil.WebsocketSecKey) != "" {
	//	return true
	//}
	if strings.Contains(strings.ToLower(r.Header.Get("Connection")), "upgrade") &&
		strings.EqualFold(r.Header.Get("Upgrade"), "websocket") {
		return true
	}
	return false
}
