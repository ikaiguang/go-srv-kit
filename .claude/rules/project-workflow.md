# 项目开发工作流

## 新服务开发流程

### 1. 创建服务目录结构

```
my-service/
├── api/
│   └── my-service/
│       └── v1/
│           ├── services/
│           ├── resources/
│           ├── errors/
│           └── enums/
├── cmd/
│   └── my-service/
│       ├── main.go
│       └── export/
│           ├── main.export.go
│           └── wire.go
├── configs/
│   └── config.yaml
├── internal/
│   ├── conf/
│   │   ├── config.conf.proto
│   │   └── main.go
│   ├── service/
│   │   ├── service/
│   │   └── dto/
│   ├── biz/
│   │   ├── biz/
│   │   ├── bo/
│   │   └── repo/
│   └── data/
│       ├── data/
│       ├── po/
│       └── repo/
└── go.mod
```

### 2. 定义 API Proto

```bash
# 创建 proto 文件
api/my-service/v1/services/my-service.proto

# 生成代码
make api-my-service
```

### 3. 实现各层代码

按照 API 开发流程实现各层

### 4. 配置服务

```yaml
# configs/config.yaml
server:
  http:
    addr: 0.0.0.0:10101
  grpc:
    addr: 0.0.0.0:10102

# 其他配置...
```

### 5. Wire 依赖注入

```go
// cmd/my-service/export/wire.go
//go:build wireinject
package exporter

func exportServices(launcher setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) {
    panic(wire.Build(
        setuputil.GetLogger,
        data.NewMyData,
        biz.NewMyBiz,
        service.NewMyService,
        service.RegisterServices,
    ))
}
```

### 6. 生成 Wire 代码

```bash
wire ./cmd/my-service/export
```

### 7. 运行服务

```bash
go run ./cmd/my-service/... -conf=./configs
```

## 日常开发流程

### 1. 拉取最新代码

```bash
git pull origin prod
```

### 2. 创建功能分支

```bash
git checkout prod
git checkout -b feature/new-feature
```

### 3. 开发功能

- 编写代码
- 运行 `make generate` 生成代码
- 运行测试

```bash
# 生成 Wire 代码
make generate

# 运行测试
go test ./...

# 代码格式化
gofmt -w .
goimports -w .
```

### 4. 提交代码

```bash
git add .
git commit -m "feat: 添加新功能"
```

### 5. 合并到 test（自测后）

```bash
git checkout test
git merge feature/new-feature
```

### 6. 合并到 pre（测试通过后）

```bash
git checkout pre
git merge feature/new-feature
```

### 7. 合并到 prod（预发布验证后）

```bash
git checkout prod
git merge feature/new-feature
```

### 8. 推送到远程

```bash
git push origin prod
```

## 代码检查

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包测试
go test ./internal/service/...

# 带覆盖率
go test -coverprofile=coverage.out ./...

# 查看覆盖率报告
go tool cover -html=coverage.out
```

### 代码格式化

```bash
# 格式化代码
gofmt -w .

# 或使用 goimports
goimports -w .
```

### 静态检查

```bash
# go vet
go vet ./...

# golangci-lint
golangci-lint run
```

## 调试技巧

### 1. 查看日志

```bash
# 实时查看日志
tail -f ./runtime/logs/ping-service_app_*.log

# 查看错误日志
tail -f ./runtime/logs/ping-service_error_*.log
```

### 2. 使用 Delve 调试

```bash
# 安装 Delve
go install github.com/go-delve/delve/cmd/dlv@latest

# 调试
dlv debug ./cmd/ping-service/... -- -conf=./configs
```

### 3. HTTP 调试

```bash
# curl
curl http://127.0.0.1:10101/api/v1/ping/say_hello

# 或使用测试脚本
make testing-service
```

## 常见问题

### Wire 生成失败

```bash
# 清理后重新生成
rm ./cmd/*/export/wire_gen.go
wire ./cmd/*/export
```

### Proto 生成失败

```bash
# 检查 protoc 是否安装
protoc --version

# 检查插件是否安装
ls $GOPATH/bin/protoc-gen-*

# 重新安装
make init
```

### 端口被占用

```bash
# Windows
netstat -ano | findstr :10101
taskkill /PID <pid> /F

# Linux/Mac
lsof -i :10101
kill -9 <pid>
```

### Go mod 依赖问题

```bash
# 清理依赖缓存
go clean -modcache

# 重新下载依赖
go mod download

# 整理依赖
go mod tidy
```

## 发布流程

### 1. 更新版本

```go
// cmd/service/main.go
var Version = "1.0.0"
```

### 2. 打 Tag

```bash
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0
```

### 3. 构建

```bash
make build
```

### 4. 部署

```bash
make deploy-on-docker
```

## 紧急修复流程

### 1. 创建 hotfix 分支

```bash
git checkout prod
git checkout -b hotfix/critical-bug
```

### 2. 修复并测试

```bash
# 修复代码
# 运行测试
go test ./...
```

### 3. 提交修复

```bash
git add .
git commit -m "fix: 修复关键问题"
```

### 4. 快速合并

```bash
# 合并到 test
git checkout test
git merge hotfix/critical-bug

# 合并到 pre
git checkout pre
git merge hotfix/critical-bug

# 合并到 prod
git checkout prod
git merge hotfix/critical-bug

# 推送到远程
git push origin prod
```

### 5. 删除 hotfix 分支

```bash
git branch -d hotfix/critical-bug
```
