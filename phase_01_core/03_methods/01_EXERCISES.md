# Methods (value vs pointer receivers) — Exercises

> Work in THIS folder, `package methodsx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 counter.go** — type Counter struct{n int}; Inc() *Counter (pointer), Value() int (value), String() string.
- **E2 bank.go** — type Account{bal float64}; Deposit/Withdraw(with error on overdraft) pointer receivers.
- **E3 builder.go** — type Sentence struct{words []string}; Add(w string) *Sentence chaining; Build() string.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package methodsx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
