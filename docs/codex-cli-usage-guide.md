# Codex CLI 使用指导

本文记录本仓库推荐的 Codex CLI 使用方式，目标是减少低风险操作的反复确认，同时保留高风险操作的人工确认。

## 推荐日常启动方式

在 `go-srv-kit` 仓库中，推荐使用：

```powershell
codex --full-auto -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"
```

含义：

- `--full-auto`：让 Codex 在沙箱内自动执行常见低风险操作，减少读文件、搜索、格式化、测试等反复确认。
- `-C <path>`：指定当前仓库作为工作目录，避免 Codex 进入错误目录。

Windows PowerShell 中建议给路径加双引号。反斜杠 `\` 在 PowerShell 字符串里不是转义符，不需要写成 `\\`。

推荐：

```powershell
codex --full-auto -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"
```

不推荐：

```powershell
codex --full-auto -C D:\\kaygrand\\golang\\src\\github.com\\ikaiguang\\go-srv-kit
```

不带双引号在当前路径没有空格时也能工作，但统一加双引号更稳。

本仓库已经通过 `AGENTS.md`、`.agents/skills/spec-workflow/` 和 `.agents/skills/go-srv-kit/` 约束工作方式，因此日常更适合把确认点放在高风险操作上，而不是低风险本地操作上。

注意：`--full-auto` 只是 Codex CLI 的工具层配置。真正的仓库协作规则仍由 `AGENTS.md`、`.agents/skills/spec-workflow/`、`.agents/skills/go-srv-kit/` 和相关 docs 共同约束。

## 固化低确认配置

如果目标是“低风险本地命令不要每次确认，但高风险操作仍然保留确认能力”，推荐把 `--full-auto` 的核心行为固化到：

```text
C:\Users\kaygrand\.codex\config.toml
```

推荐配置：

```toml
approval_policy = "on-request"
sandbox_mode = "workspace-write"

[projects.'D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit']
trust_level = "trusted"
```

配置原因：

- `approval_policy = "on-request"`：低风险命令可以由 Codex 自动执行；遇到需要越权、高风险或风险不确定的操作时，Codex 仍可以请求用户确认。
- `sandbox_mode = "workspace-write"`：允许 Codex 读取文件，并修改当前 workspace 内文件；这覆盖了本仓库常见的读文件、搜索、改文档、改源码、`gofmt`、`go test`、`wire` 等低风险本地操作。
- `trust_level = "trusted"`：把当前仓库标记为可信项目，减少对仓库内常规操作的额外摩擦。

这组配置更接近日常推荐的 `--full-auto`，但不是关闭沙箱，也不是关闭所有审批。它适合本仓库的协作目标：低风险本地操作少确认，高风险操作仍然需要明确确认。

配置后，如果已经在仓库目录中，可以直接运行：

```powershell
codex
```

或者显式指定仓库目录：

```powershell
codex -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"
```

说明：`config.toml` 可以固化审批和沙箱策略，但不等同于永久固定 `-C` 工作目录。工作目录仍由当前 shell 所在目录或 `-C <path>` 决定。

## 固化仓库入口

如果不想每次输入 `-C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"`，推荐在 PowerShell Profile 中加一个短命令。

打开 PowerShell Profile：

```powershell
notepad $PROFILE
```

加入：

```powershell
function gsk {
  codex -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit" @args
}
```

保存后重新加载：

```powershell
. $PROFILE
```

以后直接运行：

```powershell
gsk
```

如果已经在 `config.toml` 中固化了 `approval_policy = "on-request"` 和 `sandbox_mode = "workspace-write"`，`gsk` 不需要再写 `--full-auto`。

不建议直接覆盖全局 `codex` 函数，除非明确只在这个仓库中使用 Codex。覆盖 `codex` 会影响其他目录和其他项目的启动行为。

## 扩展写目录

日常需要让 Codex 额外写入某个目录时，优先使用 `--add-dir`，不要直接关闭沙箱。

临时增加一个可写目录：

```powershell
codex -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit" --add-dir "D:\some\other\dir"
```

如果某些目录长期需要写权限，可以写入 `config.toml`：

```toml
[sandbox_workspace_write]
writable_roots = [
  'D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit',
  'D:\some\other\dir',
]
```

配置原因：

- `--add-dir` 和 `writable_roots` 只扩大明确目录的写权限，影响范围可控。
- `danger-full-access` 会让 Codex 失去 workspace 沙箱边界，不适合作为“只是少确认低风险命令”的日常方案。
- 如果任务只是本仓库开发，通常不需要额外目录写权限。

## 可选的保守模式

如果希望更保守，可以使用：

```powershell
codex -s workspace-write -a on-request -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"
```

含义：

