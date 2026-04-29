# service - Service 层

实现 Proto API 定义的 HTTP/gRPC handler。

## 职责

- 接收请求，调用 DTO 转换参数
- 调用 Biz 层处理业务逻辑
- 将结果通过 DTO 转换为 Proto 响应

## 规则

- 只能调用 Biz 层，禁止直接调用 Data 层
- 参数验证在此层完成
- 错误日志在此层记录

## 命名

- 文件：`{module}.service.go`
- 构造函数：`New{Xxx}Service`
- Provider：`service_provider.service.go`
