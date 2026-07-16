# go-srv-kit v3 多 Module 开发指南

子 module 的实际发布命令见[《v3 子 Module 发布指南》](module_release.md)。

## 1. 目标与结论

本仓库现有 17 个独立 Go modules。v3 采用以下方案：

- 当前分支直接作为 v3 起点，不维护本仓库路径下的旧 v2。
- 保持现有物理目录，不为每个 module 新建 `v3/` 子目录。
- 每个 module 独立增加 `/v3` 语义导入版本后缀。
- module 内的包路径把 `/v3` 放在 module 根路径之后、包子目录之前。
- 本地联调使用 Go workspace；发布用的 `go.mod` 不提交本地路径 `replace`。
- 首次发布统一使用 `v3.0.0`，后续各 module 可以独立演进。

关键示例：

```text
物理目录: auth/
module:   github.com/ikaiguang/go-srv-kit/auth/v3
package:  github.com/ikaiguang/go-srv-kit/auth/v3/auth
tag:      auth/v3.0.0

物理目录: data/mysql/
module:   github.com/ikaiguang/go-srv-kit/data/mysql/v3
package:  github.com/ikaiguang/go-srv-kit/data/mysql/v3/mysql
tag:      data/mysql/v3.0.0
```

不要写成：

```text
github.com/ikaiguang/go-srv-kit/v3/auth
github.com/ikaiguang/go-srv-kit/data/mysql/mysql/v3
```

原因是 `auth` 和 `data/mysql` 各自有 `go.mod`，它们是独立 module。`/v3` 必须追加在各自 module path 末尾。

## 2. Module 路径与 Tag 映射

| 目录 | 当前 module | v3 module | 首次发布 tag |
| --- | --- | --- | --- |
| `.` | `github.com/ikaiguang/go-srv-kit` | `github.com/ikaiguang/go-srv-kit/v3` | `v3.0.0` |
| `auth` | `github.com/ikaiguang/go-srv-kit/auth` | `github.com/ikaiguang/go-srv-kit/auth/v3` | `auth/v3.0.0` |
| `data/consul` | `github.com/ikaiguang/go-srv-kit/data/consul` | `github.com/ikaiguang/go-srv-kit/data/consul/v3` | `data/consul/v3.0.0` |
| `data/etcd` | `github.com/ikaiguang/go-srv-kit/data/etcd` | `github.com/ikaiguang/go-srv-kit/data/etcd/v3` | `data/etcd/v3.0.0` |
| `data/gorm` | `github.com/ikaiguang/go-srv-kit/data/gorm` | `github.com/ikaiguang/go-srv-kit/data/gorm/v3` | `data/gorm/v3.0.0` |
| `data/jaeger` | `github.com/ikaiguang/go-srv-kit/data/jaeger` | `github.com/ikaiguang/go-srv-kit/data/jaeger/v3` | `data/jaeger/v3.0.0` |
| `data/mongo` | `github.com/ikaiguang/go-srv-kit/data/mongo` | `github.com/ikaiguang/go-srv-kit/data/mongo/v3` | `data/mongo/v3.0.0` |
| `data/mysql` | `github.com/ikaiguang/go-srv-kit/data/mysql` | `github.com/ikaiguang/go-srv-kit/data/mysql/v3` | `data/mysql/v3.0.0` |
| `data/postgres` | `github.com/ikaiguang/go-srv-kit/data/postgres` | `github.com/ikaiguang/go-srv-kit/data/postgres/v3` | `data/postgres/v3.0.0` |
| `data/rabbitmq` | `github.com/ikaiguang/go-srv-kit/data/rabbitmq` | `github.com/ikaiguang/go-srv-kit/data/rabbitmq/v3` | `data/rabbitmq/v3.0.0` |
| `data/redis` | `github.com/ikaiguang/go-srv-kit/data/redis` | `github.com/ikaiguang/go-srv-kit/data/redis/v3` | `data/redis/v3.0.0` |
| `kit` | `github.com/ikaiguang/go-srv-kit/kit` | `github.com/ikaiguang/go-srv-kit/kit/v3` | `kit/v3.0.0` |
| `kratos` | `github.com/ikaiguang/go-srv-kit/kratos` | `github.com/ikaiguang/go-srv-kit/kratos/v3` | `kratos/v3.0.0` |
| `ping-service` | `github.com/ikaiguang/go-srv-kit/ping-service` | `github.com/ikaiguang/go-srv-kit/ping-service/v3` | `ping-service/v3.0.0` |
| `registry/consul` | `github.com/ikaiguang/go-srv-kit/registry/consul` | `github.com/ikaiguang/go-srv-kit/registry/consul/v3` | `registry/consul/v3.0.0` |
| `registry/etcd` | `github.com/ikaiguang/go-srv-kit/registry/etcd` | `github.com/ikaiguang/go-srv-kit/registry/etcd/v3` | `registry/etcd/v3.0.0` |
| `service` | `github.com/ikaiguang/go-srv-kit/service` | `github.com/ikaiguang/go-srv-kit/service/v3` | `service/v3.0.0` |

