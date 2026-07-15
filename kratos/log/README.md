# log

`log` 包的包名为 `logpkg`，提供 Kratos `log.Logger` 实现和全局日志 helper，包括控制台日志、文件轮转日志、多路日志、no-op logger、异步 writer 和带 context 的日志方法。适合服务启动时构建统一 logger，并在中间件和业务代码中复用。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/kratos
```

```go
import logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
```

## 核心能力

- `NewStdLogger`：创建输出到 stderr 的 zap logger，支持 console/json encoder。
- `NewFileLogger`：创建 JSON 文件日志，支持按时间、大小和保留策略轮转。
- `NewMultiLogger`：将多路 `log.Logger` 组合为一个 logger。
- `NewDummyLogger`、`NewNopLogger`：创建 no-op logger。
- `Setup`、`Debug`、`Infow`、`Errorw` 等：设置和使用全局 `log.Helper`。
- `LogWithContext`、`InfowWithContext`、`ErrorwWithContext` 等：带 context 输出日志。
- `NewAsyncWriter`：将 `io.Writer` 包装为异步写入器。
- `WithWriter`、`WithLoggerKey`、`WithFilenameSuffix`、`WithTimeFormat`：定制 writer、日志字段名、文件名后缀和时间格式。

## 快速使用

```go
stdLogger, err := logpkg.NewStdLogger(&logpkg.ConfigStd{
    Level:      log.LevelInfo,
    CallerSkip: logpkg.DefaultCallerSkip,
})
if err != nil {
    return err
}
defer stdLogger.Close()

logpkg.Setup(stdLogger)
logpkg.Infow("event", "startup")
```

```go
fileLogger, err := logpkg.NewFileLogger(&logpkg.ConfigFile{
    Level:          log.LevelInfo,
    Dir:            "./bin",
    Filename:       "app",
    RotateSize:     50 << 20,
    StorageCounter: 7,
})
if err != nil {
    return err
}
defer fileLogger.Close()
```

## 测试

```bash
go test ./log
```

## 注意事项

- `NewFileLogger` 会创建和写入本地日志文件，部署前确认目录权限、磁盘容量和保留策略。
- `AsyncWriter` 在缓冲满时会主动丢弃日志并返回成功写入长度，适合可容忍丢日志的场景。
- key/value 日志参数数量为奇数时会追加 `KEYVALS UNPAIRED`，调用方应尽量保持成对传入。
- 不要记录未脱敏的密码、Token、密钥、手机号或连接串。
