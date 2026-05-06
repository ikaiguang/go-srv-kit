---
inclusion: manual
---

# Git 提交与分支规范

## Commit Message 格式

```
<type>(<scope>): <subject>
```

### Type 类型

| Type | 说明 |
|------|------|
| feat | 新功能 |
| fix | Bug 修复 |
| docs | 文档变更 |
| style | 代码格式（不影响逻辑） |
| refactor | 重构 |
| perf | 性能优化 |
| test | 测试 |
| chore | 构建/工具链变动 |

### Scope 范围

api, service, biz, data, kratos, kit, config, middleware, auth, wire

### 规范

- 使用中文
- subject 首字母小写，不加句号，不超过 50 字符
- 引用 Issue: `Closes #123`

示例: `feat(service): 实现用户登录接口`

## 分支策略

```
prod (生产) ← pre (预发布) ← test (测试) ← feature/xxx | hotfix/xxx
```

- 所有分支从 prod 创建
- 合并顺序: test → pre → prod
- 功能分支: `feature/{功能名}`
- 修复分支: `hotfix/{问题名}`
- 完成后删除 feature/hotfix 分支
