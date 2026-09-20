# parsers — CSV + Logs — Guide (`parsers`)

> Folder: `phase_05_apps/38_parsers` · Package: `parsers` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Reuse tokenizer/ini/jsonlite ideas.
- Step 1 — ParseCSV(r io.Reader) (header, []Record, malformed-row errors with line numbers).
- Step 2 — ParseLog (Common Log Format-lite) → struct{IP,Time,Method,Path,Status}.
- Step 3 — summarize: status counts, top IP.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package parsers` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_05_apps/38_parsers"` (package name `parsers`), call your funcs, `go run .`.
3. `go vet ./phase_05_apps/38_parsers/...` must be clean before moving on.

### Details
- Reuse tokenizer/ini/jsonlite ideas.
- Step 1 — ParseCSV(r io.Reader) (header, []Record, malformed-row errors with line numbers).
- Step 2 — ParseLog (Common Log Format-lite) → struct{IP,Time,Method,Path,Status}.
- Step 3 — summarize: status counts, top IP.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_05_apps/38_parsers")
func main() { fmt.Println(parsers.Demo()) }
```
```powershell
go run .
go vet ./phase_05_apps/38_parsers/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
