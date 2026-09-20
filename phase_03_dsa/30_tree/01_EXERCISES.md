# tree — BST — Exercises

> Work in THIS folder, `package treex`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 bst.go** — Insert/Search/Delete/Min/Max/Height.
- **E2 traversals.go** — InOrder/PreOrder/PostOrder/LevelOrder []int.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package treex`.

## Build-chain (reuse + banned imports)
- `LevelOrder` MUST use your queue: import `queue "practice_go/phase_02_packages/15_queue"` or mirror it + cite `// mirrors 15_queue`.
- Real-world link: this BST + your queue is the engine reused by `38_parsers` (sorted reports) and `42_final_boss` options.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
