# 创建新项目

基于`go-srv-kit`创建新项目

## 安装必要的工具

**安装必要的工具:**

```shell
# kratos
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

**克隆两个项目:**

1. 克隆`go-kratos`项目；后续生成`protobuf`需要
2. 克隆`go-srv-kit`项目；后续生成`protobuf`需要

```shell
# kratos
mkdir -p $GOPATH/src/github.com/go-kratos && cd $GOPATH/src/github.com/go-kratos
git clone https://github.com/go-kratos/kratos.git

# go-srv-kit
mkdir -p $GOPATH/src/github.com/ikaiguang && $GOPATH/src/github.com/ikaiguang
git clone https://github.com/ikaiguang/go-srv-kit.git
```

## 创建新项目

- 创建目录: `mkdir my-server-protject && cd my-server-protject`
- 初始化项目文件: 复制`go-srv-kit/example/`到`my-server-protject/`

```shell

# 复制为新项目
cp $GOPATH/src/github.com/ikaiguang/go-srv-kit/example/* path/to/my-server-protject/
# ===== 例子 =====
# 创建目录
mkdir -p $GOPATH/src/gitee.com/aircraft-group/aircraft-mall-admin && cd $GOPATH/src/gitee.com/aircraft-group/aircraft-mall-admin
# 初始化项目文件: 复制`go-srv-kit/example/`到`my-server-protject/`
cp -r $GOPATH/src/github.com/ikaiguang/go-srv-kit/example/ $GOPATH/src/gitee.com/aircraft-group/aircraft-mall-admin/
# 查看与确认文件
ls -al

```

## 编辑新项目

一： 全局替换文件路径: `github.com/ikaiguang/go-srv-kit/example`；例：

```text


`go-srv-kit/example`路径(go-mod)：github.com/ikaiguang/go-srv-kit/example

==> 替换为

`my-server-protject`路径(go-mod)：gitee.com/aircraft-group/aircraft-mall-admin

```

二：更新`go-mod`

```shell
# 在项目目录执行
go mod init
go mod tidy
```

## 运行与测试项目

```shell
# 在当前目录先运行程序
# make run
go run ./cmd/main/... -conf=./configs

# 运行测试
# make ping
curl http://127.0.0.1:8081/api/v1/ping/hello
# curl http://127.0.0.1:8081/api/v1/ping/error
```

## [额外的工具] 编译proto

```shell

# 执行生成脚本 与 编译proto
# kratos proto client api/ping/v1/ping.v1.proto
#  make proto path=api/user
#  make proto path=api/xxx
go run ./cmd/proto/... -path=./api/user
go run ./cmd/proto/... -path=./api/xxx
    
```