- `-s workspace-write`：允许读取文件，并允许修改当前 workspace 内文件。
- `-a on-request`：由 Codex 判断何时请求确认。

这种模式确认会更多，适合不确定任务风险时使用。

## 不推荐作为日常的模式

不建议日常使用：

```powershell
codex -s workspace-write -a never -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"
```

原因：

- 它不会请求审批。
- 当命令因为沙箱或权限失败时，Codex 只能拿到失败结果，不能请求升级执行。
- 对本仓库常见的 `wire`、`go test`、`gofmt` 等命令，可能更容易卡住。

也不建议日常使用：

```powershell
codex --dangerously-bypass-approvals-and-sandbox -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"
```

原因：

- 该模式会绕过审批和沙箱。
- 除非在一次性临时容器或其他外部隔离环境中运行，否则风险过高。

也不建议为了少确认而在日常配置中写：

```toml
approval_policy = "never"
sandbox_mode = "danger-full-access"
```

原因：

- `approval_policy = "never"` 会关闭审批请求能力；当命令需要用户确认或需要越权时，只会失败返回。
- `sandbox_mode = "danger-full-access"` 会关闭 workspace 沙箱边界，影响范围超过当前仓库。
- 这组配置解决的是“完全不询问”，不是“低风险不反复询问”。它和本仓库保留高风险确认点的协作规则不一致。

如果确实在外部隔离环境中临时使用这类模式，应明确这是一次性场景，不要沉淀为本机默认配置。

## 本仓库的推荐协作规则

### 低风险操作

用户确认任务规格后，Codex 可以直接执行低风险操作，不需要逐项询问：

- 读取文件，例如 `Get-Content`
- 搜索代码，例如 `rg`
- 查看工作区状态，例如 `git status`
- 修改当前任务相关的普通源码或文档
- 运行 `gofmt`
- 运行本地测试，例如 `go test`
- 运行本地生成命令，例如 `wire ./testdata/ping-service/cmd/ping-service/export`
- 运行依赖检查，例如 `go list -deps`

如果这类命令因为沙箱异常失败，Codex 应先判断是否确实需要升级执行；不要把所有低风险命令一律升级为需要用户确认。

如果失败原因是 sandbox 初始化或权限限制，应明确说明这是工具环境问题，不等同于操作本身高风险。只有确实需要越权执行时，才请求用户确认。

### 高风险操作

以下操作必须单独确认：

- 删除大量文件
- 改写 Git 历史
- 丢弃未确认修改
- 安装或修改系统级依赖
- 发布或部署
- 请求第三方或线上 API 修改数据
- 修改真实环境数据库
- 处理未脱敏敏感信息
- 任何风险不确定的操作

确认时应说明具体动作、影响范围和回滚方式。

## Specs 工作流

本仓库推荐使用规格先行流程。

需要修改文件、推进实现或执行有副作用命令时，Codex 应先创建或更新：

```text
docs/<task-name>/spec.md
```

规格文档应包含：

- 背景
- 目标
- 非目标
- 影响范围
- 方案
- 任务列表
- 验收标准
- 风险与回滚
- 执行记录

用户确认规格后，低风险操作可以继续执行，不需要再次逐项询问。

## 协作判断参考

更完整的自动执行、主动询问和文档沉淀规则见：

```text
docs/agent-collaboration-guide.md
```

简要原则：

- 能自动执行的低风险本地操作，不要反复询问。
- 会影响系统、线上、Git 历史、真实数据、敏感信息或当前 workspace 之外的操作，必须询问。
- 用户纠正了代理行为或指出应补充文档时，要主动判断并更新 `AGENTS.md`、skill 或 docs。

## 常用命令

### 启动 Codex

```powershell
codex --full-auto -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"
```

### 继续上一次会话

```powershell
codex resume --last
```

### Fork 上一次会话

```powershell
codex fork --last
```

### 查看 Codex 帮助

```powershell
codex --help
```

## 推荐结论

本仓库日常开发优先使用以下配置：

```toml
approval_policy = "on-request"
sandbox_mode = "workspace-write"

[projects.'D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit']
trust_level = "trusted"
```

如果还没有固化配置，可以临时使用：

```powershell
codex --full-auto -C "D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit"
```

如果已经固化配置，并添加了 PowerShell 函数，可以直接使用：

```powershell
gsk
```

配合仓库内的 `AGENTS.md` 和 `spec-workflow` skill：

- 低风险本地操作直接执行
- 中风险操作写进规格后执行
- 高风险操作单独确认

Codex CLI 配置项可能随版本变化，最终以 `codex --help` 和 OpenAI 官方 Codex 文档为准：

- https://developers.openai.com/codex/cli/reference
- https://developers.openai.com/codex/config-reference
