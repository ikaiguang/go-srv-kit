# go-kit

`go-kit` 是一组轻量、可按需引入的 Go 工具包，覆盖加解密、编码、HTTP、文件、分页、随机数、ID、并发、时间、压缩等常见工程场景。

项目目标是提供简单、可测试、低耦合的基础工具。每个工具都放在独立目录中，可以按模块路径单独导入。

## 特性

- 按目录拆分工具包，避免引入不需要的业务抽象。
- 每个 Go 工具目录都有本地 `README.md`，便于在 GitHub 目录页直接查看用法。
- 已补充单元测试和边界测试，覆盖常见正常路径、错误路径和安全风险。
- 尽量保持兼容新增，不轻易破坏已有 API。
- 对安全相关场景给出明确边界，例如安全随机、TLS 校验、路径穿越、密码 hash、MD5 限制等。

## 安装

```bash
go get github.com/ikaiguang/go-kit
```

要求 Go 版本：

```text
go 1.25.9
```

## 快速示例

### HTTP 请求

```go
package main

import (
	"context"
	"fmt"
	"time"

	curlpkg "github.com/ikaiguang/go-kit/curl"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := curlpkg.NewGetRequestContext(ctx, "https://example.com", nil)
	if err != nil {
		panic(err)
	}

	code, body, err := curlpkg.Do(req, curlpkg.WithTimeout(3*time.Second))
	if err != nil {
		panic(err)
	}
	if !curlpkg.IsSuccessCode(code) {
		panic(curlpkg.ErrRequestFailure(code))
	}

	fmt.Println(string(body))
}
```

### 安全随机 Token

```go
import randompkg "github.com/ikaiguang/go-kit/random"

token, err := randompkg.SecureToken(32)
if err != nil {
	return err
}
```

### 文件下载

```go
import downloadpkg "github.com/ikaiguang/go-kit/download"

reply, err := downloadpkg.StreamDownload(ctx, &downloadpkg.DownloadParam{
	URL:        "https://example.com/file.zip",
	OutputPath: "./runtime/file.zip",
	BufferSize: 32 * 1024,
})
```

## 工具目录

| 目录 | 用途 |
| --- | --- |
| [aes](./aes/README.md) | AES-CBC 加解密 |
| [base64](./base64/README.md) | Base64 编码和解码 |
| [buffer](./buffer/README.md) | `bytes.Buffer` 复用 |
| [chinese](./chinese/README.md) | 中文编码转换和检测 |
| [cmd](./cmd/README.md) | 外部命令执行 |
| [connection](./connection/README.md) | 连接和 WebSocket 请求判断 |
| [curl](./curl/README.md) | HTTP 请求辅助 |
| [download](./download/README.md) | HTTP 流式下载 |
| [email](./email/README.md) | SMTP 邮件发送 |
| [file](./file/README.md) | 文件复制、移动和 hash |
| [filepath](./filepath/README.md) | 目录遍历和目录操作 |
| [header](./header/README.md) | HTTP header 辅助 |
| [id](./id/README.md) | 分布式 ID |
| [ip](./ip/README.md) | IP 辅助 |
| [json](./json/README.md) | JSON 序列化辅助 |
| [locker](./locker/README.md) | 本地锁、缓存锁和分布式锁接口 |
| [md5](./md5/README.md) | MD5 摘要 |
| [operator](./operator/README.md) | 三元表达式辅助 |
| [os](./os/README.md) | 操作系统判断 |
| [page](./page/README.md) | 分页请求和响应辅助 |
| [password](./password/README.md) | 密码 hash 和校验 |
| [path](./path/README.md) | 源码目录辅助 |
| [ptr](./ptr/README.md) | 指针辅助 |
| [random](./random/README.md) | 随机字符串和安全随机 token |
| [reflect](./reflect/README.md) | 反射辅助 |
| [regex](./regex/README.md) | 常见格式正则校验 |
| [rsa](./rsa/README.md) | RSA 加解密和签名 |
| [slice](./slice/README.md) | 切片辅助 |
| [sort](./sort/README.md) | 排序辅助 |
| [string](./string/README.md) | 字符串转换辅助 |
| [thread](./thread/README.md) | goroutine 安全执行 |
| [time](./time/README.md) | 时间辅助 |
| [url](./url/README.md) | URL 编码和拼接 |
| [uuid](./uuid/README.md) | UUID 和 xid 辅助 |
| [writer](./writer/README.md) | writer 和日志轮转 |
| [zip](./zip/README.md) | zip 压缩和解压 |

更多索引见 [docs/README.md](../docs/README.md)。

## 测试

运行全量测试：

```bash
go test ./...
```

运行单个工具测试：

```bash
go test ./curl
go test ./random
```

如果本机默认 Go build cache 权限异常，可以临时指定 workspace 内缓存目录：

```bash
GOCACHE="$(pwd)/.gocache" go test ./...
```

Windows PowerShell：

```powershell
$env:GOCACHE="$PWD\.gocache"; go test ./...
```

## 安全说明

- `random.Token`、`random.Password` 等历史函数基于 `math/rand`，只适合非安全随机场景。安全 token 使用 `random.SecureToken`、`random.SecureHex` 或 `random.SecureBase64URL`。
- `md5` 只适合兼容摘要或非安全校验，不适合密码存储、签名安全或防篡改场景。
- `curl.WithInsecureSkipVerify` 只应在开发或测试环境使用。
- `zip.Unzip` 已做路径穿越防护，解压外部 zip 时仍应使用受控输出目录。
- 命令执行、文件写入、下载和邮件发送等能力应由调用方控制输入来源、权限、超时和错误处理。

## 文档

- [docs/README.md](../docs/README.md)：工具文档索引。
- 每个工具目录下的 `README.md`：该工具的用途、基础用法、注意事项和验证命令。

## 贡献

提交修改前建议执行：

```bash
gofmt -w <changed-go-files-or-dirs>
go test ./...
```

新增或调整工具时，请同步补充：

- 对应目录的 `README.md`
- 对应 `*_test.go`
- 必要的错误路径、边界输入和安全回归测试

## License

See [LICENSE](./LICENSE).
