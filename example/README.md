# Example

Example

## 运行项目

```shell

# 例子
# go run ./cmd/main/... -conf=./configs
# 运行 admin
# make run app=admin
go run ./app/admin/service/cmd/main/... -conf=./app/admin/service/configs
# 运行 user
# make run app=user
go run ./app/user/service/cmd/main/... -conf=./app/user/service/configs

```

## 数据库迁移

```shell

# 例子
# go run ./cmd/migration/... -conf=./configs
# 运行 admin
# make migrate app=admin
go run ./app/admin/service/cmd/migration/... -conf=./app/admin/service/configs
# 运行 user
# make migrate app=user
go run ./app/user/service/cmd/migration/... -conf=./app/user/service/configs

```

## 生成协议

```shell

# 例子
# go run ./cmd/proto/... -path=api/admin
# 运行 admin
# make proto path=api/admin
go run ./cmd/proto/... -path=api/admin
# 运行 user
# make proto path=api/user
go run ./cmd/proto/... -path=api/user

```

## 安装开发工具

**安装开发工具:**

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
