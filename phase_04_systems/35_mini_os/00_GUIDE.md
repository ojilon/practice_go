# mini-os — Files, Args, Signals — Guide (`minios`)

> Folder: `phase_04_systems/35_mini_os` · Package: `minios` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Small safe subset only — no syscalls assembly. Use os/signal/context.
- Step 1 — os.Args subcommand parse; os.Getenv with default.
- Step 2 — os.ReadFile/WriteFile/MkdirTemp lite (cat/cp-lite).
- Step 3 — signal.Notify SIGINT → context cancel graceful shutdown demo.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package minios` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_04_systems/35_mini_os"` (package name `minios`), call your funcs, `go run .`.
3. `go vet ./phase_04_systems/35_mini_os/...` must be clean before moving on.

### Details
- Small safe subset only — no syscalls assembly. Use os/signal/context.
- Step 1 — os.Args subcommand parse; os.Getenv with default.
- Step 2 — os.ReadFile/WriteFile/MkdirTemp lite (cat/cp-lite).
- Step 3 — signal.Notify SIGINT → context cancel graceful shutdown demo.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_04_systems/35_mini_os")
func main() { fmt.Println(minios.Demo()) }
```
```powershell
go run .
go vet ./phase_04_systems/35_mini_os/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
