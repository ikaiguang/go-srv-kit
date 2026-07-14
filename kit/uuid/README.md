# uuid

UUID、xid 生成和解析辅助。

## 基础用法

```go
id := uuidpkg.New()
uuid := uuidpkg.UUID()
```

## 注意事项

UUID 适合全局唯一标识，不应直接作为访问授权凭证。

## 验证

```bash
go test ./uuid
```
