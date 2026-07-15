---
name: code-audit-repair
description: Use when the user explicitly requests a whole-repository or multi-module audit covering correctness, tests, security, stability, performance, compatibility, or documentation.
---

# Code Audit Repair

Use this skill only for explicit broad audits. For a single package or bug, use `my-project` and the relevant global workflow skill.

## Workflow

1. Confirm the requested scope and whether the task is read-only or includes repairs.
2. Inventory Go modules, generated files, third-party sources, tests, and public APIs before sampling packages.
3. Prioritize correctness, panic and availability risks, resource leaks, unsafe input handling, concurrency, compatibility, and missing tests.
4. Report findings by severity with exact file and line references.
5. When repairs are authorized, make minimal compatible fixes and add focused regression tests.
6. Test affected packages from their owning module directories before widening verification.

## Boundaries

- Do not edit generated or third-party files directly.
- Do not revert existing user changes or perform unrelated refactors.
- Require confirmation for destructive changes, external writes, sensitive data, or breaking public APIs.
- Report commands run, actual results, unresolved risks, and intentionally deferred work.
