# Maps — Guide (`mapsx`)

> Folder: `phase_01_core/05_maps` · Package: `mapsx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — make(map[string]int), literals, comma-ok v,ok := m[k].
- Step 2 — delete, iteration order is random (prove by running 3x), sorted keys via sort.
- Step 3 — nested maps, struct keys (comparable requirement), map of slices.
- Step 4 — maps are references; func mutating map; nil map read vs write panic.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package mapsx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/05_maps"` (package name `mapsx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/05_maps/...` must be clean before moving on.

### Details
- Step 1 — make(map[string]int), literals, comma-ok v,ok := m[k].
- Step 2 — delete, iteration order is random (prove by running 3x), sorted keys via sort.
- Step 3 — nested maps, struct keys (comparable requirement), map of slices.
- Step 4 — maps are references; func mutating map; nil map read vs write panic.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/05_maps")
func main() { fmt.Println(mapsx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/05_maps/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
