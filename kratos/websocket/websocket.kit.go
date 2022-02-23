package websocketutil

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrade 升级http
var upgrade = &websocket.Upgrader{}

// DefaultUpgrade 默认升级
func DefaultUpgrade() *websocket.Upgrader {
	return upgrade
}

// UpgradeConn 升级链接
func UpgradeConn(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	return upgrade.Upgrade(w, r, responseHeader)
}
