# git

```shell
git config user.email "ckaiguang@outlook.com"
git config user.name "ckaiguang"
```

# 协作流程

**Git协作流程**

- 分支 dev : gitlab-ci 自动部署到开发环境。代码合并流程同 test 分支。
    - 开发工程师自测与验证
- 分支 test : gitlab-ci 自动部署到测试环境。
    - 开发工程师自测与验收
    - 测试工程师测试与验收
- 分支 pre : 设置保护分支并 gitlab-ci 自动部署到预发布环境。
    - 测试工程师再次验收
    - 产品经理验收
    - 客户演示
- 分支 prod(main) : 设置保护分支，仅允许技术负责人和运维进行合并MR，自动部署到线上环境。
    - prod分支每次发布，增加一次tag标签。用于记录版本发布和服务回滚。

```mermaid

%%---
%%title: Git 协作流程
%%---
gitGraph
    %% 分支 prod
    branch prod order: 20
    commit id: "新分支 prod" tag: "分支 prod"

    %% 分支 pre
    branch pre order: 30
    commit id: "checkout pre 分支" tag: "分支 pre"

    %% 分支 test
    checkout prod
    commit id: "prod 文档说明"
    branch test order: 40
    commit id: "checkout test 分支" tag: "分支 test"

    %% 新功能
    checkout prod
    commit id: "prod 部署脚本"
    commit
    branch feat/xxx order: 60
    commit id: "checkout feat/xxx_1 分支" tag: "新功能"
    commit id: "dev feat/xxx_1" tag:"开发中"

    %% 发现 Bug
    checkout prod
    branch hotfix/xxx order: 10
    commit id: "checkout fix/xxx_1" tag: "发现Bug"
    commit id: "fix fix/xxx_1" tag:"修复中"

    %% 完成新功能，发布到 test
    checkout feat/xxx
    commit id: "commit feat/xxx_1" tag:"完成新功能"
    commit id: "test feat/xxx_1" tag:"自测"
    checkout test
    merge feat/xxx id: "test merge feat/xxx_1" tag: "合并 feat/xxx_1"

    %% 修复 Bug，发布到 test
    checkout hotfix/xxx
    commit id: "commit fix/xxx_1" tag:"修复Bug"
    commit id: "test fix/xxx_1" tag:"自测"
    checkout test
    merge hotfix/xxx id: "test merge fix/xxx_1" tag: "合并 fix/xxx_1"

     %% 完成新功能，发布到 pre
    checkout feat/xxx
    commit
    checkout pre
    merge feat/xxx id: "pre merge feat/xxx_1" tag: "合并 feat/xxx_1"

     %% 完成新功能，发布到 prod
    checkout feat/xxx
    commit
    checkout prod
    merge feat/xxx id: "prod merge feat/xxx_1" tag: "合并 feat/xxx_1"

    %% 修复 Bug，发布到 pre
    checkout hotfix/xxx
    commit
    checkout pre
    merge hotfix/xxx id: "pre merge fix/xxx_1" tag: "合并 fix/xxx_1"

    %% 修复 Bug，发布到 prod
    checkout hotfix/xxx
    commit
    checkout prod
    merge hotfix/xxx id: "prod merge fix/xxx_1" tag: "合并 fix/xxx_1"

    %% 新功能 2
    checkout prod
    commit
    branch feat/xxx_2 order: 70
    commit id: "checkout feat/xxx_2 分支" tag: "新功能"
    commit id: "dev feat/xxx_2" tag:"开发中"

     %% 完成新功能 2，发布到 test
    checkout feat/xxx_2
    commit id: "commit feat/xxx_2" tag:"完成新功能"
    commit id: "test feat/xxx_2" tag:"自测"
    checkout test
    merge feat/xxx_2 id: "test merge feat/xxx_2" tag: "合并 feat/xxx_2"

     %% 完成新功能 2，发布到 pre
    checkout feat/xxx_2
    commit
    checkout pre
    merge feat/xxx_2 id: "pre merge feat/xxx_2" tag: "合并 feat/xxx_2"

     %% 完成新功能 2，发布到 prod
    checkout feat/xxx_2
    commit
    checkout prod
    merge feat/xxx_2 id: "prod merge feat/xxx_2" tag: "合并 feat/xxx_2"

```

## Commit Message Format

**参考文档**

- [angular : Commit Message Format](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#commit-message-format)

Each commit message consists of a **header**, a **body** and a **footer**.  The header has a special format that includes a **type**, a **scope** and a **subject**:

```SQL
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

The **header** is mandatory and the **scope** of the header is optional.

Any line of the commit message cannot be longer than 100 characters! This allows the message to be easier to read on GitHub as well as in various git tools.

### Type

Must be one of the following:

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- **refactor**: A code change that neither fixes a bug nor adds a feature
- **perf**: A code change that improves performance
- **test**: Adding missing or correcting existing tests
- **chore**: Changes to the build process or auxiliary tools and libraries such as documentation generation
