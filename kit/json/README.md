# json

JSON 序列化辅助，输出时不转义 HTML 字符。

## 基础用法

```go
data, err := jsonpkg.MarshalWithoutEscapeHTML(v)
data, err = jsonpkg.MarshalIndentWithoutEscapeHTML(v, "", "  ")
```

## 注意事项

函数会保留 HTML 字符原样输出，并复制 buffer 内容后再归还复用池。

## 验证

```bash
go test ./json
```
