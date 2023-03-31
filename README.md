# 服务工具

为微服务开发提供基础工具，通过配置按需启动服务基础组件

创建新项目请阅读: [README-NEW-PROJECT.md](./README-NEW-PROJECT.md)

## 概述

- 本工具的服务框架是： [go-kratos](https://github.com/go-kratos/kratos)
- 项目的目录结构参考： DDD(领域驱动设计)

参考链接

- [github.com/go-kratos/kratos](https://github.com/go-kratos/kratos)
- [Domain-driven design](https://en.wikipedia.org/wiki/Domain-driven_design)
- [github.com/uber-go/guide](https://github.com/uber-go/guide)
- [Go Package names](https://blog.golang.org/package-names)

## 运行程序

**Windows**系统，请使用`cmd`或`git-bash`运行。

```shell

# 启动项目
go run ./example/cmd/main/... -conf=./example/configs

# 运行测试 HTTP JSON
curl http://127.0.0.1:8081/api/v1/ping/hello
# curl http://127.0.0.1:8081/api/v1/ping/error
# curl http://127.0.0.1:8081/api/v1/ping/logger

# 运行测试 HTTP PROTOBUF
curl -X GET \
    -H "Content-Type: application/proto" \
    -H "Accept: application/proto" \
    http://127.0.0.1:8081/api/v1/ping/hello
curl -X GET \
    -H "Content-Type: application/proto" \
    -H "Accept: application/proto" \
    http://127.0.0.1:8081/api/v1/ping/error
```

## 感谢支持

| 感谢支持                                   | LOGO                                                                                                                            | 支持内容                            |
|----------------------------------------|---------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| [JETBRAINS](https://www.jetbrains.com) | <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg"  width="120" alt="JetBrains Logo"> | Open Source Development License |

## Give a star! ⭐

如果您觉得这个项目有趣，或者对您有帮助，请给个star吧！

If you think this project is interesting, or helpful to you, please give a star!
