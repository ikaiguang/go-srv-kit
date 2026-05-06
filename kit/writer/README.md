# writer

dummy writer 和按大小轮转的文件 writer。

## 基础用法

```go
w, err := writerpkg.NewRotateFile(&writerpkg.ConfigRotate{
	Dir:            "./runtime/logs",
	Filename:       "app",
	RotateSize:     100 << 20,
	StorageCounter: 30,
	Compress:       true,
})
```

## 注意事项

- 日志目录需要可写；轮转保留策略应结合磁盘容量设置。
- 当前活动文件默认为 `<Filename>.log`。
- 按大小轮转后，归档文件由 lumberjack 自动追加时间戳，便于排查。
- `RotateSize` 单位是 byte，内部会向上转换为 MB。
- 不再支持纯 `RotateTime` 自动轮转；如配置 `RotateTime` 且未配置 `RotateSize`，会返回错误。
- `Compress=true` 时会压缩归档日志文件。

## 验证

```bash
go test ./writer
```
