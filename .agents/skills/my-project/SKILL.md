---
name: my-project
description: Use when modifying, testing, debugging, or reviewing Go code, Proto files, generation workflows, or module dependencies in this repository.
---

# My Project

Read root `AGENTS.md` first. Use this skill for repository facts, then inspect the owning module and adjacent implementation before deciding how to change it.

## Repository Map

- This is a Go v3 multi-module toolkit. The root and nested `authpkg`, `kit`, `kratos`, `service`, `ping-service`, `data/*`, and `registry/*` directories have independent module boundaries.
- Toolkit packages follow their local API and construction patterns; do not impose an application-layer architecture on them.
- Proto targets are split across the root Makefile and `*makefile_protoc.mk` fragments. Inspect the include graph and confirm that the intended target is reachable before generating code; the root Makefile does not include every module fragment.
- Module release metadata lives in `devops/module-release/modules.tsv`; release guidance lives in `docs/v3_development/module_release.md`.

## Workflow

1. Locate the nearest `go.mod`, then read the package source, tests, README, and relevant Makefile or included fragment.
2. Confirm whether the target is handwritten source, generated output, testdata, or an imported definition under `third_party/`.
3. Make the smallest compatible change inside the owning module and keep `/v3` module/import semantics intact.
4. For Proto changes, edit `.proto` sources and use the established Makefile target; never patch generated `.pb.go`, validation, error, Swagger, or OpenAPI files by hand.
5. Run the affected package tests from the owning module. Widen to the module or dependent modules only when the change crosses their contract.
6. Update nearby documentation when public APIs, configuration, commands, or examples change.

For release-readiness work, verify modules with `GOWORK=off` and use `devops/module-release/module-release.sh`; do not infer that a root-module test or tag covers nested modules.
