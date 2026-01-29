# Data 层代码模板

## Data 实现模板
```go
package data

import (
    "context"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/biz/bo"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/biz/repo"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/data/po"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

type {xxx}Data struct {
    logger log.Logger
    db     *gorm.DB
}

func New{Xxx}Data(logger log.Logger, db *gorm.DB) repo.{Xxx}BizRepo {
    return &{xxx}Data{
        logger: logger,
        db:     db,
    }
}

// {ApiName} {api_description}
func (d *{xxx}Data) {ApiName}(ctx context.Context, param *bo.{ApiName}Param) (*bo.{ApiName}Result, error) {
    // 1. BO → PO
    {bo_to_po}

    // 2. 数据库操作
    var {po_var} po.{Xxx}
    err := d.db.WithContext(ctx).{operation}(&{po_var}).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errorpkg.ErrorNotFound("{xxx} not found")
        }
        return nil, errorpkg.FormatError(err)
    }

    // 3. PO → BO
    return {po_to_bo}, nil
}
```

## PO 定义模板
```go
package po

import "gorm.io/gorm"

// {Xxx} {xxx} 表
type {Xxx} struct {
    gorm.Model
    {fields}
}
```

## 数据操作模板
```go
// 查询单条
var user po.User
err := d.db.WithContext(ctx).Where("id = ?", id).First(&user).Error

// 查询列表
var users []*po.User
err := d.db.WithContext(ctx).Find(&users).Error

// 创建
user := &po.User{...}
err := d.db.WithContext(ctx).Create(user).Error

// 更新
err := d.db.WithContext(ctx).Model(&po.User{}).Where("id = ?", id).Updates(updates).Error

// 删除（软删除）
err := d.db.WithContext(ctx).Delete(&po.User{}, id).Error

// 事务
err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user).Error; err != nil {
        return err
    }
    return tx.Create(&profile).Error
})
```
