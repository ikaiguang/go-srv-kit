# connection

判断 WebSocket 请求头、检查 TCP 地址或 endpoint 是否可连接、识别连接关闭错误。

## 基础用法

```go
ok := connectionpkg.IsWebSocketConn(req)
alive, err := connectionpkg.IsValidConnection("127.0.0.1:8080")
alive, err = connectionpkg.CheckEndpointValidity("http://127.0.0.1:8080")
```

## 注意事项

连接检查会发起 TCP dial，只适合健康检查或调试路径；不要在高频请求链路中无缓存调用。

## 验证

```bash
go test ./connection
```
