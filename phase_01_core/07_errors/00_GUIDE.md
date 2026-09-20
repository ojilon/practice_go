# Error Handling — Guide (`errorsx`)

> Folder: `phase_01_core/07_errors` · Package: `errorsx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — errors.New, fmt.Errorf with %w wrapping, errors.Is/As.
- Step 2 — sentinel vars (ErrNotFound), custom type (type MyErr struct) with Error().
- Step 3 — panic vs error rule; recover in deferred func for parser demo.
- Step 4 — error chains: wrap at each layer, %w only once per layer.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package errorsx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/07_errors"` (package name `errorsx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/07_errors/...` must be clean before moving on.

### Details
- Step 1 — errors.New, fmt.Errorf with %w wrapping, errors.Is/As.
- Step 2 — sentinel vars (ErrNotFound), custom type (type MyErr struct) with Error().
- Step 3 — panic vs error rule; recover in deferred func for parser demo.
- Step 4 — error chains: wrap at each layer, %w only once per layer.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/07_errors")
func main() { fmt.Println(errorsx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/07_errors/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
