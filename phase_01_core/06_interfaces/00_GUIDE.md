# Interfaces — Guide (`interfacesx`)

> Folder: `phase_01_core/06_interfaces` · Package: `interfacesx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — interface Shape{Area() float64}; Circle/Rect impls.
- Step 2 — implicit satisfaction; empty interface any; type switch.
- Step 3 — io.Writer / fmt.Stringer: write func PrintAll(w io.Writer, v ...any).
- Step 4 — interface composition (io.ReadWriter), nil interface vs typed-nil trap.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package interfacesx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/06_interfaces"` (package name `interfacesx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/06_interfaces/...` must be clean before moving on.

### Details
- Step 1 — interface Shape{Area() float64}; Circle/Rect impls.
- Step 2 — implicit satisfaction; empty interface any; type switch.
- Step 3 — io.Writer / fmt.Stringer: write func PrintAll(w io.Writer, v ...any).
- Step 4 — interface composition (io.ReadWriter), nil interface vs typed-nil trap.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/06_interfaces")
func main() { fmt.Println(interfacesx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/06_interfaces/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
