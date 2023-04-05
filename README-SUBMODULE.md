# 添加子项目

查看帮助命令

```shell
git submodule help
```

## 添加子项目

```shell
git submodule add git@github.com:ikaiguang/go-srv-services.git
```

添加成功表现：

1. .gitmodules 文件
2. 克隆的子项目

## 克隆项目 携带submodule

```shell
git clone git@github.com:ikaiguang/go-srv-kit.git --recurse-submodules
```

## 更新 submodule

```shell
# 拉取代码
git submodule update
# 切换目录 & 拉取代码
cd go-srv-services
git pull origin master
# 或 子项目较多的时候执行以下命令
git submodule foreach 'git pull origin master'
```

## 删除 submodule

```shell
git submodule deinit go-srv-services
```
