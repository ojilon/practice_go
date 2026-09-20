# Goroutines — Exercises

> Work in THIS folder, `package goroutinesx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 hello.go** — launch 5 goroutines printing with WaitGroup (capture var correctly).
- **E2 counter.go** — 100 goroutines x 1000 inc with Mutex; expect 100000.
- **E3 race.go** — first version WITHOUT mutex (to see race via -race), then fixed version (keep both funcs).

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package goroutinesx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
