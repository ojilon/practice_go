# cache — LRU (+TTL) — Guide (`cache`)

> Folder: `phase_02_packages/20_cache` · Package: `cache` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — LRU via map[key]*list.Element + container/list.
- Step 2 — New(cap int), Get/Set/Evict oldest, Len.
- Step 3 — TTL variant with time.Now + expire on Get (lazy).
- Step 4 — hits/misses counters stretch.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package cache` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/20_cache"` (package name `cache`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/20_cache/...` must be clean before moving on.

### Details
- Step 1 — LRU via map[key]*list.Element + container/list.
- Step 2 — New(cap int), Get/Set/Evict oldest, Len.
- Step 3 — TTL variant with time.Now + expire on Get (lazy).
- Step 4 — hits/misses counters stretch.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/20_cache")
func main() { fmt.Println(cache.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/20_cache/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
