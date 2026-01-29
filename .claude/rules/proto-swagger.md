# Proto 与 Swagger 开发规范

## Protoc 插件

### 必需插件

```bash
# 安装所有 protoc 插件
make init
```

| 插件 | 版本 | 用途 |
|------|------|------|
| protoc-gen-go | v1.34.2 | 生成 Go 消息定义 |
| protoc-gen-go-grpc | v1.5.1 | 生成 gRPC 服务接口 |
| protoc-gen-go-http | latest | 生成 HTTP 服务接口 |
| protoc-gen-go-errors | v0.0.2 | 生成错误定义 |
| protoc-gen-validate | v1.1.0 | 生成验证代码 |
| protoc-gen-openapiv2 | v2.22.0 | 生成 OpenAPI v2 (Swagger) |
| protoc-gen-openapi | v0.7.0 | 生成 OpenAPI v3 |

## Proto 文件组织

### 目录结构

```
api/{service-name}/v1/
├── services/        # RPC 服务定义
├── resources/       # 请求/响应消息
├── errors/          # 错误定义
└── enums/           # 枚举类型
```

### Proto 文件布局

```protobuf
syntax = "proto3";

package api.{service}.v1;
option go_package = "github.com/ikaiguang/go-srv-kit/api/{service}/v1;v1";

// 1. 引入其他 proto
import "api/common/common.proto";
import "google/api/annotations.proto";

// 2. 定义服务
service SrvXxx {
  // 方法注释
  rpc Method(Request) returns(Response) {
    option (google.api.http) = {
      post: "/api/v1/xxx/method"
      body: "*"
    };
  }
}

// 3. 定义枚举
enum Status {
  UNSPECIFIED = 0;
  ENABLED = 1;
  DISABLED = 2;
}

// 4. 定义消息
message Request {
  string id = 1;
}

message Response {
  string message = 1;
}
```

## Proto 代码生成命令

### 完整的 protoc 命令

```bash
protoc \
  --proto_path=. \
  --proto_path=$(GOPATH)/src \
  --proto_path=./third_party \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  --go-http_out=paths=source_relative:. \
  --go-errors_out=paths=source_relative:. \
  --validate_out=paths=source_relative,lang=go:. \
  --openapiv2_out . \
  --openapiv2_opt logtostderr=true \
  --openapiv2_opt allow_delete_body=true \
  --openapiv2_opt json_names_for_fields=false \
  --openapiv2_opt enums_as_ints=true \
  --openapi_out=fq_schema_naming=true,enum_type=integer,default_response=true:. \
  {proto-files}
```

### 使用 Makefile

```bash
# 生成所有 API
make protoc-api-protobuf

# 生成配置
make protoc-config-protobuf

# 生成指定服务
make protoc-specified-api service=ping-service

# 生成 ping-service v1
make protoc-ping-v1-protobuf
```

## 生成的文件

每个 `.proto` 文件会生成以下文件：

| 文件 | 说明 |
|------|------|
| `{file}.pb.go` | Proto 消息定义（Go struct） |
| `{file}_grpc.pb.go` | gRPC 服务接口和客户端 |
| `{file}_http.pb.go` | HTTP 服务接口和路由 |
| `{file}_errors.pb.go` | 错误定义和错误码 |
| `{file}.validate.go` | 参数验证代码 |
| `{file}.swagger.json` | OpenAPI v2 文档 |
| `{file}.openapi.yaml` | OpenAPI v3 文档 |

## HTTP 注解规范

### 基本 HTTP 映射

```protobuf
service SrvUser {
  // GET 请求
  rpc GetUser(GetUserReq) returns(GetUserResp) {
    option (google.api.http) = {
      get: "/api/v1/users/{user_id}"
    };
  }

  // POST 请求
  rpc CreateUser(CreateUserReq) returns(CreateUserResp) {
    option (google.api.http) = {
      post: "/api/v1/users"
      body: "*"
    };
  }

  // PUT 请求
  rpc UpdateUser(UpdateUserReq) returns(UpdateUserResp) {
    option (google.api.http) = {
      put: "/api/v1/users/{user_id}"
      body: "*"
    };
  }

  // DELETE 请求
  rpc DeleteUser(DeleteUserReq) returns(DeleteUserResp) {
    option (google.api.http) = {
      delete: "/api/v1/users/{user_id}"
    };
  }
}
```

### 路径参数

```protobuf
message GetUserReq {
  uint64 user_id = 1;  // 路径参数: /api/v1/users/{user_id}
}
```

### Query 参数

```protobuf
message ListUsersReq {
  string name = 1;   // Query 参数: ?name=xxx
  int32 page = 2;    // Query 参数: ?page=1
  int32 size = 3;    // Query 参数: ?size=10
}

rpc ListUsers(ListUsersReq) returns(ListUsersResp) {
  option (google.api.http) = {
    get: "/api/v1/users"
  };
}
```

### Body 参数

