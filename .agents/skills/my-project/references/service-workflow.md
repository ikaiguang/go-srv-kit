# 服务开发流程

## 何时读取

当任务涉及以下内容时，读取本文件：

- Proto API 定义
- Service handler
- Biz 逻辑
- Data repo
- 配置接线
- Wire 装配

## API 与服务结构

服务专属 API 优先沿用以下目录结构：

`api/{service-name}/v1/`

常见子目录：

- `services/`：RPC 定义
- `resources/`：请求和响应消息
- `errors/`：错误定义
- `enums/`：枚举定义

## 实现流程

1. 在正确的 API 目录下新增或修改 Proto 定义
2. 使用根目录 `Makefile` 中对应的目标生成代码
3. 实现或调整 `Service` 层 handler 与 DTO 转换
4. 实现或调整 `Biz` 层逻辑
5. 需要时在 `biz/repo/` 定义仓储接口
6. 在 `Data` 层补齐仓储实现
7. 更新 `cmd/.../export/wire.go` 并重新生成 Wire 输出

## 分层职责

### Service 层

- 负责 HTTP/gRPC 入口
- 负责请求结构校验和 DTO 转换
- 只调用 `Biz` 层

### Biz 层

- 负责业务规则和编排
- 依赖仓储接口，而不是具体数据实现
- 返回业务语义明确的错误，不向上泄漏底层存储细节

### Data 层

- 负责具体持久化或远程访问实现
- 将存储和客户端错误转换为共享错误形式
- 按包内既有模式返回 PO 或业务层可接受的数据

## 配置接线

处理服务专属配置时：

1. 先在服务配置 Proto 中增加字段
2. 再通过 `cmd/*/export/main.export.go` 导出配置
3. 最后通过 launcher manager 或包内既有路径读取配置

引入新配置形态前，先对照示例服务实现。

## Wire 装配

修改 provider 或构造函数后：

1. 更新对应的 `wire.go`
2. 运行：

```bash
wire ./testdata/ping-service/cmd/ping-service/export
```

不要手工修改生成的 `wire_gen.go`。

## 首选参考路径

如果不确定一个业务服务应如何接入，先从这些位置开始：

- `testdata/ping-service/cmd/ping-service/export/wire.go`
- `testdata/ping-service/internal/service/service/`
- `testdata/ping-service/internal/biz/`
- `testdata/ping-service/internal/data/`
