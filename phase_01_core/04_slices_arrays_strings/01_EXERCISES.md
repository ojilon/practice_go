# Arrays, Slices, Strings & Runes — Exercises

> Work in THIS folder, `package slicesx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 growth.go** — start s:=make([]int,0,2); append 5 nums printing len/cap each time.
- **E2 copyclone.go** — func Clone(s []int) []int using copy; prove independence.
- **E3 strings.go** — func Reverse(s string) string rune-correct (handle é/日本); func CountRunes(s string) int.
- **E4 clobber.go** — demo append clobber on shared backing array + fix with full slice expr.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package slicesx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
