# curl

`kit/curl/` 提供轻量级 HTTP 请求辅助函数，基于标准库 `net/http` 封装基础请求创建、客户端初始化和请求执行。

## 主要内容

- `NewGetRequest()` / `NewPostRequest()` / `NewRequest()`
- `NewPutRequest()` / `NewPatchRequest()` / `NewDeleteRequest()`
- `NewGetRequestContext()` / `NewPostRequestContext()` / `NewRequestContext()`
- `NewHTTPClient()`
- `Do()`
- 常用请求头和内容类型常量

## 适用场景

- 需要轻量 HTTP 调用，而不想直接重复写标准库样板代码
- 需要显式控制超时或 TLS 验证选项

## 基础用法

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

req, err := curlpkg.NewGetRequestContext(ctx, "https://example.com", nil)
code, body, err := curlpkg.Do(req, curlpkg.WithTimeout(3*time.Second))
if !curlpkg.IsSuccessCode(code) {
	return curlpkg.ErrRequestFailure(code)
}
```

## 注意事项

- 默认不跳过 TLS 证书校验。
- `WithInsecureSkipVerify` 仅限开发或测试环境。
- 优先使用 `New*RequestContext`，让调用方能控制超时和取消。

## 验证

```bash
go test ./curl
```

## 参考

- 标准库 `net/http`
- 如需更完整的高级 HTTP 客户端能力，可参考 [go-resty/resty](https://github.com/go-resty/resty)
