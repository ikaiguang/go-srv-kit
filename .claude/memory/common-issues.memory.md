# 常见问题和解决方案

## Wire 依赖注入问题

### 问题：cycle detected
**原因**：循环依赖
**解决**：
- 检查依赖关系，找出循环引用
- 重构代码，引入中间层或接口解耦

### 问题：no provider found for *biz.XxxBiz
**原因**：返回类型不匹配或缺少 Provider
**解决**：
- 检查函数返回类型是否正确
- 检查是否在 wire.Build 中添加了 Provider

## Proto 生成问题

### 问题：protoc 命令找不到
**解决**：
```bash
# 检查 protoc 是否安装
protoc --version

# 重新运行初始化
make init
```

### 问题：生成的代码路径不对
**解决**：检查 `paths=source_relative` 参数

## GORM 问题

### 问题：记录未找到
**错误处理**：
```go
if errors.Is(err, gorm.ErrRecordNotFound) {
    return nil, errorpkg.ErrorNotFound("not found")
}
```

### 问题：软删除数据仍被查询
**解决**：使用 `Unscoped()` 查询所有数据
```go
db.Unscoped().Find(&users)
```

## Context 传递问题

### 问题：TraceID 丢失
**解决**：始终使用 `WithContext(ctx)` 传递 Context
```go
d.db.WithContext(ctx).Create(&user)
```

## 并发安全问题

### 问题：panic in goroutine
**解决**：使用 GoSafe 包装 goroutine
```go
import threadpkg "github.com/ikaiguang/go-srv-kit/kit/thread"
threadpkg.GoSafe(func() {
    // 你的代码
})
```
