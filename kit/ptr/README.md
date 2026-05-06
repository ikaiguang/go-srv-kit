# ptr

基础类型取指针、指针取值和泛型指针辅助。

## 基础用法

```go
namePtr := ptrpkg.Ptr("alice")
name := ptrpkg.Value(namePtr)
```

## 注意事项

区分零值和 nil 的业务含义，尤其是 PATCH/部分更新 DTO。

## 验证

```bash
go test ./ptr
```
