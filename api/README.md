# api - Proto 定义

`api/` 存放项目级别的 Proto 定义文件，用于定义配置结构和公共消息类型。

## 目录结构

| 目录 | 说明 |
|------|------|
| `config/` | 全局配置 Proto 定义（`Bootstrap`、数据库、Redis、日志等配置结构） |

## config/

`config.proto` 定义了 `Bootstrap` 配置结构，是所有服务的配置基础：

```protobuf
message Bootstrap {
  App app = 1;
  Server server = 2;
  Log log = 3;
  MySQL mysql = 4;
  Redis redis = 5;
  // ...
}
```

配置文件（YAML）会被解析为 `configpb.Bootstrap` 结构体，供 `LauncherManager` 使用。

## 代码生成

```bash
# 生成配置 Proto
make protoc-config-protobuf

# 或使用工具
go run cmd/proto/main.go -path=./api/config
```

生成的文件：
- `config.pb.go` - 消息定义
- `config.pb.validate.go` - 验证代码
- `config.swagger.json` - OpenAPI 文档

> 不要手动修改 `*.pb.go`、`*.pb.validate.go` 等生成文件。
