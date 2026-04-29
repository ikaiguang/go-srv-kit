# cleanup - 资源清理管理

`cleanup/` 提供资源清理管理器，确保服务关闭时按正确顺序释放资源。

## 包名

```go
import cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
```

## 使用

```go
cm := cleanuputil.NewCleanupManager()

// 注册清理函数
cm.Append(func() { db.Close() })
cm.Append(func() { redis.Close() })

// 执行清理（后进先出：先关 Redis，再关 DB）
defer cm.Cleanup()
```

## Merge 辅助函数

```go
cm, err = cleanuputil.Merge(cm, cleanup, err)
```

如果 `err != nil`，会立即执行已注册的所有清理函数并返回错误。
