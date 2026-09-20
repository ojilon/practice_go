# config — Env/Flags/Defaults — Guide (`config`)

> Folder: `phase_02_packages/21_config` · Package: `config` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Config struct{values map[string]string}.
- Step 2 — Load precedence: default < file < env < flag (document + implement Get with fallback chain).
- Step 3 — GetString/GetInt/GetBool with default + error on bad int.
- Step 4 — minimal .env file loader stretch.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package config` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/21_config"` (package name `config`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/21_config/...` must be clean before moving on.

### Details
- Step 1 — type Config struct{values map[string]string}.
- Step 2 — Load precedence: default < file < env < flag (document + implement Get with fallback chain).
- Step 3 — GetString/GetInt/GetBool with default + error on bad int.
- Step 4 — minimal .env file loader stretch.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/21_config")
func main() { fmt.Println(config.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/21_config/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
