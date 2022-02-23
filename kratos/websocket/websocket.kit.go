package websocketutil

import (
	"net/http"

	"github.com/gorilla/websocket"

	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// upgrade 升级http
var upgrade = &websocket.Upgrader{}

// DefaultUpgrade 默认升级
func DefaultUpgrade() *websocket.Upgrader {
	return upgrade
}

// UpgradeConn 升级链接
func UpgradeConn(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	cc, err := upgrade.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	headerutil.SetIsWebsocket(r.Header)
	return cc, err
}
