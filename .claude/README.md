# Claude Code 规则配置

本目录包含 `go-srv-kit` 项目的 Claude Code 智能开发规则配置。

## 配置文件说明

### 规则文件 (.claude/rules/)

| 文件 | 说明 |
|------|------|
| `coding-standards.md` | Go 通用编码规范和项目特定规范 |
| `api-development.md` | API 开发完整流程（Proto → Service → Biz → Data） |
| `testing.md` | 单元测试、集成测试规范 |
| `database.md` | 数据库操作规范（GORM、Redis、迁移） |
| `wire-dependency-injection.md` | Wire 依赖注入使用规范 |
| `error-handling.md` | 错误处理规范（kratos/error 包） |
| `logging.md` | 日志记录规范（Zap 结构化日志） |
| `authentication.md` | JWT 认证与授权规范 |
| `commit-message.md` | Git 提交信息规范和分支管理流程 |
| `code-review.md` | Code Review 审查清单和流程 |
| `go-private-setup.md` | Go 私有包配置（GOPROXY、GOPRIVATE） |
| `git-submodule.md` | Git 子模块使用和管理 |
| `project-workflow.md` | 项目开发工作流 |

### 项目设置 (.claude/project-settings.json)

包含项目元数据、架构信息、常用命令等结构化配置。

## 使用方式

这些规则会在以下场景自动加载：

1. **新会话开始** - Claude Code 自动加载 `.claude/rules/` 下的所有规则
2. **代码生成** - 根据 `api-development.md` 规范生成代码
3. **错误修复** - 根据 `error-handling.md` 规范修复错误
4. **代码审查** - 根据 `coding-standards.md` 进行代码审查

## 快速索引

### 开发新 API

查看 `api-development.md` → 完整的 API 开发流程

### 依赖注入问题

查看 `wire-dependency-injection.md` → Wire 使用规范和常见问题

### 数据库操作

查看 `database.md` → GORM、Redis、迁移规范

### 错误处理

查看 `error-handling.md` → kratos/error 包使用方式

### 测试

查看 `testing.md` → 单元测试、集成测试规范

### 认证授权

查看 `authentication.md` → JWT 认证和权限控制

## 更新规则

当项目架构或规范变化时，请同步更新对应的规则文件。

## 相关文档

- `CLAUDE.md` - 项目整体架构和开发指南（每次会话自动加载）
- `README.md` - 项目简介和快速开始
