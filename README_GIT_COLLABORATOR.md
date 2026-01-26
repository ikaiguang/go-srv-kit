# Git 协作流程规范

参考文档：

* [angular : Commit Message Format](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#commit-message-format)

## Git 协作流程

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
    branch feature/xxx order: 60
    commit id: "checkout feature/xxx_1 分支" tag: "新功能"
    commit id: "dev feature/xxx_1" tag:"开发中"

    %% 发现 Bug
    checkout prod
    branch hotfix/xxx order: 10
    commit id: "checkout fix/xxx_1" tag: "发现Bug"
    commit id: "fix fix/xxx_1" tag:"修复中"

    %% 完成新功能，发布到 test
    checkout feature/xxx
    commit id: "commit feature/xxx_1" tag:"完成新功能"
    commit id: "test feature/xxx_1" tag:"自测"
    checkout test
    merge feature/xxx tag: "合并 feature/xxx_1"

    %% 修复 Bug，发布到 test
    checkout hotfix/xxx
    commit id: "commit fix/xxx_1" tag:"修复Bug"
    commit id: "test fix/xxx_1" tag:"自测"
    checkout test
    merge hotfix/xxx tag: "合并 fix/xxx_1"

     %% 完成新功能，发布到 pre
    checkout feature/xxx
    commit
    checkout pre
    merge feature/xxx tag: "合并 feature/xxx_1"

     %% 完成新功能，发布到 prod
    checkout feature/xxx
    commit
    checkout prod
    merge feature/xxx tag: "合并 feature/xxx_1"

    %% 修复 Bug，发布到 pre
    checkout hotfix/xxx
    commit
    checkout pre
    merge hotfix/xxx tag: "合并 fix/xxx_1"

    %% 修复 Bug，发布到 prod
    checkout hotfix/xxx
    commit
    checkout prod
    merge hotfix/xxx tag: "合并 fix/xxx_1"

    %% 新功能 2
    checkout prod
    commit
    branch feature/xxx_2 order: 70
    commit id: "checkout feature/xxx_2 分支" tag: "新功能"
    commit id: "dev feature/xxx_2" tag:"开发中"

     %% 完成新功能 2，发布到 test
    checkout feature/xxx_2
    commit id: "commit feature/xxx_2" tag:"完成新功能"
    commit id: "test feature/xxx_2" tag:"自测"
    checkout test
    merge feature/xxx_2 tag: "合并 feature/xxx_2"

     %% 完成新功能 2，发布到 pre
    checkout feature/xxx_2
    commit
    checkout pre
    merge feature/xxx_2 tag: "合并 feature/xxx_2"

     %% 完成新功能 2，发布到 prod
    checkout feature/xxx_2
    commit
    checkout prod
    merge feature/xxx_2 tag: "合并 feature/xxx_2"

```

## Commit Message Format - 提交消息格式

Each commit message consists of a **header**, a **body** and a **footer**.  The header has a special format that includes a **type**, a **scope** and a **subject**:

```txt
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

* **feat**: A new feature
* **fix**: A bug fix
* **docs**: Documentation only changes
* **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
* **refactor**: A code change that neither fixes a bug nor adds a feature
* **perf**: A code change that improves performance
* **test**: Adding missing or correcting existing tests
* **chore**: Changes to the build process or auxiliary tools and libraries such as documentation generation
