# curl

`kit/curl/` 提供轻量级 HTTP 请求辅助函数，基于标准库 `net/http` 封装基础请求创建、客户端初始化和请求执行。

## 主要内容

- `NewGetRequest()` / `NewPostRequest()` / `NewRequest()`
- `NewHTTPClient()`
- `Do()`
- 常用请求头和内容类型常量

## 适用场景

- 需要轻量 HTTP 调用，而不想直接重复写标准库样板代码
- 需要显式控制超时或 TLS 验证选项

## 参考

- 标准库 `net/http`
- 如需更完整的高级 HTTP 客户端能力，可参考 [go-resty/resty](https://github.com/go-resty/resty)
