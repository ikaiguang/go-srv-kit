# header

HTTP header 常量、Request ID 辅助函数和 header value 匹配。

## 基础用法

```go
if headerpkg.ContainsValue(req.Header, "Connection", "upgrade") {
	// ...
}
req.Header.Set(headerpkg.ContentType, headerpkg.ContentTypeJSON)
headerpkg.SetRequestID(req.Header, "request-id")
headerpkg.SetContentType(req.Header, headerpkg.ContentTypeJSON)
```

## 注意事项

`ContainsValue` 支持逗号分隔 header，适合 WebSocket 和 HTTP 协议头判断。

## 验证

```bash
go test ./header
```
