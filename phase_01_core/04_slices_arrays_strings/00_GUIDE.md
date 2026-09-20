# Arrays, Slices, Strings & Runes — Guide (`slicesx`)

> Folder: `phase_01_core/04_slices_arrays_strings` · Package: `slicesx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — arrays [3]int vs slices []int; len vs cap.
- Step 2 — append growth (print len/cap each append), copy(), slice expressions s[1:3], nil vs empty.
- Step 3 — strings are []byte read-only; range gives runes; byte vs rune vs string conversions.
- Step 4 — full slice expr s[1:3:3] to control cap (append-clobber demo).

## Steps (simple → complex)
Do in order. Create `*.go` files with `package slicesx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/04_slices_arrays_strings"` (package name `slicesx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/04_slices_arrays_strings/...` must be clean before moving on.

### Details
- Step 1 — arrays [3]int vs slices []int; len vs cap.
- Step 2 — append growth (print len/cap each append), copy(), slice expressions s[1:3], nil vs empty.
- Step 3 — strings are []byte read-only; range gives runes; byte vs rune vs string conversions.
- Step 4 — full slice expr s[1:3:3] to control cap (append-clobber demo).

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/04_slices_arrays_strings")
func main() { fmt.Println(slicesx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/04_slices_arrays_strings/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
