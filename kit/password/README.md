# password

密码 hash、校验和密码相关辅助逻辑，基于 bcrypt。

## 基础用法

```go
hashedBytes, err := passwordpkg.Encrypt("plain-password")
ok := passwordpkg.Verify(string(hashedBytes), "plain-password")
err = passwordpkg.Compare(string(hashedBytes), "plain-password")
```

## 注意事项

不要记录明文密码；hash 参数和算法升级要兼容历史数据。

## 验证

```bash
go test ./password
```
