# Codex 会话续接指南

## 目的

本文档说明下次进入 `go-srv-kit` 仓库时，如何继续使用 Codex 的会话上下文，以及仓库内哪些文件会持续提供项目级上下文。

## 推荐进入方式

优先在仓库根目录执行：

```bash
cd D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit
codex resume --last
```

这会直接续上最近一次交互式会话。

如果想在续接时补一句新的任务说明，可以写成：

```bash
cd D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit
codex resume --last "继续上次在 go-srv-kit 的初始化工作"
```

## 常用命令

### 1. 继续最近一次会话

```bash
codex resume --last
```

### 2. 手动选择历史会话

```bash
codex resume
```

默认会按当前工作目录过滤历史会话。

### 3. 查看所有历史会话

```bash
codex resume --all
```

如果当前目录不对，或者想从别的目录找历史会话，用这个更稳。

### 4. 基于上一次会话开分支

```bash
codex fork --last
```

适合你想保留原会话，同时从当前进度分叉出一个新方向时使用。

### 5. 不在仓库目录时指定工作目录

```bash
codex resume --last -C D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit
```

## 项目级上下文如何续接

除了对话历史，当前仓库还提供两层持久上下文：

### 1. 常驻规则

- `AGENTS.md`

这里放的是始终生效的项目规则，例如语言、编码规范、分层边界、日志和错误处理约定。

### 2. Repo-local Codex skill

- `.agents/skills/go-srv-kit/SKILL.md`
- `.agents/skills/go-srv-kit/references/`

这里放的是按任务读取的深层上下文，例如：

- 项目结构
- 服务开发流程
- 命令与生成方式
- 编码规则

这意味着即使不开 `resume`，只要你在这个仓库里新开 Codex，会话仍然能读取到项目级规则和技能上下文；只是拿不到上一段对话本身。

## 建议用法

如果你要继续当前这类长期工作，优先使用：

```bash
cd D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit
codex resume --last
```

如果你只是新开一个任务，但仍在本仓库工作，直接：

```bash
cd D:\kaygrand\golang\src\github.com\ikaiguang\go-srv-kit
codex
```

这样会保留仓库规则，但不会强依赖上一轮对话历史。
