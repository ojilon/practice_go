# FINAL BOSS — Integrate ≥5 pkgs — Guide (`finalboss`)

> Folder: `phase_05_apps/42_final_boss` · Package: `finalboss` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Pick ONE: KV server / markdown-lite→HTML / cron daemon / chat bus / file watcher.
- Must reuse ≥5 of YOUR packages (list them in 02_QUESTIONS).
- Must include: README section (arch), graceful shutdown, errors with %w, goroutines+channels OR pool, one parser.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package finalboss` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_05_apps/42_final_boss"` (package name `finalboss`), call your funcs, `go run .`.
3. `go vet ./phase_05_apps/42_final_boss/...` must be clean before moving on.

### Details
- Pick ONE: KV server / markdown-lite→HTML / cron daemon / chat bus / file watcher.
- Must reuse ≥5 of YOUR packages (list them in 02_QUESTIONS).
- Must include: README section (arch), graceful shutdown, errors with %w, goroutines+channels OR pool, one parser.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_05_apps/42_final_boss")
func main() { fmt.Println(finalboss.Demo()) }
```
```powershell
go run .
go vet ./phase_05_apps/42_final_boss/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
