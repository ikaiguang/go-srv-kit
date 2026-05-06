---
inclusion: always
---

# Go 编码规范

参考: [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) | [Effective Go](https://go.dev/doc/effective_go)

## 格式化与风格

- 代码必须用 `gofmt` 格式化，使用 `goimports` 自动格式化 import
- 建议一行不超过 160 列
- 运算符和操作数之间留空格
- import 分组：标准包 → 第三方包 → 本地包（空行分隔），匿名包单独分组并写注释

```go
import (
    "context"
    "fmt"

    "github.com/go-kratos/kratos/v2"

    "github.com/ikaiguang/go-srv-kit/kratos/error"

    _ "github.com/go-sql-driver/mysql" // import mysql driver
)
```

## 错误处理

- error 必须是最后一个返回参数，不能用 `_` 丢弃
- 采用独立错误流，尽早 return err
- 错误信息小写开头、无句号
- 错误判断不与其他逻辑组合
- 业务逻辑中禁止使用 panic
- 异步 goroutine 必须使用 `threadpkg.GoSafe()` 包装
- 类型断言始终使用 `comma ok` 惯用法

## 函数规范

- 函数长度不超过 150 行
- 嵌套层级不超过 3 层（超过必须重构）
- 接收器命名用类名首字母小写，禁止 me/this/self
- defer 在操作成功后再调用

### 函数参数约定

参考: [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments#function-arguments) | [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

- 参数数量不超过 4 个，超过必须封装为结构体或使用 Options 模式
- 参数固定顺序：`ctx context.Context` → 请求结构体/主参数 → 其他参数 → 可选参数 `...Option`
- `context.Context` 必须作为第一个参数，且命名为 `ctx`，禁止放在结构体中
- 不需要 context 的函数不要强加 `ctx` 参数

```go
// ✅ 标准写法
func CreateUser(ctx context.Context, param *CreateUserParam) (*CreateUserResult, error)
func NewServer(conf *configpb.Bootstrap, opts ...Option) (*Server, error)
func GetLogger(launcherManager LauncherManager) (log.Logger, error)

// ❌ 错误：ctx 不在第一个
func CreateUser(param *CreateUserParam, ctx context.Context) error
// ❌ 错误：参数过多，应封装结构体
func CreateUser(ctx context.Context, name string, email string, phone string, age int, role string) error
```

### 函数返回值约定

- `error` 必须是最后一个返回值
- 返回值不超过 3 个（不含 error），超过优先使用结构体封装
- 命名返回值仅在需要 defer 中修改时使用，其他场景使用匿名返回值
- 构造函数返回值遵循：`(实例, error)` 或 `(实例, cleanup, error)`

```go
// ✅ 标准返回值
func GetUser(ctx context.Context, id uint) (*User, error)
func NewWithCleanup(path string) (LauncherManager, func(), error)

// ❌ 错误：error 不在最后
func GetUser(ctx context.Context, id uint) (error, *User)
// ❌ 错误：返回值过多，应封装结构体
func GetOrder(ctx context.Context, id uint) (string, int, float64, time.Time, error)
```

## 命名约定

| 类型 | 格式 | 示例 |
|------|------|------|
| Service | `New{Xxx}Service` | `NewPingService` |
| Biz | `New{Xxx}Biz` | `NewPingBiz` |
| Data | `New{Xxx}Data` | `NewPingData` |
| Repository 接口 | `{Xxx}BizRepo` | `PingBizRepo` |
| DTO 转换 | `ToBo{Xxx}`, `ToProto{Xxx}` | `ToBoGetPingParam` |

## 文件命名

```
internal/service/service/ping.service.go
internal/biz/biz/ping.biz.go
internal/data/data/ping.data.go
internal/service/dto/ping.dto.go
internal/biz/bo/ping.bo.go
internal/data/po/ping.po.go
```

## 禁止事项

- 禁止硬编码配置值
- 禁止业务层直接访问外部服务
- 禁止跨层调用（Service 不能直接调用 Data）
- 禁止手动修改 Proto 生成的代码
- 禁止提交注释掉的代码（除非有说明）
- 禁止保留未使用的业务函数
- 魔法数字出现超过 2 次必须定义常量
- 禁止复制粘贴第三方 API 调用代码（必须封装为可复用组件）

## 减少嵌套技巧

- 提前返回（Guard Clauses）
- 循环中使用 continue/break
- 将深层逻辑提取为独立函数
- 卫语句：`if !condition { return }` 代替 `if condition { ... }`
