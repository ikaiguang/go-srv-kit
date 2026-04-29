# biz - 业务逻辑层

Business 层，负责核心业务逻辑处理。

## 职责

- 实现业务规则和流程
- 通过 Repo 接口访问数据（不直接依赖 Data 层实现）
- 业务验证和错误处理

## 规则

- 只能调用 `biz/repo/` 中定义的仓储接口
- 不能直接访问数据库或外部服务
- 使用 BO（Business Object）作为数据载体

## 命名

- 文件：`{module}.biz.go`
- 构造函数：`New{Xxx}Biz`
