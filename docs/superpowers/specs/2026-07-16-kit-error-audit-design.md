# Kit Error 审查与修复设计

## 目标

将 `kit/error` 作为仓库自行维护的 `github.com/pkg/errors v0.9.1` fork。
保留 BSD-2-Clause 版权归属和公开行为，同时使迁移后的包能够在当前
`github.com/ikaiguang/go-srv-kit/kit/v3` 模块及 Go 1.26.3 下正确测试和维护。

## 范围

- 审查 `kit/error` 的生产代码、测试、示例、README、许可证、公开 API、
  stack 格式化、error chain、panic 边界、可用性、并发、内存分配和行尾。
- 保持以下公开函数和类型兼容：`New`、`Errorf`、`Wrap`、`Wrapf`、
  `WithStack`、`WithMessage`、`WithMessagef`、`Cause`、`Is`、`As`、
  `Unwrap`、`Frame` 和 `StackTrace`。
- 不修改用户对 `/v3` module path 的调整，也不处理无关的 `kit` 变更。
- `file-rotatelogs` 仍引用上游包时，不移除 `github.com/pkg/errors` 依赖。

## 当前状态

三个生产代码文件及相关测试均与 `github.com/pkg/errors v0.9.1` 完全一致，
但没有完成迁移适配：

- stack 格式预期仍硬编码 `github.com/pkg/errors`；
- `example_test.go` 导入并测试上游依赖，而不是本地包；
- 两处测试把非常量字符串直接传给 printf-like API，触发 Go 1.26 vet；
- Formatter 无法返回写入错误，但代码没有显式声明忽略，导致 lint 失败；
- README 仍展示上游 import path 和已经失效的 CI 链接。

因此普通测试、race test、vet 和 lint 当前均失败。所有文件已通过
`kit/.gitattributes` 使用 LF。

## 修复设计

### 来源与文档

保留 `LICENSE`。README 改为本地 import path，声明代码来源于 v0.9.1，
说明与 Go 标准 error chain 的兼容关系，并记录捕获 stack trace 的性能成本及
可能暴露构建路径的风险。

### 测试迁移

确保所有测试和示例都调用本地包。把上游 module path 预期改为当前本地包
路径，保留 stack frame 和行号断言；把动态 format 测试改为常量 format 加参数，
使 Go 1.26 vet 可以安全分析。

### 生产代码加固

- 保持现有 wrapping 和格式化输出不变。
- `fmt.Formatter` 不能返回错误，因此显式忽略 Formatter 内部写入错误；其他可
  传播错误的 API 不得静默丢弃错误。
- 在旧 Go 1.13 build tag 旁增加现代 `//go:build` 条件。
- 为循环的 legacy `Cause()` chain 增加回归测试。短链使用栈内记录，较长的
  可比较 chain 按需使用 map 检测循环，因此不截断合法深链；仅对无法比较的异常
  chain 保留防御上限，避免错误实现令进程永久占用 CPU。
- `As` 对非法 target 的 panic 行为继续与标准库一致。

### 性能

捕获 stack 是该包的核心用途，并通过现有 API 明确区分：`New`、`Errorf`、
`Wrap` 和 `WithStack` 捕获 stack，`WithMessage` 不捕获。修复前后运行 benchmark；
循环检测引入的 reflect 仅用于判断 error 是否可比较，并通过独立 `Cause`
benchmark 验证短链保持零分配；没有测量依据时不引入对象池或 stack 深度变更。

## 兼容性与命名

现有导出名称符合 Go 与 `pkg/errors` 既有惯例，不需要改名或增加 Deprecated
wrapper。内部修复不得改变合法输入下的 error string、`%s`、`%q`、`%v`、
`%+v`、`errors.Is`、`errors.As`、`errors.Unwrap`、`Cause` 或 stack trace 顺序。

## 验证

1. 每项生产行为修改都执行针对性的 red-green test。
2. 在 `kit` module 下运行 `go test ./error` 和 `go test -race ./error`。
3. 运行 `go vet ./error` 和 `golangci-lint run ./error`。
4. 检查格式、`git diff --check`、LF 和依赖一致性。
5. 对比相关 benchmark，并对最终 diff 做独立复核。

## 最终需报告的剩余风险

- stack trace 在格式化或 marshal 时会暴露函数名和本地构建路径，调用方不得将其
  直接返回给不可信客户端。
- 捕获 stack 会产生分配并遍历 runtime frame；不需要诊断 stack 的热路径应使用
  message-only wrapping。
- 采用本地 fork 后，仓库需要持续跟进未来 Go runtime 和格式化行为变化。
