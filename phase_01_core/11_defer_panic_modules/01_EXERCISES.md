# Defer, Panic/Recover, Packages & Modules — Exercises

> Work in THIS folder, `package deferpanicx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 defer.go** — func Order() []int using 3 defers appending; func Mutate() (n int) with defer n+=3.
- **E2 safe.go** — func MustParse(s string) (n int, err error) recovering from panic.
- **E3 split into two files ops_a.go / ops_b.go same package proving multi-file.** — E3 split into two files ops_a.go / ops_b.go same package proving multi-file.
- **E4 control.go** — fizzbuzz with switch + labeled loop demo.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package deferpanicx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
