# Sort & Search + Benchmarks — Guide (`sortsearch`)

> Folder: `phase_03_dsa/27_sort_search` · Package: `sortsearch` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — Bubble/Insertion (O(n²)) then Merge/Quick (O(n log n)).
- Step 2 — BinarySearch(sorted,key)(idx,bool).
- Step 3 — benchmark funcs + table comparing; use sort.Ints as oracle.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package sortsearch` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_03_dsa/27_sort_search"` (package name `sortsearch`), call your funcs, `go run .`.
3. `go vet ./phase_03_dsa/27_sort_search/...` must be clean before moving on.

### Details
- Step 1 — Bubble/Insertion (O(n²)) then Merge/Quick (O(n log n)).
- Step 2 — BinarySearch(sorted,key)(idx,bool).
- Step 3 — benchmark funcs + table comparing; use sort.Ints as oracle.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_03_dsa/27_sort_search")
func main() { fmt.Println(sortsearch.Demo()) }
```
```powershell
go run .
go vet ./phase_03_dsa/27_sort_search/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
