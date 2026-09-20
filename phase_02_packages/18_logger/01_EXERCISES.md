# logger — Leveled Structured Logger — Exercises

> Work in THIS folder, `package logger`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 logger.go** — full logger + New(out,level), SetLevel, Info/Debug/Warn/Error/Errorf with timestamp.
- **E2 fields.go** — With(key,val) chaining returning *Logger sharing out.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package logger`.

## Build-chain (reuse + banned imports)
- Formatting core should be YOURS (see `00_myfmt`): you may still `fmt` for debug, but the timestamp/level/field layout logic must be hand-rolled, not `log/slog` wrapping.
- Banned in impl: `log`, `log/slog` (importing them = thin wrapper = FAIL).

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
