# ping-service - 示例微服务

`ping-service` 是一个完整的微服务示例，展示了 go-srv-kit 推荐的 DDD 分层架构和开发模式。

新建业务服务时，建议以此为模板参考。

## 目录结构

```
ping-service/
├── api/                        # Proto API 定义
│   ├── ping-service/v1/        # Ping 服务 API
│   └── testdata-service/v1/    # 测试数据服务 API
├── cmd/                        # 启动入口
│   ├── ping-service/           # 主服务
│   ├── all-in-one/             # 多服务合并启动
│   ├── database-migration/     # 数据库迁移
│   └── store-configuration/    # 配置存储到 Consul
├── configs/                    # 配置文件（YAML）
├── devops/                     # CI/CD 和 Docker
└── internal/                   # 业务代码
    ├── service/                # Service 层
    │   ├── service/            # HTTP/gRPC handler
    │   └── dto/                # 数据传输对象（Proto ↔ BO 转换）
    ├── biz/                    # Business 层
    │   ├── biz/                # 业务逻辑
    │   ├── bo/                 # 业务对象
    │   ├── repo/               # 仓储接口定义
    │   ├── event/              # 事件处理
    │   └── scheduler/          # 定时任务
    └── data/                   # Data 层
        ├── data/               # 仓储实现
        ├── po/                 # 持久化对象（数据库模型）
        ├── repo/               # 仓储接口实现导出
        ├── cache/              # 缓存操作
        └── schema/             # 数据库 Schema
```

## 数据流

```
Proto Request → DTO → BO → PO → Database
Database → PO → BO → DTO → Proto Response
```

## 运行

```bash
# 启动服务
go run ./testdata/ping-service/cmd/ping-service/... -conf=./testdata/ping-service/configs

# 测试接口
curl http://127.0.0.1:10101/api/v1/ping/say_hello
```

## 分层规则

- **Service 层**：只能调用 Biz 层，负责参数验证和 DTO 转换
- **Biz 层**：只能调用 Repo 接口，负责业务逻辑
- **Data 层**：实现 Repo 接口，负责数据访问
- 禁止 Service 直接调用 Data
