# repo - 仓储接口定义

定义 Biz 层依赖的仓储接口（Repository Interface）。

## 规则

- 接口定义在 `biz/repo/`，实现在 `data/repo/`
- 使用 Wire 的 `wire.Bind` 进行接口绑定
- 接口命名：`{Xxx}BizRepo`

## 示例

```go
type PingBizRepo interface {
    GetPing(ctx context.Context, param *bo.GetPingParam) (*bo.PingResult, error)
}
```
