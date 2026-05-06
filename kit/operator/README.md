# operator

三元表达式辅助函数。

## 基础用法

```go
name := operatorpkg.Ternary(ok, "yes", "no")
```

## 注意事项

只适合简单值选择；复杂分支使用普通 `if` 更清晰。

## 验证

```bash
go test ./operator
```
