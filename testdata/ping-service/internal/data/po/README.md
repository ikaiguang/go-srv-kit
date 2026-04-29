# po - 持久化对象

PO（Persistent Object）对应数据库表结构，是 Data 层的数据载体。

## 命名

- 文件：`{module}.po.go`
- 结构体需实现 `TableName()` 方法

## 示例

```go
type PingPO struct {
    gorm.Model
    Message string `gorm:"type:varchar(255);not null"`
}

func (p *PingPO) TableName() string {
    return "ping"
}
```
