# 代码模式记忆

## Service 层模式

### 标准 API 实现
```go
func (s *xxxService) XxxMethod(ctx context.Context, req *pb.XxxReq) (*pb.XxxResp, error) {
    // 1. 参数验证
    if req.GetField() == "" {
        return nil, errorpkg.ErrorBadRequest("field is required")
    }

    // 2. DTO → BO
    param := dto.ToBoXxxParam(req)

    // 3. 调用业务逻辑
    result, err := s.xxxBiz.XxxMethod(ctx, param)
    if err != nil {
        logpkg.WithContext(ctx).Errorw("xxx failed", "error", err)
        return nil, err
    }

    // 4. BO → Proto
    return dto.ToProtoXxxResp(result), nil
}
```

## Business 层模式

### 带验证的业务逻辑
```go
func (b *xxxBiz) XxxMethod(ctx context.Context, param *bo.XxxParam) (*bo.XxxResult, error) {
    // 1. 业务验证
    exists, err := b.xxxRepo.CheckExists(ctx, param.Key)
    if err != nil {
        return nil, errorpkg.FormatError(err)
    }
    if exists {
        return nil, customErrors.AlreadyExists()
    }

    // 2. 调用 Data 层
    result, err := b.xxxRepo.XxxMethod(ctx, param)
    if err != nil {
        return nil, errorpkg.WrapWithMetadata(err, nil)
    }

    return result, nil
}
```

## Data 层模式

### 标准查询
```go
func (d *xxxData) Get(ctx context.Context, id uint) (*bo.Result, error) {
    var po po.Xxx
    err := d.db.WithContext(ctx).First(&po, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errorpkg.ErrorNotFound("not found")
        }
        return nil, errorpkg.FormatError(err)
    }
    return toBoResult(&po), nil
}
```

### 事务操作
```go
err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user).Error; err != nil {
        return err
    }
    profile.UserID = user.ID
    return tx.Create(&profile).Error
})
```

## 错误处理模式

### Data 层错误转换
```go
if err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errorpkg.ErrorNotFound("xxx not found")
    }
    if errors.Is(err, gorm.ErrDuplicatedKey) {
        return nil, customErrors.AlreadyExists()
    }
    return nil, errorpkg.FormatError(err)
}
```

### Service 层错误包装
```go
if err != nil {
    logpkg.WithContext(ctx).Errorw("operation failed", "error", err)
    return nil, errorpkg.WrapWithMetadata(err, metadata)
}
```
