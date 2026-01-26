# 编码规范

## 参考文档

建议阅读以下资料三遍以上：
1. [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
2. [Uber Go 编码规范（中文）](https://github.com/xxjwxc/uber_go_guide_cn)
3. [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
4. [Effective Go](https://go.dev/doc/effective_go)

## 在 GoLand 中开启 golangci-lint

1. 安装 [golangci-lint](https://golangci-lint.run/usage/install/)
2. 在 `Preferences | Tools | File Watchers` 中添加 `golangci-lint`
3. 在 `Preferences | Tools | Actions on Save` 开启 `golangci-lint`

## 代码风格

### 【强制】格式化

- 代码都必须用 `gofmt` 格式化
- 使用 `goimports` 自动格式化 import

### 【建议】换行

- 建议一行代码不要超过 160 列
- 超过的情况，使用合理的换行方法换行

### 【强制】括号和空格

- 遵循 `gofmt` 的逻辑
- 运算符和操作数之间要留空格

```go
// ❌ 错误
d := 60*60*time.Second

// ✅ 正确
d := 60 * 60 * time.Second
```

### 【强制】import 规范

- 使用 `goimports` 自动格式化引入的包名
- import 分组：标准包 → 第三方包 → 本地包（通过空行分隔）
- 不要使用相对路径引入包
- 匿名包建议使用新的分组引入，并写注释

```go
import (
	// 标准库
	"context"
	"fmt"

	// 第三方库
	"github.com/go-kratos/kratos/v2"
	"google.golang.org/protobuf/proto"

	// 项目内部
	"github.com/ikaiguang/go-srv-kit/kratos/error"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/bo"

	// 匿名包
	// import filesystem storage driver
	_ "github.com/go-sql-driver/mysql"
)
```

### 【强制】错误处理

#### error 处理

- `error` 作为函数返回值时，必须进行处理
- `error` 必须是最后一个参数
- 不能使用 `_` 丢弃任何 `err`
- 采用独立的错误流，尽早 `return err`

```go
// ❌ 错误
func do() (error, int) {}

// ✅ 正确
func do() (int, error) {}
```

```go
// ❌ 错误
if err != nil {
	// error handling
} else {
	// normal code
}

// ✅ 正确
if err != nil {
	return err
}
// normal code
```

```go
// ❌ 错误：错误信息大写开头、有句号
fmt.Errorf("Create user err,uid conflict.")

// ✅ 正确
fmt.Errorf("create user err,uid conflict")
```

```go
// ❌ 错误：错误判断与其他逻辑组合
x, y, err := f()
if err != nil || y == nil {
	return err
}

// ✅ 正确
x, y, err := f()
if err != nil {
	return err
}
if y == nil {
	return fmt.Errorf("some error")
}
```

#### panic 处理

- 在业务逻辑处理中**禁止使用 panic**
- 在 main 包中只有完全不可运行的情况才使用 panic
- 每个自行启动的 goroutine，必须在入口处捕获 panic
- 对于异步处理的逻辑，必须使用 GoSafe() 进行包装

```go
// ❌ 错误：野生协程
go foo()

// ✅ 正确
import threadpkg "github.com/ikaiguang/go-srv-kit/kit/thread"
threadpkg.GoSafe(func() {
	foo()
})
```

### 【强制】类型断言失败处理

- 始终使用 `comma ok` 惯用法

```go
// ❌ 错误
t := i.(string)

// ✅ 正确
t, ok := i.(string)
if !ok {
	// 优雅地处理错误
}
```

### 【强制】魔法数字

- 如果魔法数字出现超过 2 次，则禁止使用

```go
// ❌ 错误
func getArea(r float64) float64 {
	return 3.14 * r * r
}

// ✅ 正确
const PI = 3.14
func getArea(r float64) float64 {
	return PI * r * r
}
```

## go-srv-kit 项目规范

### 文件命名

```
internal/service/service/ping.service.go
internal/biz/biz/ping.biz.go
internal/data/data/ping.data.go
internal/service/dto/ping.dto.go
internal/biz/bo/ping.bo.go
internal/data/po/ping.po.go
```

- 采用小写，使用下划线分割单词
- 测试文件命名为 `{源文件名}_test.go`

### 接口定义位置

- Repository 接口定义在 `internal/biz/repo/`
- 由 `internal/data/repo/` 实现
- 使用 Wire 的 `wire.Bind` 进行绑定

### 数据转换

- DTO (Service) ↔ BO (Biz): `internal/service/dto/`
- BO (Biz) ↔ PO (Data): `internal/biz/bo/` 或各自层内部
- 转换函数命名：`ToBo{Xxx}`, `ToProto{Xxx}`, `ToPo{Xxx}`

### 命名约定

| 类型 | 命名 | 示例 |
|------|------|------|
| Service 结构体 | `New{Xxx}Service` | `NewPingService` |
| Biz 结构体 | `New{Xxx}Biz` | `NewPingBiz` |
| Data 结构体 | `New{Xxx}Data` | `NewPingData` |
| 接口 | `{Xxx}BizRepo` | `PingBizRepo` |
| DTO 转换 | `ToBo{Xxx}`, `ToProto{Xxx}` | `ToBoGetPingParam` |

### 注释规范

#### 包注释

```go
// Package math provides basic constants and mathematical functions.
package math
```

#### 结构体注释

```go
// User 用户结构定义了用户基础信息
type User struct {
	Name  string
	Email string
	Demographic string // 族群
}
```

#### 方法注释

```go
// GetPingMessage 获取 ping 消息
func (s *PingService) GetPingMessage(ctx context.Context, req *pb.GetPingMessageReq) (*pb.GetPingMessageResp, error) {
```

#### 变量和常量注释

```go
// AppVersion 应用程序版本号定义
const AppVersion = "1.0.0"
```

### 禁止事项

- 禁止在代码中硬编码配置值
- 禁止在业务层直接访问外部服务
- 禁止跨层调用（Service 不能直接调用 Data）
- 禁止在 Proto 文件中生成代码后再手动修改
- 禁止提交注释掉的代码（除非添加说明）
- 禁止保留未使用的业务函数

## 命名规范

### 【建议】包命名

- 保持 package 的名字和目录一致
- 包名应为小写单词，不使用下划线或混合大小写
- 不使用无意义的包名（util、common、misc、global）
- 项目名可通过中划线连接多个单词

### 【强制】结构体命名

- 采用驼峰命名方式，首字母根据访问控制决定
- 结构体名应该是名词或名词短语
- 避免使用 `Data`、`Info` 这类意义太宽泛的名称

```go
// ❌ 错误
u1 := User{"jack", "jack@gmail"}

// ✅ 正确
u1 := User{
	Name:  "jack",
	Email: "jack@gmail.com",
}
```

### 【建议】接口命名

- 单个函数的接口以 `er` 作为后缀，如 `Reader`、`Writer`
- 两个函数的接口名综合两个函数名
- 三个以上函数的接口名类似于结构体名

### 【强制】变量命名

- 变量名必须遵循驼峰式
- bool 类型变量应以 `Has`、`Is`、`Can` 或 `Allow` 开头
- 变量名更倾向于选择短命名

### 【强制】常量命名

- Golang 常量需遵循驼峰式
- Protobuf 常量定义需遵循全大写
- 枚举类型需要先创建相应类型

```go
// Scheme 传输协议
type Scheme string

const (
	// HTTP 表示 HTTP 明文传输协议
	HTTP Scheme = "http"
	// HTTPS 表示 HTTPS 加密传输协议
	HTTPS Scheme = "https"
)
```

## 控制结构

### if 语句

```go
// 接受初始化语句
if err := file.Chmod(0664); err != nil {
	return err
}

// 变量在左，常量在右
if err != nil {
	// error handling
}

// bool 类型直接判断
if allowUserLogin {
	// do something
}
if !allowUserLogin {
	// do something
}
```

## 函数规范

### 【建议】函数参数

- 参数数量不能超过 5 个
- 尽量用值传递，非指针传递
- 传入参数是 map、slice、chan、interface 不要传递指针

### 【强制】defer

```go
resp, err := http.Get(url)
if err != nil {
	return err
}
// 如果操作成功，再 defer Close()
defer resp.Body.Close()
```

### 【建议】方法的接收器

- 建议以类名第一个英文首字母的小写作为接收器的命名
- 命名不能采用 `me`、`this`、`self` 这类易混淆名称

## Proto 文件规范

### 【强制】proto 布局

```protobuf
syntax = "proto3";

package api.user.service.v1;
option go_package = "github.com/ikaiguang/go-srv-kit/api/user/v1;v1";

// 1. 先定义 rpc 方法
service UserService {
	rpc Login(LoginRequest) returns(LoginReply);
	rpc Logout(LogoutRequest) returns(LogoutReply);
}

// 2. 再定义枚举
message UserStatusEnum {
	enum UserStatus {
		UNSPECIFIED = 0;
		ENABLE = 1;
		DISABLE = 2;
	}
}

// 3. 定义通用数据结构
message UserEntity {
	int64 uid = 1;
	string username = 2;
}

// 4. 定义请求/响应
message LoginRequest {
	string account = 1;
	string password = 2;
}
message LoginReply {
	Token token = 1;
}
```

### 【强制】枚举(Enum)

- 枚举首字母大写，枚举值全部大写，用 `_` 分隔

### 【强制】字段名(Field)

- 字段名以 `snake_case` 形式命名

### 【强制】消息(Message)

- message 采用首字母大写的形式

### 【强制】列表(Repeated fields)

- 列表字段命名采用复数形式

```protobuf
repeated string keys = 1;
```

### 【强制】注释(Comment)

- 注释写在 rpc 或 message 的上方

```protobuf
// 登录凭证
message Token {
	// access_token json web token
	string access_token = 1;
}
```
