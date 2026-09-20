# Pointers — Exercises

> Work in THIS folder, `package pointersx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 basic.go** — func Inc(p *int), func Swap(a,b *int) with nil-guard error.
- **E2 node.go** — type Node{Val int; Next *Node}; func Walk(head *Node) []int.
- **E3 nil.go** — demo nil *int deref panic recovered; safe constructor NewInt(v int) *int.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package pointersx`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
