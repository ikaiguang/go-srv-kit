# logger - 日志管理

`logger/` 提供日志管理器，支持多种日志输出目标和分类日志。

## 包名

```go
import loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
```

## LoggerManager

统一管理所有日志实例：

```go
loggerManager, err := loggerutil.NewLoggerManager(logConfig, appConfig)

// 获取不同用途的日志
logger, err := loggerManager.GetLogger()              // 基础日志
mwLogger, err := loggerManager.GetLoggerForMiddleware() // 中间件日志
helperLogger, err := loggerManager.GetLoggerForHelper() // Helper 日志
gormLogger, err := loggerManager.GetLoggerForGORM()     // GORM 日志
mqLogger, err := loggerManager.GetLoggerForRabbitmq()   // RabbitMQ 日志
```

## 日志分类

| 类型 | 说明 | 配置开关 |
|------|------|----------|
| Console | 控制台输出 | `log.console.enable` |
| File | 文件输出（按天轮转） | `log.file.enable` |
| GORM | 数据库操作日志 | `log.gorm.enable` |
| RabbitMQ | 消息队列日志 | `log.rabbitmq.enable` |

## Writer

日志管理器同时提供底层 Writer，供需要自定义日志格式的场景使用：

```go
writer, err := loggerManager.GetWriter()
gormWriter, err := loggerManager.GetWriterForGORM()
```
