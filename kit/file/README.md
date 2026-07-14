# file

复制文件、移动文件到指定目录、计算文件 hash。

## 基础用法

```go
err := filepkg.CopyFile("./a.txt", "./b.txt")
target, err := filepkg.MoveFileToDir("./b.txt", "./runtime")
sum, size, err := filepkg.Hash("./b.txt")
```

## 注意事项

`CopyFile` 和 `MoveFileToDir` 会自动创建目标目录。移动文件依赖 `os.Rename`，跨设备移动可能失败。

## 验证

```bash
go test ./file
```
