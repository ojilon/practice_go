# ini — INI Parser — Guide (`ini`)

> Folder: `phase_02_packages/22_ini` · Package: `ini` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — lines: trim, skip empty + ;/# comments, [section], key=value.
- Step 2 — type File struct{sections map[string]map[string]string}.
- Step 3 — Parse(r io.Reader) + Get(sec,key)/GetDefault; quoted values + inline comment handling.
- Step 4 — Marshal back to string stretch.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package ini` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/22_ini"` (package name `ini`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/22_ini/...` must be clean before moving on.

### Details
- Step 1 — lines: trim, skip empty + ;/# comments, [section], key=value.
- Step 2 — type File struct{sections map[string]map[string]string}.
- Step 3 — Parse(r io.Reader) + Get(sec,key)/GetDefault; quoted values + inline comment handling.
- Step 4 — Marshal back to string stretch.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/22_ini")
func main() { fmt.Println(ini.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/22_ini/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
