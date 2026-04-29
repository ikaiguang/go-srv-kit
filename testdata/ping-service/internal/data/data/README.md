# data - 数据访问层

Data 层，负责数据库、缓存等数据源的访问操作。

## 职责

- 实现 `biz/repo/` 中定义的仓储接口
- 操作数据库（GORM）、缓存（Redis）等
- PO ↔ BO 转换

## 规则

- 始终使用 `db.WithContext(ctx)` 传递 Context
- GORM 错误需转换为业务错误（`errorpkg.ErrorNotFound` 等）
- 事务使用 `db.Transaction(func(tx *gorm.DB) error { ... })`

## 命名

- 文件：`{module}.data.go`
- 构造函数：`New{Xxx}Data`
