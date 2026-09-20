# strings — Builder From Scratch — Guide (`stringsx`)

> Folder: `phase_03_dsa/26_strings_builder` · Package: `stringsx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Builder{buf []byte}; WriteString/WriteRune/WriteByte/String/Reset/Len.
- Step 2 — rune-correct Reverse/Words/Count; strings.Builder comparison.
- Step 3 — KMP substring stretch.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package stringsx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_03_dsa/26_strings_builder"` (package name `stringsx`), call your funcs, `go run .`.
3. `go vet ./phase_03_dsa/26_strings_builder/...` must be clean before moving on.

### Details
- Step 1 — type Builder{buf []byte}; WriteString/WriteRune/WriteByte/String/Reset/Len.
- Step 2 — rune-correct Reverse/Words/Count; strings.Builder comparison.
- Step 3 — KMP substring stretch.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_03_dsa/26_strings_builder")
func main() { fmt.Println(stringsx.Demo()) }
```
```powershell
go run .
go vet ./phase_03_dsa/26_strings_builder/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
