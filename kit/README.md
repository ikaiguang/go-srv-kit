# kit - 通用工具库

`kit/` 提供与业务无关的通用工具函数，涵盖加密、编码、ID 生成、文件操作、网络请求等。

这是一个独立的 Go 模块（有自己的 `go.mod`），无框架依赖，可在任何 Go 项目中使用。

## 模块导入

```go
import "github.com/ikaiguang/go-srv-kit/kit"
```

## 工具目录

### 加密与编码

| 目录 | 包名 | 说明 |
|------|------|------|
| `aes/` | `aespkg` | AES 加密（CBC 模式） |
| `rsa/` | `rsapkg` | RSA 加密/解密、签名/验签 |
| `md5/` | `md5pkg` | MD5 哈希 |
| `base64/` | `base64pkg` | Base64 编解码 |
| `password/` | `passwordpkg` | 密码哈希与验证（bcrypt） |

### ID 生成

| 目录 | 包名 | 说明 |
|------|------|------|
| `id/` | `idpkg` | ID 生成器接口定义 |
| `snowflake/` | `snowflakepkg` | Snowflake 分布式 ID 生成 |
| `uuid/` | `uuidpkg` | UUID 生成 |

### 文件与路径

| 目录 | 包名 | 说明 |
|------|------|------|
| `file/` | `filepkg` | 文件操作（读写、哈希） |
| `filepath/` | `filepathpkg` | 目录遍历、文件路径工具 |
| `path/` | `pathpkg` | 路径处理工具 |
| `zip/` | `zippkg` | ZIP 压缩/解压 |
| `download/` | `downloadpkg` | 文件下载 |

### 字符串与数据处理

| 目录 | 包名 | 说明 |
|------|------|------|
| `string/` | `stringpkg` | 字符串工具（脱敏、转换等） |
| `json/` | `jsonpkg` | JSON 序列化/反序列化 |
| `regex/` | `regexpkg` | 正则表达式工具 |
| `chinese/` | `chinesepkg` | GBK 编码处理 |
| `slice/` | `slicepkg` | 切片工具（泛型） |
| `sort/` | `sortpkg` | 排序工具（泛型） |
| `reflect/` | `reflectpkg` | 反射工具 |
| `ptr/` | `ptrpkg` | 指针工具（泛型） |
| `operator/` | `operatorpkg` | 三元运算符 |

### 网络与通信

| 目录 | 包名 | 说明 |
|------|------|------|
| `curl/` | `curlpkg` | HTTP 客户端封装（基于 go-resty） |
| `ip/` | `ippkg` | IP 地址工具 |
| `url/` | `urlpkg` | URL 解析工具 |
| `header/` | `headerpkg` | HTTP Header 工具（RequestID、WebSocket） |
| `email/` | `emailpkg` | 邮件发送 |
| `connection/` | `connectionpkg` | 网络连接检测 |

### 并发与系统

| 目录 | 包名 | 说明 |
|------|------|------|
| `thread/` | `threadpkg` | 安全 goroutine（`GoSafe`，自动 recover） |
| `locker/` | `lockerpkg` | 锁接口定义（分布式锁、本地锁、缓存锁） |
| `cmd/` | `cmdpkg` | Shell 命令执行 |
| `os/` | `ospkg` | 操作系统工具 |
| `buffer/` | `bufferpkg` | Buffer 池（sync.Pool） |
| `writer/` | `writerpkg` | 日志文件轮转写入器 |

### 其他

| 目录 | 包名 | 说明 |
|------|------|------|
| `page/` | `pagepkg` | 分页工具（Proto 定义 + 解析器） |
| `random/` | `randompkg` | 随机数/字符串生成 |
| `time/` | `timepkg` | 时间格式化常量和工具 |

## 使用示例

```go
import (
    threadpkg "github.com/ikaiguang/go-srv-kit/kit/thread"
    uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
    passwordpkg "github.com/ikaiguang/go-srv-kit/kit/password"
)

// 安全 goroutine
threadpkg.GoSafe(func() {
    // 业务逻辑，panic 会被自动 recover
})

// 生成 UUID
id := uuidpkg.NewUUID()

// 密码哈希
hash, err := passwordpkg.HashPassword("mypassword")
ok := passwordpkg.CheckPasswordHash("mypassword", hash)
```
