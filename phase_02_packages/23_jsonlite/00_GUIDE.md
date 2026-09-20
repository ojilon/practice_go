# jsonlite — Mini JSON (no encoding/json in impl) — Guide (`jsonlite`)

> Folder: `phase_02_packages/23_jsonlite` · Package: `jsonlite` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- READ DOCS QUEST: go doc encoding/json first, then build subset WITHOUT importing encoding/json in non-demo code.
- Step 1 — Marshal: string/int/float/bool/nil, []any, map[string]any, flat structs via reflection-lite (or manual).
- Step 2 — Unmarshal: recursive descent parser (string with escapes, number, true/false/null, array, object).
- Step 3 — errors with byte offset; round-trip test.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package jsonlite` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/23_jsonlite"` (package name `jsonlite`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/23_jsonlite/...` must be clean before moving on.

### Details
- READ DOCS QUEST: go doc encoding/json first, then build subset WITHOUT importing encoding/json in non-demo code.
- Step 1 — Marshal: string/int/float/bool/nil, []any, map[string]any, flat structs via reflection-lite (or manual).
- Step 2 — Unmarshal: recursive descent parser (string with escapes, number, true/false/null, array, object).
- Step 3 — errors with byte offset; round-trip test.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/23_jsonlite")
func main() { fmt.Println(jsonlite.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/23_jsonlite/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
