# v3 子 Module 发布指南

本指南使用 `devops/module-release/module-release.sh` 管理多 module 仓库的独立 tag。脚本不提交代码；代码提交和分支 push 由发布人完成。

## 首次发布 `kit/v3.0.0`

先在 `kit` module 运行发布前检查：

```bash
cd kit
GOWORK=off go mod tidy
GOWORK=off go test ./...
GOWORK=off go test -race ./...
GOWORK=off go vet ./...
cd ..
git diff --check
```

确认命令全部通过并审查改动，然后提交本次修改：

```bash
git status --short
git add <本次确认需要提交的文件>
git commit -m "release: prepare kit v3.0.0"
git push origin v3
```

同步远端 tag，并查看可发布 module：

```bash
git fetch origin --tags
./devops/module-release/module-release.sh list
```

检查本次发布参数：

```bash
./devops/module-release/module-release.sh check kit v3.0.0
```

输出应包含：

```text
module:  kit
path:    github.com/ikaiguang/go-srv-kit/kit/v3
version: v3.0.0
tag:     kit/v3.0.0
```

创建 annotated tag：

```bash
./devops/module-release/module-release.sh tag kit v3.0.0
git show --no-patch kit/v3.0.0
```

确认 tag 指向正确 commit 后推送：

```bash
./devops/module-release/module-release.sh push kit v3.0.0
```

验证远端 module：

```bash
GOPROXY=direct go list -m github.com/ikaiguang/go-srv-kit/kit/v3@v3.0.0
```

## 后续独立发布

每个 module 使用自己的版本。例如只修复 `kit` 时发布：

```bash
./devops/module-release/module-release.sh check kit v3.0.1
./devops/module-release/module-release.sh tag kit v3.0.1
./devops/module-release/module-release.sh push kit v3.0.1
```

这不会改变 `data/consul`、`kratos` 或其他 module 的版本。发布其他 module 时替换 module 名称：

```bash
./devops/module-release/module-release.sh check data/consul v3.0.0
./devops/module-release/module-release.sh tag data/consul v3.0.0
./devops/module-release/module-release.sh push data/consul v3.0.0
```

## 增加 Module

module 准备完成后，在 `devops/module-release/modules.tsv` 增加一行：

```text
名称<TAB>物理目录<TAB>go.mod 中的 module path<TAB>tag 前缀
```

例如：

```text
kit<TAB>kit<TAB>github.com/ikaiguang/go-srv-kit/kit/v3<TAB>kit
```

物理目录必须与 Go 根据 module path 推导出的仓库子目录一致。当前 `authpkg/` 目录与 `github.com/ikaiguang/go-srv-kit/auth/v3` 的 `auth/` 子目录不一致，因此处理 `auth` module 时应先统一目录，再加入发布清单。

## 常用命令

查看已有版本：

```bash
git tag --list 'kit/v*' --sort=-version:refname
```

查看某个 tag 指向：

```bash
git show --no-patch kit/v3.0.0
```

公开 tag 发布错误后，不要移动同名 tag；修复代码并发布新的 patch 版本。
