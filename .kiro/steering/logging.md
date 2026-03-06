---
inclusion: fileMatch
fileMatchPattern: "**/*.go"
---

# 日志规范

## 日志级别

| 级别 | 用途 |
|------|------|
| Debug | 调试信息，开发环境 |
| Info | 关键业务流程（登录、订单创建） |
| Warn | 警告，不影响运行（重试、降级） |
| Error | 错误，需要关注（DB 连接失败） |
| Fatal | 致命错误，程序退出 |

## 使用方式

```go
// 推荐：Context 日志（携带 TraceID）
logpkg.WithContext(ctx).Infow("user created", "user_id", user.ID)
logpkg.WithContext(ctx).Errorw("operation failed", "error", err.Error())
logpkg.WithContext(ctx).Debugw("request params", "params", req)
```

## 最佳实践

- 始终使用 `WithContext(ctx)` 传递 TraceID
- 使用结构化字段（key-value），不用 `Infof` 格式化
- 敏感信息脱敏：`stringutil.MaskPassword(password)`, `stringutil.MaskPhone(phone)`
- 日志文件按天轮转，保留 30 天，单文件最大 100MB
- 分类日志：app、gorm、rabbitmq、error
