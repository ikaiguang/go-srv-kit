# 认证与授权规范

## JWT 认证机制

### Token 类型

| 类型 | 用途 | 有效期 |
|------|------|--------|
| Access Token | 访问 API | 2 小时 |
| Refresh Token | 刷新 Access Token | 7 天 |

### Token 类型（TokenType）

```go
const (
    // TokenTypeUser 普通用户
    TokenTypeUser = "USER"
    // TokenTypeAdmin 管理员
    TokenTypeAdmin = "ADMIN"
    // TokenTypeEmployee 员工
    TokenTypeEmployee = "EMPLOYEE"
    // TokenTypeDefault 默认
    TokenTypeDefault = "DEFAULT"
)
```

### Token 结构

```go
type Claims struct {
    // UserID 用户 ID
    UserID uint
    // UserUuid 用户唯一标识
    UserUuid string
    // LoginPlatform 登录平台
    LoginPlatform string
    // TokenType Token 类型
    TokenType string
    // StandardClaims 标准 Claims
    jwt.StandardClaims
}
```

## 白名单配置

### 定义白名单

在 `api/{service}/v1/auth_white_list.api.go`:

```go
package v1

import "github.com/ikaiguang/go-srv-kit/service/middleware"

// ExportAuthWhitelist 导出认证白名单
func ExportAuthWhitelist() []map[string]middleware.TransportServiceKind {
    return []map[string]middleware.TransportServiceKind{
        // Health Check
        {
            "/health": middleware.TransportServiceKindAll,
        },

        // Ping Service - 公开接口
        {
            "/api/v1/ping/say_hello": middleware.TransportServiceKindHTTP,
        },

        // 用户注册、登录
        {
            "/api/v1/user/register": middleware.TransportServiceKindHTTP,
            "/api/v1/user/login":    middleware.TransportServiceKindHTTP,
        },

        // API 文档
        {
            "/q": middleware.TransportServiceKindAll,
        },
    }
}
```

### TransportServiceKind

```go
const (
    // TransportServiceKindHTTP 仅 HTTP
    TransportServiceKindHTTP = "http"
    // TransportServiceKindGRPC 仅 gRPC
    TransportServiceKindGRPC = "grpc"
    // TransportServiceKindAll HTTP 和 gRPC
    TransportServiceKindAll = "all"
)
```

## 登录流程

### 1. 用户登录

```go
func (s *userService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
    // 1. 验证用户名密码
    user, err := s.userBiz.VerifyPassword(ctx, req.GetUsername(), req.GetPassword())
    if err != nil {
        return nil, err
    }

    // 2. 生成 Token
    tokenManager, err := s.tokenManager.(authpkg.TokenManger)
    accessToken, err := tokenManager.GenerateToken(authpkg.TokenInfo{
        UserID:        user.ID,
        UserUuid:      user.UUID,
        LoginPlatform: req.GetPlatform(),
        TokenType:     authpkg.TokenTypeUser,
    })

    refreshToken, err := tokenManager.GenerateRefreshToken(authpkg.TokenInfo{
        UserID:    user.ID,
        UserUuid:  user.UUID,
        TokenType: authpkg.TokenTypeUser,
    })

    // 3. 返回
    return &pb.LoginResp{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        ExpiresIn:    7200, // 2 小时
    }, nil
}
```

### 2. 刷新 Token

```go
func (s *userService) RefreshToken(ctx context.Context, req *pb.RefreshTokenReq) (*pb.RefreshTokenResp, error) {
    // 1. 验证 Refresh Token
    tokenManager, _ := s.tokenManager.(authpkg.TokenManger)
    claims, err := tokenManager.ParseRefreshToken(req.GetRefreshToken())
    if err != nil {
        return nil, errorpkg.ErrorUnauthorized("invalid refresh token")
    }

    // 2. 生成新的 Access Token
    accessToken, err := tokenManager.GenerateToken(authpkg.TokenInfo{
        UserID:        claims.UserID,
        UserUuid:      claims.UserUuid,
        LoginPlatform: claims.LoginPlatform,
        TokenType:     claims.TokenType,
    })

    return &pb.RefreshTokenResp{
        AccessToken: accessToken,
        ExpiresIn:   7200,
    }, nil
}
```

## 获取当前用户

### 从 Context 获取用户信息

```go
import authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"

func (s *userService) GetProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileResp, error) {
    // 1. 从 Context 获取用户信息
    userInfo, err := authpkg.GetUserInfo(ctx)
    if err != nil {
        return nil, errorpkg.ErrorUnauthorized("unauthorized")
    }

    // 2. 使用用户信息
    user, err := s.userBiz.GetUser(ctx, userInfo.UserID)
    if err != nil {
        return nil, err
    }

    return &pb.GetProfileResp{User: user}, nil
}
```

### UserInfo 结构

```go
type UserInfo struct {
    UserID        uint
    UserUuid      string
    LoginPlatform string
    TokenType     string
}
```

## 权限控制

### 基于角色的访问控制（RBAC）

```go
func (s *adminService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
    // 1. 获取当前用户
    userInfo, err := authpkg.GetUserInfo(ctx)
    if err != nil {
        return nil, errorpkg.ErrorUnauthorized("unauthorized")
    }

    // 2. 检查权限
    if userInfo.TokenType != authpkg.TokenTypeAdmin {
        return nil, errorpkg.ErrorForbidden("only admin can delete users")
    }

    // 3. 执行删除
    err = s.adminBiz.DeleteUser(ctx, req.GetUserId())
    return &pb.DeleteUserResp{}, err
}
```

### 基于资源的访问控制

```go
func (s *orderService) GetOrder(ctx context.Context, req *pb.GetOrderReq) (*pb.GetOrderResp, error) {
    userInfo, err := authpkg.GetUserInfo(ctx)
    if err != nil {
        return nil, errorpkg.ErrorUnauthorized("unauthorized")
    }

    // 获取订单
    order, err := s.orderBiz.GetOrder(ctx, req.GetOrderId())
    if err != nil {
        return nil, err
    }

    // 检查是否有权访问
    if order.UserID != userInfo.UserID && userInfo.TokenType != authpkg.TokenTypeAdmin {
        return nil, errorpkg.ErrorForbidden("access denied")
    }

    return &pb.GetOrderResp{Order: order}, nil
}
```

## 配置

### JWT 配置

```yaml
server:
  http:
    addr: 0.0.0.0:10101

setting:
  enable_auth_middleware: true

encrypt:
  jwt_signing_key: "your-secret-key"
  jwt_refresh_signing_key: "your-refresh-secret-key"
```

### Token 过期时间

在 `kratos/auth/token.go` 中配置：

```go
const (
    // AccessTokenExpire Access Token 过期时间
    AccessTokenExpire = 2 * time.Hour
    // RefreshTokenExpire Refresh Token 过期时间
    RefreshTokenExpire = 7 * 24 * time.Hour
)
```
