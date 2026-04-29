# wire - 依赖注入工具

`wire/` 提供 [Google Wire](https://github.com/google/wire) 依赖注入的辅助工具。

## 说明

当前仅包含 `wire.util.go`，引入 `wire.Bind` 以便在项目中使用 Wire 的接口绑定功能。

## Wire 在项目中的使用

Wire 用于编译期依赖注入，业务服务的 Wire 定义文件位于：

```
cmd/{service}/export/
├── wire.go           # Wire 定义文件（手写）
├── wire_gen.go       # 自动生成（不要修改）
└── main.export.go    # 导出函数
```

### 生成命令

```bash
# 示例服务
wire ./testdata/ping-service/cmd/ping-service/export

# 或使用 Makefile
make generate
```

如果你维护的是业务服务仓库中的 `cmd/{service}/export`，则按该服务自己的导出目录执行 `wire`。

### 接口绑定

```go
wire.Bind(new(biz.XxxBizRepo), new(*data.xxxData))
```

## 参考

- Wire 文档：https://github.com/google/wire
- 示例：`testdata/ping-service/cmd/ping-service/export/wire.go`
