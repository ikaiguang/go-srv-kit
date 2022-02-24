# api

接口定义

- 前置条件：安装kratos工具`go install github.com/go-kratos/kratos/cmd/kratos/v2@latest`
- 前置条件：下载kratos源代码。引入kratos的third_party

```shell

# 创建目录
mkdir -p $GOPATH/src/github.com/go-kratos
#mkdir -p %GOPATH%/src/github.com/go-kratos

# 切换目录
cd $GOPATH/src/github.com/go-kratos
cd %GOPATH%/src/github.com/go-kratos

# 克隆项目
git clone https://github.com/go-kratos/kratos.git

# 切换分支
#cd kratos
#git checkout v2.1.5

```

## 生成示例

```shell

# 生成 client 源码
kratos proto client api/ping/v1/ping.v1.proto
kratos proto client api/ping/error/ping.error.proto

# 生成 proto 模板
# kratos proto add api/ping/v1/ping.v1.proto

# 生成 server 源码
# mkdir -p internal/application/service/ping
# kratos proto server api/ping/v1/ping.v1.proto -t internal/application/service/ping
# 修改生成文件名称，符合命名规则

```