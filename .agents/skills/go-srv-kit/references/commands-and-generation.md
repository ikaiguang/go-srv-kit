# 命令与生成

## 核心命令

除非任务明确要求进入子模块目录，否则默认在仓库根目录执行：

```bash
make init
make generate
go test ./...
go vet ./...
make run-service
go run ./testdata/ping-service/cmd/ping-service/... -conf=./testdata/ping-service/configs
```

## Proto 生成

能使用根目录 `Makefile` 时，优先使用 `Makefile` 目标，不要临时手写一串 `protoc` 命令。

常见示例：

```bash
make protoc-api-protobuf
make protoc-config-protobuf
make protoc-ping-v1-protobuf
```

如果修改的是某个特定 API 区域，先检查 `api/` 和 `testdata/ping-service/` 下相关的 makefile。

## Wire 生成

默认示例服务的 Wire 生成命令是：

```bash
wire ./testdata/ping-service/cmd/ping-service/export
```

当前 `make generate` 实际也是包装这个命令。

## 根模块重要注意事项

不要在仓库根目录随手执行 `go mod tidy`。

根目录 `go.mod` 已明确说明：`testdata/` 不会被 Go 的 `./...` 包模式纳入，所以在根目录执行 `go mod tidy` 可能删除示例服务依赖。

如果任务确实需要整理依赖，先确认是否会影响 `testdata/ping-service`。

## Codex Skill 校验

校验 repo-local skill 可使用：

```bash
python C:\Users\kaygrand\.codex\skills\.system\skill-creator\scripts\quick_validate.py .agents\skills\go-srv-kit
```

如果报错 `ModuleNotFoundError: No module named 'yaml'`，说明本机 Python 缺少 `PyYAML`，可执行：

```bash
python -m pip install PyYAML
```

当前机器已经通过该方式补齐依赖。

## Windows 注意事项

- 某些 `make` 目标在 `cmd` 或 `git-bash` 下比 PowerShell 更稳定
- 如果 `make` 目标在 PowerShell 中表现异常，退回到等价的 `go`、`wire` 或 `protoc` 命令

## 校验习惯

- 使用能验证当前改动的最小命令
- 改动涉及生成代码时，运行对应生成器
- 改动只影响局部包时，优先跑定向 `go test`，再决定是否扩大验证范围
