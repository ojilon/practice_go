# Channels — Exercises

> Work in THIS folder, `package channelsx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 basic.go** — send 1..5 on goroutine, close, range-sum.
- **E2 buffered.go** — cap-2 demo showing blocking (use goroutine + prints).
- **E3 select.go** — func First(ch1,ch2 <-chan string) string using select.
- **E4 pipeline.go** — gen(n) → square → filterOdd; fan-in two gens.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package channelsx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
