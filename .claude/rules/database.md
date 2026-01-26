# 数据库操作规范

## 数据层开发

### 定义 PO (Persistent Object)

在 `internal/data/po/` 定义数据库模型：

```go
package po

import "gorm.io/gorm"

// User 用户表
type User struct {
    gorm.Model
    // UUID 用户唯一标识
    UUID string `gorm:"type:varchar(36);uniqueIndex;not null"`
    // Username 用户名
    Username string `gorm:"type:varchar(50);uniqueIndex;not null"`
    // Email 邮箱
    Email string `gorm:"type:varchar(100);uniqueIndex;not null"`
    // Status 状态
    Status int `gorm:"type:tinyint;default:1;comment:1=正常,0=禁用"`
}
```

### GORM 使用规范

```go
// 1. 使用 WithContext 传递 context
func (d *userData) CreateUser(ctx context.Context, user *po.User) error {
    return d.db.WithContext(ctx).Create(user).Error
}

// 2. 使用事务
func (d *userData) CreateUserWithProfile(ctx context.Context, user *po.User, profile *po.UserProfile) error {
    return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(user).Error; err != nil {
            return err
        }
        profile.UserID = user.ID
        return tx.Create(profile).Error
    })
}

// 3. 分页查询
func (d *userData) ListUsers(ctx context.Context, page, pageSize int) ([]*po.User, int64, error) {
    var users []*po.User
    var total int64

    if err := d.db.WithContext(ctx).Model(&po.User{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    offset := (page - 1) * pageSize
    err := d.db.WithContext(ctx).Offset(offset).Limit(pageSize).Find(&users).Error
    return users, total, err
}

// 4. 软删除
func (d *userData) DeleteUser(ctx context.Context, id uint) error {
    return d.db.WithContext(ctx).Delete(&po.User{}, id).Error
}

// 5. 预加载关联
func (d *userData) GetUserWithProfile(ctx context.Context, id uint) (*po.User, error) {
    var user po.User
    err := d.db.WithContext(ctx).Preload("Profile").First(&user, id).Error
    return &user, err
}
```

## 数据库迁移

### 创建迁移文件

```bash
# 使用项目提供的迁移工具
go run ./cmd/database-migration/... -conf=./configs
```

### 迁移文件位置

```
internal/data/schema/
├── migrations/
│   ├── 20240101_init_schema.up.sql
│   └── 20240101_init_schema.down.sql
```

### 自动迁移

```go
// 在 data 初始化时
func (d *Data) AutoMigrate() error {
    return d.db.AutoMigrate(
        &po.User{},
        &po.UserProfile{},
    )
}
```

## Redis 使用规范

```go
// 1. 使用 context
func (d *dataCache) SetUser(ctx context.Context, key string, user *bo.User) error {
    data, _ := json.Marshal(user)
    return d.redis.Set(ctx, key, data, time.Hour).Err()
}

// 2. 批量操作
func (d *dataCache) MGetUsers(ctx context.Context, keys []string) (map[string]*bo.User, error) {
    vals, err := d.redis.MGet(ctx, keys...).Result()
    // ...
}

// 3. 使用 Pipeline
func (d *dataCache) PipelineSet(ctx context.Context, items map[string]interface{}) error {
    pipe := d.redis.Pipeline()
    for key, val := range items {
        pipe.Set(ctx, key, val, time.Hour)
    }
    _, err := pipe.Exec(ctx)
    return err
}

// 4. 分布式锁
func (d *dataCache) Lock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
    return d.redis.SetNX(ctx, fmt.Sprintf("lock:%s", key), 1, expiration).Result()
}
```

## 缓存策略

### Cache-Aside 模式

```go
func (d *userData) GetUser(ctx context.Context, id uint) (*po.User, error) {
    // 1. 先查缓存
    cacheKey := fmt.Sprintf("user:%d", id)
    cached, err := d.redis.Get(ctx, cacheKey).Result()
    if err == nil {
        var user po.User
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }

    // 2. 查数据库
    var user po.User
    if err := d.db.WithContext(ctx).First(&user, id).Error; err != nil {
        return nil, err
    }

    // 3. 写缓存
    data, _ := json.Marshal(user)
    d.redis.Set(ctx, cacheKey, data, time.Hour)

    return &user, nil
}
```

## 错误处理

```go
import "gorm.io/gorm"

// 处理 GORM 错误
func (d *userData) GetUser(ctx context.Context, id uint) (*po.User, error) {
    var user po.User
    err := d.db.WithContext(ctx).First(&user, id).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errorpkg.ErrorNotFound("user not found")
        }
        return nil, errorpkg.WrapWithMetadata(err, nil)
    }

    return &user, nil
}
```

## 性能优化

```go
// 1. 使用索引
type User struct {
    Email string `gorm:"index"`
}

// 2. 只查询需要的字段
d.db.Select("id, username").Find(&users)

// 3. 批量插入
d.db.CreateInBatches(users, 100)

// 4. 使用连接池
db.DB().SetMaxIdleConns(10)
db.DB().SetMaxOpenConns(100)
db.DB().SetConnMaxLifetime(time.Hour)
```
