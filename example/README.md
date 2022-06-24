# Example

文档地址

- [kratos-github](https://github.com/go-kratos/kratos)
- [kratos-docs](https://go-kratos.dev/docs/)

## 复制为新项目

1. 编辑文件`./api/config/v1/config.v1.proto`；修改`package`和`option`定义路径；
2. 编辑文件`./api/config/README.md`；修改执行命令到对应的目录；
3. 全局替换`github.com/ikaiguang/go-srv-kit/example/internal`为`github.com/ikaiguang/go-srv-xxx/internal`；
4. 编辑文件`./README.md`；修改标题等内容
5. 编辑配置文件：app、log、server、data、...
6. 格式化项目代码；

## 创建数据库

编辑文件`./configs/config_data.yaml`；配置数据库与创建数据库

```shell
CREATE DATABASE srv_example DEFAULT CHARSET utf8mb4;
```

## 运行

在当前目录先运行

```shell
# 运行程序
# make run
go run ./cmd/main/... -conf=./configs

# 运行测试
# make ping
curl http://127.0.0.1:8081/api/v1/ping/hello
curl http://127.0.0.1:8081/api/v1/ping/error
```

## 执行生成脚本 与 编译proto

```shell

# 执行生成脚本 与 编译proto
# make proto_user
# make proto_xxx
# kratos proto client api/ping/v1/ping.v1.proto
go run ./cmd/proto/... -path=./api/user
go run ./cmd/proto/... -path=./api/xxx
    
```
