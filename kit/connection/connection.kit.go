package connectionpkg

import (
	"net"
	"net/http"
	"net/url"
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

// CheckEndpointValidity ...
func CheckEndpointValidity(endpoint string) (bool, error) {
	if !strings.Contains(endpoint, "://") && !strings.HasPrefix(endpoint, "//") {
		endpoint = "//" + endpoint
	}
	u, err := url.Parse(endpoint)
	if err != nil {
		return false, err
	}
	addr := u.Host
	if !strings.Contains(addr, ":") {
		addr += ":80"
	}
	return IsValidConnection(addr)
}

// IsValidConnection 检查链接有效性
// @param address: hostname + ":" + port
func IsValidConnection(address string) (bool, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return false, err
	}
	defer func() { _ = conn.Close() }()

	return true, nil
}
