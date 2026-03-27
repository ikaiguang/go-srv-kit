---
inclusion: auto
---

# 基础设施组件分类

在设计接口或分组组件时，必须正确区分各基础设施组件的职责：

| 组件 | 职责 | 分类 |
|------|------|------|
| Consul | 服务发现、配置中心 | ConsulProvider |
| Jaeger / OpenTelemetry | 链路追踪、可观测性 | TracerProvider |
| Redis | 缓存、分布式锁 | RedisProvider |
| MySQL / PostgreSQL | 关系型数据库 | DatabaseProvider |
| MongoDB | 文档数据库 | MongoProvider |
| RabbitMQ | 消息队列 | MessageQueueProvider |

禁止将 Consul（服务发现/配置中心）和 Jaeger（链路追踪）混为同一个 Provider 接口。
