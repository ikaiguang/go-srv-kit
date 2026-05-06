---
inclusion: manual
---

# Git 子模块管理

## 常用操作

```bash
# 添加子模块
git submodule add -b tag_v0.1.2 git@gitlab.uufff.cn:ikaiguang/go-srv-kit.git pkg/go-srv-kit

# 初始化
git submodule update --init --recursive

# 更新到最新
git submodule update --remote --init

# 切换分支
git submodule set-branch -b dev pkg/go-srv-kit
git submodule update --remote --init

# 递归克隆
git clone --recursive https://github.com/username/repo.git
```

## 删除子模块

```bash
git submodule deinit pkg/go-srv-kit
git rm --cached pkg/go-srv-kit
rm -rf pkg/go-srv-kit
rm -rf .git/modules/pkg/go-srv-kit
```

## 最佳实践

- 明确指定子模块分支（`-b` 参数）
- 定期更新子模块
- 子模块更新后运行 `go mod tidy`