根 module 的 tag 直接使用 `v3.0.0`。子目录 module 的 tag 必须带物理目录前缀。例如 `data/mysql` 的 tag 是 `data/mysql/v3.0.0`，不是 `v3.0.0`，也不是 `data/mysql/v3/v3.0.0`。

## 3. Go Import 路径规则

module path 增加 `/v3` 后，module 内所有跨包 import 都必须同步更新。

```go
import (
	filepkg "github.com/ikaiguang/go-srv-kit/kit/v3/file"
	authpkg "github.com/ikaiguang/go-srv-kit/auth/v3/auth"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm/v3/gorm"
	mysqlpkg "github.com/ikaiguang/go-srv-kit/data/mysql/v3/mysql"
	clientutil "github.com/ikaiguang/go-srv-kit/service/v3/cluster_service_api"
)
```

同一个 module 内跨目录引用也要带 `/v3`。例如 `kit/zip` 引用 `kit/file` 时，应从：

```go
github.com/ikaiguang/go-srv-kit/kit/file
```

改为：

```go
github.com/ikaiguang/go-srv-kit/kit/v3/file
```

### 旧仓库路径迁移表

当前源码仍包含拆分仓库时期的 import。迁移时统一替换为本仓库 v3 module：

| 旧路径根 | v3 路径根 |
| --- | --- |
| `github.com/ikaiguang/go-auth-kit` | `github.com/ikaiguang/go-srv-kit/auth/v3` |
| `github.com/ikaiguang/go-consul-kit` | `github.com/ikaiguang/go-srv-kit/data/consul/v3` |
| `github.com/ikaiguang/go-etcd-kit` | `github.com/ikaiguang/go-srv-kit/data/etcd/v3` |
| `github.com/ikaiguang/go-gorm-kit` | `github.com/ikaiguang/go-srv-kit/data/gorm/v3` |
| `github.com/ikaiguang/go-jaeger-kit` | `github.com/ikaiguang/go-srv-kit/data/jaeger/v3` |
| `github.com/ikaiguang/go-mongo-kit` | `github.com/ikaiguang/go-srv-kit/data/mongo/v3` |
| `github.com/ikaiguang/go-mysql-kit` | `github.com/ikaiguang/go-srv-kit/data/mysql/v3` |
| `github.com/ikaiguang/go-postgres-kit` | `github.com/ikaiguang/go-srv-kit/data/postgres/v3` |
| `github.com/ikaiguang/go-rabbitmq-kit` | `github.com/ikaiguang/go-srv-kit/data/rabbitmq/v3` |
| `github.com/ikaiguang/go-redis-kit` | `github.com/ikaiguang/go-srv-kit/data/redis/v3` |
| `github.com/ikaiguang/go-kit` | `github.com/ikaiguang/go-srv-kit/kit/v3` |
| `github.com/ikaiguang/go-kratos-kit` | `github.com/ikaiguang/go-srv-kit/kratos/v3` |
| `github.com/ikaiguang/go-service-kit` | `github.com/ikaiguang/go-srv-kit/service/v3` |
| `github.com/ikaiguang/kratos-consul-kit` | `github.com/ikaiguang/go-srv-kit/registry/consul/v3` |
| `github.com/ikaiguang/kratos-etcd-kit` | `github.com/ikaiguang/go-srv-kit/registry/etcd/v3` |

替换后还要扫描已经合并到 `go-srv-kit`、但尚未带 `/v3` 的内部 import。

