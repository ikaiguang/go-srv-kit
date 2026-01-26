# Code Review 规范

## 审查原则

1. **友好建设性** - 提出改进建议，而非批评
2. **关注代码本身** - 对事不对人
3. **解释原因** - 说明为什么需要修改
4. **承认自己的无知** - 不确定时提问而非强求
5. **及时响应** - 尽快处理评论

## 通用审查清单

### 代码质量

- [ ] 代码是否符合 `coding-standards.md` 规范
- [ ] 命名是否清晰易懂
- [ ] 函数是否单一职责
- [ ] 是否有重复代码
- [ ] 是否有魔法数字或硬编码
- [ ] 注释是否准确且必要

### 错误处理

- [ ] 是否正确处理了所有错误
- [ ] 是否使用 `kratos/error/` 包
- [ ] 错误信息是否清晰
- [ ] 是否有必要的错误日志

### 安全性

- [ ] 是否有 SQL 注入风险
- [ ] 是否有 XSS 风险
- [ ] 敏感信息是否脱敏
- [ ] 权限是否正确校验
- [ ] 输入参数是否验证

### 性能

- [ ] 是否有 N+1 查询
- [ ] 是否有不必要的循环嵌套
- [ ] 是否正确使用缓存
- [ ] 数据库查询是否可优化

### 测试

- [ ] 是否有单元测试
- [ ] 测试覆盖率是否足够
- [ ] 是否有边界情况测试
- [ ] 测试是否可独立运行

## 分层审查重点

### Proto 层

```protobuf
// ✅ 好的实践
syntax = "proto3";
package api.user.service.v1;
option go_package = "github.com/ikaiguang/go-srv-kit/api/user-service/v1;v1";

// 字段注释
message CreateUserReq {
  string username = 1 [(validate.rules).string.min_len = 1]; // 有验证规则
}
```

**审查清单：**
- [ ] 包名是否符合 `api.{service}/v1` 格式
- [ ] `go_package` option 是否正确
- [ ] 字段是否有注释
- [ ] 是否有验证规则（`validate.rules`）
- [ ] 错误定义是否在 `errors/` 目录

### Service Layer

```go
// ✅ 好的实践
func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
    // 1. 参数验证
    if req.GetUsername() == "" {
        return nil, errorpkg.ErrorBadRequest("username is required")
    }

    // 2. DTO → BO
    param := dto.ToBoCreateUserParam(req)

    // 3. 调用业务逻辑
    result, err := s.userBiz.CreateUser(ctx, param)
    if err != nil {
        log.Context(ctx).Errorw("create user failed", "error", err)
        return nil, err
    }

    // 4. BO → Proto
    return dto.ToProtoCreateUserResp(result), nil
}
```

**审查清单：**
- [ ] 是否有参数验证
- [ ] DTO 转换是否正确
- [ ] 是否调用了 Business 层而非直接调用 Data 层
- [ ] 错误是否正确处理和记录日志
- [ ] 是否正确使用 Context

### Business Layer

```go
// ✅ 好的实践
func (b *userBiz) CreateUser(ctx context.Context, param *bo.CreateUserParam) (*bo.CreateUserResult, error) {
    // 1. 业务验证
    exists, err := b.userRepo.CheckUserExists(ctx, param.Username)
    if err != nil {
        return nil, errorpkg.FormatError(err)
    }
    if exists {
        return nil, customErrors.UserAlreadyExists()
    }

    // 2. 调用 Data 层
    result, err := b.userRepo.CreateUser(ctx, param)
    if err != nil {
        return nil, errorpkg.WrapWithMetadata(err, nil)
    }

    return result, nil
}
```

**审查清单：**
- [ ] 业务逻辑是否清晰
- [ ] 是否有业务规则验证
- [ ] 是否只调用 Repository 接口
- [ ] 事务处理是否正确
- [ ] 复杂逻辑是否有注释

### Data Layer

```go
// ✅ 好的实践
func (d *userData) CreateUser(ctx context.Context, param *bo.CreateUserParam) (*bo.CreateUserResult, error) {
    user := &po.User{
        Username: param.Username,
        Email:    param.Email,
    }

    if err := d.db.WithContext(ctx).Create(user).Error; err != nil {
        // 处理唯一约束冲突
        if errors.Is(err, gorm.ErrDuplicatedKey) {
            return nil, customErrors.UserAlreadyExists()
        }
        return nil, errorpkg.FormatError(err)
    }

    return &bo.CreateUserResult{ID: user.ID}, nil
}
```

