# testdata - 测试数据与示例服务

`testdata/` 包含示例服务、测试配置和辅助工具，是理解 go-srv-kit 推荐接入方式的首选参考。

## 目录结构

| 目录 | 说明 |
|------|------|
| `ping-service/` | 完整的示例微服务，展示 DDD 分层架构和 Wire 依赖注入 |
| `configs/` | 测试用配置文件（嵌入到 Go 代码中） |
| `configuration/` | 配置文件存储工具，将 YAML 配置上传到 Consul |
| `service-api/` | 集群服务间 API 调用示例 |

## ping-service 示例服务

`ping-service` 是一个完整的微服务示例，展示了推荐的项目结构和开发模式：

```
ping-service/
├── api/                    # Proto API 定义
├── cmd/
│   ├── ping-service/       # 主服务入口
│   ├── all-in-one/         # All-In-One 多服务合并启动
│   ├── database-migration/ # 数据库迁移工具
│   └── store-configuration/ # 配置存储工具
├── configs/                # 配置文件
├── devops/                 # CI/CD 和 Docker 配置
└── internal/               # 业务代码（DDD 分层）
    ├── service/            # Service 层（HTTP/gRPC handler + DTO）
    ├── biz/                # Business 层（业务逻辑 + BO + Repo 接口）
    └── data/               # Data 层（数据访问 + PO + Repo 实现）
```

### 运行示例服务

```bash
# 方式一：make 命令
make run-service

# 方式二：直接运行
go run ./testdata/ping-service/cmd/ping-service/... -conf=./testdata/ping-service/configs
```

### 测试接口

```bash
curl http://127.0.0.1:10101/api/v1/ping/say_hello
curl http://127.0.0.1:10101/api/v1/ping/logger
curl http://127.0.0.1:10101/api/v1/ping/error
```

## configuration 配置存储

将本地 YAML 配置文件批量上传到 Consul KV：

```bash
# 上传通用配置
go run testdata/configuration/main.go \
  -consul_config consul \
  -source_dir general-configs \
  -store_dir go-micro-saas/general-configs/develop

# 上传服务配置
go run testdata/configuration/main.go \
  -consul_config consul \
  -source_dir service-configs \
  -store_dir ""
```

## 新建服务参考

创建新的业务服务时，建议以 `ping-service` 为模板，按以下步骤操作：

1. 复制 `ping-service` 目录结构
2. 修改 Proto API 定义
3. 实现 Service → Biz → Data 各层
4. 配置 Wire 依赖注入
5. 运行 `wire ./cmd/{service}/export` 生成代码

也可参考独立模板仓库：[service-layout](https://github.com/ikaiguang/service-layout)
