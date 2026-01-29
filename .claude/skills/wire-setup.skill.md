# Wire 依赖注入技能

## 描述
配置和管理 Wire 依赖注入，包括添加新的 Provider、接口绑定和依赖关系管理。

## 参数
- `layer`: 层级（data/biz/service）
- `component_name`: 组件名称（如：User）
- `dependencies`: 依赖列表

## 执行步骤
1. 在 `cmd/{service}/export/wire.go` 添加新组件
2. 确保依赖顺序正确（从底层到上层）
3. 如需接口绑定，添加 `wire.Bind`
4. 运行 `wire` 生成代码
5. 验证生成的 `wire_gen.go`

## 依赖注入顺序原则
```
基础设施（Logger, DB, Redis）
  ↓
Data 层
  ↓
Business 层
  ↓
Service 层
```

## 常见错误处理
- 循环依赖：重构代码，引入中间层
- 类型不匹配：检查返回类型和 wire.Bind
- Provider 缺失：确保所有依赖都有 Provider
