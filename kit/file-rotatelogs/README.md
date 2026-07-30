# file-rotatelogs

`file-rotatelogs` 是一个支持按时间或文件大小轮转的 `io.WriteCloser`，文件名使用
[`strftime`](https://github.com/lestrrat-go/strftime) 格式。

## 基础用法

```go
writer, err := rotatelogs.New(
	"runtime/logs/app-%Y%m%d.log",
	rotatelogs.WithRotationTime(24*time.Hour),
	rotatelogs.WithRotationSize(100<<20),
	rotatelogs.WithRotationCount(30),
	rotatelogs.WithLinkName("runtime/logs/app.log"),
)
if err != nil {
	return err
}
defer writer.Close()

log.SetOutput(writer)
```

当同一个时间片内发生按大小轮转或手动调用 `Rotate()` 时，新文件依次追加
`.1`、`.2` 等代际后缀。`WithLinkName` 创建的符号链接始终指向当前文件。

## 保留策略

- `WithMaxAge` 按修改时间删除过期文件。
- `WithRotationCount` 只保留修改时间最新的指定数量文件，包含当前文件和代际文件。
- `WithMaxAge` 与 `WithRotationCount` 不能同时启用。
- 两者都未设置时，默认保留 7 天。
- 清理发生在触发轮转的 `Write` 或 `Rotate` 调用内；调用返回时清理已经完成。

同一个 `RotateLogs` 实例支持并发写入，但多个实例或进程不应共享同一文件 pattern。
使用完毕后必须调用 `Close`，关闭后再调用 `Write` 或 `Rotate` 会返回
`os.ErrClosed`。

本包基于 [`lestrrat-go/file-rotatelogs`](https://github.com/lestrrat-go/file-rotatelogs)
移植并维护。
