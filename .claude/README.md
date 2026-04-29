# .claude - Claude Code 配置目录

`.claude/` 存放面向 Claude Code 的规则、模板和辅助资料。

## 目录说明

| 路径 | 说明 |
|------|------|
| `.claude/rules/` | Claude 专用规则文件 |
| `.claude/prompts/` | 常见任务提示词模板 |
| `.claude/templates/` | 分层代码模板 |
| `.claude/memory/` | 项目记忆与常见问题归档 |
| `.claude/variables/` | 常用代码片段和变量说明 |
| `.claude/project-settings.json` | 项目结构化元数据 |

## 与仓库其他 AI 文档的关系

当前仓库已经同时维护以下几类 AI 辅助开发上下文：

- `CLAUDE.md`
  - Claude Code 会话入口级说明
- `AGENTS.md`
  - 通用 Agent 常驻规则
- `.agents/skills/go-srv-kit/`
  - Repo-local Codex skill 与分层 references

如果项目规范发生变化，建议优先同步这些高层入口文件，再根据需要同步 `.claude/rules/` 中的细分规则。
