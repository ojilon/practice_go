# CLI — Subcommands (flag only) — Guide (`clix`)

> Folder: `phase_05_apps/41_cli` · Package: `clix` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Only flag + os.Args (no cobra). Wire YOUR packages.
- Step 1 — mytool kv set/get/list (memory or file via minios).
- Step 2 — mytool serve --port (reuse httpapi), mytool parse <csv>.
- Step 3 — --help text + unknown-command error.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package clix` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_05_apps/41_cli"` (package name `clix`), call your funcs, `go run .`.
3. `go vet ./phase_05_apps/41_cli/...` must be clean before moving on.

### Details
- Only flag + os.Args (no cobra). Wire YOUR packages.
- Step 1 — mytool kv set/get/list (memory or file via minios).
- Step 2 — mytool serve --port (reuse httpapi), mytool parse <csv>.
- Step 3 — --help text + unknown-command error.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_05_apps/41_cli")
func main() { fmt.Println(clix.Demo()) }
```
```powershell
go run .
go vet ./phase_05_apps/41_cli/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
