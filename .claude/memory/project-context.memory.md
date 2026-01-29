# 项目上下文记忆

## 项目基本信息
- **项目名称**: go-srv-kit
- **框架**: go-kratos v2.9.1
- **语言**: Go 1.24.10
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

### 2024-01-29
- 初始化 .claude 目录结构
- 创建 skills、prompts、templates、memory、variables 目录

### 待办事项
- [ ] 添加更多常用技能
- [ ] 完善 prompt 模板
- [ ] 添加更多代码模板
