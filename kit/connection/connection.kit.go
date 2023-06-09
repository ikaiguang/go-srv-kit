package connectionpkg

import (
	"net/http"
	"strings"
)

// IsWebSocketConn 是否websocket
func IsWebSocketConn(r *http.Request) bool {
	//if r.Method == http.MethodGet &&
	//	headerpkg.ContainsValue(r.Header, headerpkg.WebsocketConnection, "upgrade") &&
	//	headerpkg.ContainsValue(r.Header, headerpkg.WebsocketUpgrade, "websocket") &&
	//	r.Header.Get(headerpkg.WebsocketSecVersion) != "" &&
	//	r.Header.Get(headerpkg.WebsocketSecKey) != "" {
	//	return true
	//}
	if strings.Contains(strings.ToLower(r.Header.Get("Connection")), "upgrade") &&
		strings.EqualFold(r.Header.Get("Upgrade"), "websocket") {
		return true
	}
	return false
}
