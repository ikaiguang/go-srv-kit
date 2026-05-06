# string

字符串 snake/camel 转换和通用字符串转换辅助。

## 基础用法

```go
snake := stringpkg.ToSnake("UserID")
camel := stringpkg.ToCamel("user_id")
text := stringpkg.ToString(123)
```

## 注意事项

涉及用户隐私展示时优先使用脱敏函数，不要直接日志输出原文。

## 验证

```bash
go test ./string
```
