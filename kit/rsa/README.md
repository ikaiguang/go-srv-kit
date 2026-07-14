# rsa

RSA 加解密、签名和密钥解析辅助。

## 基础用法

```go
priKey, pubKey, err := rsapkg.GenRsaKey()
cipher, err := rsapkg.NewRsaCipher(pubKey, priKey)
ciphertext, err := cipher.EncryptToString("plain text")
plaintext, err := cipher.DecryptToString(ciphertext)
```

## 注意事项

密钥必须来自安全配置；不要把私钥写入代码或日志。

## 验证

```bash
go test ./rsa
```
