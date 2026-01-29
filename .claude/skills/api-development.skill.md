# API 开发技能

## 描述
帮助开发者使用 go-srv-kit 框架快速开发新的 API 接口，包括 Proto 定义、Service 层、Business 层、Data 层的完整实现。

## 参数
- `service_name`: 服务名称（如：user-service）
- `api_name`: API 名称（如：GetUser）
- `http_method`: HTTP 方法（GET/POST/PUT/DELETE）
- `http_path`: HTTP 路径（如：/api/v1/users/{user_id}）
- `request_fields`: 请求字段（JSON 格式）
- `response_fields`: 响应字段（JSON 格式）

## 执行步骤
1. 在 `api/{service_name}/v1/resources/` 创建请求/响应 Proto
2. 在 `api/{service_name}/v1/services/` 更新服务定义
3. 生成 Proto 代码：`make api-{service_name}`
4. 创建 DTO 转换函数
5. 实现 Service 层
6. 实现 Business 层
7. 实现 Data 层
8. 更新 Wire 依赖注入
9. 生成 Wire 代码

## 示例调用
```
使用 API 开发技能创建用户查询接口
service_name: user-service
api_name: GetUser
http_method: GET
http_path: /api/v1/users/{user_id}
request_fields: {"user_id": "uint64"}
response_fields: {"user": "UserEntity"}
```
