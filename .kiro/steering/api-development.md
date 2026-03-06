---
inclusion: manual
---

# API 开发流程

## 新增 API 完整流程（10 步）

### 1. 定义 Proto

在 `api/{service-name}/v1/` 下创建：
- `services/{service}.proto` - RPC 服务定义
- `resources/{resource}.proto` - 请求/响应消息
- `errors/{error}.proto` - 错误定义
- `enums/{enum}.proto` - 枚举类型

```protobuf
syntax = "proto3";
package api.{service}.v1;
option go_package = "github.com/ikaiguang/go-srv-kit/api/{service}/v1;v1";

import "google/api/annotations.proto";
import "validate/validate.proto";

service Srv{Name} {
  rpc {Method}({Method}Req) returns ({Method}Resp) {
    option (google.api.http) = {
      post: "/api/v1/{resource}/{action}"
      body: "*"
    };
  }
}
```

### 2. 生成 Proto 代码

```bash
make protoc-specified-api service={service-name}
```

生成文件：`*.pb.go`, `*_grpc.pb.go`, `*_http.pb.go`, `*.validate.go`, `*.swagger.json`

### 3. 实现 Service 层

`internal/service/service/{name}.service.go`:

```go
func (s *XxxService) Method(ctx context.Context, req *pb.MethodReq) (*pb.MethodResp, error) {
    // 1. 参数验证
    // 2. DTO → BO: param := dto.ToBoMethodParam(req)
    // 3. 调用业务: result, err := s.xxxBiz.Method(ctx, param)
    // 4. BO → Proto: return dto.ToProtoMethodResp(result), nil
}
```

### 4. 实现 DTO 转换

`internal/service/dto/{name}.dto.go`: `ToBoXxxParam()`, `ToProtoXxxResp()`

### 5. 实现 Business 层

`internal/biz/biz/{name}.biz.go`: 业务验证 + 调用 Repository

### 6. 定义 Repository 接口

`internal/biz/repo/{name}.repo.go`: 定义 `XxxBizRepo` 接口

### 7. 实现 Data 层

`internal/data/data/{name}.data.go`: 实现 Repository 接口，使用 GORM 操作数据库

### 8. 更新 Wire

`cmd/{service}/export/wire.go` 中添加 Provider

### 9. 注册服务

`internal/service/service/register.go` 中注册 HTTP/gRPC 服务

### 10. 生成 Wire 代码

```bash
wire ./cmd/{service}/export
```

## Proto 规范

- 包名: `api.{service}.v1`
- 字段名: `snake_case`
- 消息名: `PascalCase`
- 枚举值: `UPPER_SNAKE_CASE`，首个值必须为 `UNSPECIFIED = 0`
- 列表字段用复数形式
- 使用 `validate.rules` 添加验证规则

## HTTP 注解

```protobuf
// GET + 路径参数
rpc Get(GetReq) returns(GetResp) {
  option (google.api.http) = { get: "/api/v1/users/{user_id}" };
}
// POST + Body
rpc Create(CreateReq) returns(CreateResp) {
  option (google.api.http) = { post: "/api/v1/users" body: "*" };
}
```

## 错误定义

```protobuf
import "errors/errors.proto";
enum ErrorReason {
  USER_NOT_FOUND = 0 [(errors.code) = 404];
  USER_ALREADY_EXISTS = 1 [(errors.code) = 409];
}
```
