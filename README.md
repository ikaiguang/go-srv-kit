# 服务工具

为服务开发提供基础工具

## Overview

- 基于[go-kratos](https://github.com/go-kratos/kratos)
- Domain-Driven Design (DDD)

## 文档地址

- [kratos-github](https://github.com/go-kratos/kratos)
- [kratos-docs](https://go-kratos.dev/docs/)

## 运行Example

运行Example;

**Windows**系统，请使用`CMD`或`Git-Bash`运行。原因：某些命令行工具无法正确读取配置目录；例：`PowerShell`

```shell
go run ./example/cmd/main/... -conf=./example/configs

curl http://127.0.0.1:8081/api/v1/ping/hello
curl http://127.0.0.1:8081/api/v1/ping/error

# 运行测试 PROTOBUF
curl -X GET \
    -H "Content-Type: application/proto" \
    -H "Accept: application/proto" \
    http://127.0.0.1:8081/api/v1/ping/hello
curl -X GET \
    -H "Content-Type: application/proto" \
    -H "Accept: application/proto" \
    http://127.0.0.1:8081/api/v1/ping/error
```

## 参考链接

- [github.com/go-kratos/kratos](https://github.com/go-kratos/kratos)
- [Domain-driven design](https://en.wikipedia.org/wiki/Domain-driven_design)
- [github.com/uber-go/guide](https://github.com/uber-go/guide)
- [Go Package names](https://blog.golang.org/package-names)

## 感谢支持

| 感谢支持                                   | LOGO                                                                                                                            | 支持内容                            |
|----------------------------------------|---------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| [JETBRAINS](https://www.jetbrains.com) | <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg"  width="120" alt="JetBrains Logo"> | Open Source Development License |

## Give a star! ⭐

If you think this project is interesting, or helpful to you, please give a star!
