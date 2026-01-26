# Go 私有包配置

## 环境变量配置

添加 Go 环境变量：

```bash
# 开启 Go Modules
export GO111MODULE=on

# 设置 Go 代理（使用国内镜像）
export GOPROXY=https://goproxy.cn,direct

# 设置私有包（不通过代理）
export GOPRIVATE=gitlab.uufff.com

# 设置 GOPATH
export GOPATH="/Users/$USER/golang"  # Windows 请在环境变量 Path 中添加 $GOPATH/bin

# 添加 GOPATH/bin 到 PATH
export PATH="$PATH:$GOPATH/bin"  # Windows 请在环境变量 Path 中添加目录
```

### Windows 系统配置

在系统环境变量中添加：

| 变量名 | 值 |
|--------|-----|
| `GO111MODULE` | `on` |
| `GOPROXY` | `https://goproxy.cn,direct` |
| `GOPRIVATE` | `gitlab.uufff.com` |
| `GOPATH` | `C:\Users\{用户名}\golang` |

在 `Path` 变量中添加：`%GOPATH%\bin`

### 使用 go env 命令

```bash
# 设置环境变量
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=gitlab.uufff.com

# 查看环境变量
go env
```

## Git SSH 配置

为了避免每次拉取代码需要输入 GitLab 的账户密码，配置 Git SSH 方式：

```bash
# 配置 Git 使用 SSH 代替 HTTPS
git config --global url."ssh://git@gitlab.uufff.com/".insteadOf "https://gitlab.uufff.com/"

# 使用 .netrc 保存登录信息
touch $HOME/.netrc
echo 'machine gitlab.uufff.com login 你的用户名 password 你的TOKEN或口令' >> $HOME/.netrc
```

### Windows 配置 .netrc

```bash
# 在用户目录下创建 _netrc 文件
# 位置：C:\Users\{用户名}\_netrc

machine gitlab.uufff.com login 你的用户名 password 你的TOKEN或口令
```

## Dockerfile SSH 配置

在 Dockerfile 中配置 SSH 访问私有仓库：

### 必要条件

- [x] `~/.netrc`：密码或 oauth 认证
- [x] SSH RSA 证书
- [x] `~/.ssh/known_hosts`：主机信任

### Dockerfile 示例

```dockerfile
# 配置 Git 使用 SSH
RUN git config --global url."ssh://git@gitlab.uufff.com/".insteadOf "https://gitlab.uufff.com/" && \
    mkdir -p ~/.ssh && \
    cp ./devops/ssh/uufff_id_rsa* ~/.ssh/ && \
    chmod 600 ~/.ssh/uufff_id_rsa* && \
    mv ~/.ssh/uufff_id_rsa ~/.ssh/id_rsa && \
    mv ~/.ssh/uufff_id_rsa.pub ~/.ssh/id_rsa.pub && \
    touch ~/.netrc && \
    echo "machine gitlab.uufff.com login chenkaiguang@uufff.com password sRqVRR54wXzgLTPXxVvb" > ~/.netrc && \
    touch ~/.ssh/known_hosts && \
    chmod 600 ~/.ssh/known_hosts && \
    ssh-keyscan gitlab.uufff.com >> ~/.ssh/known_hosts
```

### 使用 SSH 密钥挂载

```dockerfile
# 复制 SSH 密钥
COPY devops/ssh/id_rsa /root/.ssh/id_rsa
COPY devops/ssh/id_rsa.pub /root/.ssh/id_rsa.pub
RUN chmod 600 /root/.ssh/id_rsa

# 配置 known_hosts
RUN ssh-keyscan gitlab.uufff.com > /root/.ssh/known_hosts

# 配置 Git
RUN git config --global url."ssh://git@gitlab.uufff.com/".insteadOf "https://gitlab.uufff.com/"
```

### 使用 .netrc 挂载

```dockerfile
# 复制 .netrc
COPY devops/.netrc /root/.netrc
RUN chmod 600 /root/.netrc
```

## 常见问题

### 1. 私有包下载失败

```bash
# 检查环境变量
go env GOPRIVATE

# 清理缓存
go clean -modcache

# 重新下载
go mod download
```

### 2. Git 认证失败

```bash
# 检查 Git 配置
git config --global --list

# 检查 SSH 密钥
ssh -T git@gitlab.uufff.com

# 检查 .netrc 文件
cat ~/.netrc
```

### 3. Docker 构建时私有包问题

确保 Dockerfile 中包含：
- SSH 密钥配置
- .netrc 文件配置
- known_hosts 配置
- Git URL 替换配置
