# Kit Error 审查修复实施计划

> **面向 agent worker：** 必须使用 `superpowers:subagent-driven-development`（推荐）或 `superpowers:executing-plans` 按 Task 执行本计划；Step 使用 checkbox（`- [ ]`）记录状态。

**目标：** 将 `kit/error` 完整迁移为可在 Go 1.26.3 下维护、测试和使用的本地 `pkg/errors v0.9.1` 兼容 fork。

**架构：** 保持生产 API 和合法输入行为兼容，先修复复制后失真的测试，再以 TDD 修复循环 `Cause` chain 的可用性问题。生产代码只做有测试或工具证据支持的加固，不重写 stack 实现。

**技术栈：** Go 1.26.3、标准库 `errors`/`fmt`/`runtime`、Testify、race detector、vet、golangci-lint。

---

### Task 1：完成测试与示例的本地迁移

**文件：**
- 修改：`kit/error/errors_test.go`
- 修改：`kit/error/example_test.go`
- 修改：`kit/error/format_test.go`
- 修改：`kit/error/json_test.go`
- 修改：`kit/error/stack_test.go`

- [x] **Step 1：让示例引用本地包**

将 `example_test.go` 的 import 改为：

```go
errors "github.com/ikaiguang/go-srv-kit/kit/v3/error"
```

- [x] **Step 2：迁移 stack path 断言**

把测试预期中的 `github.com/pkg/errors` 和对应源码目录替换为
`github.com/ikaiguang/go-srv-kit/kit/v3/error` 及本仓库 `kit/error` 路径，
保留原有函数名、frame 顺序和行号断言。

- [x] **Step 3：修正 Go 1.26 printf analyzer 用例**

```go
got := Wrapf(tt.err, "%s", tt.message).Error()
got := WithMessagef(tt.err, "%s", tt.message).Error()
```

这两个写法保持原字符串输出不变，同时让 format string 成为常量。

- [x] **Step 4：验证迁移后的既有行为**

运行：

```bash
cd kit && go test ./error
```

预期：不再出现上游 package path 或 non-constant format string 导致的失败；若生产代码 lint 尚未修复，测试本身必须通过。

### Task 2：阻止循环 Cause chain 永久占用 CPU

**文件：**
- 修改：`kit/error/errors_test.go`
- 修改：`kit/error/errors.go`

- [x] **Step 1：增加会失败的循环 chain 测试**

```go
type cyclicCauser struct{ next error }

func (e *cyclicCauser) Error() string { return "cycle" }
func (e *cyclicCauser) Cause() error  { return e.next }

func TestCauseCyclicChainTerminates(t *testing.T) {
	a := &cyclicCauser{}
	b := &cyclicCauser{}
	a.next, b.next = b, a
	done := make(chan error, 1)
	go func() { done <- Cause(a) }()
	select {
	case got := <-done:
		if got == nil {
			t.Fatal("Cause returned nil for a cyclic chain")
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("Cause did not terminate for a cyclic chain")
	}
}
```

- [x] **Step 2：运行测试并确认失败原因**

运行：

```bash
cd kit && go test -vet=off ./error -run TestCauseCyclicChainTerminates -count=1
```

预期：FAIL，错误为 `Cause did not terminate for a cyclic chain`。

- [x] **Step 3：增加兼容合法深链的循环检测**

短 chain 用固定数组在栈内记录；超过 8 层后，仅对可比较 error 按需创建 map。
不可比较 error 无法作为 map key，仅对这类异常 chain 保留 10,000 层防御上限：

```go
const (
	inlineCauseTracking        = 8
	maxNonComparableCauseDepth = 10_000
)
var inlineSeen [inlineCauseTracking]error
var seen map[error]struct{}
for depth := 0; err != nil; depth++ {
	if reflect.TypeOf(err).Comparable() {
		if seen != nil {
			if _, exists := seen[err]; exists {
				return err
			}
			seen[err] = struct{}{}
		} else if depth < len(inlineSeen) {
			for _, previous := range inlineSeen[:depth] {
				if previous == err {
					return err
				}
			}
			inlineSeen[depth] = err
		} else {
			seen = make(map[error]struct{}, depth+1)
			for _, previous := range inlineSeen {
				if previous != nil {
					seen[previous] = struct{}{}
				}
			}
			seen[err] = struct{}{}
		}
	} else if depth >= maxNonComparableCauseDepth {
		return err
	}
	cause, ok := err.(causer)
	if !ok {
		break
	}
	err = cause.Cause()
}
return err
```

