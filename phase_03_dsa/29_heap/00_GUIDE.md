# heap — Priority Queue — Guide (`heapx`)

> Folder: `phase_03_dsa/29_heap` · Package: `heapx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — slice heap with siftUp/siftDown; Less func or min-heap ints.
- Step 2 — Push/Pop(error)/Top/Len; heap invariant check func.
- Step 3 — generic with cmp func; compare to container/heap interface.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package heapx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_03_dsa/29_heap"` (package name `heapx`), call your funcs, `go run .`.
3. `go vet ./phase_03_dsa/29_heap/...` must be clean before moving on.

### Details
- Step 1 — slice heap with siftUp/siftDown; Less func or min-heap ints.
- Step 2 — Push/Pop(error)/Top/Len; heap invariant check func.
- Step 3 — generic with cmp func; compare to container/heap interface.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_03_dsa/29_heap")
func main() { fmt.Println(heapx.Demo()) }
```
```powershell
go run .
go vet ./phase_03_dsa/29_heap/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
