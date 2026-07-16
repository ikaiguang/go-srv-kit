# Kit Audit And Repair Design

## Scope

Audit and repair all code in the `kit` module except `kit/log`. Preserve the
current staged `kit/log` refactor byte-for-byte and do not include it in test or
lint conclusions when a command can exclude the package.

Normalize text files under `kit` to LF and add a repository rule that keeps
future checkouts consistent. Generated Proto code is not logically edited; its
source of truth remains `kit/page/page.kit.proto`.

## Audit Priorities

1. Module buildability and dependency correctness.
2. Panics, data races, deadlocks, stale ownership, and unsafe nil handling.
3. Path traversal, unsafe file replacement, unbounded input, weak crypto usage,
   command execution boundaries, and TLS behavior.
4. Resource lifecycle, error propagation, cancellation, partial writes, and
   cross-platform behavior.
5. Hot-path allocations, connection reuse, unnecessary I/O, and algorithmic
   overflow.
6. Public API compatibility, precise Go naming, documentation, and regression
   coverage.

## Compatibility Rules

- Public signatures stay compatible unless correctness cannot be restored
  without a new API.
- A corrected function name is introduced as the canonical implementation.
- The old function remains, delegates to the new function, and receives a Go
  `Deprecated:` doc comment.
- Existing wire formats, encryption formats, and return shapes are preserved.
- Copied `file-rotatelogs` behavior is changed only where needed to compile or
  prevent concrete panic/resource failures.

## Verification

Use focused red-green tests for every behavioral repair. Then run package tests,
`go test -race` outside `kit/log`, `go vet`, formatting checks, `golangci-lint`
when compatible with the configured Go version, and a final module-wide test.
Any unresolved security or compatibility risks are reported with exact file and
line references.
