# stack — LIFO — Guide (`stack`)

> Folder: `phase_02_packages/14_stack` · Package: `stack` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — slice-backed type Stack[T any]{items []T}.
- Step 2 — Push/Pop(error ErrEmpty)/Peek/Len/IsEmpty.
- Step 3 — MinStack stretch: O(1) min via aux stack.
- Step 4 — linked-stack alternative comment.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package stack` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/14_stack"` (package name `stack`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/14_stack/...` must be clean before moving on.

### Details
- Step 1 — slice-backed type Stack[T any]{items []T}.
- Step 2 — Push/Pop(error ErrEmpty)/Peek/Len/IsEmpty.
- Step 3 — MinStack stretch: O(1) min via aux stack.
- Step 4 — linked-stack alternative comment.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/14_stack")
func main() { fmt.Println(stack.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/14_stack/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
