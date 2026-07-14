# url

URL 编码、query 参数拼接和请求 URL 生成。

## 基础用法

```go
values := url.Values{"q": {"hello world"}}
query := urlpkg.EncodeValues(values)
requestURL := urlpkg.GenRequestURL("https://api.example.com", "/v1/users")
```

## 注意事项

拼接 URL 前确认 endpoint 和 path 是否已包含斜杠或 query。

## 验证

```bash
go test ./url
```
