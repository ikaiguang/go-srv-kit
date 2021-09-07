# DDD.Domain.Factory

- `Factory`的翻译：工厂；制造厂；
- 工厂的释义：直接进行工业生产活动的单位，通常包括不同的车间。

> 用于复杂领域对象的创建/重建。重建是指通过`DDD.Domain.Repository`重建持久化对象。

## DDD.Domain.Factory的用途

- 参数校验：检查参数是否符合业务逻辑的要求。
- 为`DDD.Domain.Repository`提供`符合业务逻辑要求`的`DDD.Domain.Entity`