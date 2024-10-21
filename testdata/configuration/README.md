# 配置文件存储到 consul

把`testdata/configuration/xxx`目录下的文件，按文件名存储到consul

如果路径不是绝对路径，则：路径统一为`./testdata/configuration/xxx`

路径前缀：${app.project_name}/${app.server_name}/${app.server_env}/${app.version}；例如：

> 使用`-path go-micro-saas/base-config`覆写路径前缀

* go-micro-saas/ping-service/production/v1.0.0/app.yaml
* go-micro-saas/ping-service/production/v1.0.0/mysql.yaml
* go-micro-saas/ping-service/production/v1.0.0/filename.yaml

```shell

# base config
go run testdata/configuration/main.go \
  -consul_config consul \
  -source_dir general-configs \
  -store_dir go-micro-saas/general-configs/develop

# service config
go run testdata/configuration/main.go \
  -consul_config consul \
  -source_dir service-configs \
  -store_dir ""

```