# 服务工具

`go-srv-kit`为微服务、业务系统开发提供开箱即用的工具；

- 按需配置启动基础组件，如：数据库、缓存、消息队列等。
- 提供基础的工具，如：日志、配置、HTTP、GRPC、JWT、SnowflakeId、...

## 概述

- 使用服务框架： [go-kratos](https://github.com/go-kratos/kratos)

**参考链接**

- [github.com/go-kratos/kratos](https://github.com/go-kratos/kratos)
- [github.com/uber-go/guide](https://github.com/uber-go/guide)
- [Go Package names](https://blog.golang.org/package-names)

## 运行程序

**Windows**系统，请使用`cmd`或`git-bash`运行。

```shell

# 启动项目
make run-service
go run ./testdata/ping-service/cmd/ping-service/... -conf=./testdata/ping-service/configs

# 运行测试 HTTP JSON
make testing-service
curl http://127.0.0.1:10101/api/v1/ping/logger && echo "\n"
curl http://127.0.0.1:10101/api/v1/ping/error && echo "\n"
curl http://127.0.0.1:10101/api/v1/ping/panic && echo "\n"
curl http://127.0.0.1:10101/api/v1/ping/say_hello && echo "\n"

```

## 创建新项目

参考： [service-layout](https://github.com/ikaiguang/service-layout)

## 感谢支持

| 感谢支持                                   | LOGO                                                                                                                           | 支持内容                            |
|----------------------------------------|--------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| [JETBRAINS](https://www.jetbrains.com) | <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg"  width="120" alt="JetBrains Logo"> | Open Source Development License |

## Give a star! ⭐

如果您觉得这个项目有趣，或者对您有帮助，请给个star吧！

If you think this project is interesting, or helpful to you, please give a star!
