---
inclusion: manual
---

# 代码模板

## Service 层模板

```go
type XxxService struct {
    logger log.Logger
    xxxBiz biz.XxxBizRepo
}

func NewXxxService(logger log.Logger, xxxBiz biz.XxxBizRepo) *XxxService {
    return &XxxService{logger: logger, xxxBiz: xxxBiz}
}

func (s *XxxService) Method(ctx context.Context, req *pb.MethodReq) (*pb.MethodResp, error) {
    // 1. 参数验证
    if req.GetField() == "" {
        return nil, errorpkg.ErrorBadRequest("field is required")
    }
    // 2. DTO → BO
    param := dto.ToBoMethodParam(req)
    // 3. 调用业务逻辑
    result, err := s.xxxBiz.Method(ctx, param)
    if err != nil {
        logpkg.WithContext(ctx).Errorw("method failed", "error", err)
        return nil, err
    }
    // 4. BO → Proto
    return dto.ToProtoMethodResp(result), nil
}
```

## Business 层模板

```go
type XxxBiz struct {
    logger  log.Logger
    xxxRepo repo.XxxBizRepo
}

func NewXxxBiz(logger log.Logger, xxxRepo repo.XxxBizRepo) repo.XxxBizRepo {
    return &XxxBiz{logger: logger, xxxRepo: xxxRepo}
}

func (b *XxxBiz) Method(ctx context.Context, param *bo.Param) (*bo.Result, error) {
    // 1. 业务验证
    // 2. 调用 Data 层
    result, err := b.xxxRepo.Method(ctx, param)
    if err != nil {
        return nil, errorpkg.WrapWithMetadata(err, nil)
    }
    return result, nil
}
```

## Repository 接口模板

```go
// internal/biz/repo/{name}.repo.go
type XxxBizRepo interface {
    Method(ctx context.Context, param *bo.Param) (*bo.Result, error)
}
```

## Data 层模板

```go
type xxxData struct {
    logger log.Logger
    db     *gorm.DB
}

func NewXxxData(logger log.Logger, db *gorm.DB) repo.XxxBizRepo {
    return &xxxData{logger: logger, db: db}
}

func (d *xxxData) Method(ctx context.Context, param *bo.Param) (*bo.Result, error) {
    var po po.Xxx
    err := d.db.WithContext(ctx).First(&po, param.ID).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errorpkg.ErrorNotFound("not found")
        }
        return nil, errorpkg.FormatError(err)
    }
    return toBoResult(&po), nil
}
```

## Wire 模板

```go
//go:build wireinject

package exporter

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (Cleanup, error) {
    panic(wire.Build(
        setuputil.GetLogger,
        data.NewXxxData,
        biz.NewXxxBiz,
        service.NewXxxService,
        service.RegisterServices,
    ))
}
```

## Proto 模板

```protobuf
syntax = "proto3";
package api.{service}.v1;
option go_package = "github.com/ikaiguang/go-srv-kit/api/{service}/v1;v1";

import "google/api/annotations.proto";
import "validate/validate.proto";

service Srv{Name} {
  rpc Method(MethodReq) returns (MethodResp) {
    option (google.api.http) = {
      post: "/api/v1/{resource}/{action}"
      body: "*"
    };
  }
}

message MethodReq {
  string field = 1 [(validate.rules).string.min_len = 1];
}

message MethodResp {
  string result = 1;
}
```
