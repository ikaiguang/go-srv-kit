# Module 发布脚本使用指南

本目录提供功能一致的 Bash 和 Windows CMD 发布脚本：

- Linux、macOS 或 Git Bash：`module-release.sh`
- Windows CMD：`module-release.bat`
- Module 清单：`modules.tsv`

脚本用于检查、创建和推送各 Go module 的独立 Git tag，不会运行 Go 测试、提交代码或推送分支。

## 使用前提

1. 安装 Go 和 Git，并确保 `go`、`git` 已加入 `PATH`；Windows CMD 脚本还会使用系统自带的 Windows PowerShell 校验版本号。
2. 在目标 module 内完成 `GOWORK=off` 的测试、静态检查和依赖校验。
3. 提交并推送待发布代码，确保工作区干净。
4. 执行 `git fetch origin --tags`，同步远端 tag 后再检查新版本。

## 命令

以下示例从仓库根目录执行。Windows CMD 使用：

```bat
devops\module-release\module-release.bat list
devops\module-release\module-release.bat check kit v3.0.0
devops\module-release\module-release.bat tag kit v3.0.0
devops\module-release\module-release.bat push kit v3.0.0
```

Bash 使用：

```bash
./devops/module-release/module-release.sh list
./devops/module-release/module-release.sh check kit v3.0.0
./devops/module-release/module-release.sh tag kit v3.0.0
./devops/module-release/module-release.sh push kit v3.0.0
```

各命令作用如下：

| 命令 | 作用 |
| --- | --- |
| `list` | 列出 `modules.tsv` 中的 module、module path 和本地最新 tag |
| `check <module> <version>` | 校验 module、`go.mod`、工作区状态、版本格式和本地 tag，并显示将要发布的 commit |
| `tag <module> <version>` | 先执行 `check`，再为当前 `HEAD` 创建 annotated tag |
| `push <module> <version>` | 将已经存在的本地 tag 推送到远端 |

版本必须是 `v3` 语义版本，例如 `v3.0.0`、`v3.1.0-rc.1` 或 `v3.1.0+build.1`。根 module 的 tag 是 `v3.0.0`，子 module 的 tag 带目录前缀，例如 `kit/v3.0.0`、`data/consul/v3.0.0`。

## 推荐发布流程

以发布 `kit/v3.0.0` 为例。

先在 `kit` module 中完成发布前验证。Bash：

```bash
cd kit
GOWORK=off go mod tidy
GOWORK=off go test -count=1 ./...
GOWORK=off go test -race -count=1 ./...
GOWORK=off go vet ./...
GOWORK=off go mod verify
cd ..
git diff --check
```

Windows CMD：

```bat
cd kit
set "GOWORK=off"
go mod tidy
go test -count=1 ./...
go test -race -count=1 ./...
go vet ./...
go mod verify
cd ..
git diff --check
```

审查 `go mod tidy` 等命令产生的改动，提交并推送发布代码，然后执行：

```bat
git fetch origin --tags
devops\module-release\module-release.bat list
devops\module-release\module-release.bat check kit v3.0.0
devops\module-release\module-release.bat tag kit v3.0.0
git show --no-patch kit/v3.0.0
devops\module-release\module-release.bat push kit v3.0.0
```

使用 Bash 时将上述 `.bat` 命令替换为对应的 `module-release.sh` 命令。推送前必须通过 `git show --no-patch <tag>` 确认 tag 指向预期 commit。

发布后可验证 Go module：

```bat
go list -m github.com/ikaiguang/go-srv-kit/kit/v3@v3.0.0
```

## 指定远端

默认推送到 `origin`。可以通过 `REMOTE` 环境变量指定其他远端。

Windows CMD：

```bat
set "REMOTE=upstream"
devops\module-release\module-release.bat push kit v3.0.0
```

Bash：

```bash
REMOTE=upstream ./devops/module-release/module-release.sh push kit v3.0.0
```

## 维护 Module 清单

新增 module 时，在 `modules.tsv` 中添加一行，字段之间使用 Tab 分隔：

```text
名称<TAB>物理目录<TAB>go.mod 中的 module path<TAB>tag 前缀
```

例如：

```text
kit<TAB>kit<TAB>github.com/ikaiguang/go-srv-kit/kit/v3<TAB>kit
```

根 module 的 tag 前缀使用 `-`。物理目录、`go.mod` 中声明的 module path 和仓库中的实际布局必须一致。

## 异常处理

- `check` 提示工作区不干净：审查并提交本次发布改动，不要用脚本绕过检查。
- 提示 tag 已在本地存在：确认是否已经创建过该版本；已公开的 tag 不应移动或覆盖，应修复后发布新的 patch 版本。
- `push` 提示本地 tag 不存在：先运行 `tag`，并用 `git show --no-patch <tag>` 检查其 commit。
- tag 尚未推送且确认创建错误时，可使用 `git tag -d <tag>` 删除本地 tag 后重新创建。

更完整的 v3 发布背景和 `authpkg/` 的目录限制见 [`../../docs/v3_development/module_release.md`](../../docs/v3_development/module_release.md)。
