# ini — INI Parser — Exercises

> Work in THIS folder, `package ini`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 ini.go** — Parse + Get/GetInt/GetBool + ErrNoSection/ErrNoKey.
- **E2 sample.ini-equivalent** — demo via raw string in demo func (no external file needed).

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package ini`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
