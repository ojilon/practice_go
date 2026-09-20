# Structs — Exercises

> Work in THIS folder, `package structsx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 user.go** — type User{Name string; Age int; Email string}, NewUser with validation error.
- **E2 embed.go** — type Address{City,Zip}; embed in User; demo promoted access.
- **E3 tags.go** — type Product with json tags; print with %+v.
- **E4 compare.go** — show which structs are comparable; func that copies struct and proves independence.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package structsx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
