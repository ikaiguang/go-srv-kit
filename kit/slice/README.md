# slice

切片反转、包含判断等基础操作，包含泛型辅助。

## 基础用法

```go
ok := slicepkg.Contains([]string{"a", "b"}, "a")
```

## 注意事项

大切片高频查找可改用 map，避免线性扫描造成性能问题。

## 验证

```bash
go test ./slice
```
