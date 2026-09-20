# Methods (value vs pointer receivers) — Guide (`methodsx`)

> Folder: `phase_01_core/03_methods` · Package: `methodsx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — func (u User) Greet() vs func (u *User) Birthday().
- Step 2 — String() string satisfies fmt.Stringer.
- Step 3 — fluent builder: func (b *B) Add(x int) *B chaining.
- Step 4 — when to use pointer: mutation + large structs + consistency rule.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package methodsx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/03_methods"` (package name `methodsx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/03_methods/...` must be clean before moving on.

### Details
- Step 1 — func (u User) Greet() vs func (u *User) Birthday().
- Step 2 — String() string satisfies fmt.Stringer.
- Step 3 — fluent builder: func (b *B) Add(x int) *B chaining.
- Step 4 — when to use pointer: mutation + large structs + consistency rule.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/03_methods")
func main() { fmt.Println(methodsx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/03_methods/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
