# bo - 业务对象

BO（Business Object）是 Biz 层的数据载体，在 Service 层和 Data 层之间传递。

## 数据流

```
DTO → BO → Biz 处理 → BO → DTO
         ↓              ↑
         PO ← Data 层 → PO
```

## 命名

- 文件：`{module}.bo.go`
