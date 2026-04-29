# 配置文件存储到 Consul

`testdata/configuration/` 提供把本地 YAML 配置批量写入 Consul KV 的示例工具。

## 存储规则

- 输入目录下的文件会按原文件名写入 Consul
- 如果 `source_dir` 不是绝对路径，则相对当前命令工作目录解析
- 默认路径前缀为：

```text
${app.project_name}/${app.server_name}/${app.server_env}/${app.server_version}
```

例如：

- `go-micro-saas/ping-service/production/v1.0.0/app.yaml`
- `go-micro-saas/ping-service/production/v1.0.0/mysql.yaml`

## 常用示例

### 上传通用配置

```bash
go run ./testdata/configuration/main.go \
  -consul_config consul \
  -source_dir general-configs \
  -store_dir go-micro-saas/general-configs/develop
```

### 上传服务配置

```bash
go run ./testdata/configuration/main.go \
  -consul_config consul \
  -source_dir service-configs \
  -store_dir ""
```

## 相关文档

- `testdata/ping-service/cmd/store-configuration/README.md`
- `testdata/README.md`
