# Git 协作流程规范

参考： [Angular Commit Message Format](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#commit-message-format)

## 分支角色

| 分支 | 用途 |
|------|------|
| `prod` | 生产环境基线 |
| `pre` | 预发布验证 |
| `test` | 集成测试、自测联调 |
| `feature/*` | 新功能开发 |
| `hotfix/*` | 线上问题修复 |

## 推荐流程

### 新功能

1. 从 `prod` 拉出 `feature/*`
2. 在功能分支完成开发和自测
3. 合并到 `test`
4. 测试通过后合并到 `pre`
5. 预发布验证通过后合并到 `prod`

### 紧急修复

1. 从 `prod` 拉出 `hotfix/*`
2. 修复并自测
3. 依次回合并到 `test`、`pre`、`prod`

## Commit Message 格式

```text
<type>(<scope>): <subject>
```

例如：

```text
feat(auth): add refresh token validation
fix(config): avoid nil consul client
docs(readme): update startup commands
```

## 常用类型

| 类型 | 说明 |
|------|------|
| `feat` | 新功能 |
| `fix` | Bug 修复 |
| `docs` | 文档修改 |
| `style` | 纯格式调整，不改变逻辑 |
| `refactor` | 重构，不新增功能也不修复缺陷 |
| `perf` | 性能优化 |
| `test` | 测试补充或修正 |
| `chore` | 构建、脚本、依赖等杂项维护 |

## 建议

- 提交标题尽量简洁明确
- 一次提交只做一类事情
- 大改动优先拆成多个可审阅提交
- 合并前至少完成本地自测或最小验证
