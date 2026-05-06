---
inclusion: fileMatch
fileMatchPattern: "**/auth*"
---

# 认证与授权规范

## JWT Token

| 类型 | 有效期 |
|------|--------|
| Access Token | 2 小时 |
| Refresh Token | 7 天 |

Token 类型: USER, ADMIN, EMPLOYEE, DEFAULT

## 白名单配置

在 `ExportAuthWhitelist()` 中定义无需认证的路径：

```go
func ExportAuthWhitelist() []map[string]middleware.TransportServiceKind {
    return []map[string]middleware.TransportServiceKind{
        {"/health": middleware.TransportServiceKindAll},
        {"/api/v1/user/login": middleware.TransportServiceKindHTTP},
    }
}
```

TransportServiceKind: `http` / `grpc` / `all`

## 获取当前用户

```go
import authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"

userInfo, err := authpkg.GetUserInfo(ctx)
// userInfo.UserID, userInfo.UserUuid, userInfo.TokenType
```

## 权限控制

```go
if userInfo.TokenType != authpkg.TokenTypeAdmin {
    return nil, errorpkg.ErrorForbidden("only admin can access")
}
```