```protobuf
message CreateUserReq {
  string name = 1;
  string email = 2;
}

rpc CreateUser(CreateUserReq) returns(CreateUserResp) {
  option (google.api.http) = {
    post: "/api/v1/users"
    body: "*"  // 整个请求作为 body
  };
}
```

## 验证规则

### 使用 protoc-gen-validate

```protobuf
import "validate/validate.proto";

message CreateUserReq {
  string name = 1 [(validate.rules).string = {
    min_len: 1,
    max_len: 50,
    pattern: "^[a-zA-Z0-9_]+$"
  }];

  string email = 2 [(validate.rules).string.email = true];

  int32 age = 3 [(validate.rules).int32 = {
    gte: 0,
    lte: 150,
    optional: true
  }];

  repeated string tags = 4 [(validate.rules).repeated = {
    min_items: 1,
    max_items: 10,
    items: {string: {min_len: 1}}
  }];
}
```

### 常用验证规则

| 类型 | 规则 | 说明 |
|------|------|------|
| string | min_len | 最小长度 |
| string | max_len | 最大长度 |
| string | pattern | 正则表达式 |
| string | email | 邮箱格式 |
| string | uri | URI 格式 |
| int32 | gte | 大于等于 |
| int32 | lte | 小于等于 |
| repeated | min_items | 最少数量 |
| repeated | max_items | 最多数量 |

## 错误定义

### 错误 Proto

```protobuf
import "errors/errors.proto";

enum ErrorReason {
  // 用户不存在
  USER_NOT_FOUND = 0 [(errors.code) = 404];

  // 用户已存在
  USER_ALREADY_EXISTS = 1 [(errors.code) = 409];

  // 参数错误
  INVALID_PARAMETER = 2 [(errors.code) = 400];
}
```

### 使用错误

```go
import v1 "github.com/ikaiguang/go-srv-kit/api/user-service/v1"

func (s *userService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserResp, error) {
    if req.GetUserId() == 0 {
        return nil, v1.ErrorInvalidParameter("user_id is required")
    }

    user, err := s.userBiz.GetUser(ctx, req.GetUserId())
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, v1.ErrorUserNotFound("user not found")
        }
        return nil, err
    }

    return &v1.GetUserResp{User: user}, nil
}
```

## Swagger/OpenAPI 文档

### 访问文档

```bash
# Swagger UI (OpenAPI v2)
http://127.0.0.1:10101/q/

# OpenAPI v3 JSON
http://127.0.0.1:10101/api/swagger/
```

### 生成的文档位置

```
api/{service}/v1/
├── {file}.swagger.json    # OpenAPI v2 (Swagger)
└── {file}.openapi.yaml    # OpenAPI v3
```

### 自定义文档

```protobuf
service SrvUser {
  // 用户服务
  option (google.api.default_host) = "api.example.com";

  rpc CreateUser(CreateUserReq) returns(CreateUserResp) {
    option (google.api.http) = {
      post: "/api/v1/users"
      body: "*"
    };

    // 方法描述
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      description: "创建新用户",
      summary: "创建用户",
      tags: "用户管理",
      security: {
        security_requirement: {
          api_key: []
        }
      }
    };
  }
}
```

## Makefile 配置

### 添加新的 Proto 生成

在 `{service}/api/{version}/makefile_protoc.mk`:

```makefile
override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

# 查找所有 proto 文件
SERVICE_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SERVICE_INTERNAL_PROTO := ""
SERVICE_PROTO_FILES := ""
ifneq ($(SERVICE_INTERNAL_PROTO), "")
	SERVICE_PROTO_FILES=$(SERVICE_API_PROTO) $(SERVICE_INTERNAL_PROTO)
else
	SERVICE_PROTO_FILES=$(SERVICE_API_PROTO)
endif

.PHONY: protoc-{service}-v{version}-protobuf
# protoc :-->: generate {service} v{version} protobuf
protoc-{service}-v{version}-protobuf:
	@echo "# generate {service} v{version} protobuf"
	$(call protoc_protobuf,$(SERVICE_PROTO_FILES))
```

然后在 `Makefile` 中引入：

```makefile
include {service}/api/{service}/v{version}/makefile_protoc.mk
```

## 常见问题

### protoc 命令找不到

```bash
# 检查 protoc 是否安装
protoc --version

# 重新运行初始化
make init
```

### 生成的代码路径不对

检查 `paths=source_relative` 参数：
- `paths=source_relative` - 在 proto 文件同目录生成
- `paths=import` - 根据 go_package 生成

### Swagger 文档未生成

检查是否安装了 `protoc-gen-openapiv2`：

```bash
# 检查插件
ls $GOPATH/bin/protoc-gen-openapiv2

# 重新安装
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.22.0
```

### HTTP 路由不生效

确保：
1. 导入了 `google/api/annotations.proto`
2. 添加了 `option (google.api.http)` 注解
3. 重新生成了 HTTP 代码
