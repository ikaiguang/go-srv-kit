# ip

IP 字符串、本机 IP 和私网 IPv4 探测辅助。

## 基础用法

```go
localIP := ippkg.LocalIP()
ok := ippkg.IsValidIP(localIP)
privateIP := ippkg.PrivateIPv4()
```

## 注意事项

`LocalIP` 会缓存首次探测结果；网络不可用时会回退到 `127.0.0.1`。

## 验证

```bash
go test ./ip
```
