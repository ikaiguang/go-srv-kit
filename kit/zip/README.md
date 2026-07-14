# zip

zip 压缩或解压相关辅助。

## 基础用法

```go
err := zippkg.Zip("./source", "./target.zip")
err = zippkg.Unzip("./target.zip", "./output")
```

## 注意事项

- 解压外部 zip 时要防止路径穿越；输出目录应使用受控路径。
- 压缩文件时使用流式复制，避免一次性读入大文件。
- 目标 zip 所在目录会自动创建。

## 验证

```bash
go test ./zip
```
