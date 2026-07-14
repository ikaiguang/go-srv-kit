# aes

AES-CBC 加解密工具，提供随机 IV 的 `NewCBCCipher` / `EncryptCBC` / `DecryptCBC`，以及兼容历史数据的 `NewAesCipher`。

## 基础用法

```go
cipher := aespkg.NewCBCCipher()
encrypted, err := cipher.EncryptToString("plain text", "1234567890123456")
plain, err := cipher.DecryptToString(encrypted, "1234567890123456")
```

## 注意事项

- AES key 长度必须是 16、24 或 32 字节。
- 密钥不要硬编码在业务代码中。
- `NewAesCipher` 使用固定 IV，只建议兼容历史数据时使用。

## 验证

```bash
go test ./aes
```
