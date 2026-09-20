# Maps — Exercises

> Work in THIS folder, `package mapsx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 freq.go** — func WordFreq(words []string) map[string]int.
- **E2 ok.go** — func Lookup(m map[string]int,k string)(int,bool) demo comma-ok.
- **E3 nested.go** — map[string]map[string]int grades; AddGrade/ Avg funcs.
- **E4 keys.go** — collect keys, sort.Strings, print sorted; show random order with 3 runs comment.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package mapsx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
