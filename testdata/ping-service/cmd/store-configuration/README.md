# 执行

参数：

* `conf`： 启动读取配置
* `source_dir`： 被存储的配置文件所在文件夹
* `store_dir`： 存储到配置中心位置

```shell
# 执行
make store-configuration
# or
go run ./testdata/ping-service/cmd/store-configuration/... -conf=./testdata/ping-service/configs
# or
go run ./testdata/ping-service/cmd/store-configuration/... -conf=./testdata/ping-service/configs \
  -source_dir=./testdata/ping-service/configs
  -store_dir=

```