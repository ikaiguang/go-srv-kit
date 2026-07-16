# debug

`debug` 包的包名为 `debugpkg`，提供一个可全局替换 logger 的调试日志 helper。默认 logger 为 no-op，调用 `Setup` 后会将 `Print`、`Debug`、`Info`、`Warn`、`Error`、`Fatal` 等方法转发到传入的 Kratos logger。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/kratos/v3
```

```go
import debugpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/debug"
```

## 核心能力

- `Setup`：设置调试日志 logger。
- `CloseDebug`：恢复默认 no-op logger。
- `Debug`、`Info`、`Warn`、`Error`、`Fatal` 及其 `f`、`w` 变体：输出调试日志。
- `DebugWithContext`、`InfoWithContext`、`ErrorWithContext` 等：带 context 输出日志。

## 快速使用

```go
logger, err := logpkg.NewStdLogger(&logpkg.ConfigStd{Level: log.LevelDebug})
if err != nil {
    return err
}
defer logger.Close()

debugpkg.Setup(logger)
defer debugpkg.CloseDebug()

debugpkg.Debugw("event", "startup")
```

## 测试

```bash
go test ./debug
```

## 注意事项

- 默认 logger 不输出任何内容，必须调用 `Setup` 才会看到日志。
- `Fatal`、`Fatalf`、`Fatalw` 会沿用底层 logger 的 fatal 行为，业务逻辑中不要把普通错误写成 fatal。
- `WithUseJSONFormat` 当前只定义了 option 类型，未被 `Setup` 使用。
