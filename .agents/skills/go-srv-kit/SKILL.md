---
name: go-srv-kit
description: go-srv-kit 仓库专用工作指引。该仓库是基于 go-kratos、DDD 分层和 Wire 的微服务工具包。当需要在本仓库中实现、重构、评审或排查 Go 代码，修改 Proto、service、biz、data、Wire 装配，或查找本仓库特有的命令、架构、日志和错误处理约定时使用。
---

# Go Srv Kit

## Overview

当任务是修改本仓库代码时使用这个 skill。
先遵循 `AGENTS.md` 中的常驻规则，再按任务需要读取对应的 `references/` 文件。

## Quick Routing

- 需要了解包职责、工作区布局、示例服务入口和关键文件时，读取 `references/project-context.md`
- 需要新增或修改 Proto、service handler、biz 逻辑、data repo、配置接线或 Wire 装配时，读取 `references/service-workflow.md`
- 需要确认函数形态、硬编码、常量、panic、日志和错误处理规范时，读取 `references/coding-rules.md`
- 需要查看 `make` 目标、`wire` 用法、生成步骤、运行命令和仓库特有命令注意事项时，读取 `references/commands-and-generation.md`

## 工作方式

- 先从要修改的具体包入手，对照相邻实现，再决定是否引入新模式
- 优先做最小改动，保持现有 `Service -> Biz -> Data` 分层和当前 Wire 装配风格
- 将 `testdata/ping-service/` 视为业务服务接入工具包的首选参考实现
- 不要拿这个 skill 代替源码阅读。它的作用是帮你选对仓库参考材料，然后回到目标代码本身

## 边界

- `AGENTS.md` 保持为简洁的常驻规则文件
- 这个 skill 只承载更深的、按任务需要才读取的仓库上下文
