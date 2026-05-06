# os

操作系统判断辅助函数。

## 基础用法

```go
if ospkg.IsWindows() {
	// windows specific logic
}
```

## 注意事项

只封装了当前运行时系统判断，复杂平台逻辑仍应由调用方显式处理。

## 验证

```bash
go test ./os
```
