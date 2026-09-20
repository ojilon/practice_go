# tree — BST — Guide (`treex`)

> Folder: `phase_03_dsa/30_tree` · Package: `treex` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Node{Val; Left,Right *Node}; Insert/Search.
- Step 2 — Delete (leaf/1-child/2-child via min-right).
- Step 3 — In/Pre/Post/Level traversals (level uses YOUR queue package logic, copy minimal).

## Steps (simple → complex)
Do in order. Create `*.go` files with `package treex` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_03_dsa/30_tree"` (package name `treex`), call your funcs, `go run .`.
3. `go vet ./phase_03_dsa/30_tree/...` must be clean before moving on.

### Details
- Step 1 — type Node{Val; Left,Right *Node}; Insert/Search.
- Step 2 — Delete (leaf/1-child/2-child via min-right).
- Step 3 — In/Pre/Post/Level traversals (level uses YOUR queue package logic, copy minimal).

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_03_dsa/30_tree")
func main() { fmt.Println(treex.Demo()) }
```
```powershell
go run .
go vet ./phase_03_dsa/30_tree/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