## 4. Proto `go_package` 规则

Proto 的 `go_package` 必须使用 v3 的实际 Go package 路径，并保留分号后的 Go package alias。

```proto
// auth/auth/auth.kit.proto
option go_package = "github.com/ikaiguang/go-srv-kit/auth/v3/auth;authpkg";

// data/mysql/mysql/config.proto
option go_package = "github.com/ikaiguang/go-srv-kit/data/mysql/v3/mysql;mysqlpkg";

// kit/page/page.kit.proto
option go_package = "github.com/ikaiguang/go-srv-kit/kit/v3/page;pagepkg";

// service/api/config/config.proto
option go_package = "github.com/ikaiguang/go-srv-kit/service/v3/api/config;configpb";
```

修改 Proto 后必须运行对应生成器。不要手改 `*.pb.go`、`*.pb.validate.go`、`*_errors.pb.go` 或其他生成文件。

## 5. 本地多 Module 开发

### 5.1 创建本地 workspace

在仓库根目录创建 `go.work`：

```bash
go work init \
  . \
  ./auth \
  ./data/consul \
  ./data/etcd \
  ./data/gorm \
  ./data/jaeger \
  ./data/mongo \
  ./data/mysql \
  ./data/postgres \
  ./data/rabbitmq \
  ./data/redis \
  ./kit \
  ./kratos \
  ./ping-service \
  ./registry/consul \
  ./registry/etcd \
  ./service
```

检查 workspace：

```bash
go env GOWORK
go work edit -json
```

本地 workspace 会让这些 module 直接使用当前工作区源码，不需要在每个 `go.mod` 写本地 `replace`。

本指南建议把 `go.work` 和 `go.work.sum` 作为本地开发文件，不纳入 v3 module 的发布契约。无论是否提交 workspace 文件，发布前都必须使用 `GOWORK=off` 单独验证每个 module，避免 workspace 掩盖缺失的 `require` 或错误版本。

### 5.2 v3 内部依赖

下游 module 的 `go.mod` 要求真实的 v3 module path 和合法的 v3 版本。例如：

```go.mod
require (
	github.com/ikaiguang/go-srv-kit/kit/v3 v3.0.0
	github.com/ikaiguang/go-srv-kit/kratos/v3 v3.0.0
	github.com/ikaiguang/go-srv-kit/data/redis/v3 v3.0.0
)
```

在这些 tag 尚未发布时，workspace 会使用本地 module。不要为了本地联调把以下内容提交到正式 `go.mod`：

```go.mod
replace github.com/ikaiguang/go-srv-kit/kit/v3 => ../kit
```

本地路径 `replace` 对其他开发者、CI 和使用方不可移植。

## 6. 从哪个 Module 开始

先改 `kit`。它是最多内部 module 依赖的底层工具 module；从根 module 或 `service` 开始不能解除下游依赖。

推荐按以下阶段做“纵向闭环”：每个 module 都完成 module path、Go imports、Proto、生成代码、README、测试，再进入下一阶段。

### 阶段 1：底层与无内部依赖 module

1. `kit`
2. `data/consul`
3. `data/etcd`
4. `data/mongo`
5. `registry/consul`
6. `registry/etcd`

其中 `kit` 应最先完成。其余无本仓库内部依赖的 module 可以并行处理。

### 阶段 2：依赖 `kit` 的基础 module

1. `data/gorm`
2. `data/jaeger`
3. `data/rabbitmq`
4. `data/redis`
5. `kratos`

### 阶段 3：组合 module

1. `data/mysql`：依赖 `data/gorm` 和 `kit`
2. `data/postgres`：依赖 `data/gorm` 和 `kit`
3. `auth`：依赖 `kit`、`kratos` 和 `data/redis`

### 阶段 4：聚合服务 module

处理 `service`。它依赖认证、数据、注册中心、Kratos 扩展和工具库，应在所有基础 module 完成后迁移。

### 阶段 5：根 module 与示例

最后处理：

1. 根 module `github.com/ikaiguang/go-srv-kit/v3`
2. `ping-service`
3. 根命令、示例、README 和最终集成验证

根 module 当前没有内部依赖，技术上可以提前改 module path，但它不应作为迁移起点，也不应早于底层 modules 发布。

