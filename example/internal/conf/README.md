# config

配置文件

## generate

```shell
protoc -I. --go_out=. --go_opt=paths=source_relative ./conf.proto
```