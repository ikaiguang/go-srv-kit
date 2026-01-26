# Git 子模块使用

## 添加子模块

添加子模块并指定分支：

```bash
git submodule add -b tag_v0.1.2 git@gitlab.uufff.cn:ikaiguang/go-srv-kit.git pkg/go-srv-kit
```

参数说明：
- `-b tag_v0.1.2`：指定子模块的分支
- `git@...`：子模块的 Git 仓库地址
- `pkg/go-srv-kit`：子模块在本地的目录

## 初始化子模块

克隆包含子模块的仓库时，子模块目录默认是空的。需要初始化子模块：

```bash
# 初始化并克隆子模块
git submodule update --init --recursive

# 或一次性初始化所有子模块
git submodule update --init --recursive --remote
```

## 更新子模块

### 更新子模块到最新代码

```bash
# 更新所有子模块到最新提交
git submodule update --remote --init

# 强制更新
git submodule update --remote --init -f
```

### 切换子模块分支

```bash
# 切换子模块到指定分支
git submodule set-branch -b dev pkg/go-srv-kit

# 更新子模块代码
git submodule update --remote --init
```

### 更新特定子模块

```bash
# 更新指定子模块
git submodule update --remote pkg/go-srv-kit
```

## 删除子模块

删除子模块的完整步骤：

```bash
# 1. 逆初始化模块（清除模块目录）
git submodule deinit pkg/go-srv-kit

# 2. 删除 .gitmodules 文件中的相关条目
vi .gitmodules

# 3. 删除 .git/config 中的子模块配置
# 可选：使用以下命令自动删除
# git config --unset --local submodule.pkg/go-srv-kit.active
# git config --unset --local submodule.pkg/go-srv-kit.url
vi .git/config

# 4. 删除 Git 中记录的模块信息
git rm --cached pkg/go-srv-kit

# 5. 删除实际的子模块目录
rm -rf pkg/go-srv-kit

# 6. 删除 .git/modules 中的缓存
rm -rf .git/modules/pkg/go-srv-kit
```

## 查看子模块状态

```bash
# 查看子模块状态
git submodule status

# 查看子模块信息
git submodule foreach 'echo $path $(git rev-parse --short HEAD)'
```

## 常用操作

### 在子模块中工作

```bash
# 进入子模块目录
cd pkg/go-srv-kit

# 正常使用 Git 命令
git status
git add .
git commit -m "message"
```

### 提交子模块的变更

当子模块有新的提交时，主仓库会显示子模块的变更：

```bash
# 在主仓库中
git status
# 显示：modified:   pkg/go-srv-kit (new commits)

# 提交子模块的更新
git add pkg/go-srv-kit
git commit -m "update submodule to latest version"
```

## 克隆包含子模块的仓库

### 递归克隆（包含子模块）

```bash
# 克隆时同时初始化和更新子模块
git clone --recursive https://github.com/username/repo.git

# 或分两步
git clone https://github.com/username/repo.git
cd repo
git submodule update --init --recursive
```

## .gitmodules 文件

.gitmodules 文件记录了子模块的配置：

```ini
[submodule "pkg/go-srv-kit"]
	path = pkg/go-srv-kit
	url = git@gitlab.uufff.cn:ikaiguang/go-srv-kit.git
	branch = tag_v0.1.2
```

字段说明：
- `path`：子模块在主仓库中的路径
- `url`：子模块的仓库地址
- `branch`：子模块跟踪的分支（可选）

## 更新 go.mod

当子模块更新后，需要更新 go.mod：

```bash
# 整理依赖
go mod tidy

# 下载依赖
go mod download
```

## 最佳实践

1. **明确指定子模块分支**：使用 `-b` 参数指定子模块的分支
2. **定期更新子模块**：使用 `git submodule update --remote` 更新子模块
3. **提交子模块变更**：当子模块有更新时，及时提交主仓库的变更
4. **使用绝对路径**：子模块 URL 建议使用绝对路径而非相对路径
5. **文档记录**：在 README 中记录子模块的用途和更新方法

## 常见问题

### 子模块目录为空

```bash
# 初始化子模块
git submodule update --init --recursive
```

### 子模块处于游离 HEAD 状态

```bash
# 切换到指定分支
cd pkg/go-srv-kit
git checkout $(git rev-parse --abbrev-ref HEAD)
```

### 子模块更新失败

```bash
# 清理子模块缓存
git submodule deinit --all
git submodule update --init --recursive
```
