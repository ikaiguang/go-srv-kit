# 添加新 API 提示

## 系统角色
你是一个 go-srv-kit 框架专家，帮助开发者添加新的 API 接口。

## 任务
在 `{service_name}` 中添加一个新的 API 接口。

## 输入参数
- `service_name`: 服务名称
- `api_name`: API 名称
- `method`: HTTP 方法（GET/POST/PUT/DELETE）
- `path`: HTTP 路径
- `request`: 请求参数（JSON 格式）
- `response`: 响应参数（JSON 格式）
- `auth_required`: 是否需要认证

## 执行流程
1. 在 `api/{service_name}/v1/resources/` 创建请求/响应 Proto
2. 在 `api/{service_name}/v1/services/` 添加 RPC 定义
3. 运行 `make api-{service_name}` 生成代码
4. 在 `internal/service/dto/` 创建 DTO 转换函数
5. 在 `internal/service/service/` 实现 Service 层
6. 在 `internal/biz/biz/` 实现 Business 层
7. 在 `internal/data/data/` 实现 Data 层
8. 更新 Wire 依赖注入
9. 如需认证，更新白名单配置

## 验证清单
- [ ] Proto 文件语法正确
- [ ] 生成的代码无编译错误
- [ ] Swagger 文档正确生成
- [ ] 单元测试通过
- [ ] HTTP/gRPC 路由正确注册
