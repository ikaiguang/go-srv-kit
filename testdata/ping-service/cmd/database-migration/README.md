# database-migration

示例服务的数据库迁移入口。

## 作用

- 加载示例服务配置
- 创建 `LauncherManager`
- 获取推荐数据库连接
- 执行 `migrate/` 下定义的迁移逻辑

## 运行方式

```bash
go run ./testdata/ping-service/cmd/database-migration/... -conf=./testdata/ping-service/configs
```

## 相关目录

- `testdata/ping-service/cmd/database-migration/migrate/`
- `service/database/`