## 7. 单个 Module 的操作模板

以 `kit` 为例：

1. 修改 `kit/go.mod`：

   ```go.mod
   module github.com/ikaiguang/go-srv-kit/kit/v3

   go 1.26.3
   ```

2. 将 `kit` module 内的自引用 import 改为 `github.com/ikaiguang/go-srv-kit/kit/v3/...`。
3. 修改 `kit` 下 Proto 的 `option go_package`。
4. 运行对应 Proto 生成器。
5. 更新 README 中的 `go get` 和 import 示例。
6. 在 `kit` 目录执行：

   ```bash
   gofmt -w <本次修改的 Go 源文件>
   go mod tidy
   go test ./...
   go vet ./...
   ```

7. 扫描旧路径：

   ```bash
   rg 'github\.com/ikaiguang/(go-kit|go-srv-kit/kit(?!/v3))' kit --pcre2
   ```

8. workspace 下通过后，再进行 `GOWORK=off` 发布验证。

其他 module 使用相同步骤，只替换 module 根路径和测试目录。

## 8. 其他仓库如何接入

### 8.1 v3 尚未发布：跨仓库本地联调

推荐在使用方仓库创建自己的 workspace，将使用方 module 和本仓库所需 modules 一起加入：

```bash
cd /path/to/consumer
go work init .
go work use \
  /path/to/go-srv-kit/kit \
  /path/to/go-srv-kit/kratos \
  /path/to/go-srv-kit/service
```

如果 `service` 的内部 v3 依赖尚未发布，还要把它依赖的其他本地 modules 加入 workspace。只加入 `service` 并不能替代所有未发布的传递依赖。

也可以在使用方 `go.mod` 临时添加 `replace`，但每个未发布的内部 module 都需要单独映射：

```go.mod
replace github.com/ikaiguang/go-srv-kit/kit/v3 => /path/to/go-srv-kit/kit
replace github.com/ikaiguang/go-srv-kit/kratos/v3 => /path/to/go-srv-kit/kratos
replace github.com/ikaiguang/go-srv-kit/service/v3 => /path/to/go-srv-kit/service
```

这些本地绝对路径不要提交。

### 8.2 v3 已发布：正常依赖

使用方按所需 module 获取版本：

```bash
go get github.com/ikaiguang/go-srv-kit/kit/v3@v3.0.0
go get github.com/ikaiguang/go-srv-kit/service/v3@v3.0.0
go mod tidy
```

代码使用 v3 import：

```go
import (
	filepkg "github.com/ikaiguang/go-srv-kit/kit/v3/file"
	serverutil "github.com/ikaiguang/go-srv-kit/service/v3/server"
)
```

Go 可以同时解析旧 module path 和 `/v3` module path，因此使用方可以逐包迁移；但同一业务边界内长期混用两套基础类型可能产生类型不兼容，应尽快统一。

## 9. 发布与 Tag 顺序

### 9.1 先发布候选版本

首次迁移建议先发布 `v3.0.0-rc.1`，让其他仓库完成集成验证，再发布 `v3.0.0`。

Tag 示例：

```bash
git tag -a kit/v3.0.0-rc.1 -m "kit v3.0.0-rc.1"
git tag -a data/gorm/v3.0.0-rc.1 -m "data/gorm v3.0.0-rc.1"
git tag -a service/v3.0.0-rc.1 -m "service v3.0.0-rc.1"
git tag -a v3.0.0-rc.1 -m "go-srv-kit v3.0.0-rc.1"
```

### 9.2 按依赖顺序发布

1. `kit` 和无内部依赖 modules。
2. `data/gorm`、`data/jaeger`、`data/rabbitmq`、`data/redis`、`kratos`。
3. `data/mysql`、`data/postgres`、`auth`。
4. `service`。
5. `ping-service` 和根 module。

每一阶段发布后，下游 `go.mod` 才能在 `GOWORK=off` 模式下引用真实版本并完成 `go mod tidy`。

稳定版 tag 使用相同顺序，将版本改为 `v3.0.0`。

### 9.3 发布后验证

```bash
GOWORK=off go list -m github.com/ikaiguang/go-srv-kit/kit/v3@v3.0.0
GOWORK=off go list -m github.com/ikaiguang/go-srv-kit/service/v3@v3.0.0
```

