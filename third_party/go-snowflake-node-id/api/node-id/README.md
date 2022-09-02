# 服务配置文件

配置文件

## generate

```shell

# 生成 client 源码
kratos proto client \
    --proto_path=. \
    --proto_path=$GOPATH/src \
    ./third_party/go-snowflake-node-id/api/node-id/v1/errors/node-id.error.v1.proto
    
kratos proto client \
    --proto_path=. \
    --proto_path=$GOPATH/src \
    ./third_party/go-snowflake-node-id/api/node-id/v1/resources/node-id.resource.v1.proto
    
kratos proto client \
    --proto_path=. \
    --proto_path=$GOPATH/src \
    ./third_party/go-snowflake-node-id/api/node-id/v1/services/node-id.service.v1.proto

```
