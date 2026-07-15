# auth

`auth` 是 `github.com/ikaiguang/go-srv-kit/auth/auth` 包，包名为 `authpkg`。它适合在 Go / Kratos 服务中处理 JWT 访问令牌、刷新令牌、认证中间件和 Redis 令牌状态管理。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/auth
```

```go
import auth "github.com/ikaiguang/go-srv-kit/auth/auth"
```

## 核心能力

- `Payload`：认证载荷，包含用户 ID、用户 UUID、登录平台、登录方式、登录限制和令牌类型。
- `Claims`：基于 `jwt.RegisteredClaims` 的项目认证声明，额外携带 `Payload`。
- `TokenItem`：Redis 中保存的访问令牌或刷新令牌摘要，包含 token ID、刷新 token ID、过期时间和载荷。
- `NewSignEncryptor`：创建 HMAC JWT 签名器，用于访问令牌签发和解析。
- `NewCBCCipher`：创建 AES-CBC 刷新令牌加解密器。
- `NewAuthRepo`：创建认证仓储，提供签发、刷新、解码和校验令牌的统一接口。
- `NewTokenManager`：创建基于 Redis 的令牌管理器，处理白名单、黑名单、登录限制和过期令牌清理。
- `Server` / `Client`：Kratos 服务端和客户端认证中间件。
- `WithSigningMethod`、`WithClaims`、`WithAccessTokenValidator`、`WithAccessTokenHeader`：中间件配置项。
- `MustUserTokenType`、`MustAdminTokenType`、`MustEmployeeTokenType`、`MustInTokenTypes`：常用 token 类型校验器。

## 快速使用

签发访问令牌和刷新令牌：

```go
ctx := context.Background()

repo, err := auth.NewAuthRepo(auth.Config{
	SignCrypto:    auth.NewSignEncryptor("access-token-sign-key"),
	RefreshCrypto: auth.NewCBCCipher("0123456789abcdef0123456789abcdef"),
}, log.DefaultLogger, nil)
if err != nil {
	return err
}

claims := auth.GenAuthClaimsByAuthPayload(&auth.Payload{
	UserID:        1001,
	LoginPlatform: auth.LoginPlatformEnum_COMPUTER,
	LoginType:     auth.LoginTypeEnum_USERNAME_AND_PASSWORD,
	LoginLimit:    auth.LoginLimitEnum_ONLY_ONE,
	TokenType:     auth.TokenTypeEnum_USER,
}, 0)

tokenResp, err := repo.SignToken(ctx, claims)
if err != nil {
	return err
}

_ = tokenResp.AccessToken
_ = tokenResp.RefreshToken
```

解码并校验令牌：

```go
claims, err := repo.DecodeAccessToken(ctx, tokenResp.AccessToken)
if err != nil {
	return err
}
if err := repo.VerifyAccessToken(ctx, claims); err != nil {
	return err
}
```

在 Kratos 服务端中间件中使用：

```go
serverAuth := auth.Server(
	repo.JWTSigningKeyFunc,
	auth.WithSigningMethod(repo.JWTSigningMethod()),
	auth.WithClaims(repo.JWTSigningClaims),
	auth.WithAccessTokenValidator(repo.VerifyAccessToken, auth.MustUserTokenType),
)
_ = serverAuth
```

启用 Redis 令牌管理：

```go
tokenManager := auth.NewTokenManager(log.DefaultLogger, redisClient, nil)
repo, err := auth.NewAuthRepo(auth.Config{
	SignCrypto:    auth.NewSignEncryptor("access-token-sign-key"),
	RefreshCrypto: auth.NewCBCCipher("0123456789abcdef0123456789abcdef"),
}, log.DefaultLogger, tokenManager)
```

## API 说明

### 令牌生成

- `GenAuthClaimsByAuthPayload(payload, expire)`：基于业务载荷生成访问令牌 claims；`expire <= 0` 时使用默认访问令牌过期时间。
- `GenAuthClaimsByAuthClaims(claims, expire)`：基于已有 claims 生成新的访问令牌 claims。
- `GenRefreshClaimsByAuthClaims(claims, expire)`：基于访问令牌 claims 生成刷新令牌 claims。
- `CheckAndCorrectAuthClaims(claims)`：补齐缺失的 token ID 和默认过期时间。

### 仓储接口

`AuthRepo` 提供以下主要方法：

- `SignToken`：签发访问令牌和刷新令牌。
- `RefreshToken`：基于刷新令牌 claims 签发新令牌，并缩短旧令牌有效期。
- `DecodeAccessToken` / `DecodeRefreshToken`：解析令牌并检查 JWT 标准过期时间。
- `VerifyAccessToken` / `VerifyRefreshToken`：结合 `TokenManager` 检查黑名单、白名单和过期状态。

### Redis TokenManager

`TokenManager` 使用 Redis hash 保存用户令牌，并使用独立 key 保存黑名单和登录限制。默认 key 前缀包括：

- `kit:auth_token:`：用户令牌集合。
- `kit:auth_black:`：黑名单令牌。
- `kit:auth_limit:`：登录限制标记。
- `kit:auth_clear:`：过期令牌清理锁。

可以通过 `AuthCacheKeyPrefix` 覆盖这些前缀。

## 测试

```bash
go test ./auth
```

全仓库测试：

```bash
go test ./...
```

## 注意事项

- `NewAuthRepo` 要求 `Config.SignCrypto` 和 `Config.RefreshCrypto` 非空；`RefreshTokenExpire` 必须大于或等于 `AccessTokenExpire`。
- `NewCBCCipher` 的 key 需要满足底层 AES-CBC 实现对密钥长度的要求；示例字符串只用于说明调用方式。
- 不传 `TokenManager` 时，`AuthRepo` 不会执行 Redis 白名单、黑名单和登录限制校验。
- `WithClaims` 用在 `Server` 时需要每次返回新的 `jwt.Claims` 对象，避免并发写问题。
- `Server` 从 Kratos server transport 上下文读取 `Authorization` 请求头；非 transport 上下文会返回 `ErrWrongContext`。
- `auth/*.pb.go` 和 `auth/*_custom.pb.go` 是生成相关文件，不应手工修改。
