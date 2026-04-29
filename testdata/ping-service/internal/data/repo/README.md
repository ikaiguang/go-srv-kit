# repo - 仓储接口实现导出

导出 Data 层对 `biz/repo/` 接口的实现，供 Wire 依赖注入使用。

## Wire 绑定

```go
wire.Bind(new(bizsrepo.PingBizRepo), new(*data.pingData))
```
