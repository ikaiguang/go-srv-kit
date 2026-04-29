# third_party - 第三方 Proto 依赖

`third_party/` 存放 Proto 代码生成所需的第三方 `.proto` 文件。

这些文件在执行 `protoc` 或 `kratos proto client` 时作为 `--proto_path` 引入，不需要手动修改。

## 目录结构

| 目录 | 说明 |
|------|------|
| `errors/` | Kratos 错误定义 Proto |
| `google/` | Google 官方 Proto（api、protobuf、rpc） |
| `openapi/` | OpenAPI v3 Proto 定义 |
| `protoc-gen-openapiv2/` | OpenAPI v2 (Swagger) Proto 定义 |
| `validate/` | protoc-gen-validate 验证规则 Proto |

## 使用

在 `Makefile` 或 Proto 生成脚本中，通过 `--proto_path` 引用：

```bash
protoc --proto_path=./third_party ...
```

> 这些文件来自上游项目，不要手动修改。如需更新，从对应的上游仓库同步。
