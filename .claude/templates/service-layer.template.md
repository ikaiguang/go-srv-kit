# Service 层代码模板

## Service 实现模板
```go
package service

import (
    "context"
    v1 "github.com/ikaiguang/go-srv-kit/api/{service_name}/v1"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/biz/biz"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/service/dto"
    "github.com/go-kratos/kratos/v2/log"
)

type {Xxx}Service struct {
    logger log.Logger
    {xxx}Biz biz.{Xxx}BizRepo
}

func New{Xxx}Service(logger log.Logger, {xxx}Biz biz.{Xxx}BizRepo) *{Xxx}Service {
    return &{Xxx}Service{
        logger:  logger,
        {xxx}Biz: {xxx}Biz,
    }
}

// {ApiName} {api_description}
func (s *{Xxx}Service) {ApiName}(ctx context.Context, req *v1.{ApiName}Req) (*v1.{ApiName}Resp, error) {
    // 1. 参数验证
    if req.Get{Field}() == "" {
        return nil, errorpkg.ErrorBadRequest("{field} is required")
    }

    // 2. DTO → BO
    param := dto.ToBo{ApiName}Param(req)

    // 3. 调用业务逻辑
    result, err := s.{xxx}Biz.{ApiName}(ctx, param)
    if err != nil {
        log.Context(ctx).Errorw("{api_name} failed", "error", err)
        return nil, err
    }

    // 4. BO → Proto
    return dto.ToProto{ApiName}Resp(result), nil
}
```

## DTO 转换模板
```go
package dto

import (
    v1 "github.com/ikaiguang/go-srv-kit/api/{service_name}/v1"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/biz/bo"
)

// ToBo{ApiName}Param 转换请求参数为 BO
func ToBo{ApiName}Param(req *v1.{ApiName}Req) *bo.{ApiName}Param {
    return &bo.{ApiName}Param{
        {fields}
    }
}

// ToProto{ApiName}Resp 转换 BO 为响应
func ToProto{ApiName}Resp(result *bo.{ApiName}Result) *v1.{ApiName}Resp {
    return &v1.{ApiName}Resp{
        {fields}
    }
}
```
