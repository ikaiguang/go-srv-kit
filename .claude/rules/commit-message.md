# Git 提交信息规范

参考文档：[angular : Commit Message Format](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#commit-message-format)

## Commit Message 格式

```
<type>(<scope>): <subject>

<body>

<footer>
```

- **header** 是必需的，**scope** 是可选的
- 任何一行不能超过 100 个字符

## Type 类型

| Type | 说明 | 示例 |
|------|------|------|
| **feat** | 新功能 | feat: 添加用户注册功能 |
| **fix** | Bug 修复 | fix: 修复登录时密码验证错误 |
| **docs** | 文档变更 | docs: 更新 API 文档 |
| **style** | 代码格式（不影响逻辑） | style: 格式化代码 |
| **refactor** | 重构（不是新功能也不是修复） | refactor: 重构用户服务层 |
| **perf** | 性能优化 | perf: 优化数据库查询性能 |
| **test** | 添加或修改测试 | test: 添加用户服务单元测试 |
| **chore** | 构建/工具链变动 | chore: 更新依赖版本 |
| **revert** | 回滚提交 | revert: 回滚用户注册功能 |

## Scope 范围

常见的 Scope：

```
api          # API 定义
service      # Service 层
biz          # Business 层
data         # Data 层
kratos       # Kratos 框架扩展
kit          # 工具库
config       # 配置
middleware   # 中间件
auth         # 认证授权
wire         # 依赖注入
```

## 示例

### 简单提交

```
feat: 添加用户注册功能
```

### 带范围的提交

```
feat(service): 实现用户登录接口
```

### 完整提交

```
feat(biz): 实现用户创建业务逻辑

- 添加用户验证逻辑
- 检查用户名是否已存在
- 密码加密存储

Closes #123
```

### 修复 Bug

```
fix(data): 修复用户查询时未正确处理软删除

导致已删除用户仍然可以被查询到
```

### 文档更新

```
docs: 更新 CLAUDE.md 添加架构说明

补充了 go-srv-kit 与业务服务的关系说明
```

### 重构

```
refactor(service): 重构 DTO 转换逻辑

将转换函数统一放在 dto 目录下
```

## 提交规范

1. **使用中文**：项目使用中文，提交信息也使用中文
2. **首字母小写**：subject 首字母小写
3. **不要句号**：subject 结尾不要句号
4. **限制长度**：subject 不超过 50 字符
5. **空行分隔**：subject 和 body 之间空一行
6. **引用 Issue**：在 footer 中使用 `Closes #123`

## 常用命令

```bash
# 添加所有变更
git add .

# 提交
git commit -m "feat: 添加新功能"

# 查看提交历史
git log --oneline

# 查看最近 5 条提交
git log -5 --pretty=format:"%h - %s (%cr)"
```

---

## Git 分支管理流程

### 分支策略

```
┌─────────────────────────────────────────────────────────────┐
│  prod (生产环境)                                              │
│    ├── pre (预发布环境)                                       │
│    │    └── test (测试环境)                                   │
│    │         ├── feature/xxx (功能分支)                      │
│    │         └── hotfix/xxx (紧急修复分支)                   │
└─────────────────────────────────────────────────────────────┘
```

### 分支说明

| 分支 | 用途 | 合并目标 |
|------|------|----------|
| `prod` | 生产环境 |不接受直接合并 |
| `pre` | 预发布环境 | 合并到 prod |
| `test` | 测试环境 | 合并到 pre |
| `feature/xxx` | 新功能开发 | 合并到 test → pre → prod |
| `hotfix/xxx` | 紧急 Bug 修复 | 合并到 test → pre → prod |

### 开发流程

#### 1. 创建功能分支

```bash
git checkout prod
git checkout -b feature/user-service
```

#### 2. 开发并提交

```bash
git add .
git commit -m "feat: 实现用户注册功能"
```

#### 3. 合并到 test（自测后）

```bash
git checkout test
git merge feature/user-service
```

#### 4. 合并到 pre（测试通过后）

```bash
git checkout pre
git merge feature/user-service
```

#### 5. 合并到 prod（预发布验证后）

```bash
git checkout prod
git merge feature/user-service
```

### 紧急修复流程

#### 1. 创建 hotfix 分支

```bash
git checkout prod
git checkout -b hotfix/login-bug
```

#### 2. 修复并提交

```bash
git add .
git commit -m "fix: 修复登录验证错误"
```

#### 3. 合并到 test

```bash
git checkout test
git merge hotfix/login-bug
```

#### 4. 合并到 pre

```bash
git checkout pre
git merge hotfix/login-bug
```

#### 5. 合并到 prod

```bash
git checkout prod
git merge hotfix/login-bug
```

### 分支命名规范

| 类型 | 命名格式 | 示例 |
|------|----------|------|
| 功能分支 | `feature/{功能名}` | `feature/user-service` |
| 修复分支 | `hotfix/{问题名}` | `hotfix/login-bug` |
| 优化分支 | `refactor/{模块名}` | `refactor/database-layer` |

### 注意事项

1. **从 prod 创建分支**：所有功能分支和修复分支都应从 prod 分支创建
2. **逐步合并**：test → pre → prod，确保每个环境都经过验证
3. **删除已合并分支**：功能完成后删除 feature 分支
4. **保持分支简洁**：feature 分支只包含相关功能，不要混入其他修改
