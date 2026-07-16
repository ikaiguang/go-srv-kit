# Kit Audit Repair Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make the non-`log` portion of the `kit` module buildable, panic-resistant, race-safe, compatibility-preserving, consistently named, tested, and LF-normalized.

**Architecture:** Repair independent packages in risk order while preserving their public contracts. Each behavior change starts with a focused failing test; corrected public names own the implementation and deprecated names are thin wrappers.

**Tech Stack:** Go 1.26.3, standard library, existing module dependencies, Testify, Go race detector, vet, golangci-lint.

---

### Task 1: Restore The Module Baseline

**Files:**
- Modify: `kit/go.mod`
- Modify: `kit/go.sum`
- Modify: imports under `kit/filepath`, `kit/id`, `kit/rsa`, `kit/zip`, and `kit/page/testdata`
- Modify: `kit/file-rotatelogs/options.go`
- Modify: `kit/file-rotatelogs/*_test.go`

- [ ] Replace stale `github.com/ikaiguang/go-kit/*` imports with the current module path.
- [ ] Point copied rotatelogs code and tests at their local package paths.
- [ ] Add only the dependencies required by the copied rotatelogs implementation.
- [ ] Run `go mod tidy` and `go test` for each previously unbuildable package.

### Task 2: Repair Concrete Panic And Concurrency Failures

**Files:**
- Modify: `kit/file-rotatelogs/rotatelogs.go`
- Modify: `kit/file-rotatelogs/rotatelogs_test.go`
- Modify: `kit/locker/locker_local.kit.go`
- Modify: `kit/locker/locker_local.kit_test.go`
- Modify: `kit/component/component.util.go`
- Modify: `kit/component/component.util_concurrency_test.go`
- Modify: `kit/download/download.go`
- Modify: `kit/download/download_test.go`
- Modify: `kit/id/id.kit.go`
- Modify: `kit/id/id.kit_test.go`

- [ ] Add a failing local rotatelogs test proving first write and write-after-close do not panic.
- [ ] Fix file handle transitions and propagate close errors.
- [ ] Add a failing lock test for stale timer deletion and concurrent idempotent unlock.
- [ ] Track lock ownership and delete map entries only when they still reference that lock.
- [ ] Add lifecycle tests for nil dependencies and re-entrant close behavior.
- [ ] Detach lifecycle closers before invoking user code and guard optional dependencies.
- [ ] Add nil download and concurrent same-target tests.
- [ ] Use unique temporary files and validate inputs before dereference or allocation.
- [ ] Add an ID race test and synchronize package-global node state.

### Task 3: Repair Input, Overflow, And Resource Handling

**Files:**
- Modify: `kit/random/random.kit.go`
- Modify: `kit/random/random.kit_test.go`
- Modify: `kit/page/page.kit.go`
- Modify: `kit/page/page.kit_test.go`
- Modify: `kit/string/string.kit.go`
- Modify: `kit/string/string.kit_test.go`
- Modify: `kit/curl/curl.kit.go`
- Modify: `kit/curl/curl.kit_test.go`
- Modify: `kit/email/email.kit.go`
- Modify: `kit/email/email_helper.kit.go`
- Modify: `kit/email/email.kit_test.go`
- Modify: `kit/slice/slice.kit.go`
- Modify: `kit/slice/slice.kit_test.go`

- [ ] Add boundary tests for full-width random integer ranges, page underflow, and unsigned string conversion.
- [ ] Implement overflow-safe range selection and page offset calculation.
- [ ] Format unsigned values without narrowing through `int`.
- [ ] Add nil argument tests for HTTP and email helpers and return errors instead of panics.
- [ ] Stop closing shared default HTTP transport idle connections after every request.
- [ ] Make the deprecated reflective reverse helper a safe no-op for invalid input.

### Task 4: Correct Public Function Names Compatibly

**Files:**
- Modify canonical source and tests in `kit/aes`, `kit/base64`, `kit/chinese`, `kit/connection`, `kit/file`, `kit/header`, `kit/id`, `kit/ip`, `kit/md5`, `kit/rsa`, `kit/url`, and `kit/zip`

- [ ] Add canonical names using Go acronym capitalization and precise verbs.
- [ ] Move implementation bodies to the canonical names.
- [ ] Keep every old name as a forwarding function with a `Deprecated:` comment.
- [ ] Compile tests against both canonical and compatibility entry points.

### Task 5: Normalize And Verify

**Files:**
- Create: `kit/.gitattributes`
- Normalize: text files under `kit`, excluding logical changes to `kit/log`

- [ ] Set `text=auto eol=lf` for the `kit` subtree.
- [ ] Convert existing CRLF text files to LF without changing content.
- [ ] Run `gofmt` on changed Go files and confirm no unformatted non-generated Go files remain.
- [ ] Run focused tests after each repair, then `go test ./...`.
- [ ] Run race tests for every non-`log` package, `go vet`, and `golangci-lint` where tool compatibility permits.
- [ ] Confirm `git diff -- kit/log` is unchanged from the pre-audit snapshot and report residual risks.
