# graph — BFS/DFS/Topo — Guide (`graphx`)

> Folder: `phase_03_dsa/31_graph` · Package: `graphx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — adjacency map[string][]string or generic; AddVertex/AddEdge.
- Step 2 — BFS (queue) + DFS (stack/recursive) returning visit order.
- Step 3 — HasCycle/TopoSort + BFS shortest path.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package graphx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_03_dsa/31_graph"` (package name `graphx`), call your funcs, `go run .`.
3. `go vet ./phase_03_dsa/31_graph/...` must be clean before moving on.

### Details
- Step 1 — adjacency map[string][]string or generic; AddVertex/AddEdge.
- Step 2 — BFS (queue) + DFS (stack/recursive) returning visit order.
- Step 3 — HasCycle/TopoSort + BFS shortest path.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_03_dsa/31_graph")
func main() { fmt.Println(graphx.Demo()) }
```
```powershell
go run .
go vet ./phase_03_dsa/31_graph/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
