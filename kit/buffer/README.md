# buffer

复用 `bytes.Buffer`，降低频繁临时 buffer 分配。

## 基础用法

```go
buf := bufferpkg.GetBuffer()
defer bufferpkg.PutBuffer(buf)
buf.WriteString("hello")
```

## 注意事项

归还后不要继续读写该 buffer；跨 goroutine 使用时由调用方保证生命周期。`PutBuffer(nil)` 会直接返回。

## 验证

```bash
go test ./buffer
```
