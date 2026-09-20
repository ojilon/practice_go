# TUI — ANSI (stdlib only) — Guide (`tuix`)

> Folder: `phase_05_apps/40_tui` · Package: `tuix` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- No deps required (bubbletea stretch only). Use ANSI escapes + goroutines/channels.
- Step 1 — Clear/HideCursor/MoveTo helpers; progress bar renderer func Bar(pct int) string.
- Step 2 — list viewer: up/down keys via raw input (bufio) OR auto-tick demo if raw too hard.
- Step 3 — tick loop with select + context cancel.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package tuix` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_05_apps/40_tui"` (package name `tuix`), call your funcs, `go run .`.
3. `go vet ./phase_05_apps/40_tui/...` must be clean before moving on.

### Details
- No deps required (bubbletea stretch only). Use ANSI escapes + goroutines/channels.
- Step 1 — Clear/HideCursor/MoveTo helpers; progress bar renderer func Bar(pct int) string.
- Step 2 — list viewer: up/down keys via raw input (bufio) OR auto-tick demo if raw too hard.
- Step 3 — tick loop with select + context cancel.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_05_apps/40_tui")
func main() { fmt.Println(tuix.Demo()) }
```
```powershell
go run .
go vet ./phase_05_apps/40_tui/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
