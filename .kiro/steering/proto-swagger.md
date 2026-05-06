---
inclusion: fileMatch
fileMatchPattern: "**/*.proto"
---

# Proto 与 Swagger 规范

## Proto 文件布局

```protobuf
syntax = "proto3";
package api.{service}.v1;
option go_package = "github.com/ikaiguang/go-srv-kit/api/{service}/v1;v1";

// 1. 引入
import "google/api/annotations.proto";
import "validate/validate.proto";

// 2. 服务定义
service SrvXxx { ... }

// 3. 枚举
enum Status { UNSPECIFIED = 0; ENABLED = 1; }

// 4. 消息
message XxxReq { ... }
message XxxResp { ... }
```

## 目录结构

```
api/{service}/v1/
├── services/    # RPC 服务定义
├── resources/   # 请求/响应消息
├── errors/      # 错误定义
└── enums/       # 枚举类型
```

## 验证规则

```protobuf
string name = 1 [(validate.rules).string = {min_len: 1, max_len: 50}];
string email = 2 [(validate.rules).string.email = true];
uint64 id = 3 [(validate.rules).uint64.gte = 1];
```

## 生成命令

```bash
make protoc-specified-api service={service-name}
make protoc-api-protobuf  # 全部
```

## 必需插件

protoc-gen-go, protoc-gen-go-grpc, protoc-gen-go-http, protoc-gen-go-errors, protoc-gen-validate, protoc-gen-openapiv2

安装: `make init`
