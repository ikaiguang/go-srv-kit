# API Proto 文件模板

## 服务定义模板
```protobuf
syntax = "proto3";

package api.{service_name}.v1;

import "api/{service_name}/v1/resources/{resource}.proto";
import "google/api/annotations.proto";

option go_package = "github.com/ikaiguang/go-srv-kit/api/{service_name}/v1;v1";

// {service_description}
service Srv{ServiceName} {
  // {api_description}
  rpc {ApiName}({ApiName}Req) returns ({ApiName}Resp) {
    option (google.api.http) = {
      {http_method}: "{http_path}"
      body: "*"
    };
  }
}
```

## 请求消息模板
```protobuf
syntax = "proto3";

package api.{service_name}.v1;

import "validate/validate.proto";

option go_package = "github.com/ikaiguang/go-srv-kit/api/{service_name}/v1;v1";

// {api_name} 请求
message {ApiName}Req {
  {fields}
}

// {api_name} 响应
message {ApiName}Resp {
  {fields}
}
```

## 字段模板
```protobuf
// 字符串字段
string field_name = 1 [(validate.rules).string = {
  min_len: 1,
  max_len: 100
}];

// 整数字段
uint64 field_id = 2 [(validate.rules).uint64 = {
  gte: 1
}];

// 邮箱字段
string email = 3 [(validate.rules).string.email = true];

// 枚举字段
EnumType status = 4;
```
