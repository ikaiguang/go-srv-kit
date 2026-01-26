# Git 提交信息规范

## 提交信息格式

```
<type>(<scope>): <subject>

<body>

<footer>
```

## Type 类型

| Type | 说明 | 示例 |
|------|------|------|
| feat | 新功能 | feat: 添加用户注册功能 |
| fix | 修复 Bug | fix: 修复登录时密码验证错误 |
| docs | 文档更新 | docs: 更新 API 文档 |
| style | 代码格式调整（不影响功能） | style: 格式化代码 |
| refactor | 重构（不是新功能也不是修复） | refactor: 重构用户服务层 |
| perf | 性能优化 | perf: 优化数据库查询性能 |
| test | 添加或修改测试 | test: 添加用户服务单元测试 |
| chore | 构建/工具链/辅助工具变动 | chore: 更新依赖版本 |
| revert | 回滚提交 | revert: 回滚用户注册功能 |

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
