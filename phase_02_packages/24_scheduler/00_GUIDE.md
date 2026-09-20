# scheduler — Tickers & One-shots — Guide (`scheduler`)

> Folder: `phase_02_packages/24_scheduler` · Package: `scheduler` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Sched struct{ctx,cancel,WG}; Every(d,fn) returns Stop func.
- Step 2 — OnceAfter(d,fn); Stop() cancels all.
- Step 3 — context + time.Ticker (not Sleep loop); no goroutine leak after Stop.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package scheduler` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/24_scheduler"` (package name `scheduler`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/24_scheduler/...` must be clean before moving on.

### Details
- Step 1 — type Sched struct{ctx,cancel,WG}; Every(d,fn) returns Stop func.
- Step 2 — OnceAfter(d,fn); Stop() cancels all.
- Step 3 — context + time.Ticker (not Sleep loop); no goroutine leak after Stop.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/24_scheduler")
func main() { fmt.Println(scheduler.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/24_scheduler/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
