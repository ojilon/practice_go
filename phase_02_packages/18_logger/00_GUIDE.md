# logger — Leveled Structured Logger — Guide (`logger`)

> Folder: `phase_02_packages/18_logger` · Package: `logger` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Level int const Debug/Info/Warn/Error; String().
- Step 2 — type Logger{out io.Writer, level Level, mu sync.Mutex}.
- Step 3 — Info/Errorf with timestamp + fields map; SetLevel; With(field) child.
- Step 4 — inject io.Writer (bytes.Buffer in tests/demos), no fmt.Print in lib.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package logger` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/18_logger"` (package name `logger`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/18_logger/...` must be clean before moving on.

### Details
- Step 1 — type Level int const Debug/Info/Warn/Error; String().
- Step 2 — type Logger{out io.Writer, level Level, mu sync.Mutex}.
- Step 3 — Info/Errorf with timestamp + fields map; SetLevel; With(field) child.
- Step 4 — inject io.Writer (bytes.Buffer in tests/demos), no fmt.Print in lib.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/18_logger")
func main() { fmt.Println(logger.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/18_logger/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
