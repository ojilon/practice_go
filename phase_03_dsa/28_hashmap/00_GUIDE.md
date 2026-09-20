# hashmap — Hash Table From Scratch — Guide (`hashmap`)

> Folder: `phase_03_dsa/28_hashmap` · Package: `hashmap` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — buckets [][]entry or []*Node; hash via fnv/maphash; mod len.
- Step 2 — Put/Get/Delete with chaining; load factor + resize x2.
- Step 3 — int keys then generic [K comparable].

## Steps (simple → complex)
Do in order. Create `*.go` files with `package hashmap` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_03_dsa/28_hashmap"` (package name `hashmap`), call your funcs, `go run .`.
3. `go vet ./phase_03_dsa/28_hashmap/...` must be clean before moving on.

### Details
- Step 1 — buckets [][]entry or []*Node; hash via fnv/maphash; mod len.
- Step 2 — Put/Get/Delete with chaining; load factor + resize x2.
- Step 3 — int keys then generic [K comparable].

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_03_dsa/28_hashmap")
func main() { fmt.Println(hashmap.Demo()) }
```
```powershell
go run .
go vet ./phase_03_dsa/28_hashmap/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
