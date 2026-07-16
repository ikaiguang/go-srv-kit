# Proto v3 与 kit 发布实施计划

> Proto 命名调整由用户自行完成，本轮不再执行；本计划继续用于记录 tag 自动化、`kit` 审查和发布交接。

**目标：** 增加子 module tag 自动化，并把 `kit` 修复到可发布 `kit/v3.0.0` 的状态。

**实施方式：** tag 工具根据静态 module 清单计算子目录 tag。正确性审查严格限定在 `kit` module，不运行其他 module 的 Go 测试。

**技术栈：** Go 1.26.3、Bash、Git、Go Modules。

---

### 任务 1：Proto v3 namespace

**文件：**
- 用户自行处理非 `third_party/` 的 `*.proto` 和对应生成文件

- [x] 用户已自行处理；本轮不再修改或验证全项目 Proto 命名。

### 任务 2：增加 module tag 自动化

**文件：**
- 新增：`devops/module-release/modules.tsv`
- 新增：`devops/module-release/module-release.sh`
- 新增：`docs/v3_development/module_release.md`
- 修改：`docs/v3_development/README.md`

- [x] 在 `modules.tsv` 中登记物理目录、v3 module path 和 tag 前缀。
- [x] 实现 `list`、`check`、`tag`、`push` 子命令。
- [x] `check kit v3.0.0` 校验 `kit/go.mod`、版本格式、工作区和 tag 冲突。
- [x] `tag kit v3.0.0` 创建 annotated tag `kit/v3.0.0`。
- [x] `push kit v3.0.0` 只推送 `kit/v3.0.0`。
- [x] 编写中文指导文档，给出首次发布和后续独立演进命令。
- [x] 用临时 Git 仓库验证脚本，不在当前仓库创建 tag。

### 任务 3：审查并修复 kit

**文件：**
- 修改：`kit/` 下确认问题所需的文件
- 测试：受影响的 `kit` package 定向测试

- [x] 核对 `kit/go.mod`、内部 import、Proto 和 README 的 v3 路径。
- [x] 运行 `GOWORK=off go test ./...`、`go test -race ./...` 和 `go vet ./...` 建立基线。
- [x] 审查 panic、边界输入、错误处理、资源释放、并发、路径与压缩安全、跨平台行为和公开 API。
- [x] 对每个确认的问题先添加或定位失败用例，再做最小修复。
- [x] 对每个修复运行定向测试，再重跑完整发布 gate。

### 任务 4：发布交接

**文件：**
- 验证：任务 2 和任务 3 修改的文件

- [x] 运行 `kit` 旧路径扫描和 `git diff --check`；全项目 Proto 声明由用户自行处理。
- [x] 确认 `kit` 发布 gate 通过，并记录环境限制。
- [x] 列出用户应执行的 `git add`、`git commit`、分支 push 和 release 脚本命令，但不代为执行。
