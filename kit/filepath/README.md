# filepath

遍历目录、读取目录内容、创建目录和重建目录。

## 基础用法

```go
paths, entries, err := filepathpkg.WalkDir("./docs")
err = filepathpkg.CreateDir("./runtime")
```

## 注意事项

路径判断受当前工作目录影响；`RenewDir` 会删除并重建目录，只对受控临时目录使用。

## 验证

```bash
go test ./filepath
```
