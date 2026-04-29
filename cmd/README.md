# cmd - 命令行工具

`cmd/` 包含项目级别的命令行工具。

## 目录结构

| 目录 | 说明 |
|------|------|
| `proto/` | Proto 代码生成工具，自动查找 `.proto` 文件并调用 `kratos proto client` 生成代码 |

## proto 工具

自动扫描指定目录下的 `.proto` 文件，生成对应的 Go 代码：

```bash
go run cmd/proto/main.go -path=./api/config
```

工具会：
1. 递归查找目录下所有 `.proto` 文件
2. 为每个文件生成 `kratos proto client` 命令
3. 生成执行脚本文件 `proto_script.sh`
4. 依次执行生成命令

> 通常不需要直接使用此工具，推荐使用 Makefile 中的 `make protoc-*` 命令。
