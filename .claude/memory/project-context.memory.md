# 项目上下文记忆

## 项目基本信息
- **项目名称**: go-srv-kit
- **框架**: go-kratos v2.9.2
- **语言**: Go 1.25.9
- **架构**: DDD 分层架构 + Wire 依赖注入
- **通信协议**: HTTP + gRPC 双协议

## 架构关系
```
业务服务 (testdata/ping-service)
  ↓ Service Layer (internal/service/)
  ↓ Business Layer (internal/biz/)
  ↓ Data Layer (internal/data/)
  ↓ go-srv-kit 基础设施
```

## 最近开发记录
<!-- 这里会记录最近开发的功能和问题 -->

### 2026-04-29
- 同步根文档、AGENTS 和 repo-local skill
- 清理过时版本号与重复说明
- 补充 Codex 会话续接与项目级上下文说明

### 待办事项
- [ ] 添加更多常用技能
- [ ] 完善 prompt 模板
- [ ] 添加更多代码模板
