# set — Generic Set — Guide (`set`)

> Folder: `phase_02_packages/16_set` · Package: `set` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Set[T comparable] map[T]struct{}.
- Step 2 — Add/Has/Remove/Size/Values.
- Step 3 — Union/Intersect/Diff as new Sets.
- Step 4 — struct-key sets; zero-size struct trick.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package set` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/16_set"` (package name `set`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/16_set/...` must be clean before moving on.

### Details
- Step 1 — type Set[T comparable] map[T]struct{}.
- Step 2 — Add/Has/Remove/Size/Values.
- Step 3 — Union/Intersect/Diff as new Sets.
- Step 4 — struct-key sets; zero-size struct trick.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/16_set")
func main() { fmt.Println(set.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/16_set/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
