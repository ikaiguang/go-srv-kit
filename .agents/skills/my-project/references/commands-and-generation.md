# 命令与生成

## 核心命令

除非任务明确要求进入子模块目录，否则默认在仓库根目录执行：

```bash
make init
make generate
go test ./...
go vet ./...
make run-service
```

## Proto 生成

能使用根目录 `Makefile` 时，优先使用 `Makefile` 目标，不要临时手写一串 `protoc` 命令。

常见示例：

```bash
make protoc-api-protobuf
make protoc-ping-v1-protobuf
```

如果修改的是某个特定 API 区域，先检查 `api/` 和 `testdata/ping-service/` 下相关的 makefile。

## Wire 生成

默认示例服务的 Wire 生成命令是：

```bash
make generate
wire ./testdata/ping-service/cmd/ping-service/export
```

当前 `make generate` 实际也是包装这个命令。

## Windows 注意事项

- 某些 `make` 目标在 `git-bash` 下比 PowerShell 更稳定
- 如果 `make` 目标在 PowerShell 中表现异常，退回到等价的 `go`、`wire` 或 `protoc` 命令

## 脚本和外部命令

- 运行脚本、生成器或外部工具前，先用轻量命令确认它能执行目标动作，而不是只确认命令存在
- 如果默认命令不可用，按当前系统尝试等价命令、版本管理器解析出的命令，或项目文档指定的命令
- 文档和规则不要写死个人机器的绝对路径；只在执行记录中记录当次实际使用的命令

## 校验习惯

- 使用能验证当前改动的最小命令
- 改动涉及生成代码时，运行对应生成器
- 改动只影响局部包时，优先跑定向 `go test`，再决定是否扩大验证范围
