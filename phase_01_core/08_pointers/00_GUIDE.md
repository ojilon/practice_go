# Pointers — Guide (`pointersx`)

> Folder: `phase_01_core/08_pointers` · Package: `pointersx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — & and *, new(T), zero value nil.
- Step 2 — func Zero(p *int), Swap(a,b *int); nil guards.
- Step 3 — pointers to structs, linked node type; method receiver recap.
- Step 4 — escape/lifetime intuition: returning &local is safe in Go (escapes to heap).

## Steps (simple → complex)
Do in order. Create `*.go` files with `package pointersx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/08_pointers"` (package name `pointersx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/08_pointers/...` must be clean before moving on.

### Details
- Step 1 — & and *, new(T), zero value nil.
- Step 2 — func Zero(p *int), Swap(a,b *int); nil guards.
- Step 3 — pointers to structs, linked node type; method receiver recap.
- Step 4 — escape/lifetime intuition: returning &local is safe in Go (escapes to heap).

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/08_pointers")
func main() { fmt.Println(pointersx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/08_pointers/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
