# taskpool — Worker Pool + Futures — Exercises

> Work in THIS folder, `package taskpool`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 pool.go** — New(workers), Submit, Shutdown; Future via done chan.
- **E2 sum demo** — E2 sum demo: 100 tasks summing squares, collect via Get.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package taskpool`.

## Build-chain (reuse + banned imports)
- MUST use your `15_queue` or `17_ringbuffer` as the job queue (import or mirror + cite), and your `18_logger` for task errors.
- Banned in impl: `golang.org/x/sync/errgroup` — build the pool from goroutines + channels yourself.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
