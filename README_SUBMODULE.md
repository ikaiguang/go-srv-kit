# Git 子模块使用指南

本文档说明如何在其他项目中以 Git 子模块方式引入仓库，或维护已有子模块。

## 查看状态

```bash
git submodule status
git submodule foreach git status --short
```

## 添加子模块

```bash
git submodule add -b <branch-or-tag> git@gitlab.uufff.cn:ikaiguang/go-srv-kit.git <path>
```

示例：

```bash
git submodule add -b tag_v0.1.2 git@gitlab.uufff.cn:ikaiguang/go-srv-kit.git pkg/go-srv-kit
```

## 初始化与更新

```bash
git submodule update --init --recursive
git submodule update --remote --init --recursive
```

如果需要切换子模块跟踪的分支：

```bash
git submodule set-branch -b <branch> <path>
git submodule update --remote --init --recursive
```

## 删除子模块

推荐顺序：

```bash
git submodule deinit -f <path>
git rm -f <path>
```

如果 `.git/modules/<path>` 仍残留，再手动清理该目录。

## Go 模块注意事项

- 如果你在**子模块项目**里维护依赖，请在对应模块目录中执行 `go mod tidy`
- 不要把“在当前仓库根目录执行 `go mod tidy`”当成默认操作
  - 本仓库根模块有 `testdata/` 依赖保留问题，详见 `README.md` 和 `docs/migration-guide.md`
