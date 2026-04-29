# debug - 调试日志

`debug/` 提供全局调试日志工具，支持多级别输出。

## 包名

```go
import debugpkg "github.com/ikaiguang/go-srv-kit/kratos/debug"
```

## 使用

```go
// 初始化（通常由 LauncherManager 自动完成）
debugpkg.Setup(logger)

// 调试日志
debugpkg.Debugw("key", "value")
debugpkg.Infow("msg", "info message")
debugpkg.Warnw("msg", "warning")
debugpkg.Errorw("msg", "error occurred")
```

支持 `Debug`/`Info`/`Warn`/`Error`/`Fatal` 五个级别，每个级别有 `Xxx`、`Xxxf`、`Xxxw` 三种调用方式。
