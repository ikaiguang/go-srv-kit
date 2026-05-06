---
name: my-project
description: 本仓库专用工作指引。用于本项目 Go、Proto、Wire、配置、测试、实现、重构、评审和问题排查。
---

# My Project

## Overview

修改本仓库代码时使用本 skill。先遵循根目录 `AGENTS.md`，再按任务读取 `references/` 中的细节。

## Quick Routing

- 仓库结构、分层、示例入口：读 `references/project-context.md`
- Proto、Service、Biz、Data、配置接线、Wire：读 `references/service-workflow.md`
- 函数形态、命名、错误、日志、安全和测试习惯：读 `references/coding-rules.md`
- Makefile、Proto/Wire 生成、Windows 命令差异：读 `references/commands-and-generation.md`

## Working Rules

- 先读目标模块和相邻实现，再决定改法。
- 优先做最小改动，保持 `Service -> Biz -> Data` 分层。
- 业务服务接入优先参考 `testdata/ping-service/`。
- 本 skill 只负责路由到仓库细节；不要用它代替源码阅读。
