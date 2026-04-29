# log - 结构化日志

`log/` 提供基于 Zap 的结构化日志，支持控制台和文件输出、Context 传递 TraceID。

## 包名

```go
import logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
```

## 使用

```go
// 带 Context 的日志（推荐，自动携带 TraceID）
logpkg.WithContext(ctx).Infow("user created", "user_id", 123)
logpkg.WithContext(ctx).Errorw("operation failed", "error", err.Error())

// 多日志输出
logger := logpkg.NewMultiLogger(consoleLogger, fileLogger)
```

## 日志级别

Debug → Info → Warn → Error → Fatal

## 日志 Key

| Key | 说明 |
|-----|------|
| `msg` | 日志消息 |
| `level` | 日志级别 |
| `time` | 时间戳 |
| `caller` | 调用位置 |
| `func` | 函数名 |
| `stack` | 堆栈信息 |

## 文件说明

| 文件 | 说明 |
|------|------|
| `log.kit.go` | 核心定义、MultiLogger、Zap 配置 |
| `log_console.kit.go` | 控制台日志 |
| `log_file.kit.go` | 文件日志（轮转） |
| `log_async.kit.go` | 异步日志 |
| `log_helper.kit.go` | 日志 Helper 封装 |
| `log_with_context.kit.go` | Context 日志（WithContext） |
| `log_dummy.kit.go` | 空日志（测试用） |
