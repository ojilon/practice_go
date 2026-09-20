# jsonlite — Mini JSON (no encoding/json in impl) — Exercises

> Work in THIS folder, `package jsonlite`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 marshal.go** — func Marshal(v any)(string,error) for listed kinds (NO encoding/json import).
- **E2 unmarshal.go** — func Unmarshal(s string)(any,error) recursive descent; objects → map[string]any.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package jsonlite`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
