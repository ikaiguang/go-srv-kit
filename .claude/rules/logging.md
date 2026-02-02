# 日志规范

## 日志级别使用

| 级别 | 用途 | 示例 |
|------|------|------|
| Debug | 调试信息，开发环境使用 | 请求参数详情 |
| Info | 关键业务流程 | 用户登录成功、订单创建 |
| Warn | 警告，但不影响运行 | 重试操作、降级服务 |
| Error | 错误，需要关注 | 数据库连接失败、外部调用失败 |
| Fatal | 致命错误，程序退出 | 无法启动服务 |

## 日志使用

### 基本用法

```go
import (
    log "github.com/go-kratos/kratos/v2/log"
)

// 1. Context 日志（推荐）
logpkg.WithContext(ctx).Info("user login", "user_id", userId)

// 2. 结构化日志
logpkg.WithContext(ctx).Infow("user created",
    "user_id", user.ID,
    "username", user.Username,
    "email", user.Email,
)

// 3. 格式化日志
log.Infof("user %s login at %s", username, time.Now())
```

### 获取 Logger

```go
import setuputil "github.com/ikaiguang/go-srv-kit/service/setup"

// 在 Wire 中注入
logger, err := launcherManager.GetLogger()

// 或使用辅助 Logger
loggerHelper, err := launcherManager.GetLoggerForHelper()
```

### 分层日志

#### Service Layer

```go
func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
    logpkg.WithContext(ctx).Infow("create user request", "username", req.GetUsername())

    result, err := s.userBiz.CreateUser(ctx, param)
    if err != nil {
        logpkg.WithContext(ctx).Errorw("create user failed",
            "error", err.Error(),
            "username", req.GetUsername(),
        )
        return nil, err
    }

    logpkg.WithContext(ctx).Infow("create user success", "user_id", result.ID)
    return result, nil
}
```

#### Business Layer

```go
func (b *userBiz) CreateUser(ctx context.Context, param *bo.CreateUserParam) (*bo.CreateUserResult, error) {
    logpkg.WithContext(ctx).Debugw("create user in biz", "param", param)

    // 业务逻辑
    result, err := b.userRepo.CreateUser(ctx, param)
    if err != nil {
        logpkg.WithContext(ctx).Errorw("create user failed in repo", "error", err)
        return nil, err
    }

    return result, nil
}
```

#### Data Layer

```go
func (d *userData) CreateUser(ctx context.Context, param *bo.CreateUserParam) (*bo.CreateUserResult, error) {
    user := &po.User{
        Username: param.Username,
    }

    log.Debugw("insert user to database", "user", user)

    if err := d.db.WithContext(ctx).Create(user).Error; err != nil {
        log.Errorw("insert user failed", "error", err)
        return nil, err
    }

    log.Infow("insert user success", "user_id", user.ID)
    return &bo.CreateUserResult{ID: user.ID}, nil
}
```

## 分类日志

### GORM 日志

```go
import gormlog "github.com/ikaiguang/go-srv-kit/data/gorm/log"

// GORM 会自动记录到 gorm 分类日志
db.Use(gormlog.NewLogger(logger))
```

### RabbitMQ 日志

```go
import "github.com/ikaiguang/go-srv-kit/data/rabbitmq/log"

// RabbitMQ 会自动记录到 rabbitmq 分类日志
rabbitmq.SetLogger(logger)
```

### 自定义分类

```go
// 创建带前缀的 Logger
customLogger := log.WithHelper(logger)
customLogger.Infow("custom log", "key", "value")
```

## 日志输出

### 开发环境

- 输出到 Console
- 级别：Debug
- 格式：JSON

### 生产环境

- 输出到文件（带轮转）
- 级别：Info
- 格式：JSON
- 文件位置：`runtime/logs/`

## 日志最佳实践

### 1. 使用 Context

```go
// 好的实践
logpkg.WithContext(ctx).Infow("user login", "user_id", userId)

// 不好的实践
log.Info("user login", userId)  // 没有 TraceID
```

### 2. 结构化字段

```go
// 好的实践
logpkg.WithContext(ctx).Infow("order created",
    "order_id", order.ID,
    "user_id", order.UserID,
    "amount", order.Amount,
)

// 不好的实践
log.Infof("order created: %+v", order)  // 不便于查询
```

### 3. 敏感信息脱敏

```go
import "github.com/ikaiguang/go-srv-kit/kit/stringutil"

logpkg.WithContext(ctx).Infow("user login",
    "user_id", userId,
    "password", stringutil.MaskPassword(password),  // *******
    "phone", stringutilMaskPhone(phone),            // 138****5678
)
```

### 4. 错误日志

```go
// 记录错误堆栈
logpkg.WithContext(ctx).Errorw("operation failed",
    "error", err,
    "stack", stringutil.GetStackTrace(2),
)
```

## 日志文件

### 日志文件命名

```
runtime/logs/
├── ping-service_app_2024-01-01.log         # 应用日志
├── ping-service_gorm_2024-01-01.log        # GORM 日志
├── ping-service_rabbitmq_2024-01-01.log    # RabbitMQ 日志
└── ping-service_error_2024-01-01.log       # 错误日志
```

### 日志轮转

- 按天轮转
- 保留 30 天
- 单个文件最大 100MB
