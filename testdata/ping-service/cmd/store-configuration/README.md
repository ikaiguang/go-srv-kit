# store-configuration

把本地配置目录写入 Consul KV。

## 参数

- `conf`：启动配置目录，默认 `../../configs`
- `source_dir`：要写入 Consul 的源目录；为空时默认使用 `conf`
- `store_dir`：写入 Consul 的目标路径；为空时根据 `app` 配置自动推导

## 常用命令

```bash
# 使用默认 source_dir 和自动推导的 store_dir
go run ./testdata/ping-service/cmd/store-configuration/... -conf=./testdata/ping-service/configs

# 显式指定 source_dir 和 store_dir
go run ./testdata/ping-service/cmd/store-configuration/... \
  -conf=./testdata/ping-service/configs \
  -source_dir=./testdata/ping-service/configs \
  -store_dir=go-micro-saas/ping-service/testing/v1.0.0
```

如需批量上传 Docker 部署相关配置，可参考：

- `testdata/configuration/README.md`
- `testdata/ping-service/devops/docker-deploy/README.md`
