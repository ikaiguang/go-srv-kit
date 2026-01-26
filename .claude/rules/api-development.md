# API 开发流程

## 新增 API 完整流程

### 1. 定义 Proto

在 `api/{service-name}/v1/` 下创建：

```
api/{service-name}/v1/
├── services/
│   └── {service}.proto      # RPC 定义
├── resources/
│   └── {resource}.proto     # 请求/响应消息
├── errors/
│   └── {error}.proto        # 错误定义
└── enums/
    └── {enum}.proto         # 枚举类型
```

**示例** (`api/ping-service/v1/services/ping.proto`):
```protobuf
syntax = "proto3";

package api.ping.service.v1;

import "api/ping-service/v1/resources/ping.proto";
import "google/api/annotations.proto";

option go_package = "github.com/ikaiguang/go-srv-kit/api/ping-service/v1;v1";

service SrvPing {
  // GetPingMessage 获取 ping 消息
  rpc GetPingMessage(GetPingMessageReq) returns (GetPingMessageResp) {
    option (google.api.http) = {
      get: "/api/v1/ping/say_hello"
    };
  }
}
```

### 2. 生成 Proto 代码

```bash
make api-{service-name}
# 或
protoc --proto_path=. --go_out=. --go-grpc_out=. api/{service-name}/v1/*.proto
```

### 3. 实现 Service Layer

在 `internal/service/service/{service}.service.go`:

```go
type XxxService struct {
    logger  log.Logger
    xxxBiz  biz.XxxBizRepo
}

func NewXxxService(logger log.Logger, xxxBiz biz.XxxBizRepo) *XxxService {
    return &XxxService{
        logger: logger,
        xxxBiz: xxxBiz,
    }
}

func (s *XxxService) XxxMethod(ctx context.Context, req *pb.XxxReq) (*pb.XxxResp, error) {
    // 1. 参数验证
    if req.GetMessage() == "" {
        return nil, errorpkg.ErrorBadRequest("message is required")
    }

    // 2. DTO → BO
    param := dto.ToBoXxxParam(req)

    // 3. 调用业务逻辑
    result, err := s.xxxBiz.XxxMethod(ctx, param)
    if err != nil {
        return nil, errorpkg.WrapWithMetadata(err, metadata)
    }

    // 4. BO → Proto Response
    return dto.ToProtoXxxResp(result), nil
}
```

### 4. 实现 DTO 转换

在 `internal/service/dto/{service}.dto.go`:

```go
func ToBoXxxParam(req *pb.XxxReq) *bo.XxxParam {
    return &bo.XxxParam{
        Message: req.GetMessage(),
    }
}

func ToProtoXxxResp(result *bo.XxxResult) *pb.XxxResp {
    return &pb.XxxResp{
        Message: result.Message,
    }
}
```

### 5. 实现 Business Layer

在 `internal/biz/biz/{service}.biz.go`:

```go
type XxxBiz struct {
    logger  log.Logger
    xxxRepo biz.XxxBizRepo
}

func NewXxxBiz(logger log.Logger, xxxRepo biz.XxxBizRepo) biz.XxxBizRepo {
    return &XxxBiz{
        logger:  logger,
        xxxRepo: xxxRepo,
    }
}

func (b *XxxBiz) XxxMethod(ctx context.Context, param *bo.XxxParam) (*bo.XxxResult, error) {
    // 业务逻辑
    return b.xxxRepo.XxxMethod(ctx, param)
}
```

### 6. 定义 Repository 接口

在 `internal/biz/repo/{service}.repo.go`:

```go
package repo

type XxxBizRepo interface {
    XxxMethod(ctx context.Context, param *bo.XxxParam) (*bo.XxxResult, error)
}
```

### 7. 实现 Data Layer

在 `internal/data/data/{service}.data.go`:

```go
type xxxData struct {
    logger log.Logger
    db     *gorm.DB
}

func NewXxxData(logger log.Logger, db *gorm.DB) biz.XxxBizRepo {
    return &xxxData{
        logger: logger,
        db:     db,
    }
}

func (d *xxxData) XxxMethod(ctx context.Context, param *bo.XxxParam) (*bo.XxxResult, error) {
    // 数据访问逻辑
    return result, nil
}
```

### 8. 更新 Wire 依赖注入

在 `cmd/{service}/export/wire.go`:

```go
func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
    panic(wire.Build(
        setuputil.GetLogger,
        data.NewXxxData,      // 添加
        biz.NewXxxBiz,        // 添加
        service.NewXxxService, // 添加
        service.RegisterServices,
    ))
}
```

### 9. 注册服务

在 `internal/service/service/register.go`:

```go
func RegisterServices(hs *http.Server, gs *grpc.Server, services ...interface{}) {
    xxxService := services[0].(*service.XxxService)
    v1.RegisterSrvXxxServer(gs, xxxService)
    v1.RegisterSrvXxxHTTPServer(hs, xxxService)
}
```

### 10. 生成 Wire 代码

```bash
wire ./cmd/{service}/export
```

## 更新 API 的流程

1. 修改 Proto 文件
2. 重新生成 Proto 代码
3. 更新对应的 Service/Biz/Data 层实现
4. 重新生成 Wire 代码

## 删除 API 的流程

1. 删除 Proto 定义
2. 删除 Service/Biz/Data 层实现
3. 从 Wire 中移除依赖
4. 重新生成 Wire 代码
