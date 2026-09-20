# reflection — Tags & Mini-Marshal — Guide (`reflectx`)

> Folder: `phase_04_systems/36_reflection` · Package: `reflectx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- DOCS QUEST: go doc reflect first (TypeOf/ValueOf/Kind).
- Step 1 — inspect struct fields/tags via reflect.
- Step 2 — GetTag(obj,field,tagKey) helper; FillMap(obj) map[string]any.
- Step 3 — mini Marshal using reflection for flat structs.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package reflectx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_04_systems/36_reflection"` (package name `reflectx`), call your funcs, `go run .`.
3. `go vet ./phase_04_systems/36_reflection/...` must be clean before moving on.

### Details
- DOCS QUEST: go doc reflect first (TypeOf/ValueOf/Kind).
- Step 1 — inspect struct fields/tags via reflect.
- Step 2 — GetTag(obj,field,tagKey) helper; FillMap(obj) map[string]any.
- Step 3 — mini Marshal using reflection for flat structs.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_04_systems/36_reflection")
func main() { fmt.Println(reflectx.Demo()) }
```
```powershell
go run .
go vet ./phase_04_systems/36_reflection/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
