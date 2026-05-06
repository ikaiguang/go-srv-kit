---
inclusion: manual
---

# Go 私有包配置

## 环境变量

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=gitlab.uufff.com
```

Windows 系统在环境变量中添加 `GOPATH` 和 `%GOPATH%\bin` 到 Path。

## Git SSH 配置

```bash
git config --global url."ssh://git@gitlab.uufff.com/".insteadOf "https://gitlab.uufff.com/"
```

## 常见问题

- 私有包下载失败：检查 `go env GOPRIVATE`，清理 `go clean -modcache`
- Git 认证失败：检查 SSH 密钥 `ssh -T git@gitlab.uufff.com`
- Docker 构建：需配置 SSH 密钥 + .netrc + known_hosts + Git URL 替换
