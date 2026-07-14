# Repo-Local Agent Context

本目录保存本仓库专用的 agent 上下文。

- `.agents/skills/my-project/`：当前 GVA 后端项目的架构、命令、编码和消息中心约定。
- `.agents/skills/spec-workflow/`：需要先规格后执行时使用的流程。
- `.agents/skills/code-audit-repair/`：全库审计、补测试、补文档和低风险修复。
- `.agents/specs/`：临时任务过程记录，已在 `.gitignore` 中忽略，不作为长期规则提交。

全局通用 skills 放在 `~/.codex/skills`。本地 skills 只补充本项目事实，不复制外部仓库的维护规则。
