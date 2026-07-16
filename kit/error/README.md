# error

`error` 提供带上下文和 stack trace 的 error wrapping 能力，并兼容 Go 标准库的
error chain。

本包基于 [`github.com/pkg/errors`](https://github.com/pkg/errors) `v0.9.1`
维护，保留原项目的 BSD-2-Clause 许可证和公开 API。仓库内的 import path 为：

```go
import errors "github.com/ikaiguang/go-srv-kit/kit/v3/error"
```

## 基本用法

```go
file, err := os.Open("config.yaml")
if err != nil {
	return errors.Wrap(err, "open config")
}
defer file.Close()
```

`Wrap`、`Wrapf` 和 `WithStack` 会在调用位置捕获 stack trace：

```go
err := errors.Wrap(io.EOF, "read response")
fmt.Printf("%v\n", err)  // read response: EOF
fmt.Printf("%+v\n", err) // 包含完整 error chain 和 stack trace
```

只需要附加消息而不需要新 stack 时，使用 `WithMessage` 或
`WithMessagef`：

```go
err := errors.WithMessage(io.EOF, "read response")
```

## Error Chain

`Is`、`As` 和 `Unwrap` 直接遵循 Go 标准库语义：

```go
if errors.Is(err, io.EOF) {
	// handle EOF
}
```

`Cause` 为兼容旧版 `pkg/errors` 保留，用于读取实现了 `Cause() error` 的
legacy chain。可比较的 error 会进行循环检测，因此不会截断合法的超深 chain；
完全由不可比较 error 组成的异常 chain 最多遍历 10,000 层，以避免错误实现永久
占用 CPU。新代码优先使用 `Is`、`As` 和 `Unwrap`。

## 性能与安全边界

- `New`、`Errorf`、`Wrap`、`Wrapf` 和 `WithStack` 会调用
  `runtime.Callers`，因此比标准库 `errors.New` 产生更多分配和 CPU 开销。
- 不需要 stack 的高频路径优先使用标准库 error 或本包的 `WithMessage`。
- `%+v`、`Frame.MarshalText` 和 `StackTrace` 可能包含函数名及本地构建路径。
  不要把完整 stack trace 直接返回给不可信客户端。
- `As` 对非法 target 的 panic 行为与标准库一致，调用方必须传入非 nil 指针。

## 许可证

BSD-2-Clause。原作者与完整条款见 [LICENSE](./LICENSE)。
