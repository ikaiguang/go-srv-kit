# 1. 前言

为形成统一的 Go 编码风格，以保障公司项目代码的易维护性和编码安全性，特制定本规范。
每项规范内容，给出了要求等级，其定义为：

* 强制（Must）：强制采用，不符合要求的不允许合入master
* 建议（Preferable）：用户理应采用，但如有特殊情况，可以不采用；
* 可选（Optional）：用户可参考，自行决定是否采用；

后续开发过程中会以golangci-lint静态代码扫描工具为主,code review为辅,帮助本规范切实执行

## 1.1 以下资料建议阅读三遍以上

1. [uber的编码规范(英文)](https://github.com/uber-go/guide/blob/master/style.md)
2. [uber的编码规范(中文)](https://github.com/xxjwxc/uber_go_guide_cn)
3. [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments#copying)
4. [EffectiveGo](https://go.dev/doc/effective_go)

## 1.2 在Goland中开启golangci-lint

1. 到[官网](https://golangci-lint.run/usage/install/)或者让goland自动下载并安装`golangci-lint`
2. 在 `Preferences | Tools | File Watchers` 中点击`+` 选择 `golangci-lint` 后使用其默认配置,然后点击保存即可
3. 在 `Preferences | Tools | Actions on Save` 开启 `File Watcher:golangci-lint`,之后的每次保存go文件都会自动检查代码

# 2. 代码风格

## 2.1 【强制】格式化

* 代码都必须用 gofmt 格式化

## 2.2 【建议】换行

* 建议一行代码不要超过160列，超过的情况，使用合理的换行方法换行。

## 2.3 【强制】括号和空格

* 遵循 gofmt 的逻辑。
* 运算符和操作数之间要留空格。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
d := 60*60*time.Scend
```

</td><td>

```go
d := 60 * 60 *time.Scend
```

</td></tr>
</tbody></table>

* 作为输入参数或者数组下标时，运算符和运算数之间不需要空格，紧凑展示。

## 2.4 【强制】import 规范

* 使用goimports自动格式化引入的包名，import 规范原则上以 goimports
  规则为准。goland可以在`Preferences | Editor | Code Style | Go | Imports`中进行配置
* goimports 会自动把依赖包按首字母排序，并对包进行分组管理，通过空行隔开，默认分为本地包（标准库、内部包）、第三方包。
* 标准包永远位于最上面的第一组。
* 内部包是指不能被外部 import 的包，如 GoPath 模式下的包名或者非域名开头的当前项目的 GoModules 包名。
* 带域名的包名都属于第三方包，如 company.code.oa.com/xxx/xxx，github.com/xxx/xxx，不用区分是否是当前项目内部的包。
* goimports 默认最少分成本地包和第三方包两大类，这两类包必须分开不能放在一起。本地包或者第三方包内部可以继续按实际情况细分不同子类。
* 不要使用相对路径引入包
* 【可选】匿名包的引用建议使用一个新的分组引入，并在匿名包上写上注释说明。

```shell
package demo

import (
	// standard package
	"encoding/json"
	"strings"

	// third-party package
	"git.obc.im/obc/utils"
	"git.obc.im/dep/beego"
	"git.obc.im/dep/mysql"
	opentracing "github.com/opentracing/opentracing-go"

	// anonymous import package
	// import filesystem storage driver
	_ "test.oa.com/org/repo/pkg/storage/filesystem
)

```

## 2.5 【强制】错误处理

### 2.5.1 【强制】error 处理

* `error` 作为函数的值返回，必须对 `error` 进行处理, 或将返回值赋值给明确忽略。对于`defer xx.Close()`可以不用显式处理。
* `error` 作为函数的值返回且有多个返回值的时候，`error` 必须是最后一个参数。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
func do() (error, int) {

}
```

</td><td>

```shell
func do() (int, error) {

}
```

</td></tr>
</tbody></table>

* 不能使用 `_` 丢弃任何 `return` 的 `err`。若不进行错误处理，要么再次向上游 `return err`，或者使用 `log` 记录下来。
* 采用独立的错误流进行处理,尽早 `return err`，函数中优先进行 `return` 检测，遇见错误则马上 `return err`。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
if err != nil {
  // error handling
} else {
  // normal code
}
```

</td><td>

```shell
if err != nil {
  // error handling
  return err
}
// normal code

```

</td></tr>
</tbody></table>

* 错误提示（Error Strings）不需要大写字母开头的单词，即使是句子的首字母也不需要。除非那是个专有名词或者缩写。同时，错误提示也不需要以句号结尾，因为通常在打印完错误提示后还需要跟随别的提示信息

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
fmt.Errorf("Create user err,uid conflict.")
```

</td><td>

```shell
fmt.Errorf("create user err,uid conflict")
```

</td></tr>
</tbody></table>

* 错误返回的判断独立处理，不与其他变量组合逻辑判断。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
x, y, err := f()
if err != nil || y == nil {
  return err // 当y与err都为空时，函数的调用者会出现错误的调用逻辑
}
```

</td><td>

```shell
x, y, err := f()
if err != nil {
  return err
}
if y == nil {
  return fmt.Errorf("some error")
}
```

</td></tr>
</tbody></table>

### 2.5.2 【强制】panic 处理

* 在业务逻辑处理中禁止使用 `panic`。
* 在main包中只有当完全不可运行的情况可使用 `panic`，例如：文件无法打开，数据库无法连接导致程序无法正常运行。
* 【建议】在 main 包中使用 `logutil.Fatal` 来记录错误，这样就可以由 `log` 来结束程序，或者将 `panic` 抛出的异常记录到日志文件中，方便排查问题。
* `panic` 捕获只能到 `goroutine` 最顶层，每个自行启动的 `goroutine`，必须在入口处捕获 `panic`，并打印详细堆栈信息或进行其它处理。
* 对于需要异步处理的逻辑,必须使用 `threadutil.GoSafe()`进行包装,避免野生协程的`panic`导致整个主进程`panic`

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
go foo()
```

</td><td>

```shell
// import "github.com/ikaiguang/go-srv-kit/kratos/thread"
threadpkg.GoSafe(func () {
    foo()
})
```

</td></tr>
</tbody></table>

## 2.6 【强制】单元测试

* 单元测试文件名命名规范为 `example_test.go`。
* 测试用例的函数名称必须以 `Test` 开头，例如 `TestExample`。
* 如果存在 `func Foo`，单测函数可以带下划线，为 `func Test_Foo`。如果存在 func (b *Bar)
  Foo，单测函数可以为 `func TestBar_Foo`。下划线不能出现在前面描述情况以外的位置。
* 单测文件行数限制是普通文件的2倍，即1600行。单测函数行数限制也是普通函数的2倍，即为160行。圈复杂度、列数限制、 import
  分组等其他规范细节和普通文件保持一致。
* 由于单测文件内的函数都是不对外的，所有可导出函数可以没有注释，但是结构体定义时尽量不要导出。
* 每个重要的可导出函数都要首先编写测试用例，测试用例和正规代码一起提交方便进行回归测试。

## 2.7 【强制】类型断言失败处理

* `type assertion` 的单个返回值形式针对不正确的类型将产生 `panic`。因此，请始终使用 `comma ok` 的惯用法。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
t := i.(string)
```

</td><td>

```shell
t, ok := i.(string)
if !ok {
  // 优雅地处理错误
}
```

</td></tr>
</tbody></table>

# 3. 注释

1. 在编码阶段同步写好变量、函数、包注释，注释可以通过 godoc 导出生成文档。
2. 程序中每一个被导出的(大写的)名字，都应该有一个文档注释。
3. 【强制】所有注释掉的代码在提交 code review 前都应该被删除，除非添加注释讲解为什么不删除， 并且标明后续处理建议（比如删除计划）。
4. 【强制】所有没有被使用的业务相关函数必须删除(一些工具函数,包的预留方法除外),除非注明下次启用的时机

## 3.1 【建议】包注释

* 每个包都应该有一个包注释,尤其是一些工具包
* 包如果有多个 go 文件，只需要出现在一个 go 文件中（一般是和包同名的文件）即可，格式为：“// Package 包名 包信息描述”。

```go
// Package math provides basic constants and mathematical functions.
package math

// 或者

/*
Package template implements data-driven templates for generating textual
output such as HTML.
....
*/
package template

```

## 3.2 【建议】结构体注释

* 每个需要导出的自定义结构体或者接口都应该有注释说明。
* 注释对结构进行简要介绍，放在结构体定义的前一行。
* 格式为："// 结构体名 结构体信息描述"。
* 结构体内的可导出成员变量名，如果是个生僻词，或者意义不明确的词，要给出注释，放在成员变量的同一行的末尾。

```shell
// User 用户结构定义了用户基础信息
type User struct {
  Name  string
  Email string
  Demographic string // 族群
}

```

## 3.3 【建议】方法注释

* 每个需要导出的函数或者方法（结构体或者接口下的函数称为方法）都应该有注释。注意，如果方法的接收器为不可导出类型，可以不注释，但需要质疑该方法可导出的必要性。
* 注释描述函数或方法功能、调用方等信息。
* 格式为："// 函数名 函数信息描述"。

```go
// NewtAttrModel 是属性数据层操作类的工厂方法
func NewAttrModel(ctx *common.Context) *AttrModel {
// code here
}
```

## 3.4 【建议】变量和常量注释

* 每个需要导出的常量和变量都应该有注释说明。
* 该注释对常量或变量进行简要介绍，放在常量或者变量定义的前一行。
* 大块常量或变量定义时，可在前面注释一个总的说明，然后每一行常量的末尾详细注释该常量的定义。
* 格式为："// 变量名 变量信息描述"，斜线后面紧跟一个空格。

```shell
// FlagConfigFile 配置文件的命令行参数名
const FlagConfigFile = "--config"

// 命令行参数
const (
  FlagConfigFile1 = "--config" // 配置文件的命令行参数名1
  FlagConfigFile2 = "--config" // 配置文件的命令行参数名2
  FlagConfigFile3 = "--config" // 配置文件的命令行参数名3
  FlagConfigFile4 = "--config" // 配置文件的命令行参数名4
)

// FullName 返回指定用户名的完整名称
var FullName = func (username string) string {
  return fmt.Sprintf("fake-%s", username)
}

```

## 3.5 空行

* 空行需要体现代码逻辑的关联，所以空行不能随意，非常严重地影响可读性。
* 保持函数内部实现的组织粒度是相近的，用空行分隔。

# 4. 命名规范

命名是代码规范中很重要的一部分，统一的命名规范有利于提高代码的可读性，好的命名仅仅通过命名就可以获取到足够多的信息。

* goland 中可以通过开启拼写检查来校验单词正确性,开启方式: `Preferences | Editor | Inspections| Spelling`

## 4.1 【建议】包命名

* 保持 package 的名字和目录一致。
* 尽量采取有意义、简短的包名，尽量不要和标准库冲突。
* 包名应该为小写单词，不要使用下划线或者混合大小写，,不使用复数,使用多级目录来划分层级。
* 项目名可以通过中划线来连接多个单词。
* 简单明了的包命名，如：time、list、http。
*

不要使用无意义的包名，如：util、common、misc、global。package名字应该追求清晰且越来越收敛，符合‘单一职责’原则。而不是像common一样，什么都能往里面放，越来越膨胀，让依赖关系变得复杂，不利于阅读、复用、重构。注意，xx/util/encryption这样的包名是允许的。

* 包名和目录名保持一致。一个目录尽量维护一个包下的所有文件。

## 4.2 【强制】文件命名

* 采用有意义，简短的文件名。
* 文件名应该采用小写，并且使用下划线分割各个单词。

## 4.3 【强制】结构体命名

* 采用驼峰命名方式，首字母根据访问控制采用大写或者小写。
* 结构体名应该是名词或名词短语，如 `Customer`、`WikiPage`、`Account`、`AddressParser`，它不应是动词。
* 避免使用 `Data`、`Info` 这类意义太宽泛的结构体名。
* 【建议】结构体初始化格式采用多行,指定key的形式,尽量避免匿名初始化

```shell
// User 多行声明
type User struct {
  Name  string
  Email string
}
```

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
// 避免匿名初始化
u1 := User{"jack","jack@gmail"}

// 无法使用Goland的`fill field`功能,且容易遗漏
u2 := User{}
u.Name = "jack"
u.Email = "jack@gmail.com"


```

</td><td>

```shell
u1 := User{
  Name : "jack"
  Email : "jack@gmail.com" 
}
```

</td></tr>
</tbody></table>

## 4.4 【建议】接口命名

* 命名规则基本保持和结构体命名规则一致。
* 单个函数的接口名以 `er` 作为后缀，例如 `Reader`，`Writer`。

```shell
// Reader 字节数组读取接口
type Reader interface {
  // Read 读取整个给定的字节数据并返回读取的长度
  Read(p []byte) (n int, err error)
}

```

* 两个函数的接口名综合两个函数名。
* 三个以上函数的接口名，类似于结构体名。

```shell
// Car 小汽车结构声明
type Car interface {
    // Start ...
    Start([]byte)
    // Stop ...
    Stop() error
    // Recover ...
    Recover()
}
```

## 4.5 【强制】变量命名

* 变量名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写
* 若变量类型为 bool 类型，则名称应以 `Has`，`Is`，`Can` 或者 `Allow` 开头。
* 私有全局变量和局部变量规范一致，均以小写字母开头。
* 代码生成工具自动生成的代码可排除此规则（如 xxx.pb.go 里面的 Id）。
* 变量名更倾向于选择短命名。特别是对于局部变量。 `c`比`lineCount`要好，`i`比`sliceIndex`
  要好。基本原则是：变量的使用和声明的位置越远，变量名就需要具备越强的描述性。

## 4.6 【强制】常量命名

* `Golang`常量均需遵循驼峰式；`Protobuf`常量定义需遵循全大写。

```shell
// AppVersion 应用程序版本号定义
const AppVersion = "1.0.0"
```

* 如果是枚举类型的常量，需要先创建相应类型：

```shell
// Scheme 传输协议
type Scheme string

const (
  // HTTP 表示HTTP明文传输协议
  HTTP Scheme = "http"
  // HTTPS 表示HTTPS加密传输协议
  HTTPS Scheme = "https"
)

```

```protobuf
// RoleTypeEnum explorer 文件权限角色类型
message RoleTypeEnum {
  // RoleType 角色类型
  enum RoleType {
    OWNER = 0;
    MANAGER = 1;
    EDITOR = 2;
    READER = 3;
  }
}
```

* 私有全局常量和局部变量规范一致，均以小写字母开头。

```shell
const appVersion = "1.0.0"

```

## 4.7 【强制】函数命名

函数名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写。

# 5. 控制结构

* if 接受初始化语句，约定如下方式建立局部变量：

```shell
if err := file.Chmod(0664); err != nil {
    return err
}
```

* if 对两个值进行判断时，约定如下顺序：变量在左，常量在右：

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
if nil != err {
  // error handling
}

if 0 == errorCode {
  // do something
}


```

</td><td>

```shell
if err != nil {
  // error handling
}

if errorCode == 0 {
  // do something
}
```

</td></tr>
</tbody></table>

* if 对于bool类型的变量，应直接进行真假判断：

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
var allowUserLogin bool
if allowUserLogin == true {
  // do something
}

if allowUserLogin == false {
  // do something
}
```

</td><td>

```shell
if allowUserLogin {
  // do something
}

if !allowUserLogin {
  // do something
}
```

</td></tr>
</tbody></table>

# 6. 函数

## 6.1 【建议】函数参数

函数返回相同类型的两个或三个参数，或者如果从上下文中不清楚结果的含义，使用命名返回，其它情况不建议使用命名返回。

```shell
// Parent1 ...
func (n *Node) Parent1() *Node

// Parent2 ...
func (n *Node) Parent2() (*Node, error)

// Location ...
func (f *Foo) Location() (lat, long float64, err error)

```

* 传入变量和返回变量以小写字母开头。
* 参数数量均不能超过3个。
* 尽量用值传递，非指针传递。
* 传入参数是 map，slice，chan，interface 不要传递指针。
* 尽可能将同种类型的参数放在相邻位置，则只需写一次类型。
* 函数、方法的顺序一般需要按照依赖关系由浅入深由上至下排序，即最底层的函数出现在最前面。例如，函数 ExecCmdDirBytes
  属于最底层的函数，它被 ExecCmdDir 函数调用，而 ExecCmdDir 又被 ExecCmd 调用。

## 6.2 【强制】defer

* 当存在资源管理时，应紧跟 defer 函数进行资源的释放。
* 判断是否有错误发生之后，再 defer 释放资源。

```shell
resp, err := http.Get(url)
if err != nil {
  return err
}
// 如果操作成功，再defer Close()
defer resp.Body.Close()

```

## 6.3 【建议】方法的接收器

* 【建议】建议以类名第一个英文首字母的小写作为接收器的命名。
* 【建议】接收器的命名在函数超过20行的时候不要用单字符。
* 【强制】命名不能采用 `me`，`this`，`self` 这类易混淆名称。

## 6.4 【建议】变量声明

* 变量声明尽量放在变量第一次使用前面，就近原则。

## 6.5 【强制】魔法数字

* 如果魔法数字出现超过2次，则禁止使用。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```shell
func getArea(r float64) float64 {
  return 3.14 * r * r
}

func getLength(r float64) float64 {
  return 3.14 * 2 * r
}
```

</td><td>

```shell
// PI ...
const PI = 3.14

func getArea(r float64) float64 {
  return PI * r * r
}

func getLength(r float64) float64 {
  return PI * 2 * r
}
```

</td></tr>
</tbody></table>

# 7 proto文件规范

proto文件规范参照官方的[proto规范](https://developers.google.com/protocol-buffers/docs/style)

## 7.1 【强制】proto 布局

```shell
// enums
// errors
// resources
// services
```

* rpc 的body中若无内容,需以`{};`结尾,参考 `rpc Login`
* 请求和响应的`message` 成对出现,根据rpc方法出现顺序依次排列
* 通用的`enums`和`message`放在service下方,根据rpc出现顺序依次排列
* 只要每次新增内容都按照以上要求进行追加,可以很大程度避免代码冲突的问题,找内容也有迹可循

```protobuf
syntax = "proto3";

package service.api.user.servicev1;

option go_package = "github.com/ikaiguang/go-srv-saas/api/user/v1/services;servicev1";

// 1. 先定义rpc方法
service UserService {
    // 用户登入接口
    rpc Login (LoginRequest) returns (LoginReply){};
    // 用户登出接口
    rpc Logout (LogoutRequest) returns (LogoutReply){};
}

// 2. 再定义枚举
// UserStatusEnum 用户状态
message UserStatusEnum {
  // UserStatus 枚举值
  enum UserStatus {
    // UNSPECIFIED 未指定
    UNSPECIFIED = 0;

    // INACTIVATED 未激活
    INACTIVATED = 1;
    // ENABLE 启用
    ENABLE = 2;
    // DISABLE 禁用
    DISABLE = 3;
    // BLACKLIST 黑名单
    BLACKLIST = 4;
    // DELETED 已删除
    DELETED = 5;
  }
}

// 3. 定义通用数据结构
message UserEntity {
    int64 uid = 1;
    string username = 2;
    string avatar = 4;
    string email = 5;
    int64 create_time = 6;
    int64 update_time = 7;
}

message Token {
    string jwt = 1;
    string refresh_token = 2;
}

// 4. 定义请求/响应
message LoginRequest {
    string account = 2;
    string password = 3;
}
message LoginReply {
    Token token = 1;
}

message LogoutRequest {
    int64 uid = 1;
}
message LogoutReply{

}
```

## 7.2【强制】枚举(Enum)

* 枚举首字母大写,枚举值全部大写,用`_`分隔

```protobuf
message FooBar {
  enum FooBarEnum {
    UNSPECIFIED = 0;
    FIRST_VALUE = 1;
    SECOND_VALUE = 2;
  }
}
```

## 7.3【强制】字段名(Field)

* 字段名以`camel_case`形式命名,与返回给前端的风格保持统一

```protobuf
message CopyFileReply{
    optional string song_name = 1;
}
```

## 7.4【强制】消息(Message)

* message采用首字母大写的形式

```protobuf
message SongServerRequest {
    optional string song_name = 1;
}
```

## 7.5【强制】列表(Repeated fields)

* 列表字段命名采用复数形式,见名知其意

```shell
repeated string keys = 1;
```

## 7.6【强制】注释(Comment)

* 注释写在rpc或message的上方,格式为 `// field comment`

```shell
// 登录凭证
message Token {
  // access_token json web token,短时间内有效
  string access_token = 1;
  // refresh_token 刷新token,长时间有效
  string refresh_token = 2;
}
```