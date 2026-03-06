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
- 参数数量不超过 5 个
- 接收器命名用类名首字母小写，禁止 me/this/self
- defer 在操作成功后再调用

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
