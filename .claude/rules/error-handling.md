# 错误处理规范

## 使用 kratos/error 包

### 基本错误类型

```go
import errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"

// 400 Bad Request
return errorpkg.ErrorBadRequest("invalid parameter")

// 401 Unauthorized
return errorpkg.ErrorUnauthorized("token is invalid")

// 403 Forbidden
return errorpkg.ErrorForbidden("access denied")

// 404 Not Found
return errorpkg.ErrorNotFound("user not found")

// 409 Conflict
return errorpkg.ErrorConflict("user already exists")

// 500 Internal Server Error
return errorpkg.ErrorInternal("database connection failed")
```

### 带元数据的错误

```go
import "github.com/go-kratos/kratos/v2/metadata"

metadata := metadata.New(map[string]string{
    "user_id": "123",
    "action":  "create_order",
})

return errorpkg.WrapWithMetadata(err, metadata)
```

### 格式化错误（带堆栈）

```go
if err != nil {
    return errorpkg.FormatError(err)
}
```

### 错误检查

```go
import "github.com/go-kratos/kratos/v2/errors"

// 检查错误码
if errors.Is(err, errorpkg.ErrorBadRequest("")) {
    // 处理 BadRequest 错误
}

// 获取错误码
code := errors.Code(err)
reason := errors.Reason(err)
```

## 业务错误定义

### 定义业务错误

在 `internal/errors/` 或 `kratos/error/errors.go`:

```go
package errors

const (
    // ErrUserNotFound 用户不存在
    ErrUserNotFound = 10001
    // ErrUserAlreadyExists 用户已存在
    ErrUserAlreadyExists = 10002
    // ErrInvalidPassword 密码错误
    ErrInvalidPassword = 10003
)

// UserNotFound 用户不存在
func UserNotFound() error {
    return errorpkg.ErrorBadRequest("user not found").
        WithCode(ErrUserNotFound).
        WithReason("USER_NOT_FOUND")
}

// UserAlreadyExists 用户已存在
func UserAlreadyExists() error {
    return errorpkg.ErrorConflict("user already exists").
        WithCode(ErrUserAlreadyExists).
        WithReason("USER_ALREADY_EXISTS")
}
```

### 使用业务错误

```go
func (b *userBiz) GetUser(ctx context.Context, id uint) (*bo.User, error) {
    user, err := b.userRepo.GetUser(ctx, id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, customErrors.UserNotFound()
        }
        return nil, errorpkg.WrapWithMetadata(err, nil)
    }
    return user, nil
}
```

## Proto 错误定义

### 定义错误 Proto

在 `api/{service}/v1/errors/`:

```protobuf
syntax = "proto3";

package api.user.service.v1;

import "errors/errors.proto";

option go_package = "github.com/ikaiguang/go-srv-kit/api/user-service/v1;v1";

// 用户服务错误
enum ErrorReason {
  // 用户不存在
  USER_NOT_FOUND = 0 [(errors.code) = 404];

  // 用户已存在
  USER_ALREADY_EXISTS = 1 [(errors.code) = 409];

  // 密码错误
  INVALID_PASSWORD = 2 [(errors.code) = 400];
}
```

### 使用 Proto 错误

```go
import v1 "github.com/ikaiguang/go-srv-kit/api/user-service/v1"

return v1.ErrorUserNotFound("user not found")
```

## 错误日志

### 记录错误

```go
func (s *service) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
    user, err := s.userBiz.CreateUser(ctx, param)
    if err != nil {
        // 记录错误日志
        logpkg.WithContext(ctx).Errorf("create user failed: %v", err)

        // 返回给用户的错误
        return nil, errorpkg.WrapWithMetadata(err, metadata)
    }
    return user, nil
}
```

## 错误处理最佳实践

### Service Layer

```go
func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
    // 1. 参数验证
    if req.GetUsername() == "" {
        return nil, errorpkg.ErrorBadRequest("username is required")
    }

    // 2. 转换为 BO
    param := dto.ToBoCreateUserParam(req)

    // 3. 调用业务逻辑
    result, err := s.userBiz.CreateUser(ctx, param)
    if err != nil {
        // 业务错误直接返回
        return nil, err
    }

    // 4. 转换为 Proto Response
    return dto.ToProtoCreateUserResp(result), nil
}
```

### Business Layer

```go
func (b *userBiz) CreateUser(ctx context.Context, param *bo.CreateUserParam) (*bo.CreateUserResult, error) {
    // 1. 业务验证
    exists, err := b.userRepo.CheckUserExists(ctx, param.Username)
    if err != nil {
        return nil, errorpkg.FormatError(err)
    }
    if exists {
        return nil, customErrors.UserAlreadyExists()
    }

    // 2. 调用 Data 层
    result, err := b.userRepo.CreateUser(ctx, param)
    if err != nil {
        return nil, errorpkg.WrapWithMetadata(err, nil)
    }

    return result, nil
}
```

### Data Layer

```go
func (d *userData) CreateUser(ctx context.Context, param *bo.CreateUserParam) (*bo.CreateUserResult, error) {
    user := &po.User{
        Username: param.Username,
        Email:    param.Email,
    }

    if err := d.db.WithContext(ctx).Create(user).Error; err != nil {
        // 判断是否是唯一约束冲突
        if errors.Is(err, gorm.ErrDuplicatedKey) {
            return nil, customErrors.UserAlreadyExists()
        }
        return nil, errorpkg.FormatError(err)
    }

    return &bo.CreateUserResult{ID: user.ID}, nil
}
```

## Panic 恢复

框架已内置 Recovery 中间件，但关键位置也需要：

```go
func (s *service) SomeMethod(ctx context.Context, req *pb.SomeReq) (*pb.SomeResp, error) {
    defer func() {
        if r := recover(); r != nil {
            logpkg.WithContext(ctx).Errorf("panic recovered: %v", r)
        }
    }()

    // 业务逻辑
}
```
