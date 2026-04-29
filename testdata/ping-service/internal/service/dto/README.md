# dto - 数据传输对象

DTO（Data Transfer Object）负责 Proto 消息和业务对象之间的转换。

## 转换方向

```
Proto Request  → DTO → BO（给 Biz 层使用）
BO（Biz 层返回）→ DTO → Proto Response
```

## 命名

- 文件：`{module}.dto.go`
- Proto → BO：`ToBo{Xxx}(req *pb.XxxRequest) *bo.Xxx`
- BO → Proto：`ToProto{Xxx}(bo *bo.Xxx) *pb.XxxReply`
