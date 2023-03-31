# DDD.Domain.Entity

领域实体是Domain的核心成员；`DDD.Domain.Entity`具有如下三个特征：

- 唯一业务标识
- 持有自己的业务属性和业务行为
- 属性可变，有着自己的生命周期

> `DDD.Domain.Entity`和`DDD.Domain.Repository`紧密联系。
> 如果说：`DDD.Domain.Repository`接口定义了`DDD.Infrastructure`的持久化层交互契约；
> 那`DDD.Domain.Entity`则是契约中条款。

## DDD.Domain.Entity 用途

- 为`DDD.Domain.Repository`定义实体