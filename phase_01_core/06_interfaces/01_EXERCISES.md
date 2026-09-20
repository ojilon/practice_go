# Interfaces — Exercises

> Work in THIS folder, `package interfacesx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 shape.go** — Shape + Circle/Rect + func TotalArea([]Shape) float64.
- **E2 stringer.go** — type countdown implementing String() string.
- **E3 writer.go** — func WriteTo(w io.Writer, s string)(int,error); test with os.Stdout and bytes.Buffer.
- **E4 anyswitch.go** — func Describe(v any) string with type switch (int,string,Shape,default).

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package interfacesx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
