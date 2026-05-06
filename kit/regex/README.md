# regex

正则匹配和常见格式校验。

## 基础用法

```go
ok := regexpkg.IsEmail("user@example.com")
```

## 注意事项

正则适合格式初筛，不代表业务真实性验证。

## 验证

```bash
go test ./regex
```
