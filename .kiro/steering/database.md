---
inclusion: fileMatch
fileMatchPattern: "**/data/**/*.go"
---

# 数据库操作规范

## GORM 使用规范

- 始终使用 `WithContext(ctx)` 传递 context
- 事务使用 `d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error { ... })`
- 分页：`Offset((page-1)*pageSize).Limit(pageSize)`
- 预加载：`Preload("Profile")`
- 只查需要的字段：`Select("id, username")`
- 批量插入：`CreateInBatches(users, 100)`

## PO 定义

```go
type User struct {
    gorm.Model
    UUID     string `gorm:"type:varchar(36);uniqueIndex;not null"`
    Username string `gorm:"type:varchar(50);uniqueIndex;not null"`
    Status   int    `gorm:"type:tinyint;default:1;comment:1=正常,0=禁用"`
}
```

## Redis 使用

- 使用 context：`r.redis.Set(ctx, key, data, time.Hour)`
- 批量操作用 Pipeline
- 分布式锁：`r.redis.SetNX(ctx, "lock:"+key, 1, expiration)`

## 缓存策略（Cache-Aside）

1. 先查缓存 → 2. 缓存未命中查数据库 → 3. 写入缓存

## 错误处理

```go
if errors.Is(err, gorm.ErrRecordNotFound) {
    return nil, errorpkg.ErrorNotFound("not found")
}
```
