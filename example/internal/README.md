# DDD

DDD是Domain driven design（领域驱动设计）的简称，

是一种软件设计和开发的方法论，特别适用于复杂业务领域软件设计和开发。

可参考wiki： [Domain-driven_design](https://en.wikipedia.org/wiki/Domain-driven_design)

## 核心

1. 将所有业务逻辑内聚到业务领域（domain）层，将设计和开发的关注点聚焦到业务领域；

2. 建立通用的‘业务领域语言（Ubiquitous Language）’， Ubiquitous Language是工程师和业务领域专家（可以是产品经理、运营、业务相关人员）的桥梁；

3. 战略上，通‘上下文（Bounded Context）’解耦各个业务系统/组件，通过‘防腐层（Anticorruption layer）’确保自有业务领域不受外界污染，通过‘开放主机服务（Open Host
   Service）’向外界公开服务；
   
4. 战术上，将业务对象建模为entity和value object，entity有唯一业务标识且在其生命周期中状态可变，value object与之相反；关联性强的entity和value
   object聚合成一个Aggregate，每个Aggregate有一个root entity，确保Aggregate内容业务规则和行为的一致性；业务行为尽量建模在entity/ value object
   上，当业务行为无法建模到任何业务entity/value object时，可以使用领域服务（domain
   service），使用factory创建复杂的业务entity，使用repository实现实体的重建和持久化操作；领域相关的通知等可以通过domain event发布出去。
