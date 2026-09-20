# taskpool — Worker Pool + Futures — Guide (`taskpool`)

> Folder: `phase_02_packages/25_taskpool` · Package: `taskpool` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Pool{workers, jobs chan Task, WG}; Task func() (any,error).
- Step 2 — Future[T] with Get()(T,error) via channel; Submit.
- Step 3 — Shutdown(wait bool); error propagation; panic-in-task → error stretch.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package taskpool` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/25_taskpool"` (package name `taskpool`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/25_taskpool/...` must be clean before moving on.

### Details
- Step 1 — type Pool{workers, jobs chan Task, WG}; Task func() (any,error).
- Step 2 — Future[T] with Get()(T,error) via channel; Submit.
- Step 3 — Shutdown(wait bool); error propagation; panic-in-task → error stretch.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/25_taskpool")
func main() { fmt.Println(taskpool.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/25_taskpool/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
