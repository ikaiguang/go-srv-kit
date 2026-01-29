# Business 层代码模板

## Business 实现模板
```go
package biz

import (
    "context"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/biz/bo"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/biz/repo"
    "github.com/go-kratos/kratos/v2/log"
)

type {Xxx}Biz struct {
    logger  log.Logger
    {xxx}Repo repo.{Xxx}BizRepo
}

func New{Xxx}Biz(logger log.Logger, {xxx}Repo repo.{Xxx}BizRepo) repo.{Xxx}BizRepo {
    return &{Xxx}Biz{
        logger:  logger,
        {xxx}Repo: {xxx}Repo,
    }
}

// {ApiName} {api_description}
func (b *{Xxx}Biz) {ApiName}(ctx context.Context, param *bo.{ApiName}Param) (*bo.{ApiName}Result, error) {
    // 1. 业务验证
    {business_validation}

    // 2. 调用 Data 层
    result, err := b.{xxx}Repo.{ApiName}(ctx, param)
    if err != nil {
        return nil, errorpkg.WrapWithMetadata(err, nil)
    }

    return result, nil
}
```

## Repository 接口模板
```go
package repo

import (
    "context"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/biz/bo"
)

// {Xxx}BizRepo {xxx} 业务仓库接口
type {Xxx}BizRepo interface {
    // {ApiName} {api_description}
    {ApiName}(ctx context.Context, param *bo.{ApiName}Param) (*bo.{ApiName}Result, error)
}
```

## BO 定义模板
```go
package bo

// {ApiName}Param {api_name} 参数
type {ApiName}Param struct {
    {fields}
}

// {ApiName}Result {api_name} 结果
type {ApiName}Result struct {
    {fields}
}
```
