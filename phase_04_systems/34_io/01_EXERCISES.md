# io — Reader/Writer Rebuilds — Exercises

> Work in THIS folder, `package iox`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 miniio.go** — Reader/Writer interfaces + ReadFull + CopyN + LimitReader + TeeReader (no io.* in impl).
- **E2 chain demo with strings.Reader/bytes.Buffer as sources (allowed in demo only).** — E2 chain demo with strings.Reader/bytes.Buffer as sources (allowed in demo only).

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package iox`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
