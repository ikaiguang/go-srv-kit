package connectionpkg

import (
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultDialTimeout = 3 * time.Second

// IsWebSocketConn 是否websocket
func IsWebSocketConn(r *http.Request) bool {
	if r == nil {
		return false
	}
	//if r.Method == http.MethodGet &&
	//	headerpkg.ContainsValue(r.Header, headerpkg.WebsocketConnection, "upgrade") &&
	//	headerpkg.ContainsValue(r.Header, headerpkg.WebsocketUpgrade, "websocket") &&
	//	r.Header.Get(headerpkg.WebsocketSecVersion) != "" &&
	//	r.Header.Get(headerpkg.WebsocketSecKey) != "" {
	//	return true
	//}
	if headerContainsToken(r.Header.Get("Connection"), "upgrade") &&
		strings.EqualFold(r.Header.Get("Upgrade"), "websocket") {
		return true
	}
	return false
}

func headerContainsToken(value, target string) bool {
	for _, token := range strings.Split(value, ",") {
		if strings.EqualFold(strings.TrimSpace(token), target) {
			return true
		}
	}
	return false
}

// IsEndpointReachable reports whether the endpoint accepts a TCP connection.
func IsEndpointReachable(endpoint string) (bool, error) {
	if !strings.Contains(endpoint, "://") && !strings.HasPrefix(endpoint, "//") {
		endpoint = "//" + endpoint
	}
	u, err := url.Parse(endpoint)
	if err != nil {
		return false, err
	}
	host := u.Hostname()
	if host == "" {
		return false, &url.Error{Op: "parse", URL: endpoint, Err: errors.New("endpoint host is empty")}
	}
	port := u.Port()
	if port == "" {
		switch strings.ToLower(u.Scheme) {
		case "https", "wss":
			port = "443"
		default:
			port = "80"
		}
	}
	return IsConnectionReachable(net.JoinHostPort(host, port))
}

// Deprecated: use IsEndpointReachable instead.
func CheckEndpointValidity(endpoint string) (bool, error) {
	return IsEndpointReachable(endpoint)
}

// IsConnectionReachable reports whether address accepts a TCP connection.
func IsConnectionReachable(address string) (bool, error) {
	conn, err := net.DialTimeout("tcp", address, defaultDialTimeout)
	if err != nil {
		return false, err
	}
	defer func() { _ = conn.Close() }()

	return true, nil
}

// Deprecated: use IsConnectionReachable instead.
func IsValidConnection(address string) (bool, error) {
	return IsConnectionReachable(address)
}
