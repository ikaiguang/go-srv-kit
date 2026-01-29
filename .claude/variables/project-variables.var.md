# 项目变量定义

## 项目结构变量
| 变量名 | 值 | 说明 |
|--------|-----|------|
| `PROJECT_NAME` | go-srv-kit | 项目名称 |
| `FRAMEWORK` | go-kratos v2.9.1 | 框架版本 |
| `GO_VERSION` | 1.24.10 | Go 版本 |
| `ARCHITECTURE` | DDD | 架构模式 |

## 路径变量
| 变量名 | 值 | 说明 |
|--------|-----|------|
| `API_DIR` | api/ | API 定义目录 |
| `SERVICE_DIR` | testdata/*/internal/service/ | Service 层目录 |
| `BIZ_DIR` | testdata/*/internal/biz/ | Business 层目录 |
| `DATA_DIR` | testdata/*/internal/data/ | Data 层目录 |
| `CMD_DIR` | testdata/*/cmd/ | 命令目录 |

## 命令变量
| 变量名 | 值 | 说明 |
|--------|-----|------|
| `PROTOC_CMD` | make protoc-api-protobuf | 生成 API 命令 |
| `WIRE_CMD` | wire ./cmd/*/export | Wire 生成命令 |
| `RUN_CMD` | go run ./cmd/*/... -conf=./configs | 运行命令 |
| `TEST_CMD` | go test ./... | 测试命令 |

## 端口变量
| 变量名 | 默认值 | 说明 |
|--------|--------|------|
| `HTTP_PORT` | 10101 | HTTP 端口 |
| `GRPC_PORT` | 10102 | gRPC 端口 |

## 命名约定
| 类型 | 格式 | 示例 |
|------|------|------|
| Service 结构体 | New{Xxx}Service | NewPingService |
| Biz 结构体 | New{Xxx}Biz | NewPingBiz |
| Data 结构体 | New{Xxx}Data | NewPingData |
| Repository 接口 | {Xxx}BizRepo | PingBizRepo |
| DTO 转换 | ToBo{Xxx}, ToProto{Xxx} | ToBoGetPingParam |

## 文件命名约定
| 类型 | 格式 | 示例 |
|------|------|------|
| Proto | {name}.proto | ping.proto |
| Service | {name}.service.go | ping.service.go |
| Biz | {name}.biz.go | ping.biz.go |
| Data | {name}.data.go | ping.data.go |
| DTO | {name}.dto.go | ping.dto.go |
| PO | {name}.po.go | ping.po.go |
| 测试 | {name}_test.go | ping.service_test.go |
