# thread

goroutine 安全执行和 recover 包装。

## 基础用法

```go
threadpkg.GoSafe(func() {
	// async work
})
threadpkg.GoSafeWithContext(ctx, func(ctx context.Context) {
	// async work with context
})
threadpkg.GoWithContext(ctx, func(ctx context.Context) {
	// alias of GoSafeWithContext
})
```

## 注意事项

异步任务仍需处理 context、超时、日志和资源释放。

## 验证

```bash
go test ./thread
```
