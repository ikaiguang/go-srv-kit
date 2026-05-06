# md5

MD5 摘要计算，支持字节内容和文件。

## 基础用法

```go
sum, err := md5pkg.Md5([]byte("hello"))
fileSum, err := md5pkg.FileMd5("./README.md")
```

## 注意事项

MD5 不适合密码存储、签名安全或防篡改场景；安全哈希优先选择 SHA-256/HMAC 等方案。

## 验证

```bash
go test ./md5
```
