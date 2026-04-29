# auth - JWT 认证

`auth/` 提供 JWT Token 的签发、验证、刷新和中间件，支持多种 Token 类型。

## 包名

```go
import authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
```

## Token 类型

| 类型 | 常量 | 说明 |
|------|------|------|
| USER | `TokenTypeEnum_USER` | 普通用户 |
| ADMIN | `TokenTypeEnum_ADMIN` | 管理员 |
| EMPLOYEE | `TokenTypeEnum_EMPLOYEE` | 员工 |
| DEFAULT | `TokenTypeEnum_DEFAULT` | 默认 |

## 服务端中间件

```go
authMiddleware := authpkg.Server(signKeyFunc,
    authpkg.WithClaims(func() jwt.Claims { return &authpkg.Claims{} }),
)
```

## 获取用户信息

```go
claims, ok := authpkg.GetAuthClaimsFromContext(ctx)
if ok {
    userID := claims.Payload.UserID
    tokenType := claims.Payload.TokenType
}
```

## Token 管理

```go
// 签发 Token
claims := authpkg.GenAuthClaimsByAuthPayload(payload, accessTokenExpire)

// 刷新 Token
newClaims := authpkg.GenAuthClaimsByAuthClaims(oldClaims, accessTokenExpire)
refreshClaims := authpkg.GenRefreshClaimsByAuthClaims(oldClaims, refreshTokenExpire)
```

## Proto 定义

```bash
kratos proto client ./kratos/auth/*.proto
```

## 参考

- JWT 库：[golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- Kratos beer-shop 示例
