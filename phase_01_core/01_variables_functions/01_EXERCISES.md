# Variables & Functions — Exercises

> Work in THIS folder, `package variablesfunctions`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 vars.go** — declare string/int/bool/float with var and :=, print zero values.
- **E2 funcs.go** — func Add(a,b int) int, func Div(a,b float64)(float64,error) returning error on /0.
- **E3 named.go** — func Split(full string)(first,last string) using named returns.
- **E4 variadic.go** — func Sum(nums ...int) int + func Apply(fn func(int)int, vals ...int) []int.
- **E5 closure.go** — func Counter() func() int returning incrementing closure; demo shadowing bug + fix.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package variablesfunctions`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
