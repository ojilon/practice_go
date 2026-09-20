# Variables & Functions — Guide (`variablesfunctions`)

> Folder: `phase_01_core/01_variables_functions` · Package: `variablesfunctions` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — var, :=, const, zero values. Print types with %T.
- Step 2 — funcs: params, multiple returns (int,error), named returns.
- Step 3 — variadic funcs Sum(nums ...int), closures capturing loop vars.
- Step 4 — blank identifier, shadowing traps, scope.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package variablesfunctions` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/01_variables_functions"` (package name `variablesfunctions`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/01_variables_functions/...` must be clean before moving on.

### Details
- Step 1 — var, :=, const, zero values. Print types with %T.
- Step 2 — funcs: params, multiple returns (int,error), named returns.
- Step 3 — variadic funcs Sum(nums ...int), closures capturing loop vars.
- Step 4 — blank identifier, shadowing traps, scope.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/01_variables_functions")
func main() { fmt.Println(variablesfunctions.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/01_variables_functions/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
