---
name: code-audit-repair
description: Use when the user explicitly requests a whole-repository or multi-module audit covering correctness, tests, security, stability, performance, compatibility, or documentation.
---

# Code Audit Repair

Use this skill only for explicit broad audits. For a single package, bug, or focused review, use `my-project` and the root instructions.

## Workflow

1. Confirm the requested scope and whether the task is read-only or includes repairs.
2. Inventory every in-scope `go.mod`; separate handwritten packages from generated output, `third_party/`, testdata, examples, and release tooling.
3. Review correctness, panic and availability risks, resource handling, concurrency, untrusted input, security, performance amplification, public compatibility, tests, and documentation.
4. Report findings by severity with exact file and line references. Do not treat missing coverage or style preferences as proven defects.
5. When repairs are authorized, make minimal compatible fixes and add focused regression tests in the owning modules.
6. Test affected packages first, then widen module by module. A root `go test ./...` is not a repository-wide result.
7. Report commands run, actual results, unresolved risks, and intentionally deferred work.

## Boundaries

- Do not edit generated output or imported `third_party/` definitions directly.
- Do not revert user changes, perform unrelated refactors, or silently break a v3 public API.
- Require explicit authorization for destructive changes, external writes, dependency installation, releases, sensitive data handling, or breaking contracts.
- Keep audit records under stable `docs/*` only when the user asks for a durable artifact; do not create a parallel task-record workflow.
