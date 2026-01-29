# Proto 代码生成技能

## 描述
根据 Proto 文件生成 Go 代码、gRPC/HTTP 服务接口、验证代码和 Swagger 文档。

## 参数
- `service`: 服务名称（如：ping-service）
- `version`: 版本（如：v1）

## 生成命令
```bash
# 生成指定服务
make protoc-specified-api service=ping-service

# 或直接使用 protoc
protoc \
  --proto_path=. \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  --go-http_out=paths=source_relative:. \
  --go-errors_out=paths=source_relative:. \
  --validate_out=paths=source_relative,lang=go:. \
  --openapiv2_out . \
  api/ping-service/v1/**/*.proto
```

## 生成的文件
- `{file}.pb.go` - Proto 消息定义
- `{file}_grpc.pb.go` - gRPC 服务接口
- `{file}_http.pb.go` - HTTP 服务接口
- `{file}.validate.go` - 参数验证代码
- `{file}.swagger.json` - Swagger 文档

## 验证生成结果
1. 检查生成的 Go 代码无编译错误
2. 检查 Swagger 文档是否正确
3. 验证 HTTP 路由是否正确注册
