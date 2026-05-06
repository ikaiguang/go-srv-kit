---
inclusion: fileMatch
fileMatchPattern: "**/*.go"
---

# 禁止魔法字符串

组件名称、标识符等在多处使用的字符串必须定义为常量（`const`），禁止直接硬编码。

示例：
```go
// 正确
const ComponentRedis = "redis"
NewComponent(ComponentRedis, factory, lc)

// 错误
NewComponent("redis", factory, lc)
```

此规则与编码规范中"魔法数字出现超过 2 次必须定义常量"一致，同样适用于字符串。
