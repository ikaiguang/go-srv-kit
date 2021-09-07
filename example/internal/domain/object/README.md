# DDD.Domain.Object

领域值对象。value object是相对于domain entity来讲的，对照起来value object有如下特征：

- 可以有唯一业务标识【区别于domain entity】
- 持有自己的业务属性和业务行为【同domain entity】
- 一旦定义，他是不可变的，它通常是短暂的；【区别于domain entity】

## DDD.Domain.Entity 用途

它是一个值，是不可变的，immutable。没有identifier，也不需要被追踪！

什么时候使用entity，什么时候使用value object；

具体问题具体分析，如下：

```text

比如我们需要对地址这个东西建模。 
如果我们关心的是地址的履历之类的信息，
过去30年前这个地址可能叫霞飞路，现在可能叫淮海路，而且需求是我们必须知道霞飞路，淮海路指的是一个地址。
那很可能我们需要的是`entity`。这个`entity`可能还要开发`change()`的方法来改变路名。

但如果我们做的是一个送货软件。
地址只是表示一个目的地而已，霞飞路和淮海路在我看来就是不同的，
那就说明，你不必对地址本身的变化进行追踪（送货地址变了，对你很重要。但霞飞路改名成淮海路对你不重要。）。
那value object就够了。

```
