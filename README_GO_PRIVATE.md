# Go 私有仓库配置

本文档说明如何让本机或 CI 正常拉取 `gitlab.uufff.com` 上的 Go 私有依赖。

## 推荐做法

优先使用 `go env -w` 写入 Go 配置：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=gitlab.uufff.com
```

如果你依赖 `$GOPATH/bin` 中安装的工具，请确认它已经在系统 `PATH` 中。

## 使用 SSH 拉取 Git 仓库

如果不想每次输入用户名和密码，可将 GitLab HTTPS 地址改写为 SSH：

```bash
git config --global url."ssh://git@gitlab.uufff.com/".insteadOf "https://gitlab.uufff.com/"
```

然后准备好：

- 可用的 SSH 私钥
- `~/.ssh/known_hosts`
- 必要时的 `~/.netrc` 或 GitLab Token

## `.netrc` 示例

请只使用占位符，不要把真实账号或口令写入仓库文档：

```text
machine gitlab.uufff.com
login <your-username>
password <your-token-or-password>
```

## Docker / CI 注意事项

- 不要把真实 SSH 私钥、用户名、密码、Token 直接写进 Dockerfile 或脚本
- 优先使用：
  - CI Secret / Variable
  - 挂载的只读密钥文件
  - 构建机预置的 `.netrc` / SSH 配置
- 如果需要在镜像构建时拉私有依赖，请确保构建环境提前注入认证信息，而不是把认证信息硬编码到仓库中

## 常见问题

### `unrecognized import path` 或 `repository not found`

先确认：

1. `GOPRIVATE` 是否包含 `gitlab.uufff.com`
2. Git 对该域名是否已配置 SSH 改写
3. SSH 密钥、Token 或 `.netrc` 是否可用

### 工具命令找不到

确认 `$GOPATH/bin` 已加入系统 `PATH`。
