# Agent Instructions

## Scope And Priority

- This file is the repository-wide rule source. More specific instructions in a nested directory take precedence for that directory.
- Read the target package, its tests, the nearest `go.mod`, and relevant README/Makefile entries before changing code.
- Follow existing package patterns, keep changes narrowly scoped, and preserve public API compatibility unless the user explicitly approves a breaking change.
- Preserve unrelated working-tree changes and do not perform releases, pushes, deployments, or destructive operations without explicit authorization.

## Repository Shape

- This is a Go v3 multi-module toolkit, not a single business service.
- Modules include the root, `authpkg/`, `kit/`, `kratos/`, `service/`, `ping-service/`, `data/*`, and `registry/*`.
- Find the nearest `go.mod` before choosing imports, dependency versions, commands, or test scope.
- A root `go test ./...` does not cover nested modules. Run tests from the module that owns the changed package.
- Treat current source, Makefile fragments, and `devops/module-release/modules.tsv` as authoritative when prose documentation disagrees.

## Changes And Generation

- Prefer the smallest compatible change in the owning module; do not impose a business-service layer model on toolkit packages.
- Do not edit files marked `Code generated` or generated Proto/OpenAPI output directly. Change the source `.proto` or generator input and run the established Makefile target.
- Before running a Make target, inspect the root Makefile include graph and use `make -n` when practical. The root Makefile includes only selected fragments and does not expose every module's Proto target.
- Treat `third_party/` as imported protocol definitions. Change it only when the task explicitly requires updating those definitions.
- For dependency changes, keep module paths and `/v3` major-version semantics correct and update only the owning module's `go.mod`/`go.sum`.
- If public behavior, configuration, commands, or examples change, update the nearest README or stable documentation in the same task.

## Verification

- Start with the affected package from its owning module, then widen to the module or dependent modules when risk warrants it.
- Use `GOWORK=off` for release-readiness checks so a local workspace cannot hide missing requirements or unpublished versions.
- For Proto changes, run the relevant generation target and verify that regenerated output is consistent.
- For agent/document-only changes, validate links, referenced paths, frontmatter, whitespace, and `git diff --check`; Go tests are not required unless code behavior changed.
- Report commands actually run, their results, and any validation intentionally skipped.

## Documentation

- Under `docs/*`, write design, specification, implementation-plan, and review prose in Chinese by default.
- Keep code identifiers, commands, paths, protocol names, library/tool names, and terms that do not translate cleanly in English.
- Keep durable rules in this file, repo-local skills, module READMEs, or stable docs; do not commit one-off agent task records as permanent guidance.

## Repo-Local Skills

- Use `.agents/skills/my-project` for repository-specific Go, module, Proto, generation, dependency, test, debugging, and review context.
- Use `.agents/skills/code-audit-repair` only when the user explicitly requests a whole-repository or multi-module audit.
- Repo-local skills add project facts; they must not duplicate this file or invent workflows and architecture that are absent from the repository.

## Context7

- Use Context7 for current documentation when the user asks about a library, framework, SDK, API, CLI tool, or cloud service.
- Resolve the library ID first, then query the selected `/org/project` ID with the user's full concept-specific question. Use separate queries for distinct concepts.
- Do not use Context7 for repository business logic, refactoring, code review, scripts written from scratch, or general programming concepts.
