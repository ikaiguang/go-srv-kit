---
inclusion: manual
---

# Code Review 审查清单

## 代码质量

- [ ] 符合编码规范（Uber Go Style）
- [ ] 命名清晰易懂
- [ ] 函数单一职责，长度不超过 150 行
- [ ] 嵌套层级不超过 3 层
- [ ] 无重复代码、无魔法数字
- [ ] 第三方 API 调用已封装为可复用组件

## 架构合规

- [ ] Service 层只调用 Business 层
- [ ] Repository 接口在 biz/repo/，实现在 data/repo/
- [ ] Wire 依赖顺序正确（基础设施 → Data → Biz → Service）
- [ ] DTO/BO/PO 转换正确

## 错误处理

- [ ] 使用 kratos/error 包
- [ ] 所有错误已处理（无 `_` 丢弃）
- [ ] 错误信息清晰
- [ ] 有必要的错误日志

## 安全性

- [ ] 无 SQL 注入风险
- [ ] 敏感信息已脱敏
- [ ] 权限正确校验
- [ ] 输入参数已验证

## 性能

- [ ] 无 N+1 查询
- [ ] 正确使用缓存
- [ ] 数据库查询可优化

## 测试

- [ ] 有单元测试
- [ ] 覆盖率足够
- [ ] 有边界情况测试

## 常见问题

```go
// ❌ Service 直接调用 Data
s.userData.GetUser(ctx, id)
// ✅ 通过 Biz 层
s.userBiz.GetUser(ctx, id)

// ❌ 忽略错误
user, _ := s.userBiz.GetUser(ctx, id)
// ✅ 处理错误
user, err := s.userBiz.GetUser(ctx, id)
if err != nil { return nil, err }

// ❌ 未传递 Context
d.db.Find(&users)
// ✅ 传递 Context
d.db.WithContext(ctx).Find(&users)

// ❌ 记录密码
log.Infow("login", "password", password)
// ✅ 脱敏
log.Infow("login", "password", stringutil.MaskPassword(password))
```