也应在一个仓库外的临时 consumer module 中执行 `go get` 和最小编译测试，避免 workspace 或本地缓存掩盖发布问题。

## 10. 验证清单

### 10.1 路径扫描

迁移完成后，不应再出现旧拆分仓库路径：

```bash
rg 'github\.com/ikaiguang/(go-auth-kit|go-consul-kit|go-etcd-kit|go-gorm-kit|go-jaeger-kit|go-mongo-kit|go-mysql-kit|go-postgres-kit|go-rabbitmq-kit|go-redis-kit|go-kit|go-kratos-kit|go-service-kit|kratos-consul-kit|kratos-etcd-kit)' \
  -g '*.go' -g '*.proto' -g 'go.mod' -g 'README.md' \
  -g '!third_party/**'
```

所有新 module path 应带 `/v3`：

```bash
rg '^module ' -g 'go.mod'
rg '^option go_package' -g '*.proto' -g '!third_party/**'
```

### 10.2 全 module 测试

不要假设在仓库根目录执行一次 `go test ./...` 就覆盖所有嵌套 modules。逐 module 运行：

```bash
for module in \
  . auth \
  data/consul data/etcd data/gorm data/jaeger data/mongo \
  data/mysql data/postgres data/rabbitmq data/redis \
  kit kratos ping-service registry/consul registry/etcd service
do
  (cd "$module" && go test ./...)
done
```

发布前关闭 workspace 再逐 module 执行同样的测试。此时所有内部 v3 版本必须已经可以从版本库解析：

```bash
for module in \
  . auth \
  data/consul data/etcd data/gorm data/jaeger data/mongo \
  data/mysql data/postgres data/rabbitmq data/redis \
  kit kratos ping-service registry/consul registry/etcd service
do
  (cd "$module" && GOWORK=off go test ./...)
done
```

### 10.3 最终检查

- 17 个 `go.mod` 的 module path 都以 `/v3` 结尾。
- 所有内部 Go imports 的 `/v3` 位置正确。
- 所有非 `third_party` Proto 的 `go_package` 与 module 和包目录一致。
- 生成文件由生成器更新，没有手工编辑。
- README 的安装命令和 import 示例使用 v3 路径。
- 正式 `go.mod` 不包含本地文件系统 `replace`。
- workspace 模式和 `GOWORK=off` 模式都完成相应验证。
- 子 module tag 使用目录前缀，根 module tag 不使用目录前缀。
- 至少一个外部 consumer 完成真实 `go get` 和编译测试。

## 11. 常见错误

### 把 `/v3` 放在仓库根后

错误：

```text
github.com/ikaiguang/go-srv-kit/v3/kit/file
```

正确：

```text
github.com/ikaiguang/go-srv-kit/kit/v3/file
```

因为 `kit` 是独立 module。

### 只修改 `go.mod`

module path 变化后，Go imports、Proto `go_package`、生成文件、README 和使用方都必须同步修改。

### 只打 `v3.0.0` tag

v2 及以上版本必须在 module path 中包含 `/vN`。只有 tag、没有 `/v3` module path，不是有效的 v3 迁移方案。

### 子 module 使用根 tag

`service` module 必须使用 `service/v3.0.0`，根 tag `v3.0.0` 不会发布 `service` module。

### 提交本地 `replace`

指向 `../kit` 或绝对路径的 `replace` 只适合本地联调，提交后会破坏 CI 和其他仓库使用。

### 只在 workspace 模式测试

workspace 会优先使用本地源码，可能掩盖缺失的 `require`、错误 tag 或未发布版本。发布前必须使用 `GOWORK=off` 验证。

### 先发布 `service`

`service` 是聚合层。它依赖的 v3 modules 尚未发布时，外部使用方无法解析完整依赖图。

## 12. 官方参考

- [Go Modules: Developing a major version update](https://go.dev/doc/modules/major-version)
- [Go Modules Reference: Major version suffixes](https://go.dev/ref/mod#major-version-suffixes)
- [Go Modules Reference: Workspaces](https://go.dev/ref/mod#workspaces)
- [Go Modules Reference: Versions and VCS tags](https://go.dev/ref/mod#vcs-version)
