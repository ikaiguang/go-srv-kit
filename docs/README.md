# 工具使用指导索引

本索引覆盖当前仓库根目录下的 Go 工具包。每个工具目录内的 `README.md` 是该工具的主要使用说明，便于 GitHub/GitLab 打开目录时直接展示，也避免后续修改某个 kit 时反复改动一个大文档。

## 通用约定

- 使用前按需导入对应工具包，不建议把工具包封装成跨层全局依赖。
- 涉及 IO、网络、命令执行、随机数、加解密时，应在调用方保留超时、权限、密钥和错误处理边界。
- 定向验证命令格式：`go test ./<tool>`。
- 全量验证命令：`go test ./...`。

## 工具目录

- [aes](../kit/aes/README.md)：AES-CBC 加解密。
- [base64](../kit/base64/README.md)：Base64 编码和解码。
- [buffer](../kit/buffer/README.md)：复用 `bytes.Buffer`。
- [chinese](../kit/chinese/README.md)：中文编码转换和检测。
- [cmd](../kit/cmd/README.md)：外部命令执行。
- [connection](../kit/connection/README.md)：连接和 WebSocket 请求判断。
- [curl](../kit/curl/README.md)：HTTP 请求辅助。
- [download](../kit/download/README.md)：HTTP 流式下载。
- [email](../kit/email/README.md)：SMTP 邮件发送。
- [file](../kit/file/README.md)：文件复制、移动和 hash。
- [filepath](../kit/filepath/README.md)：目录遍历和目录操作。
- [header](../kit/header/README.md)：HTTP header 辅助。
- [id](../kit/id/README.md)：分布式 ID。
- [ip](../kit/ip/README.md)：IP 辅助。
- [json](../kit/json/README.md)：JSON 序列化辅助。
- [locker](../kit/locker/README.md)：本地锁、缓存锁和分布式锁接口。
- [md5](../kit/md5/README.md)：MD5 摘要。
- [operator](../kit/operator/README.md)：三元表达式辅助。
- [os](../kit/os/README.md)：操作系统判断。
- [page](../kit/page/README.md)：分页请求和响应辅助。
- [password](../kit/password/README.md)：密码 hash 和校验。
- [path](../kit/path/README.md)：源码目录辅助。
- [ptr](../kit/ptr/README.md)：指针辅助。
- [random](../kit/random/README.md)：随机字符串和随机选择。
- [reflect](../kit/reflect/README.md)：反射辅助。
- [regex](../kit/regex/README.md)：常见格式正则校验。
- [rsa](../kit/rsa/README.md)：RSA 加解密和签名。
- [slice](../kit/slice/README.md)：切片辅助。
- [sort](../kit/sort/README.md)：排序辅助。
- [string](../kit/string/README.md)：字符串转换辅助。
- [thread](../kit/thread/README.md)：goroutine 安全执行。
- [time](../kit/time/README.md)：时间辅助。
- [url](../kit/url/README.md)：URL 编码和拼接。
- [uuid](../kit/uuid/README.md)：UUID 和 xid 辅助。
- [writer](../kit/writer/README.md)：writer 和日志轮转。
- [zip](../kit/zip/README.md)：zip 压缩和解压。
