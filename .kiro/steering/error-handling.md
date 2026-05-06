---
inclusion: fileMatch
fileMatchPattern: "**/*.go"
---

# 错误处理规范

## 使用 kratos/error 包

```go
import errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"

errorpkg.ErrorBadRequest("invalid parameter")     // 400
errorpkg.ErrorUnauthorized("token is invalid")     // 401
errorpkg.ErrorForbidden("access denied")           // 403
errorpkg.ErrorNotFound("user not found")           // 404
errorpkg.ErrorConflict("user already exists")      // 409
errorpkg.ErrorInternal("database connection failed") // 500

// 带元数据
errorpkg.WrapWithMetadata(err, metadata)

// 带堆栈信息
errorpkg.FormatError(err)
```

## 分层错误处理

Service 层：参数验证 + 日志记录 + 返回业务错误
```go
if req.GetField() == "" {
    return nil, errorpkg.ErrorBadRequest("field is required")
}
result, err := s.xxxBiz.Method(ctx, param)
if err != nil {
    logpkg.WithContext(ctx).Errorw("operation failed", "error", err)
    return nil, err
}
```

Business 层：业务验证 + 错误包装
```go
if exists {
    return nil, customErrors.AlreadyExists()
}
result, err := b.xxxRepo.Method(ctx, param)
if err != nil {
    return nil, errorpkg.WrapWithMetadata(err, nil)
}
```

Data 层：GORM 错误转换
```go
if errors.Is(err, gorm.ErrRecordNotFound) {
    return nil, errorpkg.ErrorNotFound("not found")
}
if errors.Is(err, gorm.ErrDuplicatedKey) {
    return nil, customErrors.AlreadyExists()
}
return nil, errorpkg.FormatError(err)
```

## Proto 错误定义

```protobuf
enum ErrorReason {
  USER_NOT_FOUND = 0 [(errors.code) = 404];
}
```

使用: `v1.ErrorUserNotFound("user not found")`
