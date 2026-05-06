# chinese

GB18030、GBK、HZGB2312 与 UTF-8 之间转换，并检测字符串是否为 UTF-8 或 GBK。

## 基础用法

```go
gbkBytes, err := chinesepkg.Utf8ToGbk([]byte("中文"))
utf8Bytes, err := chinesepkg.GbkToUtf8(gbkBytes)
ok := chinesepkg.IsGBK(string(gbkBytes))
```

## 注意事项

编码检测只能作为输入判断辅助；外部输入要按转换错误做兜底处理。

## 验证

```bash
go test ./chinese
```
