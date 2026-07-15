# websocket

`websocket` 包的包名为 `websocketpkg`，封装 `gorilla/websocket` 的默认 upgrader、HTTP 连接升级和常见关闭错误判断。适合在 Kratos HTTP handler 或标准 HTTP handler 中处理 WebSocket 连接。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/kratos
```

```go
import websocketpkg "github.com/ikaiguang/go-srv-kit/kratos/websocket"
```

## 核心能力

- `DefaultUpgrade`：返回包级默认 `*websocket.Upgrader`。
- `UpgradeConn`：将 HTTP 请求升级为 WebSocket 连接，并在请求 header 中标记 WebSocket。
- `IsCloseError`：判断 gorilla websocket 常见关闭错误和底层连接关闭错误。

## 快速使用

```go
conn, err := websocketpkg.UpgradeConn(w, r, nil)
if err != nil {
    return err
}
defer conn.Close()
```

```go
if websocketpkg.IsCloseError(err) {
    return nil
}
```

## 测试

```bash
go test ./websocket
```

## 注意事项

- 默认 upgrader 使用 `gorilla/websocket.Upgrader{}` 的默认配置；跨域校验、buffer、压缩等策略应在 `DefaultUpgrade()` 返回值上按服务需求设置。
- WebSocket 连接升级后，普通 HTTP response writer 不应再写入常规响应体。
