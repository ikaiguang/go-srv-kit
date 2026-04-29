# Update Codex CLI Config Guide

## 背景

用户希望在 `docs/codex-cli-usage-guide.md` 中补充 Codex CLI 配置固化方式，重点解决“低风险命令不要每次确认”的日常使用问题。

当前文档已经说明 `codex --full-auto -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"` 的启动方式，但对 `~/.codex/config.toml` 如何配置、为什么推荐 `workspace-write + on-request`、何时不应关闭沙箱等内容说明不足。

## 目标

- 补充 `~/.codex/config.toml` 中固化低确认模式的推荐配置。
- 说明推荐配置的原因：减少低风险命令确认，同时保留高风险操作确认能力。
- 补充 PowerShell 函数方式，用短命令固化 `-C` 仓库路径。
- 补充 `--add-dir` / `writable_roots` 的使用场景，说明它们比直接关闭沙箱更适合日常开发。
- 明确不推荐日常使用 `danger-full-access`、`approval_policy = "never"` 和 `--dangerously-bypass-approvals-and-sandbox` 的原因。

## 非目标

- 不修改真实的 `C:\Users\kaygrand\.codex\config.toml`。
- 不修改 PowerShell Profile。
- 不调整 `AGENTS.md`、skill 或 Codex CLI 安装。
- 不引入新的仓库工作流规则，只补充 Codex CLI 使用文档。

## 影响范围

- 预计修改：`docs/codex-cli-usage-guide.md`
- 本任务文档：`docs/update-codex-cli-config-guide/spec.md`

## 方案

在 `docs/codex-cli-usage-guide.md` 中增加一个“固化配置推荐”相关章节，建议内容包括：

1. 推荐在 `~/.codex/config.toml` 中设置：

   ```toml
   approval_policy = "on-request"
   sandbox_mode = "workspace-write"

   [projects.'D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit']
   trust_level = "trusted"
   ```

2. 说明该配置适合用户目标：
   - 低风险本地命令少确认。
   - 命令仍在 workspace 沙箱内执行。
   - 遇到需要越权或高风险操作时仍可确认。

3. 补充 `-C` 不能完全靠 `config.toml` 替代时的 PowerShell 函数方案：

   ```powershell
   function gsk {
     codex -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit" @args
   }
   ```

4. 补充扩展写目录的优先方案：
   - 临时使用 `--add-dir`
   - 长期使用 `[sandbox_workspace_write].writable_roots`

5. 将“关闭沙箱/关闭审批”作为明确不推荐的日常模式，只说明风险和适用边界，不作为默认建议。

## 任务列表

- [x] 阅读并确认 `docs/codex-cli-usage-guide.md` 当前结构。
- [x] 补充配置固化章节和推荐配置原因。
- [x] 补充 PowerShell 短命令方案。
- [x] 补充扩展写目录与关闭沙箱的取舍。
- [x] 检查文档表述是否符合本仓库低风险/高风险协作规则。

## 验收标准

- 文档中能直接回答：如何不用每次写 `--full-auto -C ...`。
- 文档中能直接回答：为什么推荐 `approval_policy = "on-request"` 和 `sandbox_mode = "workspace-write"`。
- 文档中明确说明：日常不建议为了少确认而关闭沙箱或完全关闭审批。
- 文档没有要求用户实际修改本机配置，只提供可选配置示例。

## 风险与回滚

- 风险：Codex CLI 配置项可能随版本变化。规避方式是在文档中说明以 `codex --help` 和官方文档为准。
- 风险：读者误以为 `danger-full-access` 是推荐日常配置。规避方式是把它放入不推荐章节，并明确风险。
- 回滚：如果内容不合适，可以删除本次新增段落或回退 `docs/codex-cli-usage-guide.md` 的对应修改。

## 执行记录

- 2026-04-30：已创建任务规格文档，等待用户确认后再修改 `docs/codex-cli-usage-guide.md`。
- 2026-04-30：用户已确认规格，开始修改 `docs/codex-cli-usage-guide.md`。
- 2026-04-30：已补充 `config.toml` 推荐配置、配置原因、PowerShell `gsk` 函数、`--add-dir` / `writable_roots` 说明，以及不推荐日常关闭沙箱和关闭审批的原因。
- 2026-04-30：已读取修改后的文档完成内容检查；`docs/codex-cli-usage-guide.md` 和本规格文件当前均为未跟踪文档文件。
