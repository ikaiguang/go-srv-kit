package bizrepo

import (
	"context"
	stdhttp "net/http"

	"github.com/gorilla/websocket"
)

type WebsocketBizRepo interface {
	Wss(ctx context.Context, w stdhttp.ResponseWriter, r *stdhttp.Request) (err error)
	ProcessWssMsg(ctx context.Context, cc *websocket.Conn) error
}
