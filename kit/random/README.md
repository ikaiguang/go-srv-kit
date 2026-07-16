# random

生成随机字符串、验证码、密码、订单号、trace ID，或随机选取/打乱切片。

## 基础用法

```go
code := randompkg.VerifyCode(6)
token := randompkg.Token(32)
password := randompkg.Password(12)
secureCode, err := randompkg.SecureVerifyCode(6)
securePassword, err := randompkg.SecurePassword(12)
secureToken, err := randompkg.SecureToken(32)
secureURLToken, err := randompkg.SecureBase64URL(32)
```

## 注意事项

- `Token`、`Password`、`VerifyCode` 等历史函数基于 `math/rand`，只适合非安全随机场景。
- 验证码和生成凭据使用 `SecureVerifyCode`、`SecurePassword`；安全 token、重置密码链接、外部认证随机串使用 `SecureToken`、`SecureHex` 或 `SecureBase64URL`。
- 安全随机函数基于 `crypto/rand`，会返回 error，调用方必须处理。

## 验证

```bash
go test ./random
```
