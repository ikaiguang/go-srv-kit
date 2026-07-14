# 子模块

```shell
git submodule help
```

## 添加子模块

```shell
git submodule add -b tag_v0.1.2 git@gitlab.uufff.cn:ikaiguang/go-srv-kit.git pkg/go-srv-kit
```

## 删除子模块

```shell
# 逆初始化模块，其中{MOD_NAME}为模块目录，执行后可发现模块目录被清空
# rm -rf .git/modules/$path_to_submodule
git submodule deinit $path_to_submodule
# 删除项目目录下.gitmodules文件中子模块相关条目
vi .gitmodules
# 删除配置项中子模块相关条目
# git config --unset --local submodule.$path_to_submodule.active
# git config --unset --local submodule.$path_to_submodule.url
vi .git/config
# 删除.gitmodules中记录的模块信息（--cached选项清除.git/modules中的缓存）
git rm --cached $path_to_submodule
```

# 切换分支 & 更新submodule

```shell
# 更新分支
git submodule set-branch -b dev pkg/go-srv-kit
# 更新代码
git submodule update --remote --init
git submodule update --remote --init -f
```

## 修改gomod

```shell
go mod tidy
```