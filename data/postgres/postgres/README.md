# postgres

`postgres` 目录提供包 `psqlpkg`，导入路径为 `github.com/ikaiguang/go-postgres-kit/postgres`。该包适合在 Go 服务中统一创建 GORM PostgreSQL 连接，并复用 protobuf 形式的数据库配置。

## 安装

```bash
go get github.com/ikaiguang/go-postgres-kit
```

```go
import psqlpkg "github.com/ikaiguang/go-postgres-kit/postgres"
```

## 核心能力

- `Config`：由 `config.proto` 生成的 PostgreSQL 配置结构，包含 DSN、日志、慢查询阈值和连接池参数。
- `NewDB(conf *Config, opts ...gormpkg.Option)`：根据 `Config` 和可选的 `go-gorm-kit/gorm` 选项创建 `*gorm.DB`。
- `NewPostgresDB(conf *Config, opts ...gormpkg.Option)`：`NewDB` 的语义化别名，便于调用方表达数据库类型。
- `IsErrDuplicatedKey(err error)`：识别 GORM 重复键错误和 PostgreSQL `23505` 唯一键冲突。
- `Validate()` / `ValidateAll()`：由 `protoc-gen-validate` 生成的配置校验方法。

## 配置字段

| 字段 | 说明 |
| --- | --- |
| `Enable` | 是否启用 PostgreSQL 配置，由调用方自行决定如何使用。 |
| `Dsn` | PostgreSQL DSN，传给 `gorm.io/driver/postgres`。 |
| `SlowThreshold` | 慢查询阈值。 |
| `LoggerEnable` | 是否启用 GORM 日志。 |
| `LoggerColorful` | 日志输出是否启用颜色。 |
| `LoggerLevel` | 日志级别，支持值由 `go-gorm-kit/gorm.ParseLoggerLevel` 解析。 |
| `ConnMaxActive` | 连接池最大打开连接数。 |
| `ConnMaxLifetime` | 连接可复用的最长时间。 |
| `ConnMaxIdle` | 连接池最大空闲连接数。 |
| `ConnMaxIdleTime` | 空闲连接最长保留时间。 |

## 快速使用

```go
package data

import (
	"time"

	psqlpkg "github.com/ikaiguang/go-postgres-kit/postgres"
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/gorm"
)

func OpenPostgres() (*gorm.DB, error) {
	conf := &psqlpkg.Config{
		Dsn:             "host=127.0.0.1 user=postgres password=postgres dbname=app port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		SlowThreshold:   durationpb.New(100 * time.Millisecond),
		LoggerEnable:    true,
		LoggerLevel:     "INFO",
		ConnMaxActive:   100,
		ConnMaxLifetime: durationpb.New(30 * time.Minute),
		ConnMaxIdle:     10,
		ConnMaxIdleTime: durationpb.New(time.Hour),
	}

	if err := conf.ValidateAll(); err != nil {
		return nil, err
	}

	return psqlpkg.NewPostgresDB(conf)
}
```

重复键错误处理：

```go
if err := db.Create(&model).Error; err != nil {
	if psqlpkg.IsErrDuplicatedKey(err) {
		return err
	}
	return err
}
```

## 生成文件

以下文件由 `postgres/config.proto` 生成，不要手动修改：

- `config.pb.go`
- `config.pb.validate.go`
- `config.swagger.json`

修改 `config.proto` 后，在仓库根目录运行：

```bash
make protoc-config-protobuf
```

## 测试

```bash
go test ./postgres
```

当前测试 `TestNewDB_Xxx` 会使用 `postgres.kit_test.go` 中的本地 DSN，并调用 `Ping()` 检查数据库连接。运行该测试前需要本机 PostgreSQL 可访问，且存在测试数据库和用户。

如果只做文档调整，也可以运行全仓库检查：

```bash
go test ./...
```

## 注意事项

- `NewDB` 当前直接读取 `conf` 及其中的 duration 字段，调用前不要传入 `nil` 配置或未初始化的 duration 指针。
- `Dsn` 可能包含密码和内网地址，不要把真实生产 DSN 写入文档、测试或日志。
- 连接生命周期和连接池参数会传给底层 `database/sql`，应按服务并发量和 PostgreSQL 资源限制配置。
- `IsErrDuplicatedKey` 只判断重复键场景；其他 PostgreSQL 错误应由调用方按业务语义处理。