**审查清单：**
- [ ] 是否使用 `WithContext(ctx)`
- [ ] PO 模型是否正确定义
- [ ] 是否正确处理 GORM 错误
- [ ] 是否有 SQL 注入风险
- [ ] 批量操作是否使用 Batch
- [ ] 是否正确使用事务

### Wire 依赖注入

```go
// ✅ 好的实践
func exportServices(launcher setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) {
    panic(wire.Build(
        // 基础设施
        setuputil.GetLogger,

        // Data 层
        data.NewUserData,

        // Business 层
        biz.NewUserBiz,

        // Service 层
        service.NewUserService,

        // 注册服务
        service.RegisterServices,
    ))
}
```

**审查清单：**
- [ ] 依赖顺序是否正确（从下到上）
- [ ] 接口是否使用 `wire.Bind`
- [ ] 是否有循环依赖
- [ ] 新增的依赖是否已添加
- [ ] 是否生成了 `wire_gen.go`

## 常见问题

### 架构违规

```go
// ❌ 错误：Service 直接调用 Data
func (s *userService) GetUser(ctx context.Context, id uint) (*bo.User, error) {
    return s.userData.GetUser(ctx, id)  // 违规！应该调用 Biz
}

// ✅ 正确
func (s *userService) GetUser(ctx context.Context, id uint) (*bo.User, error) {
    return s.userBiz.GetUser(ctx, id)
}
```

### 错误处理缺失

```go
// ❌ 错误：忽略错误
user, _ := s.userBiz.GetUser(ctx, id)

// ✅ 正确
user, err := s.userBiz.GetUser(ctx, id)
if err != nil {
    return nil, err
}
```

### Context 未传递

```go
// ❌ 错误：未传递 Context
users, _ := d.db.Find(&users).Error

// ✅ 正确
users, _ := d.db.WithContext(ctx).Find(&users).Error
```

### 硬编码配置

```go
// ❌ 错误：硬编码
timeout := 30 * time.Second

// ✅ 正确：从配置读取
timeout := time.Duration(config.GetTimeout()) * time.Second
```

### 敏感信息日志

```go
// ❌ 错误：记录密码
log.Infow("user login", "password", password)

// ✅ 正确：脱敏
log.Infow("user login", "password", stringutil.MaskPassword(password))
```

## 审查评论模板

### 建议修改

```markdown
### 建议：[标题]

**当前代码：**
\`\`\`go
// 粘贴代码
\`\`\`

**问题：**
说明具体问题

**建议修改：**
\`\`\`go
// 建议的代码
\`\`\`

**原因：**
解释为什么这样改更好
```

### 必须修改

```markdown
### 🔴 必须修复：[标题]

**问题：**
[严重问题说明]

**影响：**
[如果不修复会有什么后果]

**建议：**
[修复方案]
```

### 可以优化

```markdown
### 💡 优化建议：[标题]

**当前实现：**
[描述]

**优化方案：**
[描述]

**收益：**
[性能/可读性/维护性提升]
```

## 审查流程

### 1. 自动检查

```bash
# 运行测试
go test ./...

# 格式检查
gofmt -l .

# 静态分析
go vet ./...
golangci-lint run
```

### 2. 人工审查

按照审查清单逐项检查

### 3. 反馈

- 使用友好的语气
- 提供具体的改进建议
- 解释原因

### 4. 跟进

- 作者修改后及时重新审查
- 修改满意后批准（LGTM）
- 有大问题则请求变更（Request Changes）

## 审查响应

### 对于评论者

- 提出问题后关注作者回复
- 讨论达成共识后更新评论状态

### 对于作者

- 积极回应每条评论
- 不同意时说明理由
- 修改后标记已解决

## 审查工具

### GitHub 功能

- Review Changes - 逐文件审查
- Line Comments - 行内评论
- Suggestions - 代码建议
- Approve/Request Changes - 审查决策

### 本地工具

```bash
# 查看 PR 变更
git diff main...feature-branch

# 查看特定文件
git diff main...feature-branch -- path/to/file
```
