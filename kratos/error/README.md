# error - 统一错误处理

`error/` 提供带堆栈信息和元数据的统一错误处理，基于 Kratos errors 扩展。

## 包名

```go
import errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
```

## 快捷错误创建

```go
errorpkg.ErrorBadRequest("invalid parameter")       // 400
errorpkg.ErrorUnauthorized("token is invalid")       // 401
errorpkg.ErrorForbidden("access denied")             // 403
errorpkg.ErrorNotFound("user not found")             // 404
errorpkg.ErrorConflict("user already exists")        // 409
errorpkg.ErrorInternal("database connection failed") // 500
```

## 错误包装

```go
// 附加堆栈信息
errorpkg.WithStack(kratosErr)

// 包装原始错误（附加到 Metadata）
errorpkg.Wrap(kratosErr, originalErr)

// 格式化任意 error 为带堆栈的 Error
errorpkg.FormatError(err)

// 附加元数据
errorpkg.WrapWithMetadata(kratosErr, map[string]string{"key": "value"})
```

## 元数据 Key

| Key | 说明 |
|-----|------|
| `reason` | 错误代码 |
| `cause` | 错误原因 |
| `error` | 具体错误信息 |

## Proto 定义

```bash
kratos proto client kratos/error/error.kit.proto
```
