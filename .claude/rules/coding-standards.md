# 编码规范

## Go 通用规范

- 遵循 [Go 官方代码风格](https://go.dev/doc/effective_go)
- 使用 `gofmt` 格式化代码
- 导出的函数、类型、常量必须有注释
- 包名使用小写单词，不使用下划线或驼峰
- 错误处理优先使用 `errors.Is` 和 `errors.As`
- 避免使用 `panic`，使用 `error` 返回错误

## go-srv-kit 项目规范

### 文件命名

```
internal/service/service/ping.service.go
internal/biz/biz/ping.biz.go
internal/data/data/ping.data.go
internal/service/dto/ping.dto.go
internal/biz/bo/ping.bo.go
internal/data/po/ping.po.go
```

### 导入顺序

```go
import (
	// 1. 标准库
	"context"
	"fmt"

	// 2. 第三方库
	"github.com/go-kratos/kratos/v2"
	"google.golang.org/protobuf/proto"

	// 3. 项目内部
	"github.com/ikaiguang/go-srv-kit/kratos/error"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/bo"
)
```

### 接口定义位置

- Repository 接口定义在 `internal/biz/repo/`
- 由 `internal/data/repo/` 实现
- 使用 Wire 的 `wire.Bind` 进行绑定

### 数据转换

- DTO (Service) ↔ BO (Biz): `internal/service/dto/`
- BO (Biz) ↔ PO (Data): `internal/biz/bo/` 或各自层内部
- 转换函数命名：`ToBo{Xxx}`, `ToProto{Xxx}`, `ToPo{Xxx}`

### 命名约定

| 类型 | 命名 | 示例 |
|------|------|------|
| Service 结构体 | `New{Xxx}Service` | `NewPingService` |
| Biz 结构体 | `New{Xxx}Biz` | `NewPingBiz` |
| Data 结构体 | `New{Xxx}Data` | `NewPingData` |
| 接口 | `{Xxx}Repo` | `PingBizRepo` |
| DTO 转换 | `ToBo{Xxx}`, `ToProto{Xxx}` | `ToBoGetPingParam` |

### 注释规范

```go
// PingService ping 服务
type PingService struct {
    // logger 日志
    logger log.Logger
    // pingBiz 业务逻辑
    pingBiz biz.PingBizRepo
}

// GetPingMessage 获取 ping 消息
func (s *PingService) GetPingMessage(ctx context.Context, req *pb.GetPingMessageReq) (*pb.GetPingMessageResp, error) {
```

## 禁止事项

- 禁止在代码中硬编码配置值
- 禁止在业务层直接访问外部服务
- 禁止跨层调用（Service 不能直接调用 Data）
- 禁止在 Proto 文件中生成代码后再手动修改
