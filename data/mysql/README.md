# mysql

`mysql` 目录提供包 `mysqlpkg`，用于根据 protobuf 配置创建 GORM MySQL 连接，并识别 MySQL 唯一键冲突错误。它适合在业务服务的基础设施初始化阶段使用，向 Biz/Data 层提供统一的 `*gorm.DB`。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/data/mysql/v3
```

```go
import mysqlpkg "github.com/ikaiguang/go-srv-kit/data/mysql/v3"
```

## 核心能力

- `Config`：由 `config.proto` 生成的 MySQL 配置结构，包含 `dsn`、慢查询阈值、日志开关、日志级别和连接池参数。
- `NewDB(conf *Config, opts ...gormpkg.Option) (*gorm.DB, error)`：使用 `gorm.io/driver/mysql` 和本仓库的 GORM 工具包创建 GORM DB。
- `NewMysqlDB(conf *Config, opts ...gormpkg.Option) (*gorm.DB, error)`：`NewDB` 的同义入口，保留更明确的 MySQL 命名。
- `IsErrDuplicatedKey(err error) bool`：判断错误是否为 GORM `ErrDuplicatedKey` 或 MySQL driver 错误码 `1062`。

## 快速使用

```go
package main

import (
	"time"

	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm/v3"
	mysqlpkg "github.com/ikaiguang/go-srv-kit/data/mysql/v3"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	conf := &mysqlpkg.Config{
		Dsn:             "user:password@tcp(127.0.0.1:3306)/app?charset=utf8mb4&parseTime=True&loc=Local",
		SlowThreshold:   durationpb.New(200 * time.Millisecond),
		LoggerEnable:    true,
		LoggerLevel:     "INFO",
		LoggerColorful:  false,
		ConnMaxActive:   100,
		ConnMaxLifetime: durationpb.New(30 * time.Minute),
		ConnMaxIdle:     10,
		ConnMaxIdleTime: durationpb.New(time.Hour),
	}

	db, err := mysqlpkg.NewMysqlDB(conf, gormpkg.WithWriters(gormpkg.NewStdWriter()))
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
}
```

重复键错误判断：

```go
if mysqlpkg.IsErrDuplicatedKey(err) {
	// handle duplicate key
}
```

## 配置字段

| 字段 | 说明 |
| --- | --- |
| `enable` | 是否启用 MySQL 配置。当前包不读取该字段，通常由上层服务启动逻辑判断。 |
| `dsn` | MySQL DSN。可能包含账号和密码，不应提交真实生产值。 |
| `slow_threshold` | GORM 慢查询阈值。 |
| `logger_enable` | 是否启用 GORM 日志。 |
| `logger_colorful` | 是否启用彩色日志。 |
| `logger_level` | 日志级别，由 `gormpkg.ParseLoggerLevel` 解析。 |
| `conn_max_active` | 最大打开连接数。 |
| `conn_max_lifetime` | 连接可复用的最长时间。 |
| `conn_max_idle` | 最大空闲连接数。 |
| `conn_max_idle_time` | 连接最大空闲时间。 |

## 测试

```bash
go test ./...
```

当前单元测试覆盖配置转换、nil 配置和重复键错误识别，不依赖真实 MySQL。真实连接与 `Ping` 应在集成环境使用受控凭据另行验证。

## 生成

`Config` 来自 `data/mysql/config.proto`，相关 Go 文件是生成文件。修改 proto 后必须在仓库根目录执行，保留唯一的仓库相对 descriptor path：

```bash
protoc --proto_path=. --proto_path=./third_party \
  --go_out=paths=source_relative:. \
  --validate_out=paths=source_relative,lang=go:. \
  data/mysql/config.proto
```

不要手动修改 `*.pb.go` 或 `*.validate.go`。

## 注意事项

- `NewDB` 对 nil 配置返回错误；未配置的 duration 字段按零值处理。
- DSN、日志和错误信息可能包含敏感信息，生产环境应使用脱敏后的配置和日志策略。
- `NewDB` 只负责构建 DB 实例；业务层仍应自行管理迁移、事务边界、超时上下文和关闭流程。
