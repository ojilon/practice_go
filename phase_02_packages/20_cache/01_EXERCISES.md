# cache — LRU (+TTL) — Exercises

> Work in THIS folder, `package cache`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 lru.go** — generic LRU[K comparable,V any] with New(cap), Get, Set, Len, EvictOldest.
- **E2 ttl.go** — SetWithTTL + Get respecting expiry (time.Sleep demo).

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package cache`.

## Build-chain (reuse + banned imports)
- MUST reuse your `13_list` logic for LRU order: either `import list "practice_go/phase_02_packages/13_list"` or re-implement + cite `// mirrors 13_list` in a comment. Verifier checks for one of the two.
- Banned in impl: `container/list` (use yours), `sync.Map` as the store (plain map + Mutex).

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
