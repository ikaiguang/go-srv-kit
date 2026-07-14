# base64

Base64 编码和解码工具，提供字节函数与 `Encryptor` 接口。

## 基础用法

```go
encoded := base64pkg.Encode([]byte("hello"))
decoded, err := base64pkg.Decode(encoded)
```

## 注意事项

Base64 不是加密，只适合传输编码；解码外部输入时必须处理错误。

## 验证

```bash
go test ./base64
```
