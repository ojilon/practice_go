# graph — BFS/DFS/Topo — Exercises

> Work in THIS folder, `package graphx`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 graph.go** — AddVertex/AddEdge/BFS/DFS/ShortestPath(from,to)([]string,int).
- **E2 topo.go** — TopoSort()([]string,error) with cycle error.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package graphx`.

## Build-chain (reuse + banned imports)
- BFS MUST use your queue, DFS your stack (import `15_queue`/`14_stack` or mirror + cite).
- Real-world link: crawl order in `39_scraper` and dependency order in `42_final_boss` (cron/file-watcher) are graph traversals.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
