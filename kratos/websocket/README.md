# websocket - WebSocket 工具

`websocket/` 提供 WebSocket 连接升级工具。

## 包名

```go
import websocketpkg "github.com/ikaiguang/go-srv-kit/kratos/websocket"
```

## 使用

```go
// 升级 HTTP 连接为 WebSocket
conn, err := websocketpkg.UpgradeConn(w, r, responseHeader)

// 获取默认 Upgrader
upgrader := websocketpkg.DefaultUpgrade()
```

基于 [gorilla/websocket](https://github.com/gorilla/websocket)。
