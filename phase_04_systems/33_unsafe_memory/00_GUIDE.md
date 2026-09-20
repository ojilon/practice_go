# unsafe & Memory Layout — Guide (`unsafex`)

> Folder: `phase_04_systems/33_unsafe_memory` · Package: `unsafex` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- DOCS QUEST: go doc unsafe + unsafe.Sizeof/Alignof/Offsetof BEFORE coding. SAFETY comments required.
- Step 1 — Sizeof int/int32/struct/pointer/slice/map on your arch.
- Step 2 — unsafe.Pointer uintptr round-trip; SliceHeader/StringHeader zero-copy (read-only demo).
- Step 3 — why unsafe is dangerous: GC + aliasing notes.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package unsafex` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_04_systems/33_unsafe_memory"` (package name `unsafex`), call your funcs, `go run .`.
3. `go vet ./phase_04_systems/33_unsafe_memory/...` must be clean before moving on.

### Details
- DOCS QUEST: go doc unsafe + unsafe.Sizeof/Alignof/Offsetof BEFORE coding. SAFETY comments required.
- Step 1 — Sizeof int/int32/struct/pointer/slice/map on your arch.
- Step 2 — unsafe.Pointer uintptr round-trip; SliceHeader/StringHeader zero-copy (read-only demo).
- Step 3 — why unsafe is dangerous: GC + aliasing notes.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_04_systems/33_unsafe_memory")
func main() { fmt.Println(unsafex.Demo()) }
```
```powershell
go run .
go vet ./phase_04_systems/33_unsafe_memory/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
