# Add Spec Workflow Skill

## 背景

当前仓库已经在 `AGENTS.md` 中规定“先文档后执行”，并在 `.agents/skills/go-srv-kit/` 中沉淀了仓库架构、编码规范、生成命令和工作流参考。

这套机制已经覆盖了 Kiro specs 模式的核心，但还缺少一个专门负责“规格先行流程”的 repo-local skill。现在的规则主要依赖 `AGENTS.md` 的常驻文字，缺少可复用、可扩展、可触发的流程说明。

## 目标

新增一个 repo-local skill：`.agents/skills/spec-workflow/`。

该 skill 用于在 Codex 中模拟并增强 Kiro specs 模式，明确以下事项：

- 什么任务需要先创建规格文档
- 规格文档写在哪里、用什么结构
- 什么时候必须停下来等待用户确认
- 执行过程中发现方案变化时如何记录和重新确认
- 低风险、中风险、高风险操作如何区分
- 与 `AGENTS.md` 和 `.agents/skills/go-srv-kit/` 的边界如何划分

## 非目标

- 不修改业务代码、生成代码、Wire 装配或 Go 模块依赖
- 不调整现有 `AGENTS.md` 的规则内容
- 不替代 `.agents/skills/go-srv-kit/` 中已有的仓库专用上下文
- 不引入 Codex plugin；本次只新增 skill
- 不建立完整自动化评测体系；本次先完成可用的初版 skill

## 影响范围

预计新增文件：

- `.agents/skills/spec-workflow/SKILL.md`

可选新增文件：

- `.agents/skills/spec-workflow/references/spec-template.md`

本次优先将模板写在 `SKILL.md` 中，只有当正文过长或后续需要多种模板时，再拆到 `references/`。

## 设计方案

### 分层边界

`AGENTS.md` 保持为常驻宪法，只放最高优先级的简短规则。

`.agents/skills/go-srv-kit/` 继续负责仓库技术上下文，包括 Go、Kratos、DDD 分层、Wire、Proto、生成命令、日志和错误处理。

`.agents/skills/spec-workflow/` 只负责规格先行流程，包括文档模板、确认节点、执行记录和风险分级。

### 触发方式

skill 的 description 需要写得偏主动，确保以下任务会触发：

- 实现功能
- 修复 bug
- 重构代码
- 修改配置
- 修改 Wire 或 Proto
- 新增测试
- 执行有副作用命令
- 用户明确提到 specs、规格、计划、先文档、Kiro 模式

### 文档形态

默认使用单文件：

```text
docs/<task-name>/spec.md
```

文档结构：

```md
# <Task Title>

## 背景

## 目标

## 非目标

## 影响范围

## 方案

## 任务列表

## 验收标准

## 风险与回滚

## 执行记录
```

大任务后续可以拆成：

```text
docs/<task-name>/
├── 00-context.md
├── 01-requirements.md
├── 02-design.md
├── 03-tasks.md
└── 04-log.md
```

初版 skill 只要求默认单文件，避免流程过重。

### 确认节点

创建或更新任务规格后，需要停止实际执行并等待用户确认。

用户确认后，低风险操作可直接执行，包括：

- 读取和搜索文件
- 修改当前仓库内与任务相关的普通文件
- 运行本地格式化
- 运行本地测试
- 运行约定内代码生成

高风险操作必须再次确认，包括：

- 删除大量文件
- 改写 Git 历史
- 丢弃未确认修改
- 安装或修改系统级依赖
- 发布部署
- 请求第三方或线上 API 修改数据
- 处理未脱敏敏感信息
- 风险不确定的操作

### 执行记录

执行阶段应在 `spec.md` 的“执行记录”中记录：

- 开始执行的时间或阶段说明
- 已完成的关键步骤
- 执行过的验证命令
- 验证结果
- 偏离原方案的原因和新方案

如果发现原方案不适合，需要先更新文档说明：

- 发现的问题
- 原方案为什么不合适
- 准备怎么改
- 为什么这样改

如果变更会扩大影响范围或引入中高风险，需要再次等待用户确认。

## 任务列表

- [x] 新增 `.agents/skills/spec-workflow/` 目录
- [x] 编写 `.agents/skills/spec-workflow/SKILL.md`
- [x] 在 skill frontmatter 中定义清晰、主动的触发 description
- [x] 在 skill 正文中定义工作流、文档模板、确认节点、风险分级和执行记录规则
- [x] 检查 skill 与 `AGENTS.md`、`.agents/skills/go-srv-kit/` 的边界是否清晰
- [x] 读取新增文件，确认格式和内容无明显问题

## 验收标准

- 存在 `.agents/skills/spec-workflow/SKILL.md`
- skill 明确要求修改类任务先创建 `docs/<task-name>/spec.md`
- skill 明确要求创建规格后等待用户确认
- skill 明确区分 `AGENTS.md`、`go-srv-kit` skill 和 `spec-workflow` skill 的职责
- skill 包含可直接复用的 `spec.md` 模板
- skill 包含低风险和高风险操作说明
- skill 包含执行中方案变更的记录和确认规则

## 风险与回滚

风险较低。本次只新增 repo-local skill，不修改业务代码或生成代码。

如果新增 skill 不符合预期，可以删除 `.agents/skills/spec-workflow/` 目录，或继续迭代 `SKILL.md` 内容。

## 执行记录

- 已创建本任务规格文档。
- 用户已确认任务规格。
- 已新增 `.agents/skills/spec-workflow/SKILL.md`，初版内容集中在单个 skill 文件中，未拆分 `references/`。
- 已在 skill 中定义触发场景、工作流、规格模板、风险分级、执行记录规则，以及与 `AGENTS.md` 和 `.agents/skills/go-srv-kit/` 的边界。
- 已读取检查 `.agents/skills/spec-workflow/SKILL.md` 和本任务文档，格式和内容无明显问题。
- 本次任务只新增文档和 skill，不涉及 Go 代码、生成代码或测试命令。
