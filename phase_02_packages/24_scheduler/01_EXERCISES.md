# scheduler — Tickers & One-shots — Exercises

> Work in THIS folder, `package scheduler`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 sched.go** — New(), Every, OnceAfter, Stop; uses context.WithCancel + Ticker.
- **E2 count demo** — E2 count demo: Every(10ms) x ~55ms → count in [4..7]; OnceAfter fires once.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package scheduler`.

## Build-chain (reuse + banned imports)
- MUST log ticks/errors via your `18_logger` (import it or accept a minimal `Logger` interface you define + note `// satisfied by 18_logger`).
- Banned in impl: `github.com/robfig/cron` and friends — hand-roll with `time.Ticker` + `context`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
