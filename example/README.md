# Example

Example

## 运行项目

```shell

go run ./cmd/main/... -conf=./configs

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
