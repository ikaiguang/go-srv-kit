# time

时间格式化、解析、时间范围和常用时间计算。

## 基础用法

```go
today := timepkg.Today()
end := timepkg.EndOfDay(time.Now())
text := timepkg.FormatRFC3339(time.Now())
```

## 注意事项

跨时区业务要明确 location；不要隐式依赖本机时区。

## 验证

```bash
go test ./time
```
