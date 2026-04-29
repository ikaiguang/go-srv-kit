# 编码规则

## 函数形态

- 函数尽量控制在 150 行以内
- 嵌套尽量不超过 3 层；优先使用卫语句或拆 helper 降低复杂度
- 参数数量不超过 4 个；超出时改为参数结构体或 `Options` 模式
- `context.Context` 必须放第一个参数，并命名为 `ctx`
- 参数顺序保持为：`ctx` -> 主请求或主参数 -> 辅助参数 -> `...Option`
- `error` 必须放在返回值最后
- 非 `error` 返回值不超过 3 个；超过时优先封装成结构体
- 只有在 `defer` 需要修改最终返回值时才使用命名返回值

## 命名

- 沿用现有构造函数命名，例如 `NewXxxService`、`NewXxxBiz`、`NewXxxData`
- receiver 使用类型名首字母小写，不使用 `me`、`this`、`self`
- Go 常量使用驼峰命名
- 私有常量使用小写开头
- `.proto` 枚举值保持全大写

## 硬编码与常量

- 不要硬编码配置值、密码、Token、连接串或其他敏感信息
- 重复出现的魔法数字要提取为常量
- 复用的标识符、组件名字符串也要提取为常量，不要到处写裸字符串
- 导出常量和变量应有注释；成组常量可写总注释，并为关键项补充行尾说明

## 安全与控制流

- 业务逻辑中不要使用 `panic`
- 异步 goroutine 使用 `threadpkg.GoSafe()` 包装
- 类型断言使用 `comma ok`
- 不要手工修改生成的 Proto 或 Wire 输出文件

## 错误处理

优先使用共享错误包：

- `errorpkg.ErrorBadRequest(...)`
- `errorpkg.ErrorUnauthorized(...)`
- `errorpkg.ErrorForbidden(...)`
- `errorpkg.ErrorNotFound(...)`
- `errorpkg.ErrorConflict(...)`
- `errorpkg.ErrorInternal(...)`
- `errorpkg.WrapWithMetadata(err, metadata)`
- `errorpkg.FormatError(err)`

分层要求：

- `Service` 层做入参校验、记录必要错误日志、返回业务可理解的错误
- `Biz` 层做业务校验，并统一包装下游错误
- `Data` 层做存储和驱动错误转换，例如把 `gorm.ErrRecordNotFound` 转成 `ErrorNotFound`

## 日志

- 优先使用 `logpkg.WithContext(ctx)`，保留 trace 上下文
- 优先使用结构化日志，如 `Infow`、`Warnw`、`Errorw`
- 不要默认写成 `Infof`、`Errorf` 这种格式化日志
- 错误日志至少带上 `"error", err`
- 只有在堆栈信息对排查有实际帮助时，才补 `"stack", stringutil.GetStackTrace(2)`
- 不要记录未脱敏的密码、手机号、Token、密钥等敏感信息；先用现有脱敏工具处理
