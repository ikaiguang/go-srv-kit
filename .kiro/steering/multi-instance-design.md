---
inclusion: auto
---

# 多实例设计原则

设计基础设施组件（数据库、缓存、消息队列等）时，必须考虑多实例场景：

- 同类型组件可能需要多个实例（如多个 MySQL 数据源：order_db、user_db）
- 使用命名实例模式（`GetNamed{Xxx}(name)`）支持多数据源
- 保留默认单实例方法向后兼容
- Proto 配置使用 `map<string, T>` 字段支持命名实例配置
