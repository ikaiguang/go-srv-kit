# example

文档地址

- [kratos-github](https://github.com/go-kratos/kratos)
- [kratos-docs](https://go-kratos.dev/docs/)

## JETBRAINS Golang 配置

配置项：setting -> tool -> terminal -> application settings -> shell path

配置值：path/to/Git/bin/bash.exe

## 复制为新项目

1. 编辑文件`./api/config/v1/config.v1.proto`；修改`package`和`option`定义路径；
2. 编辑文件`./api/config/README.md`；修改执行命令到对应的目录；
3. 全局替换`github.com/ikaiguang/go-srv-kit/example/internal`为`github.com/ikaiguang/go-srv-xxx/internal`；
4. 格式化项目代码；

## 运行

在当前目录先运行

```shell
go run ./cmd/main/... -conf=./configs

curl http://127.0.0.1:8080/api/v1/ping/hello
curl http://127.0.0.1:8080/api/v1/ping/error
```
