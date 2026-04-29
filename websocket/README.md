# websocket - WebSocket 支持

`websocket/` 提供 WebSocket 的基础依赖引入。

## 说明

当前仅引入 `github.com/gorilla/websocket` 依赖，实际的 WebSocket 工具实现位于 `kratos/websocket/` 包中。

## 使用

WebSocket 连接升级和管理请使用 `kratos/websocket/` 包：

```go
import websocketpkg "github.com/ikaiguang/go-srv-kit/kratos/websocket"

// 升级 HTTP 连接为 WebSocket
conn, err := websocketpkg.UpgradeConn(w, r, responseHeader)
```

## 参考

- gorilla/websocket：https://github.com/gorilla/websocket
- WebSocket 业务实现示例：`testdata/ping-service/internal/service/service/websocket.service.go`
