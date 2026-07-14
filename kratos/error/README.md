# error

`error` 包的包名为 `errorpkg`，基于 Kratos `errors.Error` 提供统一错误创建、包装、metadata、枚举错误、错误判断和堆栈格式化能力。适合在 Service、Biz、Data 分层中返回可识别的业务错误。

## 安装

```bash
go get github.com/ikaiguang/go-kratos-kit
```

```go
import errorpkg "github.com/ikaiguang/go-kratos-kit/error"
```

## 核心能力

- `New`、`Newf`、`Errorf`、`NewWithMetadata`：创建带 code、reason、message 和 metadata 的错误。
- `WithStack`、`Wrap`、`WrapWithMetadata`、`FormatError`：包装 Kratos 错误并保留堆栈或下游错误信息。
- `FromError`、`Cause`、`Code`、`Reason`、`Message`、`Metadata`：读取错误详情。
- `Is`、`IsCode`、`IsReason`、`IsCustomError`：判断错误类型。
- `Err`、`Errf`、`ErrWithMetadata`：基于实现了 `Enum` 接口的枚举创建错误。
- `ErrorBadRequest`、`ErrorInternalServer`、`ErrorTooManyRequests` 等：由 `error.kit.proto` 生成的常用错误构造函数。
- `DefaultErrorBadRequest`、`DefaultErrorInternalServer` 等：生成的默认错误构造函数。

## 快速使用

```go
err := errorpkg.New(
    http.StatusBadRequest,
    "INVALID_ARGUMENT",
    "invalid request",
)

if errorpkg.IsCode(err, http.StatusBadRequest) {
    reason := errorpkg.Reason(err)
    _ = reason
}
```

```go
dbErr := errors.New("record not found")
wrapped := errorpkg.Wrap(
    errorpkg.ErrorRecordNotFound("record not found"),
    dbErr,
)
_ = wrapped
```

## 测试

```bash
go test ./error
```

## 注意事项

- `error.kit.pb.go`、`error.kit_errors.pb.go`、`error.kit_errors_custom.pb.go`、`error.kit_message.pb.go` 由 proto 或生成流程产生，不要手工修改。
- metadata 会进入响应和日志，避免写入未脱敏的密码、Token、手机号、连接串等敏感信息。
- 使用 `%+v` 格式化 `*errorpkg.Error` 会输出堆栈，生产日志中注意体积和敏感路径暴露风险。
