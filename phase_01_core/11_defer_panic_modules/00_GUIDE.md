# Defer, Panic/Recover, Packages & Modules — Guide (`deferpanicx`)

> Folder: `phase_01_core/11_defer_panic_modules` · Package: `deferpanicx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — defer LIFO order, named-return + defer mutation.
- Step 2 — panic/recover: only at top-level handler, not for normal errors.
- Step 3 — multi-file package: split logic across a.go/b.go same package; exported vs unexported.
- Step 4 — if/switch/for deep-dive: init stmt, fallthrough, labeled break.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package deferpanicx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/11_defer_panic_modules"` (package name `deferpanicx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/11_defer_panic_modules/...` must be clean before moving on.

### Details
- Step 1 — defer LIFO order, named-return + defer mutation.
- Step 2 — panic/recover: only at top-level handler, not for normal errors.
- Step 3 — multi-file package: split logic across a.go/b.go same package; exported vs unexported.
- Step 4 — if/switch/for deep-dive: init stmt, fallthrough, labeled break.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/11_defer_panic_modules")
func main() { fmt.Println(deferpanicx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/11_defer_panic_modules/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
