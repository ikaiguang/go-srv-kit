# reflect

反射辅助函数，用于默认值判断、空值判断、对象复制和创建同类型实例。

## 基础用法

```go
empty := reflectpkg.IsEmpty(v)
zero := reflectpkg.IsDefaultValue(v)
```

## 注意事项

反射逻辑要控制在边界层或通用工具层；业务主流程优先使用显式类型。nil 输入会按空值处理或返回 nil/false。

## 验证

```bash
go test ./reflect
```
