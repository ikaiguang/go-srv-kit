# locker

本地锁、缓存锁和分布式锁接口辅助。

## 基础用法

```go
locker := lockerpkg.NewLocalLocker()
unlocker, err := locker.Mutex(ctx, "resource-key")
if err == nil {
	defer unlocker.Unlock(ctx)
}
```

## 注意事项

锁 key 要稳定且粒度清晰；业务代码必须保证成功加锁后释放。

## 验证

```bash
go test ./locker
```
