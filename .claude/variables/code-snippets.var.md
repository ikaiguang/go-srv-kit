# 代码片段变量

## 导入语句模板
```go
import (
    // 标准库
    "context"
    "fmt"

    // 第三方库
    "github.com/go-kratos/kratos/v2/log"
    "google.golang.org/protobuf/proto"

    // 项目内部
    "github.com/ikaiguang/go-srv-kit/kratos/error"
    "github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/bo"

    // 匿名包
    _ "github.com/go-sql-driver/mysql"
)
```

## 错误处理片段
```go
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

## 日志记录片段
```go
// Info 日志
log.Context(ctx).Infow("user created",
    "user_id", user.ID,
    "username", user.Username,
)

// Error 日志
log.Context(ctx).Errorw("operation failed",
    "error", err.Error(),
    "user_id", userId,
)

// Debug 日志
log.Context(ctx).Debugw("request params", "params", req)
```

## GORM 操作片段
```go
// 创建
err := d.db.WithContext(ctx).Create(&user).Error

// 查询单条
err := d.db.WithContext(ctx).First(&user, id).Error

// 查询列表
err := d.db.WithContext(ctx).Find(&users).Error

// 更新
err := d.db.WithContext(ctx).Model(&user).Updates(updates).Error

// 删除（软删除）
err := d.db.WithContext(ctx).Delete(&user, id).Error

// 事务
err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
    // 事务操作
    return nil
})
```

## Redis 操作片段
```go
// 设置
err := r.redis.Set(ctx, key, value, expiration).Err()

// 获取
val, err := r.redis.Get(ctx, key).Result()

// 删除
err := r.redis.Del(ctx, key).Err()

// 批量获取
vals, err := r.redis.MGet(ctx, keys...).Result()

// 分布式锁
locked, err := r.redis.SetNX(ctx, lockKey, 1, expiration).Result()
```