- [x] **Step 4：验证循环和正常 chain**

运行：

```bash
cd kit && go test -vet=off ./error -run 'TestCause(CyclicChainTerminates)?$' -count=1
```

预期：PASS，既有 `Cause`、循环 chain 和超过 10,000 层的合法 chain 测试均通过。

### Task 3：完成 Go 1.26 与静态检查适配

**文件：**
- 修改：`kit/error/errors.go`
- 修改：`kit/error/stack.go`
- 修改：`kit/error/go113.go`

- [x] **Step 1：显式处理 Formatter 无法返回的写入错误**

将 Formatter 内的调用写为：

```go
_, _ = io.WriteString(s, value)
_, _ = fmt.Fprintf(s, format, args...)
```

仅在 `fmt.Formatter` 方法内部使用该形式，因为接口没有 error 返回值。

- [x] **Step 2：增加现代 build tag**

`go113.go` 文件头改为：

```go
//go:build go1.13
// +build go1.13
```

- [x] **Step 3：运行静态检查**

运行：

```bash
cd kit && go vet ./error
cd kit && golangci-lint run ./error
```

预期：两个命令均退出 0 且无诊断。

### Task 4：更新本地 fork 文档

**文件：**
- 修改：`kit/error/README.md`
- 保留：`kit/error/LICENSE`

- [x] **Step 1：用中文重写 README**

README 必须包含：本地 import path、v0.9.1 来源、BSD-2-Clause、API 示例、
标准库 `Is`/`As`/`Unwrap` 兼容说明、stack 性能成本和构建路径暴露风险。

- [x] **Step 2：检查许可证和错误引用**

运行：

```bash
rg -n 'go get github.com/pkg/errors|travis-ci|AppVeyor' kit/error
```

预期：README 和示例中无陈旧安装说明或 CI 链接；许可证中的原作者信息保持不变。

### Task 5：性能与完整验证

**文件：**
- 检查：`kit/error/bench_test.go`
- 检查：`kit/error/*.go`
- 检查：`kit/go.mod`
- 检查：`kit/go.sum`

- [x] **Step 1：运行 benchmark**

运行：

```bash
cd kit && go test -run '^$' -bench 'BenchmarkErrors|BenchmarkStackFormatting|BenchmarkCause' -benchmem ./error
```

预期：benchmark 完成；短 `Cause` chain 为 `0 B/op, 0 allocs/op`。

- [x] **Step 2：运行完整测试与 race test**

```bash
cd kit && go test ./error
cd kit && go test -race ./error
cd kit && go test -trimpath ./error
```

预期：全部 PASS；stack path 断言同时兼容普通 checkout 和 trimpath。

- [x] **Step 3：检查格式、依赖与 LF**

```bash
cd kit && gofmt -l error
cd kit && go mod tidy -diff
git diff --check -- kit/error docs/superpowers
git ls-files --eol kit/error
```

预期：无格式、依赖或 whitespace 差异；所有文本为 `w/lf`。

### Task 6：补充 Superpowers 文档语言规则

**文件：**
- 修改：`AGENTS.md`

- [x] **Step 1：增加中文正文规则**

在 `Working Rules` 中增加：

```markdown
- 在 `docs/superpowers/` 下新增或更新的设计、规格、实施计划和审查记录，正文默认使用中文；代码标识符、命令、路径、协议名、库/工具名称及不宜翻译的专业术语保留英文。
```

- [ ] **Step 2：验证规则和本次文档**

检查 `AGENTS.md`、本设计文档和本实施计划，确认规则措辞明确，本次新增文档正文以中文为主且专业术语未被生硬翻译。
