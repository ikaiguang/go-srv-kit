# api - 服务 API 定义

`api/` 存放 ping-service 的 Proto API 定义文件。

## 目录结构

| 目录 | 说明 |
|------|------|
| `ping-service/v1/` | Ping 服务 API（健康检查、日志测试、错误测试等） |
| `testdata-service/v1/` | 测试数据服务 API |

## 代码生成

```bash
# 生成 ping-service API
make protoc-ping-v1-protobuf

# 或使用通用命令
make protoc-specified-api service=ping-service
```

每个 `.proto` 文件会生成：
- `*.pb.go` - 消息定义
- `*_grpc.pb.go` - gRPC 服务接口
- `*_http.pb.go` - HTTP 服务接口
- `*.pb.validate.go` - 参数验证
- `*.swagger.json` - OpenAPI v2 文档

> 不要手动修改生成的代码文件。
