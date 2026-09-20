# Error Handling — Exercises

> Work in THIS folder, `package errorsx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 sentinel.go** — var ErrEmpty; func First(s []int)(int,error) returning ErrEmpty.
- **E2 wrap.go** — low-level open/parse funcs wrapping with %w; demo errors.Is.
- **E3 custom.go** — type ValidationError{Field,Msg string}; Error() string; func Validate age>=0.
- **E4 recover.go** — func SafeRun(fn func())(err error) with defer+recover converting panic to error.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package errorsx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
