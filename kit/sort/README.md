# sort

排序辅助函数，包含基础类型和泛型排序。

## 基础用法

```go
sortpkg.Sort(values)
sortpkg.SortFunc(users, func(a, b User) int {
	return cmp.Compare(a.ID, b.ID)
})
```

## 注意事项

确认函数是否原地修改切片，避免调用方误用共享切片。

## 验证

```bash
go test ./sort
```
