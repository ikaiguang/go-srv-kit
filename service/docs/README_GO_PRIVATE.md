# service pkg

私有包

## 添加私有包

添加环境变量

```shell

# golang 环境变量
# windows系统，请在系统环境变量中添加
# 也可使用 go env -w 写入；例如 go env -w GO111MODULE=on
# 如果在环境变量中添加过了了。go env -w 会包系统变量冲突错误
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOPRIVATE=gitlab.uufff.com
export GOPATH="/Users/$USER/golang" # windows请在环境变量Path中添加目录$GOPATH/bin
export GOPATH_BIN_PATH="$GOPATH/bin" # windows请在环境变量Path中添加目录$GOPATH/bin
export PATH="$PATH:$GOPATH_BIN_PATH" # windows请在环境变量Path中添加目录$GOPATH/bin

```

## 将下载代码方式由https改为ssh

为了避免每次拉取代码需要输入gitlab的账户密码，执行修改为git-ssh方式链接gitlab

```shell
# 配置git-ssh
git config --global url."ssh://git@gitlab.uufff.com/".insteadOf "https://gitlab.uufff.com/"
# 使用git的netrc保存登录信息
touch $HOME/.netrc
echo 'machine gitlab.uufff.com login 你的用户名 password 你的TOKEN或口令' >> $HOME/.netrc
```

## dockerfile ssh

必须条件

* [x] ~/.netrc : 密码 or oauth认证
* [x] ssh rsa证书
* [x] ~/.ssh/known_hosts : 主机信任

```dockerfile
RUN git config --global url."ssh://git@gitlab.uufff.com/".insteadOf "https://gitlab.uufff.com/" && \
    mkdir -p ~/.ssh && \
    cp ./devops/ssh/uufff_id_rsa* ~/.ssh/ && \
    chmod 600 ~/.ssh/uufff_id_rsa* && \
    mv ~/.ssh/uufff_id_rsa ~/.ssh/id_rsa && \
    mv ~/.ssh/uufff_id_rsa.pub ~/.ssh/id_rsa.pub && \
    touch ~/.netrc && \
    echo "machine gitlab.uufff.com login chenkaiguang@uufff.com password sRqVRR54wXzgLTPXxVvb" > ~/.netrc && \
    touch ~/.ssh/known_hosts && \
    chmod 600  ~/.ssh/known_hosts && \
    ssh-keyscan gitlab.uufff.com >> ~/.ssh/known_hosts
```
