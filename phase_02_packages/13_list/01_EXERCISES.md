# list — Linked List (generics) — Exercises

> Work in THIS folder, `package list`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 list.go** — generic List[T comparable or any] with PushFront/Back, PopFront/PopBack(error on empty), Len, ToSlice.
- **E2 extra.go** — Reverse(), Find(v T) bool, Clear().

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package list`.

## Build-chain (reuse + banned imports)
- This package IS the building block: `20_cache`, `30_tree` (level-order), `31_graph` (BFS/DFS) will reuse it. Design `ToSlice`/`Len` cleanly for them.
- Banned in impl: `container/list` (this exercise replaces it).

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
