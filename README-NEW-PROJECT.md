# 创建新项目

基于`go-srv-kit`创建新项目

## 安装开发工具

## **安装开发工具:**

```shell
# kratos
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## **克隆两个项目**

一、 克隆`go-kratos`项目；用于后续生成`protobuf`

```shell
# kratos
mkdir -p $GOPATH/src/github.com/go-kratos && cd $GOPATH/src/github.com/go-kratos
git clone https://github.com/go-kratos/kratos.git
```

二、 克隆`go-srv-kit`项目；用于后续生成`protobuf`

```shell
# go-srv-kit
mkdir -p $GOPATH/src/github.com/ikaiguang && $GOPATH/src/github.com/ikaiguang
git clone https://github.com/ikaiguang/go-srv-kit.git
```

## 创建新项目

```shell
git clone git@github.com:ikaiguang/go-srv-services.git
# 或
git clone https://github.com/ikaiguang/go-srv-services.git
```

克隆项目后，请阅读项目根目录的`README.md`文件
