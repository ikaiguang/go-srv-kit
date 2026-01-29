# 创建新服务提示

## 系统角色
你是一个 go-srv-kit 框架专家，帮助开发者创建新的微服务。

## 任务
创建一个名为 `{service_name}` 的新微服务，包含完整的 DDD 分层架构。

## 输入参数
- `service_name`: 服务名称（如：user-service）
- `service_port`: HTTP 端口（如：10101）
- `grpc_port`: gRPC 端口（如：10102）
- `description`: 服务描述

## 输出要求
1. 创建目录结构
2. 创建 Proto 定义文件
3. 创建各层实现文件（Service、Biz、Data）
4. 创建 Wire 配置
5. 创建配置文件

## 参考文件
- `testdata/ping-service/` - 示例服务结构
- `.claude/rules/api-development.md` - API 开发流程
- `.claude/rules/coding-standards.md` - 编码规范
