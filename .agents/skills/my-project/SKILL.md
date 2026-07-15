---
name: my-project
description: Use when modifying, testing, debugging, or reviewing Go code, Proto files, generation workflows, or module dependencies in this repository.
---

# My Project

Read root `AGENTS.md` first. Use this skill for repository facts; use matching global skills for general engineering workflows.

## Repository Shape

- This is a multi-module Go toolkit, not a single business service.
- Modules include the root plus `auth`, `kit`, `kratos`, `service`, `ping-service`, `data/*`, and `registry/*`.
- Find the nearest `go.mod` before selecting imports, commands, or test scope.
- Follow the target package's existing architecture, naming, and construction patterns.

## Workflow

1. Read the owning `go.mod`, package files, tests, README, and relevant Makefile targets.
2. Make the smallest compatible change in the owning module.
3. For Proto or Wire changes, edit source definitions and run the established generator; never hand-edit generated output.
4. Run targeted tests from the owning module directory, then widen verification only when the change crosses module boundaries.
5. Check documentation when public behavior, configuration, commands, or examples change.
