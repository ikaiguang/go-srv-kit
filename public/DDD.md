# 服务工具

# DDD

## 1 DDD是什么？

参考博客：https://www.cnblogs.com/zcqiand/p/13686640.html

DDD是领域驱动设计，是Eric Evans于2003年提出的，离现在有17年。

## 2 为什么需要DDD

当软件越来越复杂，实际开发中，大量的业务逻辑堆积在一个巨型类中的例子屡见不鲜，代码的复用性和扩展性无法得到保证。
为了解决这样的问题，DDD提出了清晰的分层架构和领域对象的概念，
让面向对象的分析和设计进入了一个新的阶段，对企业级软件开发起到了巨大的推动作用。

### 2.1 POP，OOP，DDD是如何解决问题

面向过程编程（POP），接触到需求第一步考虑把需求自顶向下分解成一个一个函数。
并且在这个过程中考虑分层，模块化等具体的组织方式，从而分解软件的复杂度。
当软件的复杂度不是很大，POP也能得到很好的效果。

面向对象编程（OOP），接触到需求第一步考虑把需求分解成一个一个对象，然后每个对象添加一个一个方法和属性，
程序通过各种对象之间的调用以及协作，从而实现计算机软件的功能。跟很多工程方法一样，
OOP的初衷就是一种处理软件复杂度的设计方法。

领域驱动设计（DDD），接触到需求第一步考虑把需求分解成一个一个问题域，然后再把每个问题域分解成一个一个对象，
程序通过各种问题域之间的调用以及协作，从而实现计算机软件的功能。
DDD是解决复杂中大型软件的一套行之有效方式，现已成为主流。

### 2.2 POP，OOP，DDD的特点

POP，无边界，软件复杂度小适用，例如“盖房子”。

OOP，以“对象”为边界，软件复杂度中适用，例如“盖小区”。

DDD，以“问题域”为边界，软件复杂度大适用，例如“盖城市”。

## 3 DDD的分层架构和构成要素

### 3.1 分层架构

![分层架构1](ddd/ddd_1.png)

![分层架构2](ddd/ddd_2.png)

整个架构分为四层，其核心就是领域层（Domain），所有的业务逻辑应该在领域层实现，具体描述如下：

用户界面/展现层，负责向用户展现信息以及解释用户命令。

应用层，很薄的一层,用来协调应用的活动。它不包含业务逻辑。它不保留业务对象的状态,但它保有应用任务的进度状态。

领域层，本层包含关于领域的信息。这是业务软件的核心所在。在这里保留业务对象的状态,对业务对象和它们状态的持久化被委托给了基础设施层。

基础设施层，本层作为其他层的支撑库存在。它提供了层间的通信,实现对业务对象的持久化,包含对用户界面层的支撑库等作用。

### 3.2 构成要素

实体（Entity），具备唯一ID，能够被持久化，具备业务逻辑，对应现实世界业务对象。

值对象（Value Object），不具有唯一ID，由对象的属性描述，一般为内存中的临时对象，可以用来传递参数或对实体进行补充描述。

领域服务（Domain Service），为上层建筑提供可操作的接口，负责对领域对象进行调度和封装，同时可以对外提供各种形式的服务。

聚合根（Aggregate Root），聚合根属于实体对象，聚合根具有全局唯一ID，而实体只有在聚合内部有唯一的本地ID，值对象没有唯一ID

工厂（Factories），主要用来创建聚合根，目前架构实践中一般采用IOC容器来实现工厂的功能。

仓储（Repository），封装了基础设施来提供查询和持久化聚合操作。