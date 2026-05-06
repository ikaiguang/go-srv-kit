# path

返回当前工具包源码所在目录。

## 基础用法

```go
p := pathpkg.Path()
```

## 注意事项

该函数基于 `runtime.Caller`，适合调试和测试辅助，不建议作为业务运行目录配置。

## 验证

```bash
go test ./path
```
