# math2 — Vectors & Matrices — Exercises

> Work in THIS folder, `package math2`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 vec2.go** — Vec2 + Add/Sub/Scale/Dot/Len/Norm (Norm returns error on zero).
- **E2 vec3.go** — Vec3 + Cross; Mat2 + MulVec + MulMat + Det.
- **E3 doc example func ExampleVec2 in example-style func (plain func Demo() printing).** — E3 doc example func ExampleVec2 in example-style func (plain func Demo() printing).

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package math2`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
